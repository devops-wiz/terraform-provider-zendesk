package models

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"

	"github.com/JacobPotter/go-zendesk/zendesk"
)

func TestGetApiModelFromTfModelTicket(t *testing.T) {
	ctx := t.Context()
	cases := []struct {
		testName string
		input    TicketFieldResourceModel
		expected zendesk.TicketField
	}{
		{
			testName: "basic tf model should create basic api model",
			input:    testTicketFieldTf,
			expected: testTicketFieldApiExpected,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			out, diags := c.input.GetApiModelFromTfModel(ctx)
			if diags.HasError() {
				t.Errorf("got error diags: %v", diags.Errors())
			}
			assert.Equal(t, out, c.expected)
		})
	}
}

func TestGetTfModelFromApiModelTicket(t *testing.T) {
	ctx := t.Context()
	cases := []struct {
		testName   string
		existingTf TicketFieldResourceModel
		input      zendesk.TicketField
		expected   TicketFieldResourceModel
	}{
		{
			testName: "basic api model should create basic tf model",
			input:    testTicketFieldApiInput,
			expected: testTicketFieldTf,
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
