package models

import (
	"context"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strconv"
)

func UpgradeConditionsV1(priorAll []ConditionResourceModelV0, priorAny []ConditionResourceModelV0) ConditionsResourceModel {
	var conditions ConditionsResourceModel

	conditions.All = make([]ConditionResourceModel, len(priorAll))

	for i, priorCondition := range priorAll {
		conditions.All[i] = ConditionResourceModel{
			Field:         priorCondition.Field,
			Operator:      priorCondition.Operator,
			Value:         priorCondition.Value,
			Values:        types.ListNull(types.StringType),
			CustomFieldID: priorCondition.CustomFieldID,
		}
	}

	if len(priorAny) > 0 {
		conditions.Any = make([]ConditionResourceModel, len(priorAny))
		for i, priorCondition := range priorAny {
			conditions.Any[i] = ConditionResourceModel{
				Field:         priorCondition.Field,
				Operator:      priorCondition.Operator,
				Value:         priorCondition.Value,
				Values:        types.ListNull(types.StringType),
				CustomFieldID: priorCondition.CustomFieldID,
			}
		}
	}

	return conditions
}

func UpgradeFormConditionsV1(ctx context.Context, conditionSlice []FormConditionsV0) (conditionMap map[string]FormConditionsSet, diags diag.Diagnostics) {
	diags = diag.Diagnostics{}

	if len(conditionSlice) == 0 {
		return conditionMap, diags
	}

	conditionMap = make(map[string]FormConditionsSet)

	groupedConditions := utils.GroupSlice(conditionSlice, func(f FormConditionsV0) string {
		return strconv.FormatInt(f.ParentFieldId.ValueInt64(), 10)
	})

	for ticketFieldId, conditions := range groupedConditions {

		fieldValueMap := make(map[string]FormConditions, len(conditions))

		var tfFieldValueMap types.Map

		for _, condition := range conditions {

			var childFieldConditions []FormChildFieldConditions

			diags.Append(condition.ChildFields.ElementsAs(ctx, &childFieldConditions, true)...)

			if diags.HasError() {
				return conditionMap, diags
			}

			convertedChildren, tempDiag := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: FormChildFieldConditions{}.AttributeTypes()}, childFieldConditions)

			diags = append(diags, tempDiag...)

			if diags.HasError() {
				return conditionMap, diags
			}

			fieldValueMap[condition.Value.ValueString()] = FormConditions{ChildConditions: convertedChildren}
		}

		if len(fieldValueMap) > 0 {
			var tempDiag diag.Diagnostics
			tfFieldValueMap, tempDiag = types.MapValueFrom(ctx, types.ObjectType{AttrTypes: FormConditions{}.AttributeTypes()}, fieldValueMap)

			diags = append(diags, tempDiag...)

			if diags.HasError() {
				return conditionMap, diags
			}
		} else {
			tfFieldValueMap = types.MapNull(types.ObjectType{AttrTypes: FormConditions{}.AttributeTypes()})
		}

		conditionMap[ticketFieldId] = FormConditionsSet{
			FieldValueMap: tfFieldValueMap,
		}

	}

	return conditionMap, diags
}
