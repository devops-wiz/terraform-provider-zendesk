package models

import (
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strconv"
	"strings"
)

type ConditionsResourceModel struct {
	All []ConditionResourceModel `tfsdk:"all"`
	Any []ConditionResourceModel `tfsdk:"any"`
}

type ConditionsResourceModelV0 struct {
	All []ConditionResourceModelV0 `tfsdk:"all"`
	Any []ConditionResourceModelV0 `tfsdk:"any"`
}
type ConditionResourceModel struct {
	Field         types.String `tfsdk:"field"`
	Operator      types.String `tfsdk:"operator"`
	Value         types.String `tfsdk:"value"`
	Values        types.List   `tfsdk:"values"`
	CustomFieldID types.Int64  `tfsdk:"custom_field_id"`
}

type ConditionResourceModelV0 struct {
	Field         types.String `tfsdk:"field"`
	Operator      types.String `tfsdk:"operator"`
	Value         types.String `tfsdk:"value"`
	CustomFieldID types.Int64  `tfsdk:"custom_field_id"`
}

func getApiConditionsFromTf(ctx context.Context, conditionModels ConditionsResourceModel) (conditions zendesk.Conditions, diags diag.Diagnostics) {

	var newAllConditions, newAnyConditions []zendesk.Condition

	newAllConditions, diags = mapApiConditionsFromTf(ctx, conditionModels.All)

	if diags.HasError() {
		return conditions, diags
	}

	newAnyConditions, diags = mapApiConditionsFromTf(ctx, conditionModels.Any)

	if diags.HasError() {
		return conditions, diags
	}

	conditions = zendesk.Conditions{
		All: newAllConditions,
		Any: newAnyConditions,
	}

	return conditions, diags
}

func mapApiConditionsFromTf(ctx context.Context, conditions []ConditionResourceModel) (newConditions []zendesk.Condition, diags diag.Diagnostics) {
	newConditions = make([]zendesk.Condition, len(conditions))

	var value zendesk.ParsedValue

	for index, condition := range conditions {
		fieldName := condition.Field.ValueString()

		if condition.Field.ValueString() == "custom_field" {
			fieldName = fmt.Sprintf("%s%d", zendesk.ConditionFieldCustomField, condition.CustomFieldID.ValueInt64())
		}

		if condition.Field.ValueString() == "ticket_field" {
			fieldName = fmt.Sprintf("%s%d", zendesk.ConditionFieldCustomFieldAlt, condition.CustomFieldID.ValueInt64())
		}

		if condition.Field.ValueString() == "" {
			diags = append(diags, diag.NewAttributeErrorDiagnostic(path.Root("conditions"), "invalid field value", "field should not be empty"))
			return newConditions, diags
		}

		if !condition.Values.IsUnknown() && !condition.Values.IsNull() {
			channelsRaw := make([]types.String, len(condition.Values.Elements()))
			diags = append(diags, condition.Values.ElementsAs(ctx, &channelsRaw, false)...)
			if diags.HasError() {
				return newConditions, diags
			}
			channels := make([]string, len(channelsRaw))
			for i, channel := range channelsRaw {
				channels[i] = channel.ValueString()
			}
			value = zendesk.ParsedValue{ListData: channels}
		} else {
			value = zendesk.ParsedValue{Data: condition.Value.ValueString()}
		}

		newConditions[index] = zendesk.Condition{
			Field:    fieldName,
			Operator: condition.Operator.ValueString(),
			Value:    value,
		}
	}
	return newConditions, diags
}

func getTfConditionsFromApi(ctx context.Context, conditions zendesk.Conditions) (ConditionsResourceModel, diag.Diagnostics) {
	diags, newTfAllConditions := mapTFConditionFromAPI(ctx, conditions.All)

	if diags.HasError() {
		return ConditionsResourceModel{}, diags
	}

	if len(newTfAllConditions) <= 0 {
		newTfAllConditions = nil
	}

	diags, newTfAnyConditions := mapTFConditionFromAPI(ctx, conditions.Any)

	if diags.HasError() {
		return ConditionsResourceModel{}, diags
	}

	if len(newTfAnyConditions) <= 0 {
		newTfAnyConditions = nil

	}

	return ConditionsResourceModel{
		All: newTfAllConditions,
		Any: newTfAnyConditions,
	}, nil
}

func mapTFConditionFromAPI(ctx context.Context, conditions []zendesk.Condition) (diags diag.Diagnostics, newTfConditions []ConditionResourceModel) {
	newTfConditions = make([]ConditionResourceModel, len(conditions))

	for index, condition := range conditions {
		fieldName := condition.Field
		cid := types.Int64Unknown()
		if strings.HasPrefix(condition.Field, string(zendesk.ConditionFieldCustomField)) || strings.HasPrefix(condition.Field, string(zendesk.ConditionFieldCustomFieldAlt)) {
			var id string
			switch {
			case strings.HasPrefix(condition.Field, string(zendesk.ConditionFieldCustomField)):
				fieldName = "custom_field"
				id, _ = strings.CutPrefix(condition.Field, string(zendesk.ConditionFieldCustomField))
			case strings.HasPrefix(condition.Field, string(zendesk.ConditionFieldCustomFieldAlt)):
				fieldName = "ticket_field"
				id, _ = strings.CutPrefix(condition.Field, string(zendesk.ConditionFieldCustomFieldAlt))
			}
			convertedId, _ := strconv.ParseInt(id, 10, 64)
			cid = types.Int64Value(convertedId)
		}

		newTfConditions[index] = ConditionResourceModel{
			Field:    types.StringValue(fieldName),
			Operator: types.StringValue(condition.Operator),
		}

		if condition.Value.ListData == nil {
			newTfConditions[index].Value = types.StringValue(condition.Value.Data)
			newTfConditions[index].Values = types.ListNull(types.StringType)

		} else {
			values, tempDiags := types.ListValueFrom(ctx, types.StringType, condition.Value.ListData)
			if tempDiags.HasError() {
				diags = append(diags, tempDiags...)
				return diags, newTfConditions
			}
			newTfConditions[index].Values = values
			newTfConditions[index].Value = types.StringNull()
		}

		if !cid.IsUnknown() && !cid.IsNull() {
			newTfConditions[index].CustomFieldID = cid
		}

	}
	return diags, newTfConditions
}
