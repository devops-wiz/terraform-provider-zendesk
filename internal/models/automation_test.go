package models

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/JacobPotter/go-zendesk/zendesk"
)

func TestAutomationResourceModel_GetApiModelFromTfModel(t *testing.T) {
	ctx := t.Context()
	cases := []struct {
		testName string
		input    AutomationResourceModel
		expected zendesk.Automation
	}{
		{
			testName: "should create api model from tf model",
			input:    automationModelInput,
			expected: apiAutomationModelExpected,
		},
		{
			testName: "should create api model from tf model w/o position",
			input:    automationModelNoPositionInput,
			expected: apiAutomationModelNoPositionExpected,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			out, _ := c.input.GetApiModelFromTfModel(ctx)
			assert.Equal(t, out, c.expected)
		})
	}

}

func TestAutomationResourceModel_GetTfModelFromApiModel(t *testing.T) {

	ctx := t.Context()
	cases := []struct {
		testName   string
		existingTf AutomationResourceModel
		input      zendesk.Automation
		expected   AutomationResourceModel
	}{
		{
			testName:   "should get fill empty tf model from api model",
			existingTf: AutomationResourceModel{},
			input:      apiAutomationModelInput,
			expected:   automationModelExpected,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			c.existingTf.GetTfModelFromApiModel(ctx, c.input)
			assert.Equal(t, c.expected, c.existingTf)
		})
	}
}
