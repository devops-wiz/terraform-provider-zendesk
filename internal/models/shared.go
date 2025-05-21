package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func getApiCustomFieldOptionsFromTf(ctx context.Context, fieldOptions types.List) (customFieldOptions []zendesk.CustomFieldOption, diags diag.Diagnostics) {

	if !fieldOptions.IsNull() && !fieldOptions.IsUnknown() {
		tfCustomFieldOptions := make([]CustomFieldOptionResourceModel, len(fieldOptions.Elements()))
		customFieldOptions = make([]zendesk.CustomFieldOption, len(fieldOptions.Elements()))
		diags = fieldOptions.ElementsAs(ctx, &tfCustomFieldOptions, true)

		if diags.HasError() {
			return customFieldOptions, diags
		}

		for i, option := range tfCustomFieldOptions {
			optionId := option.ID.ValueInt64Pointer()
			if *optionId == 0 {
				optionId = nil
			}
			customFieldOptions[i] = zendesk.CustomFieldOption{
				ID:    optionId,
				Name:  option.Name.ValueString(),
				Value: option.Value.ValueString(),
			}
		}
	} else {
		customFieldOptions = []zendesk.CustomFieldOption(nil)
	}

	return customFieldOptions, diags
}

func getApiRelationshipFilterFromTf(ctx context.Context, filterObject types.Object) (relationshipFilter zendesk.RelationshipFilter, diags diag.Diagnostics) {
	var tfRelationshipFilter RelationshipFilterResourceModel

	var tfRelationshipFilterAllObjs []RelationshipFilterObjectResourceModel
	var tfRelationshipFilterAnyObjs []RelationshipFilterObjectResourceModel

	var allRelationshipFilterObjects []zendesk.RelationshipFilterObject
	var anyRelationshipFilterObjects []zendesk.RelationshipFilterObject

	if !filterObject.IsNull() && !filterObject.IsUnknown() {
		diags = filterObject.As(ctx, &tfRelationshipFilter, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})

		if diags.HasError() {
			return relationshipFilter, diags
		}

		if !tfRelationshipFilter.All.IsNull() && !tfRelationshipFilter.All.IsUnknown() {
			diags = tfRelationshipFilter.All.ElementsAs(ctx, &tfRelationshipFilterAllObjs, true)

			if diags.HasError() {
				return relationshipFilter, diags
			}

			allRelationshipFilterObjects = getApiRelationshipObjectsFromTf(tfRelationshipFilterAllObjs)

		}

		if !tfRelationshipFilter.Any.IsNull() && !tfRelationshipFilter.Any.IsUnknown() {
			diags = tfRelationshipFilter.Any.ElementsAs(ctx, &tfRelationshipFilterAnyObjs, true)

			if diags.HasError() {
				return relationshipFilter, diags
			}

			anyRelationshipFilterObjects = getApiRelationshipObjectsFromTf(tfRelationshipFilterAnyObjs)

		}

	}

	relationshipFilter = zendesk.RelationshipFilter{
		All: allRelationshipFilterObjects,
		Any: anyRelationshipFilterObjects,
	}

	return relationshipFilter, diags
}

func getApiRelationshipObjectsFromTf(tfObjs []RelationshipFilterObjectResourceModel) (filterObjs []zendesk.RelationshipFilterObject) {
	filterObjs = make([]zendesk.RelationshipFilterObject, len(tfObjs))
	for i, obj := range tfObjs {
		filterObjs[i] = zendesk.RelationshipFilterObject{
			Field:    obj.Field.ValueString(),
			Operator: obj.Operator.ValueString(),
			Value:    obj.Value.ValueString(),
		}
	}

	return filterObjs
}

func getTfCustomFieldOptionsFromApi(ctx context.Context, options []zendesk.CustomFieldOption) (tfCustomOptionsList types.List, diags diag.Diagnostics) {

	if len(options) > 0 {
		tfCustomOptions := make([]CustomFieldOptionResourceModel, len(options))

		for i, option := range options {
			tfCustomOptions[i] = CustomFieldOptionResourceModel{
				ID: types.Int64Value(*option.ID),
				CustomFieldOptionResourceBase: CustomFieldOptionResourceBase{
					Name:  types.StringValue(option.Name),
					Value: types.StringValue(option.Value),
				},
			}
		}

		tfCustomOptionsList, diags = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: CustomFieldOptionResourceModel{}.AttributeTypes()}, tfCustomOptions)

		if diags.HasError() {
			return tfCustomOptionsList, diags
		}
	} else {
		tfCustomOptionsList = types.ListNull(types.ObjectType{AttrTypes: CustomFieldOptionResourceModel{}.AttributeTypes()})
	}
	return tfCustomOptionsList, diags
}

type CustomFieldOptionResourceBase struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

func (c CustomFieldOptionResourceBase) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":  types.StringType,
		"value": types.StringType,
	}
}

type CustomFieldOptionResourceModel struct {
	ID types.Int64 `tfsdk:"id"`
	CustomFieldOptionResourceBase
}

func (c CustomFieldOptionResourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":    types.Int64Type,
		"name":  types.StringType,
		"value": types.StringType,
	}
}
