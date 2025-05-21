package models

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"testing"
)

func getTestContext(t *testing.T) context.Context {
	t.Helper()

	ctx := context.Background()

	return ctx
}

func diagnosticErrorHelper(t *testing.T, diags diag.Diagnostics, errorMsgHeader string) {
	t.Helper()
	t.Log(errorMsgHeader)
	for i, diagnostic := range diags.Errors() {
		t.Logf("Error %d:\nSummary: %s\nDetails: %s", i+1, diagnostic.Summary(), diagnostic.Detail())
	}
	t.Error("Unable to continue test due to previous errors")
}
