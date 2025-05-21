package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"reflect"
	"testing"
)

func TestSLAPolicyResourceModel_GetApiModelFromTfModel(t *testing.T) {
	cases := []struct {
		testName string
		input    SLAPolicyResourceModel
		expected zendesk.SLAPolicy
	}{
		{
			testName: "should generate api model from tf resource",
			input:    testSlaPolicyModelInput,
			expected: testSlaPolicyExpected,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			out, _ := c.input.GetApiModelFromTfModel(context.Background())
			if !reflect.DeepEqual(out, c.expected) {
				t.Fatalf(errorOutputMismatch, c.testName, out, c.expected)
			}
		})
	}
}

func TestSLAPolicyResourceModel_GetTfModelFromApiModel(t *testing.T) {
	cases := []struct {
		testName string
		target   SLAPolicyResourceModel
		input    zendesk.SLAPolicy
		expected SLAPolicyResourceModel
	}{
		{
			testName: "should generate tf model from api response",
			target:   SLAPolicyResourceModel{},
			input:    testSlaPolicyInput,
			expected: testSlaPolicyModelExpected,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			diags := c.target.GetTfModelFromApiModel(context.Background(), c.input)
			if diags.HasError() {
				t.Fatalf("errors: %s", diags)
			}
			if !reflect.DeepEqual(c.target, c.expected) {
				t.Fatalf(errorOutputMismatch, c.testName, c.target, c.expected)
			}
		})
	}
}
