// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"github.com/JacobPotter/go-zendesk/credentialtypes"
	"os"
	"testing"

	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"zendesk": providerserver.NewProtocol6WithError(New("test")()),
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.

	if v := os.Getenv("ZENDESK_SUBDOMAIN"); v == "" {
		t.Fatal("ZENDESK_SUBDOMAIN must be set for acceptance tests")
	}

	if v := os.Getenv("ZENDESK_USERNAME"); v == "" {
		t.Fatal("ZENDESK_USERNAME must be set for acceptance tests")
	}

	if v := os.Getenv("ZENDESK_API_TOKEN"); v == "" {
		t.Fatal("ZENDESK_API_TOKEN must be set for acceptance tests")
	}

}

func getZdTestClient() (*zendesk.Client, error) {
	var subdomain = os.Getenv("ZENDESK_SUBDOMAIN")
	var username = os.Getenv("ZENDESK_USERNAME")
	var apiToken = os.Getenv("ZENDESK_API_TOKEN")

	client, err := zendesk.NewClient(nil)
	if err != nil {
		return nil, err
	}
	err = client.SetSubdomain(subdomain)
	if err != nil {
		return nil, err
	}
	client.SetCredential(credentialtypes.NewAPITokenCredential(username, apiToken))

	return client, nil

}
