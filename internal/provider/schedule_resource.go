package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/schema/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type ScheduleResource struct {
	client *zendesk.Client
}

func NewScheduleResource() resource.Resource {
	return &ScheduleResource{}
}

func (s *ScheduleResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_schedule"

}

func (s *ScheduleResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	client, ok := request.ProviderData.(*zendesk.Client)

	if !ok {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *zendesk.Client, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)

		return
	}

	s.client = client
}

func (s *ScheduleResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = tfschema.ScheduleSchema
}

func (s *ScheduleResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var data models.ScheduleResourceModel

	response.Diagnostics.Append(request.Plan.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	newSchedule, diags := data.GetApiModelFromTfModel(ctx)

	response.Diagnostics.Append(diags...)

	if response.Diagnostics.HasError() {
		return
	}

	scheduleResp, err := s.client.CreateSchedule(ctx, newSchedule)

	if err != nil {
		response.Diagnostics.AddError("Error while creating schedule", "Error creating schedule: "+err.Error())
		return
	}

	intervalResp, err := s.client.UpdateScheduleIntervals(ctx, int64(scheduleResp.Id), newSchedule.Intervals)

	if err != nil {
		response.Diagnostics.AddError("Error while creating schedule", "Error updating schedule intervals: "+err.Error())
	}

	scheduleResp.Intervals = intervalResp

	response.Diagnostics.Append(data.GetTfModelFromApiModel(ctx, scheduleResp)...)

	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, data)...)
}

func (s *ScheduleResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	models.ReadResource(ctx, request, response, &models.ScheduleResourceModel{}, s.client.GetSchedule)
}

func (s *ScheduleResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var data models.ScheduleResourceModel

	response.Diagnostics.Append(request.Plan.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	updatedSchedule, diags := data.GetApiModelFromTfModel(ctx)

	response.Diagnostics.Append(diags...)

	if response.Diagnostics.HasError() {
		return
	}

	scheduleResp, err := s.client.UpdateSchedule(ctx, data.ID.ValueInt64(), updatedSchedule)
	if err != nil {
		response.Diagnostics.AddError("Error while updating schedule", "Error updating schedule: "+err.Error())
		return
	}

	intervalResp, err := s.client.UpdateScheduleIntervals(ctx, int64(scheduleResp.Id), updatedSchedule.Intervals)

	if err != nil {
		response.Diagnostics.AddError("Error while creating schedule", "Error updating schedule intervals: "+err.Error())
	}

	scheduleResp.Intervals = intervalResp

	response.Diagnostics.Append(data.GetTfModelFromApiModel(ctx, scheduleResp)...)
	if response.Diagnostics.HasError() {
		return
	}
	response.Diagnostics.Append(response.State.Set(ctx, data)...)
}

func (s *ScheduleResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	models.DeleteResource[zendesk.Schedule](ctx, request, response, &models.ScheduleResourceModel{}, s.client.DeleteSchedule)

}

func (s *ScheduleResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	models.ImportResource(ctx, request, response, &models.ScheduleResourceModel{}, s.client.GetSchedule)
}
