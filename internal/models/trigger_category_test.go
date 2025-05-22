package models

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	testCategoryId             = "123"
	testCategoryIdInt, _       = strconv.ParseInt(testCategoryId, 10, 64)
	testCategoryName           = "test category"
	testCategoryPosition int64 = 2
	testCategoryTime           = time.Now()
)

func TestGetApiModelFromTfModelTriggerCategory(t *testing.T) {
	ctx := t.Context()
	cases := []struct {
		testName string
		input    TriggerCategoryResourceModel
		expected zendesk.TriggerCategory
	}{
		{
			testName: "should create basic trigger category api object from tf schema with position",
			input: TriggerCategoryResourceModel{
				ID:        types.Int64Value(testCategoryIdInt),
				Name:      types.StringValue(testCategoryName),
				Position:  types.Int64Value(testCategoryPosition),
				CreatedAt: types.StringValue(testCategoryTime.UTC().String()),
				UpdatedAt: types.StringValue(testCategoryTime.UTC().String()),
			},
			expected: zendesk.TriggerCategory{
				Name:     testCategoryName,
				Position: testCategoryPosition,
			},
		},
		{
			testName: "should create basic trigger api object from tf schema without position",
			input: TriggerCategoryResourceModel{
				Name: types.StringValue(testCategoryName),
			},
			expected: zendesk.TriggerCategory{
				Name: testCategoryName,
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

func TestGetTfModelFromApiModelTriggerCategory(t *testing.T) {
	ctx := t.Context()
	cases := []struct {
		testName   string
		existingTf TriggerCategoryResourceModel
		input      zendesk.TriggerCategory
		expected   TriggerCategoryResourceModel
	}{
		{
			testName:   "should create basic trigger category tf object from api response",
			existingTf: TriggerCategoryResourceModel{},
			input: zendesk.TriggerCategory{
				ID:        testCategoryId,
				Name:      testCategoryName,
				Position:  testCategoryPosition,
				CreatedAt: testCategoryTime,
				UpdatedAt: testCategoryTime,
			},
			expected: TriggerCategoryResourceModel{
				ID:        types.Int64Value(testCategoryIdInt),
				Name:      types.StringValue(testCategoryName),
				Position:  types.Int64Value(testCategoryPosition),
				CreatedAt: types.StringValue(testCategoryTime.UTC().String()),
				UpdatedAt: types.StringValue(testCategoryTime.UTC().String()),
			},
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
