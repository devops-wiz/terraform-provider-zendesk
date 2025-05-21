package models

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"strconv"
)

type RestrictionResourceModel struct {
	Type types.String `tfsdk:"type"`
	IDS  types.Set    `tfsdk:"ids"`
}

func (r RestrictionResourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"type": types.StringType,
		"ids":  types.SetType{ElemType: types.Int64Type},
	}
}

func getApiRestrictionFromTf(ctx context.Context, restrictionObject types.Object) (zendesk.Restriction, diag.Diagnostics) {
	var restriction RestrictionResourceModel

	diags := restrictionObject.As(ctx, &restriction, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})

	newRestIds := make([]int64, len(restriction.IDS.Elements()))

	for index, id := range restriction.IDS.Elements() {
		newRestIds[index] = id.(types.Int64).ValueInt64()
	}

	newRestriction := zendesk.Restriction{
		Type: restriction.Type.ValueString(),
		IDS:  newRestIds,
	}

	return newRestriction, diags
}

func getTfRestrictionFromApi(ctx context.Context, apiRestriction interface{}) (restrictionObject types.Object, diag diag.Diagnostics) {

	restrictions := apiRestriction.(map[string]interface{})["ids"].([]interface{})
	macRestIds := make([]types.Int64, len(restrictions))
	for index, id := range restrictions {
		switch v := id.(type) {
		case float64:
			macRestIds[index] = types.Int64Value(int64(v))
		case int64:
			macRestIds[index] = types.Int64Value(v)
		case string:
			convId, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				diag.AddAttributeError(
					path.Root("restriction"),
					"Error converting restriction id from string to int64",
					fmt.Sprintf("Error: %s", err.Error()),
				)
				break
			}
			macRestIds[index] = types.Int64Value(convId)

		}
	}

	restIdList, diagNew := types.SetValueFrom(ctx, types.Int64Type, macRestIds)
	diag.Append(diagNew...)

	if diag.HasError() {
		return restrictionObject, diag
	}
	restriction := RestrictionResourceModel{
		Type: types.StringValue(apiRestriction.(map[string]interface{})["type"].(string)),
		IDS:  restIdList,
	}

	restrictionObject, diag = types.ObjectValueFrom(ctx, restriction.AttributeTypes(), restriction)

	return restrictionObject, diag
}
