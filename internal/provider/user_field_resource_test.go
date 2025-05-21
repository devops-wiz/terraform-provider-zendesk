package provider

import (
	"context"
	"fmt"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"regexp"
	"testing"
)

var dummyUserFieldResourceName = "zendesk_user_field.test"

func TestAccUserField(t *testing.T) {

	t.Run("basic user field", func(t *testing.T) {
		fullResourceName := fmt.Sprintf("test_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyUserFieldResourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(fullResourceName),
						),
					},
				},
			},
		})
	})
	t.Run("basic user field change desc", func(t *testing.T) {
		fullResourceName := fmt.Sprintf("test_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestStepFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyUserFieldResourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(fullResourceName),
						),
					},
				},
				{
					ConfigFile: config.TestStepFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyUserFieldResourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(fullResourceName),
						),
					},
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(dummyUserFieldResourceName, plancheck.ResourceActionUpdate),
						},
					},
				},
			},
		})
	})
	t.Run("user field dropdown", func(t *testing.T) {
		fullResourceName := fmt.Sprintf("test_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyUserFieldResourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(fullResourceName),
						),
					},
				},
			},
		})
	})
	t.Run("user field dropdown missing", func(t *testing.T) {
		fullResourceName := fmt.Sprintf("test_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ExpectError: regexp.MustCompile(`.*Error: 422:*`),
				},
			},
		})
	})
}

func TestUserFieldSchema(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	schemaRequest := fwresource.SchemaRequest{}
	schemaResponse := &fwresource.SchemaResponse{}

	NewUserFieldResource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	// Validate the schema
	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}

}
