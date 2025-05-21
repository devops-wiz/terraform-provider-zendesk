package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"reflect"
	"testing"
)

func TestUserFieldResourceModel_GetApiModelFromTfModel(t *testing.T) {
	cases := []struct {
		testName string
		input    UserFieldResourceModel
		expected zendesk.UserField
	}{
		{
			testName: "should get api model from basic tf model",
			input: UserFieldResourceModel{
				UserOrgFieldResourceModel{
					Key:         types.StringValue(testKey),
					Type:        types.StringValue(testTicketFieldType),
					Title:       types.StringValue(testTicketFieldTitle),
					Description: types.StringValue(testTicketFieldDesc),
				},
			},
			expected: zendesk.UserField{
				Key:         testKey,
				Type:        testTicketFieldType,
				Title:       testTicketFieldTitle,
				Description: testTicketFieldDesc,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			out, diags := c.input.GetApiModelFromTfModel(context.Background())
			if diags.HasError() {
				t.Fatalf("%s. got diags error: %s", c.testName, diags.Errors())
			}
			if !reflect.DeepEqual(out, c.expected) {
				t.Fatalf(errorOutputMismatch, c.testName, out, c.expected)
			}
		})
	}
}

func TestUserFieldResourceModel_GetTfModelFromApiModel(t *testing.T) {
	cases := []struct {
		testName string
		target   UserFieldResourceModel
		input    zendesk.UserField
		expected UserFieldResourceModel
	}{
		{
			testName: "should get tf model from basic api model",
			target:   UserFieldResourceModel{},
			input: zendesk.UserField{
				ID:          testTicketFieldId,
				URL:         testTicketFieldUrl,
				Key:         testKey,
				Type:        testTicketFieldType,
				Title:       testTicketFieldTitle,
				Description: testTicketFieldDesc,
				Position:    testTicketFieldPosition,
				Active:      true,
				System:      false,
				CreatedAt:   testCreatedAt,
				UpdatedAt:   testUpdatedAt,
			},
			expected: UserFieldResourceModel{
				UserOrgFieldResourceModel{
					ID:                     types.Int64Value(testTicketFieldId),
					URL:                    types.StringValue(testTicketFieldUrl),
					Key:                    types.StringValue(testKey),
					Type:                   types.StringValue(testTicketFieldType),
					Title:                  types.StringValue(testTicketFieldTitle),
					Description:            types.StringValue(testTicketFieldDesc),
					Position:               types.Int64Value(testTicketFieldPosition),
					Active:                 types.BoolValue(true),
					System:                 types.BoolValue(false),
					RegexpForValidation:    types.StringNull(),
					Tag:                    types.StringNull(),
					CustomFieldOptions:     types.ListNull(types.ObjectType{AttrTypes: CustomFieldOptionResourceModel{}.AttributeTypes()}),
					CreatedAt:              types.StringValue(testCreatedAt.UTC().String()),
					UpdatedAt:              types.StringValue(testUpdatedAt.UTC().String()),
					RelationshipTargetType: types.StringNull(),
					RelationshipFilter:     types.ObjectNull(RelationshipFilterResourceModel{}.AttributeTypes()),
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			diags := c.target.GetTfModelFromApiModel(context.Background(), c.input)
			if diags.HasError() {
				t.Fatalf("%s. got diags error: %s", c.testName, diags.Errors())
			}
			if !reflect.DeepEqual(c.expected, c.target) {
				t.Fatalf(errorOutputMismatch, c.testName, c.target, c.expected)
			}
		})
	}
}
