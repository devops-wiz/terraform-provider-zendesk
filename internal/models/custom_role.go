package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ ResourceTransformWithID[zendesk.CustomRole] = &CustomRoleResourceModel{}

type CustomRoleResourceModel struct {
	Description   types.String `tfsdk:"description"`
	ID            types.Int64  `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	Configuration types.Object `tfsdk:"configuration"`
	RoleType      types.Int64  `tfsdk:"role_type"`
}

type ScopesResourceModel struct {
	Scopes []types.String `tfsdk:"scopes"`
}

func (s ScopesResourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"scopes": types.SetType{ElemType: types.StringType},
	}
}

type ConfigurationResourceModel struct {
	AssignTicketsToAnyBrand      types.Bool   `tfsdk:"assign_tickets_to_any_brand"`
	AssignTicketsToAnyGroup      types.Bool   `tfsdk:"assign_tickets_to_any_group"`
	ChatAccess                   types.Bool   `tfsdk:"chat_access"`
	CustomObjects                types.Map    `tfsdk:"custom_objects"`
	EndUserListAccess            types.String `tfsdk:"end_user_list_access"`
	EndUserProfileAccess         types.String `tfsdk:"end_user_profile_access"`
	ExploreAccess                types.String `tfsdk:"explore_access"`
	ForumAccess                  types.String `tfsdk:"forum_access"`
	ForumAccessRestrictedContent types.Bool   `tfsdk:"forum_access_restricted_content"`
	LightAgent                   types.Bool   `tfsdk:"light_agent"`
	MacroAccess                  types.String `tfsdk:"macro_access"`
	ManageAutomations            types.Bool   `tfsdk:"manage_automations"`
	ManageBusinessRules          types.Bool   `tfsdk:"manage_business_rules"`
	ManageContextualWorkspaces   types.Bool   `tfsdk:"manage_contextual_workspaces"`
	ManageDynamicContent         types.Bool   `tfsdk:"manage_dynamic_content"`
	ManageExtensionsAndChannels  types.Bool   `tfsdk:"manage_extensions_and_channels"`
	ManageFacebook               types.Bool   `tfsdk:"manage_facebook"`
	ManageGroupMemberships       types.Bool   `tfsdk:"manage_group_memberships"`
	ManageGroups                 types.Bool   `tfsdk:"manage_groups"`
	ManageOrganizationFields     types.Bool   `tfsdk:"manage_organization_fields"`
	ManageOrganizations          types.Bool   `tfsdk:"manage_organizations"`
	ManageRoles                  types.String `tfsdk:"manage_roles"`
	ManageSkills                 types.Bool   `tfsdk:"manage_skills"`
	ManageSlas                   types.Bool   `tfsdk:"manage_slas"`
	ManageSuspendedTickets       types.Bool   `tfsdk:"manage_suspended_tickets"`
	ManageTeamMembers            types.String `tfsdk:"manage_team_members"`
	ManageTicketFields           types.Bool   `tfsdk:"manage_ticket_fields"`
	ManageTicketForms            types.Bool   `tfsdk:"manage_ticket_forms"`
	ManageTriggers               types.Bool   `tfsdk:"manage_triggers"`
	ManageUserFields             types.Bool   `tfsdk:"manage_user_fields"`
	OrganizationEditing          types.Bool   `tfsdk:"organization_editing"`
	OrganizationNotesEditing     types.Bool   `tfsdk:"organization_notes_editing"`
	ReportAccess                 types.String `tfsdk:"report_access"`
	SideConversationCreate       types.Bool   `tfsdk:"side_conversation_create"`
	TicketAccess                 types.String `tfsdk:"ticket_access"`
	TicketCommentAccess          types.String `tfsdk:"ticket_comment_access"`
	TicketDeletion               types.Bool   `tfsdk:"ticket_deletion"`
	TicketRedaction              types.Bool   `tfsdk:"ticket_redaction"`
	ViewDeletedTickets           types.Bool   `tfsdk:"view_deleted_tickets"`
	TicketEditing                types.Bool   `tfsdk:"ticket_editing"`
	TicketMerge                  types.Bool   `tfsdk:"ticket_merge"`
	TicketTagEditing             types.Bool   `tfsdk:"ticket_tag_editing"`
	TwitterSearchAccess          types.Bool   `tfsdk:"twitter_search_access"`
	ViewAccess                   types.String `tfsdk:"view_access"`
	VoiceAccess                  types.Bool   `tfsdk:"voice_access"`
	VoiceDashboardAccess         types.Bool   `tfsdk:"voice_dashboard_access"`
}

func (c *CustomRoleResourceModel) GetID() int64 {
	return c.ID.ValueInt64()
}

func (m ConfigurationResourceModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"assign_tickets_to_any_brand":     types.BoolType,
		"assign_tickets_to_any_group":     types.BoolType,
		"chat_access":                     types.BoolType,
		"custom_objects":                  types.MapType{ElemType: types.ObjectType{AttrTypes: ScopesResourceModel{}.AttributeTypes()}},
		"end_user_list_access":            types.StringType,
		"end_user_profile_access":         types.StringType,
		"explore_access":                  types.StringType,
		"forum_access":                    types.StringType,
		"forum_access_restricted_content": types.BoolType,
		"light_agent":                     types.BoolType,
		"macro_access":                    types.StringType,
		"manage_automations":              types.BoolType,
		"manage_business_rules":           types.BoolType,
		"manage_contextual_workspaces":    types.BoolType,
		"manage_dynamic_content":          types.BoolType,
		"manage_extensions_and_channels":  types.BoolType,
		"manage_facebook":                 types.BoolType,
		"manage_group_memberships":        types.BoolType,
		"manage_groups":                   types.BoolType,
		"manage_organization_fields":      types.BoolType,
		"manage_organizations":            types.BoolType,
		"manage_roles":                    types.StringType,
		"manage_skills":                   types.BoolType,
		"manage_slas":                     types.BoolType,
		"manage_suspended_tickets":        types.BoolType,
		"manage_team_members":             types.StringType,
		"manage_ticket_fields":            types.BoolType,
		"manage_ticket_forms":             types.BoolType,
		"manage_triggers":                 types.BoolType,
		"manage_user_fields":              types.BoolType,
		"organization_editing":            types.BoolType,
		"organization_notes_editing":      types.BoolType,
		"report_access":                   types.StringType,
		"side_conversation_create":        types.BoolType,
		"ticket_access":                   types.StringType,
		"ticket_comment_access":           types.StringType,
		"ticket_deletion":                 types.BoolType,
		"ticket_redaction":                types.BoolType,
		"view_deleted_tickets":            types.BoolType,
		"ticket_editing":                  types.BoolType,
		"ticket_merge":                    types.BoolType,
		"ticket_tag_editing":              types.BoolType,
		"twitter_search_access":           types.BoolType,
		"view_access":                     types.StringType,
		"voice_access":                    types.BoolType,
		"voice_dashboard_access":          types.BoolType,
	}
}

func (c *CustomRoleResourceModel) GetApiModelFromTfModel(ctx context.Context) (apiRole zendesk.CustomRole, diags diag.Diagnostics) {
	apiRole = zendesk.CustomRole{
		Description: c.Description.ValueString(),
		ID:          c.ID.ValueInt64(),
		Name:        c.Name.ValueString(),
		RoleType:    c.RoleType.ValueInt64(),
	}

	var configModel ConfigurationResourceModel

	diags = c.Configuration.As(ctx, &configModel, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})

	if diags.HasError() {
		return zendesk.CustomRole{}, diags
	}

	var customObjects zendesk.Scopes

	if !configModel.CustomObjects.IsUnknown() && !configModel.CustomObjects.IsNull() {

		customObjects = make(zendesk.Scopes, len(configModel.CustomObjects.Elements()))

		elements := make(map[string]types.Object, len(configModel.CustomObjects.Elements()))

		diags = configModel.CustomObjects.ElementsAs(ctx, &elements, false)

		if diags.HasError() {
			return zendesk.CustomRole{}, diags
		}

		for key, object := range elements {

			var scopeObj ScopesResourceModel

			diags = object.As(ctx, &scopeObj, basetypes.ObjectAsOptions{
				UnhandledNullAsEmpty:    true,
				UnhandledUnknownAsEmpty: true,
			})

			customObjects[key] = struct {
				Scopes []string `json:"scopes"`
			}{Scopes: make([]string, len(scopeObj.Scopes))}

			for i, scope := range scopeObj.Scopes {
				customObjects[key].Scopes[i] = scope.ValueString()
			}
		}

	} else {
		customObjects = nil
	}
	apiRole.Configuration = zendesk.Configuration{
		AssignTicketsToAnyBrand:      configModel.AssignTicketsToAnyBrand.ValueBool(),
		AssignTicketsToAnyGroup:      configModel.AssignTicketsToAnyGroup.ValueBool(),
		ChatAccess:                   configModel.ChatAccess.ValueBool(),
		EndUserListAccess:            configModel.EndUserListAccess.ValueString(),
		EndUserProfileAccess:         configModel.EndUserProfileAccess.ValueString(),
		ExploreAccess:                configModel.ExploreAccess.ValueString(),
		ForumAccess:                  configModel.ForumAccess.ValueString(),
		ForumAccessRestrictedContent: configModel.ForumAccessRestrictedContent.ValueBool(),
		LightAgent:                   configModel.LightAgent.ValueBool(),
		MacroAccess:                  configModel.MacroAccess.ValueString(),
		ManageAutomations:            configModel.ManageAutomations.ValueBool(),
		ManageBusinessRules:          configModel.ManageBusinessRules.ValueBool(),
		ManageContextualWorkspaces:   configModel.ManageContextualWorkspaces.ValueBool(),
		ManageDynamicContent:         configModel.ManageDynamicContent.ValueBool(),
		ManageExtensionsAndChannels:  configModel.ManageExtensionsAndChannels.ValueBool(),
		ManageFacebook:               configModel.ManageFacebook.ValueBool(),
		ManageGroupMemberships:       configModel.ManageGroupMemberships.ValueBool(),
		ManageGroups:                 configModel.ManageGroups.ValueBool(),
		ManageOrganizationFields:     configModel.ManageOrganizationFields.ValueBool(),
		ManageOrganizations:          configModel.ManageOrganizations.ValueBool(),
		ManageRoles:                  configModel.ManageRoles.ValueString(),
		ManageSkills:                 configModel.ManageSkills.ValueBool(),
		ManageSlas:                   configModel.ManageSlas.ValueBool(),
		ManageSuspendedTickets:       configModel.ManageSuspendedTickets.ValueBool(),
		ManageTeamMembers:            configModel.ManageTeamMembers.ValueString(),
		ManageTicketFields:           configModel.ManageTicketFields.ValueBool(),
		ManageTicketForms:            configModel.ManageTicketForms.ValueBool(),
		ManageTriggers:               configModel.ManageTriggers.ValueBool(),
		ManageUserFields:             configModel.ManageUserFields.ValueBool(),
		OrganizationEditing:          configModel.OrganizationEditing.ValueBool(),
		OrganizationNotesEditing:     configModel.OrganizationNotesEditing.ValueBool(),
		ReportAccess:                 configModel.ReportAccess.ValueString(),
		SideConversationCreate:       configModel.SideConversationCreate.ValueBool(),
		TicketAccess:                 configModel.TicketAccess.ValueString(),
		TicketCommentAccess:          configModel.TicketCommentAccess.ValueString(),
		TicketDeletion:               configModel.TicketDeletion.ValueBool(),
		TicketRedaction:              configModel.TicketRedaction.ValueBool(),
		ViewDeletedTickets:           configModel.ViewDeletedTickets.ValueBool(),
		TicketEditing:                configModel.TicketEditing.ValueBool(),
		TicketMerge:                  configModel.TicketMerge.ValueBool(),
		TicketTagEditing:             configModel.TicketTagEditing.ValueBool(),
		TwitterSearchAccess:          configModel.TwitterSearchAccess.ValueBool(),
		ViewAccess:                   configModel.ViewAccess.ValueString(),
		VoiceAccess:                  configModel.VoiceAccess.ValueBool(),
		VoiceDashboardAccess:         configModel.VoiceDashboardAccess.ValueBool(),
		CustomObjects:                customObjects,
	}

	return apiRole, diags
}

func (c *CustomRoleResourceModel) GetTfModelFromApiModel(ctx context.Context, apiRole zendesk.CustomRole) (diags diag.Diagnostics) {

	var configObject types.Object

	var customObjectsMap types.Map

	if len(apiRole.Configuration.CustomObjects) > 0 {

		customObjectsMap, diags = types.MapValueFrom(ctx, types.ObjectType{AttrTypes: ScopesResourceModel{}.AttributeTypes()}, apiRole.Configuration.CustomObjects)

		if diags.HasError() {
			return diags
		}

	} else {
		customObjectsMap = types.MapNull(types.ObjectType{AttrTypes: ScopesResourceModel{}.AttributeTypes()})
	}

	rawConfig := ConfigurationResourceModel{
		AssignTicketsToAnyBrand:      types.BoolValue(apiRole.Configuration.AssignTicketsToAnyBrand),
		AssignTicketsToAnyGroup:      types.BoolValue(apiRole.Configuration.AssignTicketsToAnyGroup),
		ChatAccess:                   types.BoolValue(apiRole.Configuration.ChatAccess),
		CustomObjects:                customObjectsMap,
		EndUserListAccess:            types.StringValue(apiRole.Configuration.EndUserListAccess),
		EndUserProfileAccess:         types.StringValue(apiRole.Configuration.EndUserProfileAccess),
		ExploreAccess:                types.StringValue(apiRole.Configuration.ExploreAccess),
		ForumAccess:                  types.StringValue(apiRole.Configuration.ForumAccess),
		ForumAccessRestrictedContent: types.BoolValue(apiRole.Configuration.ForumAccessRestrictedContent),
		LightAgent:                   types.BoolValue(apiRole.Configuration.LightAgent),
		MacroAccess:                  types.StringValue(apiRole.Configuration.MacroAccess),
		ManageAutomations:            types.BoolValue(apiRole.Configuration.ManageAutomations),
		ManageBusinessRules:          types.BoolValue(apiRole.Configuration.ManageBusinessRules),
		ManageContextualWorkspaces:   types.BoolValue(apiRole.Configuration.ManageContextualWorkspaces),
		ManageDynamicContent:         types.BoolValue(apiRole.Configuration.ManageDynamicContent),
		ManageExtensionsAndChannels:  types.BoolValue(apiRole.Configuration.ManageExtensionsAndChannels),
		ManageFacebook:               types.BoolValue(apiRole.Configuration.ManageFacebook),
		ManageGroupMemberships:       types.BoolValue(apiRole.Configuration.ManageGroupMemberships),
		ManageGroups:                 types.BoolValue(apiRole.Configuration.ManageGroups),
		ManageOrganizationFields:     types.BoolValue(apiRole.Configuration.ManageOrganizationFields),
		ManageOrganizations:          types.BoolValue(apiRole.Configuration.ManageOrganizations),
		ManageRoles:                  types.StringValue(apiRole.Configuration.ManageRoles),
		ManageSkills:                 types.BoolValue(apiRole.Configuration.ManageSkills),
		ManageSlas:                   types.BoolValue(apiRole.Configuration.ManageSlas),
		ManageSuspendedTickets:       types.BoolValue(apiRole.Configuration.ManageSuspendedTickets),
		ManageTeamMembers:            types.StringValue(apiRole.Configuration.ManageTeamMembers),
		ManageTicketFields:           types.BoolValue(apiRole.Configuration.ManageTicketFields),
		ManageTicketForms:            types.BoolValue(apiRole.Configuration.ManageTicketForms),
		ManageTriggers:               types.BoolValue(apiRole.Configuration.ManageTriggers),
		ManageUserFields:             types.BoolValue(apiRole.Configuration.ManageUserFields),
		OrganizationEditing:          types.BoolValue(apiRole.Configuration.OrganizationEditing),
		OrganizationNotesEditing:     types.BoolValue(apiRole.Configuration.OrganizationNotesEditing),
		ReportAccess:                 types.StringValue(apiRole.Configuration.ReportAccess),
		SideConversationCreate:       types.BoolValue(apiRole.Configuration.SideConversationCreate),
		TicketAccess:                 types.StringValue(apiRole.Configuration.TicketAccess),
		TicketCommentAccess:          types.StringValue(apiRole.Configuration.TicketCommentAccess),
		TicketDeletion:               types.BoolValue(apiRole.Configuration.TicketDeletion),
		TicketRedaction:              types.BoolValue(apiRole.Configuration.TicketRedaction),
		ViewDeletedTickets:           types.BoolValue(apiRole.Configuration.ViewDeletedTickets),
		TicketEditing:                types.BoolValue(apiRole.Configuration.TicketEditing),
		TicketMerge:                  types.BoolValue(apiRole.Configuration.TicketMerge),
		TicketTagEditing:             types.BoolValue(apiRole.Configuration.TicketTagEditing),
		TwitterSearchAccess:          types.BoolValue(apiRole.Configuration.TwitterSearchAccess),
		ViewAccess:                   types.StringValue(apiRole.Configuration.ViewAccess),
		VoiceAccess:                  types.BoolValue(apiRole.Configuration.VoiceAccess),
		VoiceDashboardAccess:         types.BoolValue(apiRole.Configuration.VoiceDashboardAccess),
	}

	configObject, diags = types.ObjectValueFrom(ctx, ConfigurationResourceModel{}.AttributeTypes(), rawConfig)

	*c = CustomRoleResourceModel{
		Description:   types.StringValue(apiRole.Description),
		ID:            types.Int64Value(apiRole.ID),
		Name:          types.StringValue(apiRole.Name),
		Configuration: configObject,
		RoleType:      types.Int64Value(apiRole.RoleType),
	}

	return diags
}
