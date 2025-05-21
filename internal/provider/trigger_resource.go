package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/schema/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.Resource = &TriggerResource{}
var _ resource.ResourceWithImportState = &TriggerResource{}
var _ resource.ResourceWithUpgradeState = &TriggerResource{}

type TriggerResource struct {
	client *zendesk.Client
}

func NewTriggerResource() resource.Resource {
	return &TriggerResource{}
}

// Metadata implements resource.Resource.
func (t *TriggerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_trigger"
}

// Schema implements resource.Resource.
func (t *TriggerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = tfschema.TriggerSchema
}

func (t *TriggerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	t.client = client
}

// Create implements resource.Resource.
func (t *TriggerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	models.CreateResource(ctx, req, resp, &models.TriggerResourceModel{}, t.client.CreateTrigger)
}

// Read implements resource.Resource.
func (t *TriggerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	models.ReadResource(ctx, req, resp, &models.TriggerResourceModel{}, t.client.GetTrigger)
}

// Update implements resource.Resource.
func (t *TriggerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	models.UpdateResource(ctx, req, resp, &models.TriggerResourceModel{}, t.client.UpdateTrigger)
}

// Delete implements resource.Resource.
func (t *TriggerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	models.DeleteResource[zendesk.Trigger](ctx, req, resp, &models.TriggerResourceModel{}, t.client.DeleteTrigger)
}

// ImportState implements resource.ResourceWithImportState.
func (t *TriggerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	models.ImportResource(ctx, req, resp, &models.TriggerResourceModel{}, t.client.GetTrigger)
}

func (t *TriggerResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// State upgrade implementation from 0 (prior state version) to 1 (Schema.Version)
		0: {
			PriorSchema: &tfschema.TriggerSchemaV0,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var priorStateData models.TriggerResourceModelV0

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)

				if resp.Diagnostics.HasError() {
					return
				}

				priorAll := priorStateData.Conditions.All
				priorAny := priorStateData.Conditions.Any

				conditions := models.UpgradeConditionsV1(priorAll, priorAny)

				upgradedStateData := models.TriggerResourceModel{
					ID:          priorStateData.ID,
					Actions:     priorStateData.Actions,
					Active:      priorStateData.Active,
					Description: priorStateData.Description,
					CategoryID:  priorStateData.CategoryID,
					Conditions:  conditions,
					CreatedAt:   priorStateData.CreatedAt,
					Position:    priorStateData.Position,
					Title:       priorStateData.Title,
					UpdatedAt:   priorStateData.UpdatedAt,
					URL:         priorStateData.URL,
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, upgradedStateData)...)
			},
		},
	}
}
