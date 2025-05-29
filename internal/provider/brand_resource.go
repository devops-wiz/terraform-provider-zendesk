package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithConfigure = &BrandResource{}
var _ resource.ResourceWithImportState = &BrandResource{}

type BrandResource struct {
	client *zendesk.Client
}

func NewBrandResource() resource.Resource {
	return &BrandResource{}
}

func (b *BrandResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

	b.client = client
}

func (b *BrandResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_brand"
}

func (b *BrandResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = BrandSchema
}

func (b *BrandResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	models.CreateResource(ctx, request, response, &models.BrandResourceModel{}, b.client.CreateBrand)
}

func (b *BrandResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	models.ReadResource(ctx, request, response, &models.BrandResourceModel{}, b.client.GetBrand)
}

func (b *BrandResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	models.UpdateResource(ctx, request, response, &models.BrandResourceModel{}, b.client.UpdateBrand)

}

func (b *BrandResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	models.DeleteResource[zendesk.Brand](ctx, request, response, &models.BrandResourceModel{}, b.client.DeleteBrand)
}

func (b *BrandResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	models.ImportResource(ctx, request, response, &models.BrandResourceModel{}, b.client.GetBrand)
}
