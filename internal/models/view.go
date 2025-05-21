package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ ResourceTransformWithID[zendesk.View] = &ViewResourceModel{}

type ViewResourceModel struct {
	ID          types.Int64              `tfsdk:"id"`
	URL         types.String             `tfsdk:"url"`
	Title       types.String             `tfsdk:"title"`
	Description types.String             `tfsdk:"description"`
	Active      types.Bool               `tfsdk:"active"`
	UpdatedAt   types.String             `tfsdk:"updated_at"`
	CreatedAt   types.String             `tfsdk:"created_at"`
	Position    types.Int64              `tfsdk:"position"`
	Conditions  *ConditionsResourceModel `tfsdk:"conditions"`
	Output      types.Object             `tfsdk:"output"`
	Restriction types.Object             `tfsdk:"restriction"`
}

type ViewResourceModelV0 struct {
	ID          types.Int64                `tfsdk:"id"`
	URL         types.String               `tfsdk:"url"`
	Title       types.String               `tfsdk:"title"`
	Description types.String               `tfsdk:"description"`
	Active      types.Bool                 `tfsdk:"active"`
	UpdatedAt   types.String               `tfsdk:"updated_at"`
	CreatedAt   types.String               `tfsdk:"created_at"`
	Position    types.Int64                `tfsdk:"position"`
	Conditions  *ConditionsResourceModelV0 `tfsdk:"conditions"`
	Output      types.Object               `tfsdk:"output"`
	Restriction types.Object               `tfsdk:"restriction"`
}

func (v *ViewResourceModel) GetID() int64 {
	return v.ID.ValueInt64()
}

// GetApiModelFromTfModel implements ResourceTransform.
func (v *ViewResourceModel) GetApiModelFromTfModel(ctx context.Context) (newUpdatedView zendesk.View, diags diag.Diagnostics) {
	newConditions, diags := getApiConditionsFromTf(ctx, *v.Conditions)

	if diags.HasError() {
		return zendesk.View{}, diags
	}

	var newRestriction zendesk.Restriction

	newOutput, diags := getApiOutputFromTf(ctx, v.Output)

	if diags.HasError() {
		return zendesk.View{}, diags
	}

	newUpdatedView = zendesk.View{
		Title:       v.Title.ValueString(),
		Description: v.Description.ValueString(),
		Active:      v.Active.ValueBool(),
		Conditions:  newConditions,
		All:         newConditions.All,
		Any:         newConditions.Any,
		Output:      newOutput,
	}

	if !v.Position.IsNull() && !v.Position.IsUnknown() {
		newUpdatedView.Position = v.Position.ValueInt64()
	}

	if !v.Restriction.IsNull() && !v.Restriction.IsUnknown() {
		newRestriction, diags = getApiRestrictionFromTf(ctx, v.Restriction)
		diags.Append(diags...)

		if diags.HasError() {
			return zendesk.View{}, diags
		}

		newUpdatedView.Restriction = newRestriction
	}

	return newUpdatedView, diags

}

// GetTfModelFromApiModel implements ResourceTransform.
func (v *ViewResourceModel) GetTfModelFromApiModel(ctx context.Context, apiView zendesk.View) (diags diag.Diagnostics) {

	newConditions, diags := getTfConditionsFromApi(ctx, apiView.Conditions)

	if diags.HasError() {
		return diags
	}

	var newTfViewOutput types.Object

	if apiView.Execution != nil {
		viewExec := getViewExecFromInterfaceVals(apiView.Execution)

		diags = getTfOutputFromApi(ctx, viewExec, &newTfViewOutput)
		if diags.HasError() {
			return diags
		}

	} else {
		newTfViewOutput = types.ObjectUnknown(newTfViewOutput.AttributeTypes(ctx))
	}

	*v = ViewResourceModel{
		ID:          types.Int64Value(apiView.ID),
		Title:       types.StringValue(apiView.Title),
		Active:      types.BoolValue(apiView.Active),
		Description: types.StringValue(apiView.Description),
		Conditions:  &newConditions,
		Position:    types.Int64Value(apiView.Position),
		Output:      newTfViewOutput,
		URL:         types.StringValue(apiView.URL),
		UpdatedAt:   types.StringValue(apiView.UpdatedAt),
		CreatedAt:   types.StringValue(apiView.CreatedAt),
	}

	var newTfRestriction types.Object
	if apiView.Restriction != nil {

		newTfRestriction, diags = getTfRestrictionFromApi(ctx, apiView.Restriction)

		if diags.HasError() {
			return diags
		}

	} else {
		newTfRestriction = types.ObjectNull(RestrictionResourceModel{}.AttributeTypes())
	}
	v.Restriction = newTfRestriction

	return diags
}

func getViewExecFromInterfaceVals(execution interface{}) zendesk.ViewExecution {

	viewExecRaw := execution.(map[string]interface{})
	viewExecColumnsRaw := viewExecRaw["columns"].([]interface{})

	viewExecColumns := make([]zendesk.ViewExecColumn, len(viewExecColumnsRaw))

	for i, vec := range viewExecColumnsRaw {
		viewExecColumns[i] = zendesk.ViewExecColumn{
			ID: vec.(map[string]interface{})["id"].(string),
		}
	}

	viewExecGroupBy := viewExecRaw["group_by"].(string)
	viewExecGroupOrder := viewExecRaw["group_order"].(string)
	viewExecSortBy := viewExecRaw["sort_by"].(string)
	viewExecSortOrder := viewExecRaw["sort_order"].(string)

	viewExec := zendesk.ViewExecution{
		Columns:    viewExecColumns,
		GroupBy:    viewExecGroupBy,
		GroupOrder: viewExecGroupOrder,
		SortBy:     viewExecSortBy,
		SortOrder:  viewExecSortOrder,
	}

	return viewExec
}

type ViewOutputResourceModel struct {
	Columns    types.List   `tfsdk:"columns"`
	GroupBy    types.String `tfsdk:"group_by"`
	GroupOrder types.String `tfsdk:"group_order"`
	SortBy     types.String `tfsdk:"sort_by"`
	SortOrder  types.String `tfsdk:"sort_order"`
}

func (v *ViewOutputResourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"columns":     types.ListType{ElemType: types.StringType},
		"group_by":    types.StringType,
		"group_order": types.StringType,
		"sort_by":     types.StringType,
		"sort_order":  types.StringType,
	}
}
