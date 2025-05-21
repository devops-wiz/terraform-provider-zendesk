package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SystemFieldOption struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

func (s SystemFieldOption) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":  types.StringType,
		"value": types.StringType,
	}
}

var _ ResourceTransformWithID[zendesk.TicketField] = &TicketFieldResourceModel{}

type TicketFieldResourceModel struct {
	ID                  types.Int64  `tfsdk:"id"`
	Title               types.String `tfsdk:"title"`
	Type                types.String `tfsdk:"type"`
	Active              types.Bool   `tfsdk:"active"`
	AgentDescription    types.String `tfsdk:"agent_description"`
	CreatedAt           types.String `tfsdk:"created_at"`
	Description         types.String `tfsdk:"portal_description"`
	EditableInPortal    types.Bool   `tfsdk:"editable_in_portal"`
	Position            types.Int64  `tfsdk:"position"`
	RegexpForValidation types.String `tfsdk:"regexp_for_validation"`
	Required            types.Bool   `tfsdk:"required"`
	RequiredInPortal    types.Bool   `tfsdk:"required_in_portal"`
	Tag                 types.String `tfsdk:"tag"`
	TitleInPortal       types.String `tfsdk:"title_in_portal"`
	UpdatedAt           types.String `tfsdk:"updated_at"`
	URL                 types.String `tfsdk:"url"`
	VisibleInPortal     types.Bool   `tfsdk:"visible_in_portal"`
	CustomFieldOptions  types.List   `tfsdk:"custom_field_options"`
	SystemFieldOptions  types.List   `tfsdk:"system_field_options"`
}

func (t *TicketFieldResourceModel) GetID() int64 {
	return t.ID.ValueInt64()
}

func (t *TicketFieldResourceModel) GetApiModelFromTfModel(ctx context.Context) (ticketField zendesk.TicketField, diags diag.Diagnostics) {

	var cfos []zendesk.CustomFieldOption
	cfos, diags = getApiCustomFieldOptionsFromTf(ctx, t.CustomFieldOptions)

	if diags.HasError() {
		return ticketField, diags
	}

	// Generate API request body from plan

	if !t.Position.IsNull() || !t.Position.IsUnknown() {
		ticketField = zendesk.TicketField{
			ID:                  types.Int64Null().ValueInt64(),
			Title:               t.Title.ValueString(),
			TitleInPortal:       t.TitleInPortal.ValueString(),
			Type:                t.Type.ValueString(),
			Required:            t.Required.ValueBool(),
			RequiredInPortal:    t.RequiredInPortal.ValueBool(),
			Description:         t.Description.ValueString(),
			AgentDescription:    t.AgentDescription.ValueString(),
			Tag:                 t.Tag.ValueString(),
			RegexpForValidation: t.RegexpForValidation.ValueString(),
			Active:              t.Active.ValueBool(),
			VisibleInPortal:     t.VisibleInPortal.ValueBool(),
			EditableInPortal:    t.EditableInPortal.ValueBool(),
			Position:            t.Position.ValueInt64(),
			CustomFieldOptions:  cfos,
		}
	} else {
		ticketField = zendesk.TicketField{
			ID:                  types.Int64Null().ValueInt64(),
			Title:               t.Title.ValueString(),
			TitleInPortal:       t.TitleInPortal.ValueString(),
			Type:                t.Type.ValueString(),
			Required:            t.Required.ValueBool(),
			RequiredInPortal:    t.RequiredInPortal.ValueBool(),
			Description:         t.Description.ValueString(),
			AgentDescription:    t.AgentDescription.ValueString(),
			Tag:                 t.Tag.ValueString(),
			RegexpForValidation: t.RegexpForValidation.ValueString(),
			Active:              t.Active.ValueBool(),
			VisibleInPortal:     t.VisibleInPortal.ValueBool(),
			EditableInPortal:    t.EditableInPortal.ValueBool(),
			CustomFieldOptions:  cfos,
		}
	}

	return ticketField, diags

}

func (t *TicketFieldResourceModel) GetTfModelFromApiModel(ctx context.Context, apiTicketField zendesk.TicketField) (diags diag.Diagnostics) {
	cfos := apiTicketField.CustomFieldOptions

	var cfoList types.List

	cfoList, diags = getTfCustomFieldOptionsFromApi(ctx, cfos)

	if diags.HasError() {
		return diags
	}

	sfos := apiTicketField.SystemFieldOptions
	var sfoList types.List
	sfoObjectType := types.ObjectType{AttrTypes: SystemFieldOption{}.AttributeTypes()}

	if len(sfos) > 0 {
		tfsfos := make([]SystemFieldOption, len(sfos))
		for index, sfo := range sfos {
			tfsfos[index] = SystemFieldOption{
				Name:  types.StringValue(sfo.Name),
				Value: types.StringValue(sfo.Value),
			}
		}
		sfoList, diags = types.ListValueFrom(ctx, sfoObjectType, tfsfos)
		if diags.HasError() {
			return diags
		}
	} else {

		sfoList = types.ListNull(sfoObjectType)
	}

	*t = TicketFieldResourceModel{
		ID:                  types.Int64Value(apiTicketField.ID),
		Title:               types.StringValue(apiTicketField.Title),
		TitleInPortal:       types.StringValue(apiTicketField.TitleInPortal),
		Type:                types.StringValue(apiTicketField.Type),
		Required:            types.BoolValue(apiTicketField.Required),
		RequiredInPortal:    types.BoolValue(apiTicketField.RequiredInPortal),
		Description:         types.StringValue(apiTicketField.Description),
		AgentDescription:    types.StringValue(apiTicketField.AgentDescription),
		Tag:                 types.StringValue(apiTicketField.Tag),
		RegexpForValidation: types.StringValue(apiTicketField.RegexpForValidation),
		CreatedAt:           types.StringValue(apiTicketField.CreatedAt.UTC().String()),
		UpdatedAt:           types.StringValue(apiTicketField.UpdatedAt.UTC().String()),
		URL:                 types.StringValue(apiTicketField.URL),
		Active:              types.BoolValue(apiTicketField.Active),
		VisibleInPortal:     types.BoolValue(apiTicketField.VisibleInPortal),
		EditableInPortal:    types.BoolValue(apiTicketField.EditableInPortal),
		Position:            types.Int64Value(apiTicketField.Position),
		CustomFieldOptions:  cfoList,
		SystemFieldOptions:  sfoList,
	}
	return diags
}
