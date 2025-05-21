package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTicketFormResourceModel(t *testing.T) {

	var (
		diags = diag.Diagnostics{}

		ctx = getTestContext(t)

		testTicketFormResourceModelSimpleInput = TicketFormResourceModel{
			Name: types.StringValue(testTitle),
			TicketFieldIds: types.ListValueMust(types.Int64Type, []attr.Value{
				types.Int64Value(12345),
				types.Int64Value(56789),
			}),
			AgentConditions:   types.MapNull(types.ObjectType{AttrTypes: FormConditions{}.AttributeTypes()}),
			EndUserConditions: types.MapNull(types.ObjectType{AttrTypes: FormConditions{}.AttributeTypes()}),
		}
		testTicketFormSimpleExpected = zendesk.TicketForm{
			Name:              testTitle,
			TicketFieldIds:    []int64{12345, 56789},
			AgentConditions:   []zendesk.ConditionalTicketField(nil),
			EndUserConditions: []zendesk.ConditionalTicketField(nil),
		}
		testRequiredOnStatusResourceModel = RequiredOnStatusesResourceModel{
			Statuses: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("open"),
				types.StringValue("pending"),
			}),
			Type: types.StringValue(string(zendesk.SomeStatuses)),
		}
	)

	testRequiredOnStatusObj, diag1 := types.ObjectValueFrom(context.Background(), testRequiredOnStatusResourceModel.AttributeTypes(), testRequiredOnStatusResourceModel)

	diags.Append(diag1...)

	if diags.HasError() {
		diagnosticErrorHelper(t, diags, "Errors from converting testRequiredOnStatusResourceModel")
	}

	testChildFieldsModel := []FormChildFieldConditions{
		{
			Id:                 types.Int64Value(56789),
			IsRequired:         types.BoolValue(true),
			RequiredOnStatuses: testRequiredOnStatusObj,
		},
	}

	testChildFieldSet, diag2 := types.SetValueFrom(context.Background(), types.ObjectType{AttrTypes: FormChildFieldConditions{}.AttributeTypes()}, testChildFieldsModel)

	diags.Append(diag2...)

	if diags.HasError() {
		diagnosticErrorHelper(t, diags, "Errors from converting testChildFieldsModel")
	}

	testFieldValueMap := map[string]FormConditions{
		"some_value": {
			ChildConditions: testChildFieldSet,
		},
	}

	testFieldValueMapTf, diag3 := types.MapValueFrom(ctx, types.ObjectType{AttrTypes: FormConditions{}.AttributeTypes()}, testFieldValueMap)

	diags.Append(diag3...)

	if diags.HasError() {
		diagnosticErrorHelper(t, diags, "Errors from converting testFieldValueMap")
	}

	testAgentConditionObject := map[string]FormConditionsSet{
		"12345": {
			FieldValueMap: testFieldValueMapTf,
		},
	}

	testAgentConditions, diag4 := types.MapValueFrom(context.Background(),
		types.ObjectType{AttrTypes: FormConditionsSet{}.AttributeTypes()},
		testAgentConditionObject,
	)

	diags.Append(diag4...)

	if diags.HasError() {
		diagnosticErrorHelper(t, diags, "Errors from converting testAgentConditions")
	}

	var (
		testTicketFormResourceModelConditionsInput = TicketFormResourceModel{
			Name: types.StringValue(testTitle),
			TicketFieldIds: types.ListValueMust(types.Int64Type, []attr.Value{
				types.Int64Value(12345),
				types.Int64Value(56789),
			}),
			AgentConditions:   testAgentConditions,
			EndUserConditions: types.MapNull(types.ObjectType{AttrTypes: FormConditionsSet{}.AttributeTypes()}),
		}
		testTicketFormConditionsExpected = zendesk.TicketForm{
			Name:           testTitle,
			TicketFieldIds: []int64{12345, 56789},
			AgentConditions: []zendesk.ConditionalTicketField{
				{
					ParentFieldId: 12345,
					Value:         "some_value",
					ChildFields: []zendesk.ChildField{
						{
							Id:         56789,
							IsRequired: true,
							RequiredOnStatuses: zendesk.RequiredOnStatuses{
								Statuses: []string{"open", "pending"},
								Type:     zendesk.SomeStatuses,
							},
						},
					},
				},
			},
			EndUserConditions: []zendesk.ConditionalTicketField(nil),
		}

		testTicketFormSimpleInput = zendesk.TicketForm{
			Active:            true,
			AgentConditions:   nil,
			CreatedAt:         testCreatedAt,
			Default:           false,
			DisplayName:       testTitle,
			EndUserConditions: nil,
			EndUserVisible:    false,
			ID:                testId,
			Name:              testTitle,
			Position:          testPosition,
			TicketFieldIds:    []int64{12345, 56789},
			UpdatedAt:         testUpdatedAt,
			Url:               testUrl,
		}

		testTicketFormResourceModelSimpleExpected = TicketFormResourceModel{
			ID:          types.Int64Value(testId),
			Name:        types.StringValue(testTitle),
			DisplayName: types.StringValue(testTitle),
			TicketFieldIds: types.ListValueMust(types.Int64Type, []attr.Value{
				types.Int64Value(12345),
				types.Int64Value(56789),
			}),
			AgentConditions:   types.MapNull(types.ObjectType{AttrTypes: FormConditionsSet{}.AttributeTypes()}),
			EndUserConditions: types.MapNull(types.ObjectType{AttrTypes: FormConditionsSet{}.AttributeTypes()}),
			Active:            types.BoolValue(true),
			Position:          types.Int64Value(testPosition),
			Default:           types.BoolValue(false),
			EndUserVisible:    types.BoolValue(false),
			CreatedAt:         types.StringValue(testCreatedAt.UTC().String()),
			UpdatedAt:         types.StringValue(testUpdatedAt.UTC().String()),
			Url:               types.StringValue(testUrl),
		}

		testTicketFormConditionsInput = zendesk.TicketForm{
			Active: true,
			AgentConditions: []zendesk.ConditionalTicketField{
				{
					ParentFieldId: 12345,
					Value:         "some_value",
					ChildFields: []zendesk.ChildField{
						{
							Id:         56789,
							IsRequired: true,
							RequiredOnStatuses: zendesk.RequiredOnStatuses{
								Statuses: []string{"open", "pending"},
								Type:     zendesk.SomeStatuses,
							},
						},
					},
				},
			},
			EndUserConditions: nil,
			CreatedAt:         testCreatedAt,
			Default:           false,
			DisplayName:       testTitle,
			EndUserVisible:    false,
			ID:                testId,
			Name:              testTitle,
			Position:          testPosition,
			TicketFieldIds:    []int64{12345, 56789},
			UpdatedAt:         testUpdatedAt,
			Url:               testUrl,
		}

		testTicketFormResourceModelConditionsExpected = TicketFormResourceModel{
			ID:          types.Int64Value(testId),
			Name:        types.StringValue(testTitle),
			DisplayName: types.StringValue(testTitle),
			TicketFieldIds: types.ListValueMust(types.Int64Type, []attr.Value{
				types.Int64Value(12345),
				types.Int64Value(56789),
			}),
			AgentConditions:   testAgentConditions,
			EndUserConditions: types.MapNull(types.ObjectType{AttrTypes: FormConditionsSet{}.AttributeTypes()}),
			Active:            types.BoolValue(true),
			Position:          types.Int64Value(testPosition),
			Default:           types.BoolValue(false),
			EndUserVisible:    types.BoolValue(false),
			CreatedAt:         types.StringValue(testCreatedAt.UTC().String()),
			UpdatedAt:         types.StringValue(testUpdatedAt.UTC().String()),
			Url:               types.StringValue(testUrl),
		}
	)

	t.Run("GetApiModelFromTfModel", func(t *testing.T) {
		cases := []struct {
			testName string
			input    TicketFormResourceModel
			expected zendesk.TicketForm
		}{
			{
				testName: "should return a ticket form api resource from terraform resource with no conditional fields",
				input:    testTicketFormResourceModelSimpleInput,
				expected: testTicketFormSimpleExpected,
			}, {
				testName: "should return a ticket form api resource from TF " +
					"resource with agent conditional fields",
				input:    testTicketFormResourceModelConditionsInput,
				expected: testTicketFormConditionsExpected,
			},
		}

		for _, c := range cases {
			t.Run(c.testName, func(t *testing.T) {
				out, diagTest := c.input.GetApiModelFromTfModel(context.Background())

				diags.Append(diagTest...)

				if diags.HasError() {
					diagnosticErrorHelper(t, diags, "Diagnostic Error found running 'GetApiModelFromTfModel'")
				} else {
					assert.Equal(t, out, c.expected)
				}

			})
		}
	})

	t.Run("GetTfModelFromApiModel", func(t *testing.T) {
		cases := []struct {
			testName string
			target   TicketFormResourceModel
			input    zendesk.TicketForm
			expected TicketFormResourceModel
		}{
			{
				testName: "should return a ticket form terraform resource from api resource with no conditional fields",
				target:   TicketFormResourceModel{},
				input:    testTicketFormSimpleInput,
				expected: testTicketFormResourceModelSimpleExpected,
			},
			{
				testName: "should return a ticket form TF resource from api " +
					"with agent conditional fields",
				target:   TicketFormResourceModel{},
				input:    testTicketFormConditionsInput,
				expected: testTicketFormResourceModelConditionsExpected,
			},
		}

		for _, c := range cases {
			t.Run(c.testName, func(t *testing.T) {
				diags.Append(c.target.GetTfModelFromApiModel(context.Background(), c.input)...)
				if diags.HasError() {
					diagnosticErrorHelper(t, diags, "Diagnostic Error found running 'GetTfModelFromApiModel'")
				} else {
					assert.Equal(t, c.expected, c.target)
				}

			})
		}
	})

}
