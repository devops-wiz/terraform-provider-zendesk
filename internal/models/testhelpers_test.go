package models

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"testing"
)

func diagnosticErrorHelper(t *testing.T, diags diag.Diagnostics, errorMsgHeader string) {
	t.Helper()
	t.Log(errorMsgHeader)
	for i, diagnostic := range diags.Errors() {
		t.Logf("Error %d:\nSummary: %s\nDetails: %s", i+1, diagnostic.Summary(), diagnostic.Detail())
	}
	t.Error("Unable to continue test due to previous errors")
}
