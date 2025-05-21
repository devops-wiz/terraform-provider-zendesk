package models

import (
	"context"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"reflect"
	"testing"
)

func TestWebhookResourceModel_GetApiModelFromTfModel(t *testing.T) {

	cases := []struct {
		testName string
		input    WebhookResourceModel
		expected zendesk.Webhook
	}{
		{testName: "should get api webhook from TF model",
			input: WebhookResourceModel{
				ID:             types.StringValue(testWebhookId),
				Name:           types.StringValue(testWebhookName),
				Description:    types.StringValue(testWebhookDesc),
				Endpoint:       types.StringValue(testWebhookEndpoint),
				HttpMethod:     types.StringValue(testWebhookHttpMethod),
				RequestFormat:  types.StringValue(testWebhookRequestFormat),
				Status:         types.StringValue(testWebhookStatus),
				CreatedBy:      types.StringValue(testWebhookUser),
				CreatedAt:      types.StringValue(testWebhookTime.UTC().String()),
				UpdatedBy:      types.StringValue(testWebhookUser),
				UpdatedAt:      types.StringValue(testWebhookTime.UTC().String()),
				Authentication: testAuthObj,
				Subscriptions:  types.ListNull(types.StringType),
				CustomHeaders:  testWebhookCustomHeadersMap,
			},
			expected: zendesk.Webhook{
				Name:          testWebhookName,
				Description:   testWebhookDesc,
				Endpoint:      testWebhookEndpoint,
				HTTPMethod:    testWebhookHttpMethod,
				RequestFormat: testWebhookRequestFormat,
				Status:        testWebhookStatus,
				Subscriptions: []string{},
				Authentication: &zendesk.WebhookAuthentication{
					Type:        testWebhookBasicAuth,
					AddPosition: testWebhookAddPosition,
					Data: zendesk.WebhookCredentials{
						Username: testWebhookUsername,
						Password: testWebhookPassword,
					},
				},
				CustomHeaders: testWebhookCustomHeadersApi,
			},
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

func TestGetApiWebhookAuthenticationFromTf(t *testing.T) {
	cases := []struct {
		testName string
		input    types.Object
		expected *zendesk.WebhookAuthentication
	}{
		{
			testName: "should get api auth model from tf for basic auth",
			input:    testAuthObjPass,
			expected: &zendesk.WebhookAuthentication{
				Type:        testWebhookBasicAuth,
				AddPosition: testWebhookAddPosition,
				Data: zendesk.WebhookCredentials{
					Username: testWebhookUsername,
					Password: testWebhookPassword,
				},
			},
		},
		{
			testName: "should get api auth model from tf for api key",
			input:    testAuthObjHeader,
			expected: &zendesk.WebhookAuthentication{
				Type:        testWebhookApiKey,
				AddPosition: testWebhookAddPosition,
				Data: zendesk.WebhookCredentials{
					HeaderName:  testWebhookApiHeaderKey,
					HeaderValue: testWebhookApiHeaderValue,
				},
			},
		},
		{
			testName: "should get api auth model from tf for bearer token",
			input:    testAuthObjToken,
			expected: &zendesk.WebhookAuthentication{
				Type:        testWebhookBearerToken,
				AddPosition: testWebhookAddPosition,
				Data: zendesk.WebhookCredentials{
					Token: testWebhookBearerTokenValue,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			out, diags := GetApiWebhookAuthenticationFromTf(context.Background(), c.input)
			if diags.HasError() {
				t.Fatalf("unexpected error: %s", diags.Errors())
			}
			if !reflect.DeepEqual(out, c.expected) {
				t.Fatalf(errorOutputMismatch, c.testName, out, c.expected)
			}
		})
	}
}

func TestWebhookResourceModel_GetTfModelFromApiModel(t *testing.T) {

	cases := []struct {
		testName string
		target   WebhookResourceModel
		input    zendesk.Webhook
		expected WebhookResourceModel
	}{
		{
			testName: "should populate tf resource with api data",
			target:   WebhookResourceModel{},
			input: zendesk.Webhook{
				ID:            testWebhookId,
				Name:          testWebhookName,
				Description:   testWebhookDesc,
				Endpoint:      testWebhookEndpoint,
				HTTPMethod:    testWebhookHttpMethod,
				RequestFormat: testWebhookRequestFormat,
				Status:        testWebhookStatus,
				Authentication: &zendesk.WebhookAuthentication{
					Type:        testWebhookBasicAuth,
					AddPosition: testWebhookAddPosition,
					Data: zendesk.WebhookCredentials{
						Username: testWebhookUsername,
						Password: testWebhookPassword,
					},
				},
				SigningSecret: &zendesk.WebhookSigningSecret{
					Secret:    testWebhookSigningSecret,
					Algorithm: testWebhookSigningAlgorithm,
				},
				CustomHeaders: testWebhookCustomHeadersApi,
				CreatedBy:     testWebhookUser,
				CreatedAt:     testWebhookTime,
				UpdatedBy:     testWebhookUser,
				UpdatedAt:     testWebhookTime,
				Subscriptions: nil,
			},
			expected: WebhookResourceModel{
				ID:             types.StringValue(testWebhookId),
				Name:           types.StringValue(testWebhookName),
				Description:    types.StringValue(testWebhookDesc),
				Endpoint:       types.StringValue(testWebhookEndpoint),
				HttpMethod:     types.StringValue(testWebhookHttpMethod),
				RequestFormat:  types.StringValue(testWebhookRequestFormat),
				Status:         types.StringValue(testWebhookStatus),
				CreatedBy:      types.StringValue(testWebhookUser),
				CreatedAt:      types.StringValue(testWebhookTime.UTC().String()),
				UpdatedBy:      types.StringValue(testWebhookUser),
				UpdatedAt:      types.StringValue(testWebhookTime.UTC().String()),
				Authentication: testAuthObj,
				CustomHeaders:  testWebhookCustomHeadersMap,
				Secret:         types.StringValue(testWebhookSigningSecret),
				Subscriptions:  types.ListValueMust(types.StringType, []attr.Value{}),
			},
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			_ = c.target.GetTfModelFromApiModel(context.Background(), c.input)
			if !reflect.DeepEqual(c.target, c.expected) {
				t.Fatalf(errorOutputMismatch, c.testName, c.target, c.expected)
			}
		})
	}
}

func TestGetTfWebhookAuthenticationFromApi(t *testing.T) {

	cases := []struct {
		testName string
		input    *zendesk.WebhookAuthentication
		expected types.Object
	}{
		{
			testName: "should get tf auth model from api basic auth",
			input: &zendesk.WebhookAuthentication{
				Type:        testWebhookBasicAuth,
				AddPosition: testWebhookAddPosition,
				Data: zendesk.WebhookCredentials{
					Username: testWebhookUsername,
					Password: testWebhookPassword,
				},
			},
			expected: testAuthObjPass,
		},
		{
			testName: "should get tf auth model from api for api key",
			input: &zendesk.WebhookAuthentication{
				Type:        testWebhookApiKey,
				AddPosition: testWebhookAddPosition,
				Data: zendesk.WebhookCredentials{
					HeaderName:  testWebhookApiHeaderKey,
					HeaderValue: testWebhookApiHeaderValue,
				},
			},
			expected: testAuthObjHeader,
		},
		{
			testName: "should get tf auth model from api for bearer token",
			input: &zendesk.WebhookAuthentication{
				Type:        testWebhookBearerToken,
				AddPosition: testWebhookAddPosition,
				Data: zendesk.WebhookCredentials{
					Token: testWebhookBearerTokenValue,
				},
			},
			expected: testAuthObjToken,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			out, diags := getTfWebhookAuthenticationFromApi(context.Background(), c.input)
			if diags.HasError() {
				t.Fatalf("unexpected error: %s", diags.Errors())
			}
			if !reflect.DeepEqual(out, c.expected) {
				t.Fatalf(errorOutputMismatch, c.testName, out, c.expected)
			}
		})
	}
}
