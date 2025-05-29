package provider

import (
	"context"
	"fmt"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/hashicorp/terraform-plugin-framework/path"

	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                   = &TicketFieldResource{}
	_ resource.ResourceWithConfigure      = &TicketFieldResource{}
	_ resource.ResourceWithImportState    = &TicketFieldResource{}
	_ resource.ResourceWithValidateConfig = &TicketFieldResource{}
)

// NewTicketFieldResource is a helper function to simplify the provider implementation.
func NewTicketFieldResource() resource.Resource {
	return &TicketFieldResource{}
}

// TicketFieldResource is the resource implementation.
type TicketFieldResource struct {
	client *zendesk.Client
}

// Metadata returns the resource type name.
func (r *TicketFieldResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ticket_field"
}

// Configure adds the provider configured client to the resource.
func (r *TicketFieldResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = client
}

// Schema defines the schema for the resource.
func (r *TicketFieldResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = TicketFieldSchema
}

// Create creates the resource and sets the initial Terraform state.
// Create a new resource.
func (r *TicketFieldResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	models.CreateResource(ctx, req, resp, &models.TicketFieldResourceModel{}, r.client.CreateTicketField)
}

// Read resource information.
func (r *TicketFieldResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	models.ReadResource(ctx, req, resp, &models.TicketFieldResourceModel{}, r.client.GetTicketField)
}

func (r *TicketFieldResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	models.UpdateResource(ctx, req, resp, &models.TicketFieldResourceModel{}, r.client.UpdateTicketField)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *TicketFieldResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	models.DeleteResource[zendesk.TicketField](ctx, req, resp, &models.TicketFieldResourceModel{}, r.client.DeleteTicketField)
}

func (r *TicketFieldResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	models.ImportResource(ctx, req, resp, &models.TicketFieldResourceModel{}, r.client.GetTicketField)

}

func (r *TicketFieldResource) ValidateConfig(ctx context.Context, request resource.ValidateConfigRequest, response *resource.ValidateConfigResponse) {
	var data models.TicketFieldResourceModel

	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	if data.VisibleInPortal.IsNull() && (!data.EditableInPortal.IsNull() || !data.RequiredInPortal.IsNull()) {
		response.Diagnostics.AddAttributeError(
			path.Root("visible_in_portal"),
			"Invalid attribute combination",
			"For a ticket field to be Editable or Required in portal, visible_in_portal must be set to true",
		)
		return
	}

	if data.VisibleInPortal.ValueBool() && data.EditableInPortal.IsNull() && !data.RequiredInPortal.IsNull() {
		response.Diagnostics.AddAttributeError(
			path.Root("editable_in_portal"),
			"Invalid attribute combination",
			"For a ticket field to be Required in portal, editable_in_portal must be set to true",
		)
		return
	}
}
