package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DynamicContentVariantModel struct {
	ID       types.Int64  `tfsdk:"id"`
	Content  types.String `tfsdk:"content"`
	LocaleID types.Int64  `tfsdk:"locale_id"`
	Active   types.Bool   `tfsdk:"active"`
	Default  types.Bool   `tfsdk:"default"`
}

type DynamicContentItemResourceModel struct {
	ID              types.Int64                  `tfsdk:"id"`
	Name            types.String                 `tfsdk:"name"`
	Placeholder     types.String                 `tfsdk:"placeholder"`
	DefaultLocaleID types.Int64                  `tfsdk:"default_locale_id"`
	Variants        []DynamicContentVariantModel `tfsdk:"variants"`
}

func (d *DynamicContentItemResourceModel) GetID() int64 {
	return d.ID.ValueInt64()
}

func (d *DynamicContentItemResourceModel) GetApiModelFromTfModel(_ context.Context) (dci zendesk.DynamicContentItem, diags diag.Diagnostics) {

	dciVariants := make([]zendesk.DynamicContentVariant, len(d.Variants))

	for i, variant := range d.Variants {
		dciVariants[i] = zendesk.DynamicContentVariant{
			ID:       variant.ID.ValueInt64(),
			Content:  variant.Content.ValueString(),
			LocaleID: variant.LocaleID.ValueInt64(),
			Active:   variant.Active.ValueBool(),
			Default:  variant.Default.ValueBool(),
		}
	}

	dci = zendesk.DynamicContentItem{
		ID:              d.ID.ValueInt64(),
		Name:            d.Name.ValueString(),
		Variants:        dciVariants,
		DefaultLocaleID: d.DefaultLocaleID.ValueInt64(),
	}

	return dci, diags
}

func (d *DynamicContentItemResourceModel) GetTfModelFromApiModel(_ context.Context, dci zendesk.DynamicContentItem) (diags diag.Diagnostics) {

	dciTfVariants := make([]DynamicContentVariantModel, len(dci.Variants))

	for i, variant := range dci.Variants {
		dciTfVariants[i] = DynamicContentVariantModel{
			ID:       types.Int64Value(variant.ID),
			Content:  types.StringValue(variant.Content),
			LocaleID: types.Int64Value(variant.LocaleID),
			Active:   types.BoolValue(variant.Active),
			Default:  types.BoolValue(variant.Default),
		}
	}

	*d = DynamicContentItemResourceModel{
		ID:              types.Int64Value(dci.ID),
		Name:            types.StringValue(dci.Name),
		Placeholder:     types.StringValue(dci.Placeholder),
		DefaultLocaleID: types.Int64Value(dci.DefaultLocaleID),
		Variants:        dciTfVariants,
	}

	return diags
}
