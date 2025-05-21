// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/schema/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MacroResource{}
var _ resource.ResourceWithImportState = &MacroResource{}

func NewMacroResource() resource.Resource {
	return &MacroResource{}
}

// MacroResource defines the resource implementation.
type MacroResource struct {
	client *zendesk.Client
}

func (r *MacroResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_macro"
}

func (r *MacroResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MacroResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = tfschema.MacroSchema
}

func (r *MacroResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	models.CreateResource(ctx, req, resp, &models.MacroResourceModel{}, r.client.CreateMacro)
}

func (r *MacroResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	models.ReadResource(ctx, req, resp, &models.MacroResourceModel{}, r.client.GetMacro)
}

func (r *MacroResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	models.UpdateResource(ctx, req, resp, &models.MacroResourceModel{}, r.client.UpdateMacro)
}

func (r *MacroResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	models.DeleteResource[zendesk.Macro](ctx, req, resp, &models.MacroResourceModel{}, r.client.DeleteMacro)
}

func (r *MacroResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	models.ImportResource(ctx, req, resp, &models.MacroResourceModel{}, r.client.GetMacro)
}
