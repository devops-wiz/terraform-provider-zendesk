package provider

import (
	"context"
	"fmt"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.Resource = &WebhookResource{}
var _ resource.ResourceWithImportState = &WebhookResource{}

type WebhookResource struct {
	client *zendesk.Client
}

func NewWebhookResource() resource.Resource {
	return &WebhookResource{}
}

// Metadata implements resource.Resource.
func (w *WebhookResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhook"
}

// Schema implements resource.Resource.
func (w *WebhookResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = WebhookSchema
}

func (w *WebhookResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*zendesk.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	w.client = client
}

// Create implements resource.Resource.
func (w *WebhookResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.WebhookResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newWebhook, diags := data.GetApiModelFromTfModel(ctx)

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	webhookResp, err := w.client.CreateWebhook(ctx, newWebhook)

	if err != nil {
		resp.Diagnostics.AddError("Error creating webhook", err.Error())
		return
	}

	secretResp, err := w.client.GetWebhookSigningSecret(ctx, webhookResp.ID)

	if err != nil {
		resp.Diagnostics.AddError("Error getting webhook secret", err.Error())
		return
	}
	// in order to make sure API response doesn't overwrite with empty credential values
	webhookResp.Authentication = newWebhook.Authentication
	webhookResp.SigningSecret = &secretResp

	resp.Diagnostics.Append(data.GetTfModelFromApiModel(ctx, webhookResp)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// save data into tf state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read implements resource.Resource.
func (w *WebhookResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.WebhookResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	webhookResp, err := w.client.GetWebhook(ctx, data.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Unable to read Zendesk Webhook", err.Error())
		return
	}

	secretResp, err := w.client.GetWebhookSigningSecret(ctx, webhookResp.ID)

	if err != nil {
		resp.Diagnostics.AddError("Error getting webhook secret", err.Error())
		return
	}

	var diags diag.Diagnostics

	if !data.Authentication.IsNull() && !data.Authentication.IsUnknown() {
		webhookResp.Authentication, diags = models.GetApiWebhookAuthenticationFromTf(ctx, data.Authentication)
	}

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	webhookResp.SigningSecret = &secretResp

	resp.Diagnostics.Append(data.GetTfModelFromApiModel(ctx, webhookResp)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

// Update implements resource.Resource.
func (w *WebhookResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.WebhookResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	updatedWebhook, _ := data.GetApiModelFromTfModel(ctx)

	err := w.client.UpdateWebhook(ctx, data.ID.ValueString(), updatedWebhook)

	if err != nil {
		resp.Diagnostics.AddError("Error updating webhook", err.Error())
		return
	}

	webhookResp, err := w.client.GetWebhook(ctx, data.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Error retrieving updated webhook", err.Error())
		return
	}

	secretResp, err := w.client.GetWebhookSigningSecret(ctx, webhookResp.ID)

	if err != nil {
		resp.Diagnostics.AddError("Error getting webhook secret", err.Error())
		return
	}

	webhookResp.Authentication = updatedWebhook.Authentication
	webhookResp.SigningSecret = &secretResp

	resp.Diagnostics.Append(data.GetTfModelFromApiModel(ctx, webhookResp)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// save data into tf state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete implements resource.Resource.
func (w *WebhookResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.WebhookResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	err := w.client.DeleteWebhook(ctx, data.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Webhook",
			"Could not delete Webhook, unexpected error: "+err.Error(),
		)
		return
	}
}

// ImportState implements resource.ResourceWithImportState.
func (w *WebhookResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data models.WebhookResourceModel

	webhookResp, err := w.client.GetWebhook(ctx, req.ID)

	if err != nil {
		resp.Diagnostics.AddError("Error retrieving updated webhook", err.Error())
		return
	}

	secretResp, err := w.client.GetWebhookSigningSecret(ctx, webhookResp.ID)

	if err != nil {
		resp.Diagnostics.AddError("Error getting webhook secret", err.Error())
		return
	}

	webhookResp.SigningSecret = &secretResp

	resp.Diagnostics.Append(data.GetTfModelFromApiModel(ctx, webhookResp)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
