package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"reflect"
)

type UserOrgFieldResourceModel struct {
	ID                     types.Int64  `tfsdk:"id"`
	URL                    types.String `tfsdk:"url"`
	Key                    types.String `tfsdk:"key"`
	Type                   types.String `tfsdk:"type"`
	Title                  types.String `tfsdk:"name"`
	Description            types.String `tfsdk:"description"`
	Position               types.Int64  `tfsdk:"position"`
	Active                 types.Bool   `tfsdk:"active"`
	System                 types.Bool   `tfsdk:"system"`
	RegexpForValidation    types.String `tfsdk:"regexp_for_validation"`
	Tag                    types.String `tfsdk:"tag"`
	CustomFieldOptions     types.List   `tfsdk:"custom_field_options"`
	CreatedAt              types.String `tfsdk:"created_at"`
	UpdatedAt              types.String `tfsdk:"updated_at"`
	RelationshipTargetType types.String `tfsdk:"relationship_target_type"`
	RelationshipFilter     types.Object `tfsdk:"relationship_filter"`
}

var _ ResourceTransformWithID[zendesk.UserField] = &UserFieldResourceModel{}

type UserFieldResourceModel struct {
	UserOrgFieldResourceModel
}

func (u *UserOrgFieldResourceModel) GetID() int64 {
	return u.ID.ValueInt64()
}

func (u *UserFieldResourceModel) GetApiModelFromTfModel(ctx context.Context) (userField zendesk.UserField, diags diag.Diagnostics) {

	fieldOptions := u.CustomFieldOptions
	var customFieldOptions []zendesk.CustomFieldOption

	customFieldOptions, diags = getApiCustomFieldOptionsFromTf(ctx, fieldOptions)

	if diags.HasError() {
		return zendesk.UserField{}, diags
	}

	filterObject := u.RelationshipFilter

	relationshipFilter, diags := getApiRelationshipFilterFromTf(ctx, filterObject)

	if diags.HasError() {
		return zendesk.UserField{}, diags
	}

	userField = zendesk.UserField{
		Key:                    u.Key.ValueString(),
		Type:                   u.Type.ValueString(),
		Title:                  u.Title.ValueString(),
		Description:            u.Description.ValueString(),
		Active:                 u.Active.ValueBool(),
		RegexpForValidation:    u.RegexpForValidation.ValueString(),
		Tag:                    u.Tag.ValueString(),
		CustomFieldOptions:     customFieldOptions,
		RelationshipTargetType: u.RelationshipTargetType.ValueString(),
		RelationshipFilter:     relationshipFilter,
	}

	if !u.Position.IsUnknown() && !u.Position.IsNull() {
		userField.Position = u.Position.ValueInt64()
	}

	return userField, diags

}

func getTfRelationshipFiltersObjectsFromApi(ctx context.Context, relationshipFilterObjects []zendesk.RelationshipFilterObject) (filterList types.List, diags diag.Diagnostics) {
	var filterObjectResourceModels []RelationshipFilterObjectResourceModel

	if len(relationshipFilterObjects) > 0 {
		for i, object := range relationshipFilterObjects {
			filterObjectResourceModels[i] = RelationshipFilterObjectResourceModel{
				Field:    types.StringValue(object.Field),
				Operator: types.StringValue(object.Operator),
				Value:    types.StringValue(object.Value),
			}
		}

		filterList, diags = types.ListValueFrom(
			ctx,
			types.ObjectType{AttrTypes: filterObjectResourceModels[0].AttributeTypes()},
			filterObjectResourceModels,
		)

		if diags.HasError() {
			return filterList, diags
		}

	} else {
		filterList = types.ListNull(types.ObjectType{AttrTypes: RelationshipFilterObjectResourceModel{}.AttributeTypes()})
	}
	return filterList, diags
}

func (u *UserFieldResourceModel) GetTfModelFromApiModel(ctx context.Context, userField zendesk.UserField) (diags diag.Diagnostics) {

	var tfCustomOptionsList types.List

	tfCustomOptionsList, diags = getTfCustomFieldOptionsFromApi(ctx, userField.CustomFieldOptions)

	if diags.HasError() {
		return diags
	}

	var tfRelationshipFilterObject types.Object

	if !reflect.DeepEqual(userField.RelationshipFilter, zendesk.RelationshipFilter{}) {

		allObjects := userField.RelationshipFilter.All

		var tfAllRelationshipFilterList types.List

		tfAllRelationshipFilterList, diags = getTfRelationshipFiltersObjectsFromApi(ctx, allObjects)

		if diags.HasError() {
			return diags
		}

		anyObjects := userField.RelationshipFilter.Any

		var tfAnyRelationshipFilterList types.List

		tfAnyRelationshipFilterList, diags = getTfRelationshipFiltersObjectsFromApi(ctx, anyObjects)

		if diags.HasError() {
			return diags
		}

		tfRelationshipFilter := RelationshipFilterResourceModel{
			All: tfAllRelationshipFilterList,
			Any: tfAnyRelationshipFilterList,
		}

		tfRelationshipFilterObject, diags = types.ObjectValueFrom(ctx, tfRelationshipFilter.AttributeTypes(), tfRelationshipFilter)
		if diags.HasError() {
			return diags
		}

	} else {
		tfRelationshipFilterObject = types.ObjectNull(RelationshipFilterResourceModel{}.AttributeTypes())
	}

	var tfDescription types.String

	if userField.Description != "" {
		tfDescription = types.StringValue(userField.Description)
	} else {
		tfDescription = types.StringNull()
	}

	var tfRegexpForValidation types.String

	if userField.RegexpForValidation != "" {
		tfRegexpForValidation = types.StringValue(userField.RegexpForValidation)
	} else {
		tfRegexpForValidation = types.StringNull()
	}

	var tfTag types.String

	if userField.Tag != "" {
		tfTag = types.StringValue(userField.Tag)
	} else {
		tfTag = types.StringNull()
	}

	var tfRelationshipTargetType types.String

	if userField.RelationshipTargetType != "" {
		tfRelationshipTargetType = types.StringValue(userField.RelationshipTargetType)
	} else {
		tfRelationshipTargetType = types.StringNull()
	}

	*u = UserFieldResourceModel{
		UserOrgFieldResourceModel{
			ID:                     types.Int64Value(userField.ID),
			URL:                    types.StringValue(userField.URL),
			Key:                    types.StringValue(userField.Key),
			Type:                   types.StringValue(userField.Type),
			Title:                  types.StringValue(userField.Title),
			Description:            tfDescription,
			Position:               types.Int64Value(userField.Position),
			Active:                 types.BoolValue(userField.Active),
			System:                 types.BoolValue(userField.System),
			RegexpForValidation:    tfRegexpForValidation,
			Tag:                    tfTag,
			CustomFieldOptions:     tfCustomOptionsList,
			CreatedAt:              types.StringValue(userField.CreatedAt.UTC().String()),
			UpdatedAt:              types.StringValue(userField.UpdatedAt.UTC().String()),
			RelationshipTargetType: tfRelationshipTargetType,
			RelationshipFilter:     tfRelationshipFilterObject,
		},
	}

	return diags
}

var _ ResourceTransform[zendesk.OrganizationField] = &OrganizationFieldResourceModel{}

type OrganizationFieldResourceModel struct {
	UserOrgFieldResourceModel
}

func (o *OrganizationFieldResourceModel) GetApiModelFromTfModel(ctx context.Context) (organizationField zendesk.OrganizationField, diags diag.Diagnostics) {
	fieldOptions := o.CustomFieldOptions
	var customFieldOptions []zendesk.CustomFieldOption

	customFieldOptions, diags = getApiCustomFieldOptionsFromTf(ctx, fieldOptions)

	if diags.HasError() {
		return organizationField, diags
	}

	filterObject := o.RelationshipFilter

	relationshipFilter, diags := getApiRelationshipFilterFromTf(ctx, filterObject)

	if diags.HasError() {
		return organizationField, diags
	}

	organizationField = zendesk.OrganizationField{
		Key:                    o.Key.ValueString(),
		Type:                   o.Type.ValueString(),
		Title:                  o.Title.ValueString(),
		Description:            o.Description.ValueString(),
		Active:                 o.Active.ValueBool(),
		RegexpForValidation:    o.RegexpForValidation.ValueString(),
		Tag:                    o.Tag.ValueString(),
		CustomFieldOptions:     customFieldOptions,
		RelationshipTargetType: o.RelationshipTargetType.ValueString(),
		RelationshipFilter:     relationshipFilter,
	}

	if !o.Position.IsUnknown() && !o.Position.IsNull() {
		organizationField.Position = o.Position.ValueInt64()
	}

	return organizationField, diags

}

func (o *OrganizationFieldResourceModel) GetTfModelFromApiModel(ctx context.Context, organizationField zendesk.OrganizationField) (diags diag.Diagnostics) {
	var tfCustomOptionsList types.List

	tfCustomOptionsList, diags = getTfCustomFieldOptionsFromApi(ctx, organizationField.CustomFieldOptions)

	if diags.HasError() {
		return diags
	}

	var tfRelationshipFilterObject types.Object

	if !reflect.DeepEqual(organizationField.RelationshipFilter, zendesk.RelationshipFilter{}) {

		allObjects := organizationField.RelationshipFilter.All

		var tfAllRelationshipFilterList types.List

		tfAllRelationshipFilterList, diags = getTfRelationshipFiltersObjectsFromApi(ctx, allObjects)

		if diags.HasError() {
			return diags
		}

		anyObjects := organizationField.RelationshipFilter.Any

		var tfAnyRelationshipFilterList types.List

		tfAnyRelationshipFilterList, diags = getTfRelationshipFiltersObjectsFromApi(ctx, anyObjects)

		if diags.HasError() {
			return diags
		}

		tfRelationshipFilter := RelationshipFilterResourceModel{
			All: tfAllRelationshipFilterList,
			Any: tfAnyRelationshipFilterList,
		}

		tfRelationshipFilterObject, diags = types.ObjectValueFrom(ctx, tfRelationshipFilter.AttributeTypes(), tfRelationshipFilter)
		if diags.HasError() {
			return diags
		}

	} else {
		tfRelationshipFilterObject = types.ObjectNull(RelationshipFilterResourceModel{}.AttributeTypes())
	}

	var tfDescription types.String

	if organizationField.Description != "" {
		tfDescription = types.StringValue(organizationField.Description)
	} else {
		tfDescription = types.StringNull()
	}

	var tfRegexpForValidation types.String

	if organizationField.RegexpForValidation != "" {
		tfRegexpForValidation = types.StringValue(organizationField.RegexpForValidation)
	} else {
		tfRegexpForValidation = types.StringNull()
	}

	var tfTag types.String

	if organizationField.Tag != "" {
		tfTag = types.StringValue(organizationField.Tag)
	} else {
		tfTag = types.StringNull()
	}

	var tfRelationshipTargetType types.String

	if organizationField.RelationshipTargetType != "" {
		tfRelationshipTargetType = types.StringValue(organizationField.RelationshipTargetType)
	} else {
		tfRelationshipTargetType = types.StringNull()
	}

	*o = OrganizationFieldResourceModel{
		UserOrgFieldResourceModel{
			ID:                     types.Int64Value(organizationField.ID),
			URL:                    types.StringValue(organizationField.URL),
			Key:                    types.StringValue(organizationField.Key),
			Type:                   types.StringValue(organizationField.Type),
			Title:                  types.StringValue(organizationField.Title),
			Description:            tfDescription,
			Position:               types.Int64Value(organizationField.Position),
			Active:                 types.BoolValue(organizationField.Active),
			System:                 types.BoolValue(organizationField.System),
			RegexpForValidation:    tfRegexpForValidation,
			Tag:                    tfTag,
			CustomFieldOptions:     tfCustomOptionsList,
			CreatedAt:              types.StringValue(organizationField.CreatedAt.UTC().String()),
			UpdatedAt:              types.StringValue(organizationField.UpdatedAt.UTC().String()),
			RelationshipTargetType: tfRelationshipTargetType,
			RelationshipFilter:     tfRelationshipFilterObject,
		},
	}

	return diags
}

// RelationshipFilterResourceModel is struct for value of `relationship_filter`
type RelationshipFilterResourceModel struct {
	All types.List `tfsdk:"all"`
	Any types.List `tfsdk:"any"`
}

func (m RelationshipFilterResourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"all": types.ListType{ElemType: types.ObjectType{AttrTypes: RelationshipFilterObjectResourceModel{}.AttributeTypes()}},
		"any": types.ListType{ElemType: types.ObjectType{AttrTypes: RelationshipFilterObjectResourceModel{}.AttributeTypes()}},
	}
}

type RelationshipFilterObjectResourceModel struct {
	Field    types.String `tfsdk:"field"`
	Operator types.String `tfsdk:"operator"`
	Value    types.String `tfsdk:"value"`
}

func (m RelationshipFilterObjectResourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"field":    types.StringType,
		"operator": types.StringType,
		"value":    types.StringType,
	}
}
