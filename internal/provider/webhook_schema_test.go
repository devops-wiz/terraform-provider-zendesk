package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"testing"
)

func TestWebhookResourceSchema(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	schemaRequest := resource.SchemaRequest{}
	schemaResponse := &resource.SchemaResponse{}

	// Instantiate the resource.Resource and call its Schema method
	NewWebhookResource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	// Validate the schema
	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}
