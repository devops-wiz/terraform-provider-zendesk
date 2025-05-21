package models

import (
	"context"
	"strconv"

	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ ResourceTransformWithID[zendesk.TriggerCategory] = &TriggerCategoryResourceModel{}

type TriggerCategoryResourceModel struct {
	ID        types.Int64  `tfsdk:"id"`
	CreatedAt types.String `tfsdk:"created_at"`
	Name      types.String `tfsdk:"name"`
	Position  types.Int64  `tfsdk:"position"`
	UpdatedAt types.String `tfsdk:"updated_at"`
}

func (t *TriggerCategoryResourceModel) GetID() int64 {
	return t.ID.ValueInt64()
}

// GetApiModelFromTfModel implements ResourceTransform.
func (t *TriggerCategoryResourceModel) GetApiModelFromTfModel(context.Context) (newTriggerCategory zendesk.TriggerCategory, diag diag.Diagnostics) {

	if !t.Position.IsNull() && !t.Position.IsUnknown() {
		newTriggerCategory = zendesk.TriggerCategory{
			Name:     t.Name.ValueString(),
			Position: t.Position.ValueInt64(),
		}
	} else {
		newTriggerCategory = zendesk.TriggerCategory{
			Name: t.Name.ValueString(),
		}
	}

	return newTriggerCategory, diag
}

// GetTfModelFromApiModel implements ResourceTransform.
func (t *TriggerCategoryResourceModel) GetTfModelFromApiModel(ctx context.Context, apiTriggerCategory zendesk.TriggerCategory) (diag diag.Diagnostics) {
	convertedId, _ := strconv.ParseInt(apiTriggerCategory.ID, 10, 64)

	*t = TriggerCategoryResourceModel{
		ID:        types.Int64Value(convertedId),
		Name:      types.StringValue(apiTriggerCategory.Name),
		Position:  types.Int64Value(apiTriggerCategory.Position),
		CreatedAt: types.StringValue(apiTriggerCategory.CreatedAt.UTC().String()),
		UpdatedAt: types.StringValue(apiTriggerCategory.UpdatedAt.UTC().String()),
	}
	return diag
}
