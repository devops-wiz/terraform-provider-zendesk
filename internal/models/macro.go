package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ ResourceTransformWithID[zendesk.Macro] = &MacroResourceModel{}

// MacroResourceModel describes the resource m model.
type MacroResourceModel struct {
	ID          types.Int64           `tfsdk:"id"`
	Actions     []ActionResourceModel `tfsdk:"actions"`
	Title       types.String          `tfsdk:"title"`
	Restriction types.Object          `tfsdk:"restriction"`
	Active      types.Bool            `tfsdk:"active"`
	Description types.String          `tfsdk:"description"`
	CreatedAt   types.String          `tfsdk:"created_at"`
	UpdatedAt   types.String          `tfsdk:"updated_at"`
	URL         types.String          `tfsdk:"url"`
	Position    types.Int64           `tfsdk:"position"`
}

func (m *MacroResourceModel) GetID() int64 {
	return m.ID.ValueInt64()
}

// GetApiModelFromTfModel implements ResourceTransform.
func (m *MacroResourceModel) GetApiModelFromTfModel(ctx context.Context) (zendesk.Macro, diag.Diagnostics) {
	newMacroActions, diags := getApiActionsFromTf(m.Actions)

	var newMacro zendesk.Macro

	if diags.HasError() {
		return newMacro, diags
	}

	var newRestriction zendesk.Restriction

	hasRestriction := !m.Restriction.IsNull() && !m.Restriction.IsUnknown()
	hasPosition := !m.Position.IsNull() && !m.Position.IsUnknown()

	newMacro = zendesk.Macro{
		Title:       m.Title.ValueString(),
		Description: m.Description.ValueString(),
		Active:      m.Active.ValueBool(),
		Actions:     newMacroActions,
	}

	if hasRestriction {
		newRestriction, diags = getApiRestrictionFromTf(ctx, m.Restriction)
		newMacro.Restriction = newRestriction
	}
	if hasPosition {
		newMacro.Position = int(m.Position.ValueInt64())
	}

	return newMacro, diags
}

// GetTfModelFromApiModel implements ResourceTransform.
func (m *MacroResourceModel) GetTfModelFromApiModel(ctx context.Context, apiMacro zendesk.Macro) diag.Diagnostics {
	newTfMacroActions, diags := getTfActionsFromApi(apiMacro.Actions)

	if diags.HasError() {
		return diags
	}

	*m = MacroResourceModel{
		ID:          types.Int64Value(apiMacro.ID),
		Title:       types.StringValue(apiMacro.Title),
		Actions:     newTfMacroActions,
		Description: types.StringValue(apiMacro.Description),
		Active:      types.BoolValue(apiMacro.Active),
		CreatedAt:   types.StringValue(apiMacro.CreatedAt.UTC().String()),
		UpdatedAt:   types.StringValue(apiMacro.UpdatedAt.UTC().String()),
		URL:         types.StringValue(apiMacro.URL),
	}

	var objectValue = types.ObjectNull(RestrictionResourceModel{}.AttributeTypes())

	if apiMacro.Restriction != nil {
		objectValue, diags = getTfRestrictionFromApi(ctx, apiMacro.Restriction)
		if diags.HasError() {
			return diags
		}
	}

	m.Restriction = objectValue

	return diags
}
