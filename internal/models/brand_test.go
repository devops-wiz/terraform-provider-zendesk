package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"reflect"
	"testing"
)

func TestBrandResource_GetApiModelFromTfModel(t *testing.T) {
	cases := []struct {
		testName string
		input    BrandResourceModel
		expected zendesk.Brand
	}{
		{
			testName: "should generate api model from tf resource",
			input: BrandResourceModel{
				Name:              types.StringValue(testTitle),
				BrandURL:          types.StringValue(testUrl),
				HasHelpCenter:     types.BoolValue(true),
				HelpCenterState:   types.StringValue("test"),
				Active:            types.BoolValue(true),
				TicketFormIDs:     types.ListValueMust(types.Int64Type, []attr.Value{types.Int64Value(123)}),
				Subdomain:         types.StringValue("dynatrace"),
				HostMapping:       types.StringValue("host-mapping"),
				SignatureTemplate: types.StringValue("signature-template"),
			},
			expected: zendesk.Brand{
				Name:              testTitle,
				BrandURL:          testUrl,
				HasHelpCenter:     true,
				HelpCenterState:   "test",
				Active:            true,
				TicketFormIDs:     []int64{123},
				Subdomain:         "dynatrace",
				HostMapping:       "host-mapping",
				SignatureTemplate: "signature-template",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			out, diags := c.input.GetApiModelFromTfModel(context.Background())
			if diags.HasError() {
				t.Errorf("%s: got error %s", c.testName, diags.Errors())
			}
			if !reflect.DeepEqual(out, c.expected) {
				t.Fatalf(errorOutputMismatch, c.testName, out, c.expected)
			}
		})
	}
}

func TestBrandResource_GetTfModelFromApiModel(t *testing.T) {
	cases := []struct {
		testName string
		target   BrandResourceModel
		input    zendesk.Brand
		expected BrandResourceModel
	}{
		{
			testName: "should generate tf resource from api model",
			target:   BrandResourceModel{},
			input: zendesk.Brand{
				ID:                testId,
				URL:               testUrl,
				Name:              testTitle,
				BrandURL:          testUrl,
				HasHelpCenter:     true,
				HelpCenterState:   "test",
				Active:            true,
				Default:           false,
				TicketFormIDs:     []int64{123},
				Subdomain:         "dynatrace",
				HostMapping:       "host-mapping",
				SignatureTemplate: "signature-template",
				CreatedAt:         testCreatedAt,
				UpdatedAt:         testUpdatedAt,
			},
			expected: BrandResourceModel{
				ID:                types.Int64Value(testId),
				URL:               types.StringValue(testUrl),
				Name:              types.StringValue(testTitle),
				BrandURL:          types.StringValue(testUrl),
				HasHelpCenter:     types.BoolValue(true),
				HelpCenterState:   types.StringValue("test"),
				Active:            types.BoolValue(true),
				Default:           types.BoolValue(false),
				TicketFormIDs:     types.ListValueMust(types.Int64Type, []attr.Value{types.Int64Value(123)}),
				Subdomain:         types.StringValue("dynatrace"),
				HostMapping:       types.StringValue("host-mapping"),
				SignatureTemplate: types.StringValue("signature-template"),
				CreatedAt:         types.StringValue(testCreatedAt.UTC().String()),
				UpdatedAt:         types.StringValue(testUpdatedAt.UTC().String()),
			},
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			diags := c.target.GetTfModelFromApiModel(context.Background(), c.input)
			if diags.HasError() {
				t.Errorf("%s: got error %s", c.testName, diags.Errors())
			}
			if !reflect.DeepEqual(c.target, c.expected) {
				t.Fatalf(errorOutputMismatch, c.testName, c.target, c.expected)
			}
		})
	}

}
