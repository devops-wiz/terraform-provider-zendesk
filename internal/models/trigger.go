package models

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"strconv"

	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ ResourceTransformWithID[zendesk.Trigger] = &TriggerResourceModel{}

type TriggerResourceModel struct {
	ID          types.Int64             `tfsdk:"id"`
	Actions     []ActionResourceModel   `tfsdk:"actions"`
	Active      types.Bool              `tfsdk:"active"`
	Description types.String            `tfsdk:"description"`
	CategoryID  types.Int64             `tfsdk:"category_id"`
	Conditions  ConditionsResourceModel `tfsdk:"conditions"`
	CreatedAt   types.String            `tfsdk:"created_at"`
	Position    types.Int64             `tfsdk:"position"`
	Title       types.String            `tfsdk:"title"`
	UpdatedAt   types.String            `tfsdk:"updated_at"`
	URL         types.String            `tfsdk:"url"`
}

type TriggerResourceModelV0 struct {
	ID          types.Int64               `tfsdk:"id"`
	Actions     []ActionResourceModel     `tfsdk:"actions"`
	Active      types.Bool                `tfsdk:"active"`
	Description types.String              `tfsdk:"description"`
	CategoryID  types.Int64               `tfsdk:"category_id"`
	Conditions  ConditionsResourceModelV0 `tfsdk:"conditions"`
	CreatedAt   types.String              `tfsdk:"created_at"`
	Position    types.Int64               `tfsdk:"position"`
	Title       types.String              `tfsdk:"title"`
	UpdatedAt   types.String              `tfsdk:"updated_at"`
	URL         types.String              `tfsdk:"url"`
}

func (t *TriggerResourceModel) GetID() int64 {
	return t.ID.ValueInt64()
}

// GetApiModelFromTfModel Maps API object from TF model
func (t *TriggerResourceModel) GetApiModelFromTfModel(ctx context.Context) (zendesk.Trigger, diag.Diagnostics) {

	var newTrigger zendesk.Trigger

	newActions, diags := getApiActionsFromTf(t.Actions)

	if diags.HasError() {
		return newTrigger, diags
	}

	newConditions, diags := getApiConditionsFromTf(ctx, t.Conditions)

	if diags.HasError() {
		return newTrigger, diags
	}

	newTrigger = zendesk.Trigger{
		Title:       t.Title.ValueString(),
		Description: t.Description.ValueString(),
		Active:      t.Active.ValueBool(),
		CategoryID:  fmt.Sprintf("%d", t.CategoryID.ValueInt64()),
		Actions:     newActions,
		Conditions:  newConditions,
	}

	if !t.Position.IsNull() && !t.Position.IsUnknown() {
		newTrigger.Position = t.Position.ValueInt64()
	}

	return newTrigger, diags
}

// GetTfModelFromApiModel implements ResourceTransform.
func (t *TriggerResourceModel) GetTfModelFromApiModel(ctx context.Context, apiTrigger zendesk.Trigger) diag.Diagnostics {
	newTfActions, diags := getTfActionsFromApi(apiTrigger.Actions)

	if diags.HasError() {
		return diags
	}

	newTfConditions, diags := getTfConditionsFromApi(ctx, apiTrigger.Conditions)

	if diags.HasError() {
		return diags
	}

	convertedCatId, err := strconv.ParseInt(apiTrigger.CategoryID, 10, 64)

	if err != nil {
		diags.AddAttributeError(path.Root("category_id"), "Error converting trigger category ID", fmt.Sprintf("%v cannot be converted to an int64", apiTrigger.CategoryID))
		return diags
	}

	*t = TriggerResourceModel{
		ID:          types.Int64Value(apiTrigger.ID),
		Title:       types.StringValue(apiTrigger.Title),
		Actions:     newTfActions,
		Conditions:  newTfConditions,
		CategoryID:  types.Int64Value(convertedCatId),
		Position:    types.Int64Value(apiTrigger.Position),
		Description: types.StringValue(apiTrigger.Description),
		Active:      types.BoolValue(apiTrigger.Active),
		CreatedAt:   types.StringValue(apiTrigger.CreatedAt.UTC().String()),
		UpdatedAt:   types.StringValue(apiTrigger.UpdatedAt.UTC().String()),
		URL:         types.StringValue(apiTrigger.URL),
	}

	return diags
}
