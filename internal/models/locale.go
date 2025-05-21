package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LocaleDatasourceModel struct {
	Code   types.String `tfsdk:"code"`
	Locale *LocaleModel `tfsdk:"locale"`
}

type LocaleModel struct {
	ID         types.Int64  `tfsdk:"id"`
	LocaleCode types.String `tfsdk:"locale_code"`
	Name       types.String `tfsdk:"name"`
}

func (l *LocaleModel) GetTfModelFromApiModel(_ context.Context, locale zendesk.Locale) (diags diag.Diagnostics) {
	*l = LocaleModel{
		ID:         types.Int64Value(locale.ID),
		LocaleCode: types.StringValue(locale.Locale),
		Name:       types.StringValue(locale.Name),
	}
	return diags
}
