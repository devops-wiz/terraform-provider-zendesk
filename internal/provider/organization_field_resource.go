package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/schema/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &OrganizationFieldResource{}
	_ resource.ResourceWithConfigure   = &OrganizationFieldResource{}
	_ resource.ResourceWithImportState = &OrganizationFieldResource{}
)

// NewOrganizationFieldResource is a helper function to simplify the provider implementation.
func NewOrganizationFieldResource() resource.Resource {
	return &OrganizationFieldResource{}
}

// OrganizationFieldResource is the resource implementation.
type OrganizationFieldResource struct {
	client *zendesk.Client
}

// Metadata returns the resource type name.
func (r *OrganizationFieldResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_field"
}

// Configure adds the provider configured client to the resource.
func (r *OrganizationFieldResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *OrganizationFieldResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = tfschema.GetUserOrgFieldSchema("org")
}

// Create creates the resource and sets the initial Terraform state.
// Create a new resource.
func (r *OrganizationFieldResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan models.OrganizationFieldResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	newField, diag := plan.GetApiModelFromTfModel(ctx)

	resp.Diagnostics.Append(diag...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create new org field
	field, err := r.client.CreateOrganizationField(ctx, newField)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Organization Field",
			"Could not create Organization Field, unexpected error: "+err.Error(),
		)
		return
	}

	// Bug on ZD side returns 9999 on new org creation
	if field.Position == 9999 {
		field.Position = newField.Position
	}

	// Map the response body to schema and populate Computed attribute values
	resp.Diagnostics.Append(plan.GetTfModelFromApiModel(ctx, field)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state to fully populated data
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *OrganizationFieldResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	models.ReadResource(ctx, req, resp, &models.OrganizationFieldResourceModel{}, r.client.GetOrganizationField)
}

func (r *OrganizationFieldResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	models.UpdateResource(ctx, req, resp, &models.OrganizationFieldResourceModel{}, r.client.UpdateOrganizationField)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *OrganizationFieldResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	models.DeleteResource[zendesk.OrganizationField](ctx, req, resp, &models.OrganizationFieldResourceModel{}, r.client.DeleteOrganizationField)
}

func (r *OrganizationFieldResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	models.ImportResource(ctx, req, resp, &models.OrganizationFieldResourceModel{}, r.client.GetOrganizationField)

}
