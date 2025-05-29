package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithImportState = &GroupResource{}
var _ resource.ResourceWithConfigure = &GroupResource{}

type GroupResource struct {
	client *zendesk.Client
}

func NewGroupResource() resource.Resource {
	return &GroupResource{}
}

func (g *GroupResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

	g.client = client
}

func (g *GroupResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_group"
}

func (g *GroupResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = GroupSchema
}

func (g *GroupResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	models.CreateResource(ctx, request, response, &models.GroupResourceModel{}, g.client.CreateGroup)
}

func (g *GroupResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	models.ReadResource(ctx, request, response, &models.GroupResourceModel{}, g.client.GetGroup)
}

func (g *GroupResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	models.UpdateResource(ctx, request, response, &models.GroupResourceModel{}, g.client.UpdateGroup)
}

func (g *GroupResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	models.DeleteResource[zendesk.Group](ctx, request, response, &models.GroupResourceModel{}, g.client.DeleteGroup)
}

func (g *GroupResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	models.ImportResource(ctx, request, response, &models.GroupResourceModel{}, g.client.GetGroup)
}
