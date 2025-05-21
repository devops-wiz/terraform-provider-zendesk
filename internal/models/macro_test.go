package models

import (
	"github.com/JacobPotter/go-zendesk/zendesk"
	"reflect"
	"testing"
)

func TestGetApiModelFromTfModelMacro(t *testing.T) {
	ctx := getTestContext(t)
	cases := []struct {
		testName string
		input    MacroResourceModel
		expected zendesk.Macro
	}{
		{
			testName: "should create api model from tf model with position and w/o restriction",
			input:    testMacroResourceModelWithPositionWithoutRestriction,
			expected: testMacroWithPositionWithoutRestriction,
		},
		{
			testName: "should create api model from tf model w/o position or restriction",
			input:    testMacroResourceModelWithoutPositionRestriction,
			expected: testMacroWithoutPositionRestriction,
		},
		{
			testName: "should create api model from tf model with restriction and without position",
			input:    testMacroResourceModelWithRestrictionWithoutPosition,
			expected: testMacroWithoutPositionWithRestriction,
		},
		{
			testName: "should create api model from tf model with restriction and with position",
			input:    testMacroResourceModelWithRestrictionWithPosition,
			expected: testMacroWithPositionWithRestriction,
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

func TestGetTfModelFromApiModelMacro(t *testing.T) {
	t.Skip()
	ctx := getTestContext(t)
	cases := []struct {
		testName   string
		existingTf MacroResourceModel
		input      zendesk.Macro
		expected   MacroResourceModel
	}{
		{
			testName:   "should get fill empty tf model from api model",
			existingTf: MacroResourceModel{},
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
