package models

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"strconv"
)

type ResourceTransform[M any] interface {
	GetApiModelFromTfModel(context.Context) (M, diag.Diagnostics)
	GetTfModelFromApiModel(context.Context, M) diag.Diagnostics
}
type ResourceTransformWithID[M any] interface {
	GetID() int64
	ResourceTransform[M]
}

type DatasourceTransform[M any] interface {
	GetTfModelFromApiModel(context.Context, M) diag.Diagnostics
}

func CreateResource[M any](ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse, resourceModel ResourceTransformWithID[M], createFunc func(ctx context.Context, newResource M) (M, error)) {
	response.Diagnostics.Append(request.Plan.Get(ctx, resourceModel)...)

	if response.Diagnostics.HasError() {
		return
	}
	newResource, diags := resourceModel.GetApiModelFromTfModel(ctx)

	response.Diagnostics.Append(diags...)

	if response.Diagnostics.HasError() {
		return
	}

	resp, err := createFunc(ctx, newResource)

	if err != nil {
		response.Diagnostics.AddError("Error creating resource", fmt.Sprintf("Error: %s", err))
		return
	}

	response.Diagnostics.Append(resourceModel.GetTfModelFromApiModel(ctx, resp)...)

	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, resourceModel)...)
}

func ReadResource[M any](ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse, resourceModel ResourceTransformWithID[M], readFunc func(ctx context.Context, id int64) (M, error)) {

	response.Diagnostics.Append(request.State.Get(ctx, resourceModel)...)

	if response.Diagnostics.HasError() {
		return
	}

	resp, err := readFunc(ctx, resourceModel.GetID())

	if err != nil {
		response.Diagnostics.AddError("Error reading resource", "Error reading resource: "+err.Error())
		return
	}

	response.Diagnostics.Append(resourceModel.GetTfModelFromApiModel(ctx, resp)...)

	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, resourceModel)...)

}

func UpdateResource[M any](ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse, resourceModel ResourceTransformWithID[M], updateFunc func(ctx context.Context, updatedResourceId int64, updatedResource M) (M, error)) {
	response.Diagnostics.Append(request.Plan.Get(ctx, resourceModel)...)

	if response.Diagnostics.HasError() {
		return
	}

	updatedResource, diags := resourceModel.GetApiModelFromTfModel(ctx)

	response.Diagnostics.Append(diags...)

	if response.Diagnostics.HasError() {
		return
	}

	resp, err := updateFunc(ctx, resourceModel.GetID(), updatedResource)

	if err != nil {
		response.Diagnostics.AddError("Error updating resource", fmt.Sprintf("Error: %s", err))
		return
	}

	response.Diagnostics.Append(resourceModel.GetTfModelFromApiModel(ctx, resp)...)

	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, resourceModel)...)

}

func DeleteResource[M any](ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse, resourceModel ResourceTransformWithID[M], deleteFunc func(ctx context.Context, id int64) error) {
	response.Diagnostics.Append(request.State.Get(ctx, resourceModel)...)

	if response.Diagnostics.HasError() {
		return
	}

	err := deleteFunc(ctx, resourceModel.GetID())
	if err != nil {
		response.Diagnostics.AddError("Error deleting resource", fmt.Sprintf("Error: %s", err))
		return
	}
}

func ImportResource[M any](ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse, resourceModel ResourceTransformWithID[M], getFunc func(ctx context.Context, id int64) (M, error)) {
	importId, err := strconv.ParseInt(request.ID, 10, 64)

	if err != nil {
		response.Diagnostics.AddError("Unable to convert import id", fmt.Sprintf("error converting value %s to int64", request.ID))
		return
	}

	resp, err := getFunc(ctx, importId)

	if err != nil {
		response.Diagnostics.AddError("Error importing resource", fmt.Sprintf("Error importing resource: %s", err))
		return
	}

	response.Diagnostics.Append(resourceModel.GetTfModelFromApiModel(ctx, resp)...)

	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, resourceModel)...)
}
