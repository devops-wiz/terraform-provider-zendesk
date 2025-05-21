package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/schema/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"time"
)

type DynamicContentResource struct {
	client *zendesk.Client
}

func NewDynamicContentResource() resource.Resource {
	return &DynamicContentResource{}
}

func (d *DynamicContentResource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_dynamic_content"
}

func (d *DynamicContentResource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = tfschema.DynamicContentSchema
}

func (d *DynamicContentResource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

	d.client = client
}

func (d *DynamicContentResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	models.CreateResource(ctx, request, response, &models.DynamicContentItemResourceModel{}, d.client.CreateDynamicContentItem)

}

func (d *DynamicContentResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	models.ReadResource(ctx, request, response, &models.DynamicContentItemResourceModel{}, d.client.GetDynamicContentItem)
}

func (d *DynamicContentResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var data models.DynamicContentItemResourceModel

	response.Diagnostics.Append(request.Plan.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	updatedDci, diags := data.GetApiModelFromTfModel(ctx)

	response.Diagnostics.Append(diags...)

	if response.Diagnostics.HasError() {
		return
	}

	dciResp, err := d.client.UpdateDynamicContentItem(ctx, data.ID.ValueInt64(), updatedDci)

	if err != nil {
		response.Diagnostics.AddError("Error updating dynamic content item", "Error updating dynamic content item: "+err.Error())
		return
	}

	_, err = d.client.UpdateDynamicContentVariants(ctx, data.ID.ValueInt64(), updatedDci.Variants)

	if err != nil {
		response.Diagnostics.AddError("Error updating dynamic content variants", "Error updating dynamic content variants: "+err.Error())
		return
	}

	dciResp.Variants = updatedDci.Variants

	response.Diagnostics.Append(data.GetTfModelFromApiModel(ctx, dciResp)...)

	if response.Diagnostics.HasError() {
		return
	}

	time.Sleep(10 * time.Second)

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)

}

func (d *DynamicContentResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	models.DeleteResource[zendesk.DynamicContentItem](ctx, request, response, &models.DynamicContentItemResourceModel{}, d.client.DeleteDynamicContentItem)
}

func (d *DynamicContentResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	models.ImportResource(ctx, request, response, &models.DynamicContentItemResourceModel{}, d.client.GetDynamicContentItem)
}
