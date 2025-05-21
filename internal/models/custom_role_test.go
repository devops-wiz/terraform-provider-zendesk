package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"reflect"
	"testing"
)

var testConfigModel = ConfigurationResourceModel{
	AssignTicketsToAnyBrand:      types.BoolValue(false),
	AssignTicketsToAnyGroup:      types.BoolValue(false),
	ChatAccess:                   types.BoolValue(false),
	CustomObjects:                types.MapNull(types.ObjectType{AttrTypes: ScopesResourceModel{}.AttributeTypes()}),
	EndUserListAccess:            types.StringValue(""),
	EndUserProfileAccess:         types.StringValue(""),
	ExploreAccess:                types.StringValue(""),
	ForumAccess:                  types.StringValue(""),
	ForumAccessRestrictedContent: types.BoolValue(false),
	LightAgent:                   types.BoolValue(false),
	MacroAccess:                  types.StringValue(""),
	ManageAutomations:            types.BoolValue(false),
	ManageBusinessRules:          types.BoolValue(false),
	ManageContextualWorkspaces:   types.BoolValue(false),
	ManageDynamicContent:         types.BoolValue(false),
	ManageExtensionsAndChannels:  types.BoolValue(false),
	ManageFacebook:               types.BoolValue(false),
	ManageGroupMemberships:       types.BoolValue(false),
	ManageGroups:                 types.BoolValue(false),
	ManageOrganizationFields:     types.BoolValue(false),
	ManageOrganizations:          types.BoolValue(false),
	ManageRoles:                  types.StringValue(""),
	ManageSkills:                 types.BoolValue(false),
	ManageSlas:                   types.BoolValue(false),
	ManageSuspendedTickets:       types.BoolValue(false),
	ManageTeamMembers:            types.StringValue(""),
	ManageTicketFields:           types.BoolValue(false),
	ManageTicketForms:            types.BoolValue(false),
	ManageTriggers:               types.BoolValue(false),
	ManageUserFields:             types.BoolValue(false),
	OrganizationEditing:          types.BoolValue(false),
	OrganizationNotesEditing:     types.BoolValue(false),
	ReportAccess:                 types.StringValue(""),
	SideConversationCreate:       types.BoolValue(false),
	TicketAccess:                 types.StringValue(""),
	TicketCommentAccess:          types.StringValue(""),
	TicketDeletion:               types.BoolValue(false),
	TicketRedaction:              types.BoolValue(false),
	ViewDeletedTickets:           types.BoolValue(false),
	TicketEditing:                types.BoolValue(false),
	TicketMerge:                  types.BoolValue(false),
	TicketTagEditing:             types.BoolValue(false),
	TwitterSearchAccess:          types.BoolValue(false),
	ViewAccess:                   types.StringValue(""),
	VoiceAccess:                  types.BoolValue(false),
	VoiceDashboardAccess:         types.BoolValue(false),
}

var testConfigObject, _ = types.ObjectValueFrom(context.Background(), testConfigModel.AttributeTypes(), testConfigModel)

var testConfig = zendesk.Configuration{
	AssignTicketsToAnyBrand:      false,
	AssignTicketsToAnyGroup:      false,
	ChatAccess:                   false,
	CustomObjects:                nil,
	EndUserListAccess:            "",
	EndUserProfileAccess:         "",
	ExploreAccess:                "",
	ForumAccess:                  "",
	ForumAccessRestrictedContent: false,
	LightAgent:                   false,
	MacroAccess:                  "",
	ManageAutomations:            false,
	ManageBusinessRules:          false,
	ManageContextualWorkspaces:   false,
	ManageDynamicContent:         false,
	ManageExtensionsAndChannels:  false,
	ManageFacebook:               false,
	ManageGroupMemberships:       false,
	ManageGroups:                 false,
	ManageOrganizationFields:     false,
	ManageOrganizations:          false,
	ManageRoles:                  "",
	ManageSkills:                 false,
	ManageSlas:                   false,
	ManageSuspendedTickets:       false,
	ManageTeamMembers:            "",
	ManageTicketFields:           false,
	ManageTicketForms:            false,
	ManageTriggers:               false,
	ManageUserFields:             false,
	OrganizationEditing:          false,
	OrganizationNotesEditing:     false,
	ReportAccess:                 "",
	SideConversationCreate:       false,
	TicketAccess:                 "",
	TicketCommentAccess:          "",
	TicketDeletion:               false,
	TicketRedaction:              false,
	ViewDeletedTickets:           false,
	TicketEditing:                false,
	TicketMerge:                  false,
	TicketTagEditing:             false,
	TwitterSearchAccess:          false,
	ViewAccess:                   "",
	VoiceAccess:                  false,
	VoiceDashboardAccess:         false,
}

var testRole = zendesk.CustomRole{
	Description:   testDescription,
	ID:            testId,
	Name:          testTitle,
	Configuration: testConfig,
	RoleType:      1,
}

var testRoleModel = CustomRoleResourceModel{
	Description:   types.StringValue(testDescription),
	ID:            types.Int64Value(testId),
	Name:          types.StringValue(testTitle),
	Configuration: testConfigObject,
	RoleType:      types.Int64Value(1),
}

func TestCustomRoleResourceModel_GetApiModelFromTfModel(t *testing.T) {

	cases := []struct {
		testName string
		input    CustomRoleResourceModel
		expected zendesk.CustomRole
	}{
		{
			testName: "Get Custom Role API Model from tf",
			input:    testRoleModel,
			expected: testRole,
		},
	}
	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			out, diags := c.input.GetApiModelFromTfModel(context.Background())
			if diags.HasError() {
				t.Error(diags)
			}
			if !reflect.DeepEqual(out, c.expected) {
				t.Fatalf(errorOutputMismatch, c.testName, out, c.expected)
			}
		})
	}
}

func TestCustomRoleResourceModel_GetTfModelFromApiModel(t *testing.T) {
	cases := []struct {
		testName string
		target   CustomRoleResourceModel
		input    zendesk.CustomRole
		expected CustomRoleResourceModel
	}{
		{
			testName: "Get Custom Role TF Model from api",
			target:   CustomRoleResourceModel{},
			input:    testRole,
			expected: testRoleModel,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			c.target.GetTfModelFromApiModel(context.Background(), c.input)
			if !reflect.DeepEqual(c.target, c.expected) {
				t.Fatalf(errorOutputMismatch, c.testName, c.target, c.expected)
			}
		})
	}
}
