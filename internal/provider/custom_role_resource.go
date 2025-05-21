package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/schema/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithConfigure = &CustomRoleResource{}
var _ resource.ResourceWithImportState = &CustomRoleResource{}

type CustomRoleResource struct {
	client *zendesk.Client
}

func NewCustomRoleResource() resource.Resource {
	return &CustomRoleResource{}
}

func (c *CustomRoleResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_custom_role"
}

func (c *CustomRoleResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = tfschema.CustomRoleSchema
}

func (c *CustomRoleResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	client, ok := request.ProviderData.(*zendesk.Client)

	if !ok {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)

		return
	}

	c.client = client
}

func (c *CustomRoleResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	models.CreateResource(ctx, request, response, &models.CustomRoleResourceModel{}, c.client.CreateCustomRole)
}

func (c *CustomRoleResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	models.ReadResource(ctx, request, response, &models.CustomRoleResourceModel{}, c.client.GetCustomRole)
}

func (c *CustomRoleResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	models.UpdateResource(ctx, request, response, &models.CustomRoleResourceModel{}, c.client.UpdateCustomRole)
}

func (c *CustomRoleResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	models.DeleteResource[zendesk.CustomRole](ctx, request, response, &models.CustomRoleResourceModel{}, c.client.DeleteCustomRole)
}

func (c *CustomRoleResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	models.ImportResource(ctx, request, response, &models.CustomRoleResourceModel{}, c.client.GetCustomRole)
}
