package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ ResourceTransformWithID[zendesk.Automation] = &AutomationResourceModel{}

type AutomationResourceModel struct {
	ID          types.Int64             `tfsdk:"id"`
	Actions     []ActionResourceModel   `tfsdk:"actions"`
	Active      types.Bool              `tfsdk:"active"`
	Description types.String            `tfsdk:"description"`
	Conditions  ConditionsResourceModel `tfsdk:"conditions"`
	CreatedAt   types.String            `tfsdk:"created_at"`
	Position    types.Int64             `tfsdk:"position"`
	Title       types.String            `tfsdk:"title"`
	UpdatedAt   types.String            `tfsdk:"updated_at"`
	URL         types.String            `tfsdk:"url"`
}
type AutomationResourceModelV0 struct {
	ID          types.Int64               `tfsdk:"id"`
	Actions     []ActionResourceModel     `tfsdk:"actions"`
	Active      types.Bool                `tfsdk:"active"`
	Description types.String              `tfsdk:"description"`
	Conditions  ConditionsResourceModelV0 `tfsdk:"conditions"`
	CreatedAt   types.String              `tfsdk:"created_at"`
	Position    types.Int64               `tfsdk:"position"`
	Title       types.String              `tfsdk:"title"`
	UpdatedAt   types.String              `tfsdk:"updated_at"`
	URL         types.String              `tfsdk:"url"`
}

func (a *AutomationResourceModel) GetID() int64 {
	return a.ID.ValueInt64()
}

// GetApiModelFromTfModel implements ResourceTransform.
func (a *AutomationResourceModel) GetApiModelFromTfModel(ctx context.Context) (newAutomation zendesk.Automation, diags diag.Diagnostics) {

	var conditions zendesk.Conditions

	conditions, diags = getApiConditionsFromTf(ctx, a.Conditions)

	if diags.HasError() {
		return newAutomation, diags
	}

	actionsTf, diags := getApiActionsFromTf(a.Actions)

	if diags.HasError() {
		return newAutomation, diags
	}

	newAutomation = zendesk.Automation{
		Title:       a.Title.ValueString(),
		Description: a.Description.ValueString(),
		Active:      a.Active.ValueBool(),
		Actions:     actionsTf,
		Conditions:  conditions,
	}

	if !a.Position.IsNull() && !a.Position.IsUnknown() {
		newAutomation.Position = a.Position.ValueInt64()
	}

	return newAutomation, diags
}

// GetTfModelFromApiModel implements ResourceTransform.
func (a *AutomationResourceModel) GetTfModelFromApiModel(ctx context.Context, apiAutomation zendesk.Automation) diag.Diagnostics {
	newTfActions, diags := getTfActionsFromApi(apiAutomation.Actions)

	if diags.HasError() {
		return diags
	}

	newTfConditions, diags := getTfConditionsFromApi(ctx, apiAutomation.Conditions)

	if diags.HasError() {
		return diags
	}

	*a = AutomationResourceModel{
		ID:          types.Int64Value(apiAutomation.ID),
		Title:       types.StringValue(apiAutomation.Title),
		Actions:     newTfActions,
		Conditions:  newTfConditions,
		Position:    types.Int64Value(apiAutomation.Position),
		Description: types.StringValue(apiAutomation.Description),
		Active:      types.BoolValue(apiAutomation.Active),
		CreatedAt:   types.StringValue(apiAutomation.CreatedAt.UTC().String()),
		UpdatedAt:   types.StringValue(apiAutomation.UpdatedAt.UTC().String()),
		URL:         types.StringValue(apiAutomation.URL),
	}

	return diags
}
