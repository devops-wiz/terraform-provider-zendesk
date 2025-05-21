package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/devops-wiz/terraform-provider-zendesk/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"strconv"
)

var _ ResourceTransformWithID[zendesk.TicketForm] = &TicketFormResourceModel{}

type TicketFormResourceModel struct {
	ID                types.Int64  `tfsdk:"id"`
	Name              types.String `tfsdk:"form_name"`
	DisplayName       types.String `tfsdk:"end_user_display_name"`
	TicketFieldIds    types.List   `tfsdk:"ticket_field_ids"`
	AgentConditions   types.Map    `tfsdk:"agent_conditions"`
	EndUserConditions types.Map    `tfsdk:"end_user_conditions"`
	Active            types.Bool   `tfsdk:"active"`
	Position          types.Int64  `tfsdk:"position"`
	Default           types.Bool   `tfsdk:"default"`
	EndUserVisible    types.Bool   `tfsdk:"end_user_visible"`
	CreatedAt         types.String `tfsdk:"created_at"`
	UpdatedAt         types.String `tfsdk:"updated_at"`
	Url               types.String `tfsdk:"url"`
}

type TicketFormResourceModelV0 struct {
	ID                types.Int64  `tfsdk:"id"`
	Name              types.String `tfsdk:"form_name"`
	DisplayName       types.String `tfsdk:"end_user_display_name"`
	TicketFieldIds    types.List   `tfsdk:"ticket_field_ids"`
	AgentConditions   types.Set    `tfsdk:"agent_conditions"`
	EndUserConditions types.Set    `tfsdk:"end_user_conditions"`
	Active            types.Bool   `tfsdk:"active"`
	Position          types.Int64  `tfsdk:"position"`
	Default           types.Bool   `tfsdk:"default"`
	EndUserVisible    types.Bool   `tfsdk:"end_user_visible"`
	CreatedAt         types.String `tfsdk:"created_at"`
	UpdatedAt         types.String `tfsdk:"updated_at"`
	Url               types.String `tfsdk:"url"`
}

type FormConditionsSet struct {
	FieldValueMap types.Map `tfsdk:"field_value_map"`
}

type FormConditions struct {
	ChildConditions types.Set `tfsdk:"child_field_conditions"`
}

func (c FormConditions) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"child_field_conditions": types.SetType{ElemType: types.ObjectType{AttrTypes: FormChildFieldConditions{}.AttributeTypes()}},
	}
}

type FormConditionsV0 struct {
	ParentFieldId types.Int64  `tfsdk:"parent_field_id"`
	Value         types.String `tfsdk:"value"`
	ChildFields   types.Set    `tfsdk:"child_fields"`
}

func (c FormConditionsSet) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"field_value_map": types.MapType{ElemType: types.ObjectType{AttrTypes: FormConditions{}.AttributeTypes()}},
	}
}

type FormChildFieldConditions struct {
	Id                 types.Int64  `tfsdk:"id"`
	IsRequired         types.Bool   `tfsdk:"is_required"`
	RequiredOnStatuses types.Object `tfsdk:"required_on_statuses"`
}

func (c FormChildFieldConditions) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                   types.Int64Type,
		"is_required":          types.BoolType,
		"required_on_statuses": types.ObjectType{AttrTypes: RequiredOnStatusesResourceModel{}.AttributeTypes()},
	}
}

type RequiredOnStatusesResourceModel struct {
	Statuses types.List   `tfsdk:"statuses"`
	Type     types.String `tfsdk:"type"`
}

func (c RequiredOnStatusesResourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"statuses": types.ListType{ElemType: types.StringType},
		"type":     types.StringType,
	}
}

func (t *TicketFormResourceModel) GetID() int64 {
	return t.ID.ValueInt64()
}

func (t *TicketFormResourceModel) GetApiModelFromTfModel(ctx context.Context) (form zendesk.TicketForm, diags diag.Diagnostics) {

	var fieldIds = make([]int64, len(t.TicketFieldIds.Elements()))
	for i, id := range t.TicketFieldIds.Elements() {
		fieldIds[i] = id.(types.Int64).ValueInt64()
	}

	agentConditions, diags := getApiFormConditionsFromTf(ctx, t.AgentConditions)

	if diags.HasError() {
		return form, diags
	}

	endUserConditions, diags := getApiFormConditionsFromTf(ctx, t.EndUserConditions)

	if diags.HasError() {
		return form, diags
	}

	form = zendesk.TicketForm{
		Active:            t.Active.ValueBool(),
		AgentConditions:   agentConditions,
		Default:           t.Default.ValueBool(),
		DisplayName:       t.DisplayName.ValueString(),
		EndUserConditions: endUserConditions,
		EndUserVisible:    t.EndUserVisible.ValueBool(),
		Name:              t.Name.ValueString(),
		TicketFieldIds:    fieldIds,
	}

	if !t.Position.IsUnknown() && !t.Position.IsNull() {
		form.Position = t.Position.ValueInt64()
	}

	return form, diags
}

func getApiFormConditionsFromTf(ctx context.Context, conditionsMap types.Map) (apiConditions []zendesk.ConditionalTicketField, diags diag.Diagnostics) {

	diags = diag.Diagnostics{}

	if len(conditionsMap.Elements()) > 0 {

		tfConditionsSet := make(map[string]FormConditionsSet, len(conditionsMap.Elements()))

		diags.Append(conditionsMap.ElementsAs(ctx, &tfConditionsSet, true)...)

		if diags.HasError() {
			diags.AddError("Map Conversion Issue", "Error Converting conditionsMap types.Map to conditionSet map")
			return apiConditions, diags
		}

		for parentFieldId, conditionsSet := range tfConditionsSet {
			fieldValueMap := make(map[string]FormConditions, len(conditionsSet.FieldValueMap.Elements()))

			diags.Append(conditionsSet.FieldValueMap.ElementsAs(ctx, &fieldValueMap, true)...)

			if diags.HasError() {
				diags.AddError("Map Conversion Issue", "Error Converting conditionsSet.FieldValueMap types.Map to fieldValueMap map")
				return apiConditions, diags
			}

			for fieldTag, formCondition := range fieldValueMap {

				var childConditions []zendesk.ChildField

				var childFields []FormChildFieldConditions

				diags.Append(formCondition.ChildConditions.ElementsAs(ctx, &childFields, false)...)

				if diags.HasError() {
					diags.AddError("Map Conversion Issue", "Error Converting formCondition.ChildConditions types.Set to childFields slice")
					return apiConditions, diags
				}

				for _, field := range childFields {
					requiredOnStatusesObj := field.RequiredOnStatuses

					if !requiredOnStatusesObj.IsNull() && !requiredOnStatusesObj.IsUnknown() {
						var requiredOnStatuses RequiredOnStatusesResourceModel

						diags.Append(requiredOnStatusesObj.As(ctx, &requiredOnStatuses, basetypes.ObjectAsOptions{
							UnhandledNullAsEmpty:    true,
							UnhandledUnknownAsEmpty: true,
						})...)

						if diags.HasError() {
							diags.AddError("Object Conversion Issue", "Error Converting requiredOnStatusesObj types.Object to requiredOnStatuses struct")
							return apiConditions, diags
						}

						var statuses []string

						if len(requiredOnStatuses.Statuses.Elements()) > 0 {
							statuses = make([]string, len(requiredOnStatuses.Statuses.Elements()))

							for k, status := range requiredOnStatuses.Statuses.Elements() {
								statuses[k] = status.(types.String).ValueString()
							}
						}

						childConditions = append(childConditions, zendesk.ChildField{
							Id:         field.Id.ValueInt64(),
							IsRequired: field.IsRequired.ValueBool(),
							RequiredOnStatuses: zendesk.RequiredOnStatuses{
								Statuses: statuses,
								Type:     zendesk.RequirementType(requiredOnStatuses.Type.ValueString()),
							},
						})
					} else {
						childConditions = append(childConditions, zendesk.ChildField{
							Id:         field.Id.ValueInt64(),
							IsRequired: field.IsRequired.ValueBool(),
						})

					}

				}

				parentFieldIdStr, err := strconv.ParseInt(parentFieldId, 10, 64)

				if err != nil {
					diags.AddError("Type Conversion Error", "Error parsing map key parentFieldId for FormConditionsSet map")
					return apiConditions, diags
				}

				apiConditions = append(apiConditions, zendesk.ConditionalTicketField{
					ParentFieldId: parentFieldIdStr,
					Value:         fieldTag,
					ChildFields:   childConditions,
				})

			}

		}

	}

	return apiConditions, diags
}

func (t *TicketFormResourceModel) GetTfModelFromApiModel(ctx context.Context, apiModel zendesk.TicketForm) (diags diag.Diagnostics) {

	var fieldIds = make([]attr.Value, len(apiModel.TicketFieldIds))
	for i, id := range apiModel.TicketFieldIds {
		fieldIds[i] = types.Int64Value(id)
	}

	fieldList, diags := types.ListValue(types.Int64Type, fieldIds)

	if diags.HasError() {
		return diags
	}

	agentConditions, diags := getTfFormConditionsFromApi(ctx, apiModel.AgentConditions)

	if diags.HasError() {
		return diags
	}

	endUserConditions, diags := getTfFormConditionsFromApi(ctx, apiModel.EndUserConditions)

	if diags.HasError() {
		return diags
	}

	var displayName types.String

	if apiModel.DisplayName != "" {
		displayName = types.StringValue(apiModel.DisplayName)
	} else {
		displayName = types.StringNull()
	}

	*t = TicketFormResourceModel{
		ID:                types.Int64Value(apiModel.ID),
		Name:              types.StringValue(apiModel.Name),
		DisplayName:       displayName,
		TicketFieldIds:    fieldList,
		AgentConditions:   agentConditions,
		EndUserConditions: endUserConditions,
		Active:            types.BoolValue(apiModel.Active),
		Position:          types.Int64Value(apiModel.Position),
		Default:           types.BoolValue(apiModel.Default),
		EndUserVisible:    types.BoolValue(apiModel.EndUserVisible),
		CreatedAt:         types.StringValue(apiModel.CreatedAt.UTC().String()),
		UpdatedAt:         types.StringValue(apiModel.UpdatedAt.UTC().String()),
		Url:               types.StringValue(apiModel.Url),
	}
	return diags
}

func getTfFormConditionsFromApi(ctx context.Context, apiConditions []zendesk.ConditionalTicketField) (tfConditionSets types.Map, diags diag.Diagnostics) {

	diags = diag.Diagnostics{}

	if len(apiConditions) > 0 {

		tfConditionsSet := make(map[string]FormConditionsSet)

		groupedConditions := utils.GroupSlice(apiConditions, func(t zendesk.ConditionalTicketField) string {
			return strconv.FormatInt(t.ParentFieldId, 10)
		})

		for ticketFieldId, conditions := range groupedConditions {

			tfConditions := make(map[string]FormConditions, len(conditions))

			var tfConditionsMap types.Map

			for _, condition := range conditions {
				tfChildFields := make([]FormChildFieldConditions, len(condition.ChildFields))
				for j, field := range condition.ChildFields {
					var statusList types.List

					if len(field.RequiredOnStatuses.Statuses) > 0 {

						var diag1 diag.Diagnostics

						statusList, diag1 = types.ListValueFrom(ctx, types.StringType, field.RequiredOnStatuses.Statuses)

						diags.Append(diag1...)

						if diags.HasError() {
							diags.AddError("List Conversion Issue", "Error Converting field.RequiredOnStatuses.Statuses slice to statusList types.List")
							return tfConditionSets, diags
						}
					} else {
						statusList = types.ListNull(types.StringType)
					}

					var requiredStatusesObj types.Object

					if field.RequiredOnStatuses.Type != "" {
						requiredOnStatusesResourceModel := RequiredOnStatusesResourceModel{
							Statuses: statusList,
							Type:     types.StringValue(string(field.RequiredOnStatuses.Type)),
						}

						var diag1 diag.Diagnostics

						requiredStatusesObj, diag1 = types.ObjectValueFrom(ctx, requiredOnStatusesResourceModel.AttributeTypes(), requiredOnStatusesResourceModel)

						diags.Append(diag1...)

						if diags.HasError() {
							diags.AddError("Object Conversion Issue", "Error Converting requiredOnStatusesResourceModel struct to requiredStatusesObj types.Object")
							return tfConditionSets, diags
						}
					} else {
						requiredStatusesObj = types.ObjectNull(RequiredOnStatusesResourceModel{}.AttributeTypes())
					}

					tfChildFields[j] = FormChildFieldConditions{
						Id:                 types.Int64Value(field.Id),
						IsRequired:         types.BoolValue(field.IsRequired),
						RequiredOnStatuses: requiredStatusesObj,
					}
				}

				tfChildFieldsSet, diag1 := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: FormChildFieldConditions{}.AttributeTypes()}, tfChildFields)

				diags.Append(diag1...)

				if diags.HasError() {
					diags.AddError("Set Conversion Issue", "Error Converting tfChildFields slice to tfChildFieldsSet types.Set")
					return tfConditionSets, diags
				}

				tfConditions[condition.Value] = FormConditions{ChildConditions: tfChildFieldsSet}
			}

			if len(tfConditions) > 0 {
				var diag1 diag.Diagnostics
				tfConditionsMap, diag1 = types.MapValueFrom(ctx, types.ObjectType{AttrTypes: FormConditions{}.AttributeTypes()}, tfConditions)

				diags.Append(diag1...)

				if diags.HasError() {
					diags.AddError("Map Conversion Issue", "Error Converting tfConditions map to tfConditionsMap types.Map")
					return tfConditionSets, diags
				}
			} else {
				tfConditionsMap = types.MapNull(types.ObjectType{AttrTypes: FormConditions{}.AttributeTypes()})
			}

			tfConditionsSet[ticketFieldId] = FormConditionsSet{
				FieldValueMap: tfConditionsMap,
			}
		}

		var diag1 diag.Diagnostics

		tfConditionSets, diag1 = types.MapValueFrom(ctx, types.ObjectType{AttrTypes: FormConditionsSet{}.AttributeTypes()}, tfConditionsSet)

		diags.Append(diag1...)

		if diags.HasError() {
			diags.AddError("Map Conversion Issue", "Error Converting tfConditionsSet map to tfConditionSets types.Map")
			return tfConditionSets, diags
		}
	} else {
		tfConditionSets = types.MapNull(types.ObjectType{AttrTypes: FormConditionsSet{}.AttributeTypes()})
	}
	return tfConditionSets, diags
}
