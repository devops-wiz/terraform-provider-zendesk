package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"testing"
)

func TestTicketFormSchema(t *testing.T) {
	t.Parallel()

	schemaRequest := resource.SchemaRequest{}
	schemaResponse := &resource.SchemaResponse{}

	NewTicketFormResource().Schema(t.Context(), schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	// Validate the schema
	diagnostics := schemaResponse.Schema.ValidateImplementation(t.Context())

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}

}
