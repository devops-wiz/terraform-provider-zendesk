package provider

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"slices"
)

var _ resource.Resource = &AutomationResource{}
var _ resource.ResourceWithImportState = &AutomationResource{}
var _ resource.ResourceWithValidateConfig = &AutomationResource{}
var _ resource.ResourceWithUpgradeState = &AutomationResource{}

type AutomationResource struct {
	client *zendesk.Client
}

func NewAutomationResource() resource.Resource {
	return &AutomationResource{}
}

// Metadata implements resource.Resource.
func (t *AutomationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_automation"
}

// Schema implements resource.Resource.
func (t *AutomationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = AutomationSchema
}

func (t *AutomationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	t.client = client
}

// Create implements resource.Resource.
func (t *AutomationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	models.CreateResource(ctx, req, resp, &models.AutomationResourceModel{}, t.client.CreateAutomation)
}

// Read implements resource.Resource.
func (t *AutomationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	models.ReadResource(ctx, req, resp, &models.AutomationResourceModel{}, t.client.GetAutomation)
}

// Update implements resource.Resource.
func (t *AutomationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	models.UpdateResource(ctx, req, resp, &models.AutomationResourceModel{}, t.client.UpdateAutomation)
}

// Delete implements resource.Resource.
func (t *AutomationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	models.DeleteResource[zendesk.Automation](ctx, req, resp, &models.AutomationResourceModel{}, t.client.DeleteAutomation)
}

// ImportState implements resource.ResourceWithImportState.
func (t *AutomationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	models.ImportResource(ctx, req, resp, &models.AutomationResourceModel{}, t.client.GetAutomation)
}

// ValidateConfig Validates config for Automation Resource.
func (t *AutomationResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AutomationResourceModel

	// Secondary required conditions for automations
	automationSecondaryRequiredConditions := []zendesk.ConditionField{
		zendesk.ConditionFieldStatus,
		zendesk.ConditionFieldType,
		zendesk.ConditionFieldGroupID,
		zendesk.ConditionFieldAssigneeID,
		zendesk.ConditionFieldRequesterID,
	}

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	conditions := data.Conditions

	timeBasedConditions := utils.SliceFilter(conditions.All, func(c models.ConditionResourceModel) bool {

		field := zendesk.ConditionField(c.Field.ValueString())

		if field == "custom_field" {
			field = "custom_fields_"
		}

		return slices.Contains(zendesk.TimeBasedConditions, field)
	})

	if len(timeBasedConditions) == 0 {
		resp.Diagnostics.AddAttributeError(
			path.Root("conditions"),
			"All conditions must contain at least one time-based condition",
			fmt.Sprintf("Automations require at least one \"All\" condition, with at least one time based condition: %v ", zendesk.TimeBasedConditions),
		)
		return
	}

	// verify at least one secondary automation "All" condition is present
	secondaryAutoIdx := slices.IndexFunc(conditions.All, func(c models.ConditionResourceModel) bool {
		tflog.Info(ctx, "Validating automation", map[string]interface{}{
			"field": c.Field.ValueString(),
		})
		return slices.Contains(automationSecondaryRequiredConditions, zendesk.ConditionField(c.Field.ValueString()))
	})

	if secondaryAutoIdx == -1 {
		resp.Diagnostics.AddAttributeError(
			path.Root("conditions"),
			"Missing condition from \"All\" Conditions",
			fmt.Sprintf(
				"Automations require at least one of the following fields in the \"All\" conditions: %v",
				automationSecondaryRequiredConditions),
		)
		return
	}
}

func (t *AutomationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// State upgrade implementation from 0 (prior state version) to 1 (Schema.Version)
		0: {
			PriorSchema: &AutomationSchemaV0,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var priorStateData models.AutomationResourceModelV0

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)

				if resp.Diagnostics.HasError() {
					return
				}

				priorAll := priorStateData.Conditions.All
				priorAny := priorStateData.Conditions.Any

				conditions := models.UpgradeConditionsV1(priorAll, priorAny)

				upgradedStateData := models.AutomationResourceModel{
					ID:          priorStateData.ID,
					Actions:     priorStateData.Actions,
					Active:      priorStateData.Active,
					Description: priorStateData.Description,
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
