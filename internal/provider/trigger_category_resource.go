package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/schema/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.Resource = &TriggerCategoryResource{}
var _ resource.ResourceWithImportState = &TriggerCategoryResource{}

type TriggerCategoryResource struct {
	client *zendesk.Client
}

func NewTriggerCategoryResource() resource.Resource {
	return &TriggerCategoryResource{}
}

// Metadata implements resource.Resource.
func (t *TriggerCategoryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_trigger_category"
}

// Schema implements resource.Resource.
func (t *TriggerCategoryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = tfschema.TriggerCategorySchema
}

func (t *TriggerCategoryResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*zendesk.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *zendesk.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	t.client = client
}

// Create implements resource.Resource.
func (t *TriggerCategoryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	models.CreateResource(ctx, req, resp, &models.TriggerCategoryResourceModel{}, t.client.CreateTriggerCategory)
}

// Read implements resource.Resource.
func (t *TriggerCategoryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	models.ReadResource(ctx, req, resp, &models.TriggerCategoryResourceModel{}, t.client.GetTriggerCategory)
}

// Update implements resource.Resource.
func (t *TriggerCategoryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	models.UpdateResource(ctx, req, resp, &models.TriggerCategoryResourceModel{}, t.client.UpdateTriggerCategory)
}

// Delete implements resource.Resource.
func (t *TriggerCategoryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	models.DeleteResource[zendesk.TriggerCategory](ctx, req, resp, &models.TriggerCategoryResourceModel{}, t.client.DeleteTriggerCategory)
}

// ImportState implements resource.ResourceWithImportState.
func (t *TriggerCategoryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	models.ImportResource(ctx, req, resp, &models.TriggerCategoryResourceModel{}, t.client.GetTriggerCategory)
}
