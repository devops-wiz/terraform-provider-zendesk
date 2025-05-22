package models

import (
	"reflect"
	"testing"

	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestGetApiModelFromTfModelView(t *testing.T) {

	ctx := t.Context()

	testTitle := "Test View"
	testDesc := "Test View Desc"
	testActive := true
	var testPosition int64 = 5
	testConditionField := "status"
	testConditionOperator := "is"
	testConditionValue := "open"
	testAll := []ConditionResourceModel{
		{
			Field:         types.StringValue(testConditionField),
			Operator:      types.StringValue(testConditionOperator),
			Value:         types.StringValue(testConditionValue),
			CustomFieldID: types.Int64Unknown(),
		},
	}

	testGroupColumn := "status"
	testSortColumn := "assignee_id"
	testDirection := "asc"
	testOutputRaw := ViewOutputResourceModel{
		Columns: types.ListValueMust(types.StringType, []attr.Value{
			types.StringValue(testGroupColumn),
			types.StringValue(testSortColumn),
		}),
		GroupBy:    types.StringValue(testGroupColumn),
		GroupOrder: types.StringValue(testDirection),
		SortBy:     types.StringValue(testSortColumn),
		SortOrder:  types.StringValue(testDirection),
	}

	testOutput, _ := types.ObjectValueFrom(ctx, testOutputRaw.AttributeTypes(), testOutputRaw)

	cases := []struct {
		testName string
		input    ViewResourceModel
		expected zendesk.View
	}{
		{
			testName: "should generate create or update api body from tf schema with no position",
			input: ViewResourceModel{
				Title:       types.StringValue(testTitle),
				Description: types.StringValue(testDesc),
				Active:      types.BoolValue(testActive),
				Position:    types.Int64Null(),
				Conditions: &ConditionsResourceModel{
					All: testAll,
					Any: []ConditionResourceModel{},
				},
				Output:      testOutput,
				Restriction: types.ObjectNull(RestrictionResourceModel{}.AttributeTypes()),
			},
			expected: zendesk.View{
				Title:       testTitle,
				Description: testDesc,
				Active:      testActive,
				Conditions: zendesk.Conditions{
					All: []zendesk.Condition{{
						Field:    testConditionField,
						Operator: testConditionOperator,
						Value:    zendesk.ParsedValue{Data: testConditionValue},
					}},
					Any: []zendesk.Condition{},
				},
				All: []zendesk.Condition{{
					Field:    testConditionField,
					Operator: testConditionOperator,
					Value:    zendesk.ParsedValue{Data: testConditionValue},
				}},
				Any: []zendesk.Condition{},
				Output: zendesk.ViewOutput{
					Columns:    []string{testGroupColumn, testSortColumn},
					GroupBy:    testGroupColumn,
					GroupOrder: testDirection,
					SortBy:     testSortColumn,
					SortOrder:  testDirection,
				},
			},
		},
		{
			testName: "should generate create or update api body from tf schema",
			input: ViewResourceModel{
				Title:       types.StringValue(testTitle),
				Description: types.StringValue(testDesc),
				Active:      types.BoolValue(testActive),
				Position:    types.Int64Value(testPosition),
				Conditions: &ConditionsResourceModel{
					All: testAll,
					Any: []ConditionResourceModel{},
				},
				Output:      testOutput,
				Restriction: types.ObjectNull(RestrictionResourceModel{}.AttributeTypes()),
			},
			expected: zendesk.View{
				Title:       testTitle,
				Description: testDesc,
				Active:      testActive,
				Position:    testPosition,
				Conditions: zendesk.Conditions{
					All: []zendesk.Condition{{
						Field:    testConditionField,
						Operator: testConditionOperator,
						Value:    zendesk.ParsedValue{Data: testConditionValue},
					}},
					Any: []zendesk.Condition{},
				},
				All: []zendesk.Condition{{
					Field:    testConditionField,
					Operator: testConditionOperator,
					Value:    zendesk.ParsedValue{Data: testConditionValue},
				}},
				Any: []zendesk.Condition{},
				Output: zendesk.ViewOutput{
					Columns:    []string{testGroupColumn, testSortColumn},
					GroupBy:    testGroupColumn,
					GroupOrder: testDirection,
					SortBy:     testSortColumn,
					SortOrder:  testDirection,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			out, _ := c.input.GetApiModelFromTfModel(ctx)
			if !reflect.DeepEqual(out, c.expected) {
				t.Fatalf(errorOutputMismatch, c.testName, out, c.expected)
			}
		})
	}
}

func TestGetTfModelFromApiModelView(t *testing.T) {

	ctx := t.Context()
	cases := []struct {
		testName   string
		existingTf ViewResourceModel
		input      zendesk.View
		expected   ViewResourceModel
	}{
		{
			testName:   "should create tf view from api model",
			existingTf: ViewResourceModel{},
			input:      apiViewInput,
			expected:   testViewModelExpected,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			c.existingTf.GetTfModelFromApiModel(ctx, c.input)
			if !reflect.DeepEqual(c.existingTf, c.expected) {
				t.Fatalf(errorOutputMismatch, c.testName, c.existingTf, c.expected)
			}
		})
	}

}
