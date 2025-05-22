package models

import (
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"reflect"
	"testing"
)

func TestGroupResourceModel_GetApiModelFromTfModel(t *testing.T) {
	cases := []struct {
		testName string
		input    GroupResourceModel
		expected zendesk.Group
	}{
		{
			testName: "should get a api model from a tf resource",
			input: GroupResourceModel{
				Name:        types.StringValue(testTitle),
				Description: types.StringValue(testDescription),
				IsPublic:    types.BoolValue(false),
			},
			expected: zendesk.Group{
				Name:        testTitle,
				Description: testDescription,
				IsPublic:    false,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			out, _ := c.input.GetApiModelFromTfModel(t.Context())
			if !reflect.DeepEqual(out, c.expected) {
				t.Fatalf(errorOutputMismatch, c.testName, out, c.expected)
			}
		})
	}
}

func TestGroupResourceModel_GetTfModelFromApiModel(t *testing.T) {
	cases := []struct {
		testName string
		target   GroupResourceModel
		input    zendesk.Group
		expected GroupResourceModel
	}{
		{
			testName: "should generate TF resource model from api model",
			target:   GroupResourceModel{},
			input: zendesk.Group{
				ID:          testId,
				URL:         testUrl,
				Name:        testTitle,
				Default:     false,
				Deleted:     false,
				IsPublic:    false,
				Description: testDescription,
				CreatedAt:   testCreatedAt,
				UpdatedAt:   testUpdatedAt,
			},
			expected: GroupResourceModel{
				ID:          types.Int64Value(testId),
				URL:         types.StringValue(testUrl),
				Name:        types.StringValue(testTitle),
				Default:     types.BoolValue(false),
				Deleted:     types.BoolValue(false),
				IsPublic:    types.BoolValue(false),
				Description: types.StringValue(testDescription),
				CreatedAt:   types.StringValue(testCreatedAt.UTC().String()),
				UpdatedAt:   types.StringValue(testUpdatedAt.UTC().String()),
			},
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			c.target.GetTfModelFromApiModel(t.Context(), c.input)
			if !reflect.DeepEqual(c.target, c.expected) {
				t.Fatalf(errorOutputMismatch, c.testName, c.target, c.expected)
			}
		})
	}
}
