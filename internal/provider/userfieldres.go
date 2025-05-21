package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &UserFieldResource{}
	_ resource.ResourceWithConfigure   = &UserFieldResource{}
	_ resource.ResourceWithImportState = &UserFieldResource{}
)

// NewUserFieldResource is a helper function to simplify the provider implementation.
func NewUserFieldResource() resource.Resource {
	return &UserFieldResource{}
}

// UserFieldResource is the resource implementation.
type UserFieldResource struct {
	client *zendesk.Client
}

// Metadata returns the resource type name.
func (r *UserFieldResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_field"
}

// Configure adds the provider configured client to the resource.
func (r *UserFieldResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *UserFieldResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = GetUserOrgFieldSchema("user")
}

// Create creates the resource and sets the initial Terraform state.
// Create a new resource.
func (r *UserFieldResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	models.CreateResource(ctx, req, resp, &models.UserFieldResourceModel{}, r.client.CreateUserField)
}

// Read resource information.
func (r *UserFieldResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	models.ReadResource(ctx, req, resp, &models.UserFieldResourceModel{}, r.client.GetUserField)
}

func (r *UserFieldResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	models.UpdateResource(ctx, req, resp, &models.UserFieldResourceModel{}, r.client.UpdateUserField)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *UserFieldResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	models.DeleteResource[zendesk.UserField](ctx, req, resp, &models.UserFieldResourceModel{}, r.client.DeleteUserField)
}

func (r *UserFieldResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	models.ImportResource(ctx, req, resp, &models.UserFieldResourceModel{}, r.client.GetUserField)
}
