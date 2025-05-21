package models

import (
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetApiActionsFromTf_function(t *testing.T) {

	cases := []struct {
		testName string
		input    []ActionResourceModel
		expected []zendesk.Action
	}{
		{
			testName: "should create trigger actions api object from tf object without custom fields in conditions or actions",
			input:    testActionModels,
			expected: testApiActionModels,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			out, diags := getApiActionsFromTf(c.input)
			if diags.HasError() {
				t.Fatalf("diagnostics: %+v", diags)
			}
			assert.Equal(t, c.expected, out)
		})
	}
}
