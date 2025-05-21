package models

import (
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"reflect"
	"testing"
)

func TestGetApiConditionsFromTf(t *testing.T) {
	cases := []struct {
		testName string
		input    ConditionsResourceModel
		expected zendesk.Conditions
	}{
		{
			testName: "should create trigger conditions api object from tf object without custom fields in conditions or actions",
			input: ConditionsResourceModel{
				All: []ConditionResourceModel{
					{
						Field:         types.StringValue(testField),
						Operator:      types.StringValue(testOperator),
						Value:         types.StringValue(testValue),
						CustomFieldID: types.Int64Null(),
					},
				},
			},
			expected: zendesk.Conditions{
				All: []zendesk.Condition{
					{
						Field:    testField,
						Operator: testOperator,
						Value:    zendesk.ParsedValue{Data: testValue},
					},
				},
				Any: []zendesk.Condition{},
			},
		},
	}

	for _, c := range cases {
		out, _ := getApiConditionsFromTf(t.Context(), c.input)
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf(errorOutputMismatch, c.testName, out, c.expected)
		}
	}
}
