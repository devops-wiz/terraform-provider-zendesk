package models

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	testTriggerId          int64 = 123
	testTriggerTitle             = "test title"
	testTriggerDesc              = "test desc"
	testTriggerCatId             = "123"
	testTriggerUrl               = "https://example.org"
	testTriggerPosition    int64 = 0
	testTriggerField             = "status"
	testTriggerOperator          = "is"
	testTriggerValue             = "open"
	testTriggerTime              = time.Now()
	testTriggerCatIdInt, _       = strconv.ParseInt(testTriggerCatId, 10, 64)
)

func TestGetApiModelFromTfModelTrigger(t *testing.T) {
	ctx := getTestContext(t)
	cases := []struct {
		testName string
		input    TriggerResourceModel
		expected zendesk.Trigger
	}{
		{
			testName: "should create trigger api object from tf object without custom fields in conditions or actions",
			input: TriggerResourceModel{
				ID:          types.Int64Value(123),
				Title:       types.StringValue(testTriggerTitle),
				Description: types.StringValue(testTriggerDesc),
				Active:      types.BoolValue(true),
				Position:    types.Int64Value(testTriggerPosition),
				CategoryID:  types.Int64Value(testTriggerCatIdInt),
				CreatedAt:   types.StringValue(testTriggerTime.UTC().String()),
				UpdatedAt:   types.StringValue(testTriggerTime.UTC().String()),
				URL:         types.StringValue(testTriggerUrl),
				Actions: []ActionResourceModel{
					{
						Field:         types.StringValue(testTriggerField),
						Value:         types.StringValue(testTriggerValue),
						Target:        types.StringNull(),
						CustomFieldID: types.Int64Null(),
					},
				},
				Conditions: ConditionsResourceModel{
					All: []ConditionResourceModel{
						{
							Field:         types.StringValue(testTriggerField),
							Operator:      types.StringValue(testTriggerOperator),
							Value:         types.StringValue(testTriggerValue),
							CustomFieldID: types.Int64Null(),
						},
					},
				},
			},
			expected: zendesk.Trigger{
				Title:       testTriggerTitle,
				Description: testTriggerDesc,
				Active:      true,
				Position:    testTriggerPosition,
				CategoryID:  testTriggerCatId,
				Actions: []zendesk.Action{
					{
						Field: testTriggerField,
						Value: zendesk.ParsedValue{Data: testValue},
					},
				},
				Conditions: zendesk.Conditions{
					All: []zendesk.Condition{
						{
							Field:    testTriggerField,
							Operator: testTriggerOperator,
							Value:    zendesk.ParsedValue{Data: testTriggerValue},
						},
					},
					Any: []zendesk.Condition{},
				},
			},
		},
	}

	for _, c := range cases {
		out, _ := c.input.GetApiModelFromTfModel(ctx)
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf(errorOutputMismatch, c.testName, out, c.expected)
		}
	}
}

func TestGetTfModelFromApiModelTrigger(t *testing.T) {
	ctx := getTestContext(t)
	cases := []struct {
		testName   string
		existingTf TriggerResourceModel
		input      zendesk.Trigger
		expected   TriggerResourceModel
	}{
		{
			testName:   "should create trigger tf object from api object without custom fields in conditions or actions",
			existingTf: TriggerResourceModel{},
			input: zendesk.Trigger{
				ID:          testTriggerId,
				Title:       testTriggerTitle,
				Description: testTriggerDesc,
				Active:      true,
				Position:    testTriggerPosition,
				CategoryID:  testTriggerCatId,
				Actions: []zendesk.Action{
					{
						Field: testTriggerField,
						Value: zendesk.ParsedValue{Data: testValue},
					},
				},
				Conditions: zendesk.Conditions{
					All: []zendesk.Condition{
						{
							Field:    testTriggerField,
							Operator: testTriggerOperator,
							Value:    zendesk.ParsedValue{Data: testTriggerValue},
						},
					},
				},
				CreatedAt: &testTriggerTime,
				UpdatedAt: &testTriggerTime,
				URL:       testTriggerUrl,
			},
			expected: TriggerResourceModel{
				ID:          types.Int64Value(123),
				Title:       types.StringValue(testTriggerTitle),
				Description: types.StringValue(testTriggerDesc),
				Active:      types.BoolValue(true),
				Position:    types.Int64Value(testTriggerPosition),
				CategoryID:  types.Int64Value(testTriggerCatIdInt),
				CreatedAt:   types.StringValue(testTriggerTime.UTC().String()),
				UpdatedAt:   types.StringValue(testTriggerTime.UTC().String()),
				URL:         types.StringValue(testTriggerUrl),
				Actions: []ActionResourceModel{
					{
						Field:               types.StringValue(testTriggerField),
						Value:               types.StringValue(testTriggerValue),
						NotificationSubject: types.StringNull(),
						Target:              types.StringNull(),
						CustomFieldID:       types.Int64Null(),
					},
				},
				Conditions: ConditionsResourceModel{
					All: []ConditionResourceModel{
						{
							Field:         types.StringValue(testTriggerField),
							Operator:      types.StringValue(testTriggerOperator),
							Value:         types.StringValue(testTriggerValue),
							Values:        types.ListNull(types.StringType),
							CustomFieldID: types.Int64Null(),
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		c.existingTf.GetTfModelFromApiModel(ctx, c.input)
		assert.Equal(t, c.expected, c.existingTf, c.testName)
	}
}

func TestGetTfActionsFromApi(t *testing.T) {

	cases := []struct {
		testName string
		input    []zendesk.Action
		expected []ActionResourceModel
	}{
		{
			testName: "should create trigger actions tf object from api object without custom fields in conditions or actions",
			input: []zendesk.Action{
				{
					Field: testTriggerField,
					Value: zendesk.ParsedValue{Data: testValue},
				},
			},
			expected: []ActionResourceModel{
				{
					Field:         types.StringValue(testTriggerField),
					Value:         types.StringValue(testTriggerValue),
					Target:        types.StringNull(),
					CustomFieldID: types.Int64Null(),
				},
			},
		},
	}

	for _, c := range cases {
		out, diags := getTfActionsFromApi(c.input)
		if diags.HasError() {
			t.Fatal(diags.Errors())
		}
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf(errorOutputMismatch, c.testName, out, c.expected)
		}
	}
}
