package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.Resource = &ViewResource{}
var _ resource.ResourceWithImportState = &ViewResource{}
var _ resource.ResourceWithValidateConfig = &ViewResource{}

type ViewResource struct {
	client *zendesk.Client
}

func NewViewResource() resource.Resource {
	return &ViewResource{}
}

func (v *ViewResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	v.client = client
}

// Metadata implements resource.ResourceWithImportState.
func (v *ViewResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_view"

}

// Schema implements resource.ResourceWithImportState.
func (v *ViewResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ViewSchema
}

// Create implements resource.ResourceWithImportState.
func (v *ViewResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	models.CreateResource(ctx, req, resp, &models.ViewResourceModel{}, v.client.CreateView)
}

// Read implements resource.ResourceWithImportState.
func (v *ViewResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	models.ReadResource(ctx, req, resp, &models.ViewResourceModel{}, v.client.GetView)
}

// Update implements resource.ResourceWithImportState.
func (v *ViewResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	models.UpdateResource(ctx, req, resp, &models.ViewResourceModel{}, v.client.UpdateView)
}

// Delete implements resource.ResourceWithImportState.
func (v *ViewResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	models.DeleteResource[zendesk.View](ctx, req, resp, &models.ViewResourceModel{}, v.client.DeleteView)
}

// ImportState implements resource.ResourceWithImportState.
func (v *ViewResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	models.ImportResource(ctx, req, resp, &models.ViewResourceModel{}, v.client.GetView)
}

func (v *ViewResource) ValidateConfig(ctx context.Context, request resource.ValidateConfigRequest, response *resource.ValidateConfigResponse) {
	var data models.ViewResourceModel

	validAllConditions := []zendesk.ConditionField{
		zendesk.ConditionFieldStatus,
		zendesk.ConditionFieldType,
		zendesk.ConditionFieldGroupID,
		zendesk.ConditionFieldAssigneeID,
		zendesk.ConditionFieldRequesterID,
	}

	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	all := data.Conditions.All

	requiredConditions := utils.SliceFilter(all, func(model models.ConditionResourceModel) bool {
		return slices.Contains(validAllConditions, zendesk.ConditionField(model.Field.ValueString()))
	})

	if len(requiredConditions) == 0 {
		response.Diagnostics.AddAttributeError(
			path.Root("all"),
			"Missing condition from \"All\" Conditions",
			fmt.Sprintf("Views require at least one of the following fields in the \"All\" conditions: %v", validAllConditions),
		)
	}

}

func (v *ViewResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// State upgrade implementation from 0 (prior state version) to 1 (Schema.Version)
		0: {
			PriorSchema: &ViewSchemaV0,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var priorStateData models.ViewResourceModelV0

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)

				if resp.Diagnostics.HasError() {
					return
				}

				upgradedStateData := models.ViewResourceModel{
					ID:          priorStateData.ID,
					URL:         priorStateData.URL,
					Title:       priorStateData.Title,
					Description: priorStateData.Description,
					Active:      priorStateData.Active,
					UpdatedAt:   priorStateData.UpdatedAt,
					CreatedAt:   priorStateData.CreatedAt,
					Position:    priorStateData.Position,
					Output:      priorStateData.Output,
					Restriction: priorStateData.Restriction,
				}

				if priorStateData.Conditions != nil {
					priorAll := priorStateData.Conditions.All
					priorAny := priorStateData.Conditions.Any

					conditions := models.UpgradeConditionsV1(priorAll, priorAny)

					upgradedStateData.Conditions = &conditions
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, upgradedStateData)...)
			},
		},
	}
}
