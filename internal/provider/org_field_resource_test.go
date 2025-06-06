package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"regexp"
	"testing"
)

var dummyOrganizationFieldResourceName = "zendesk_organization_field.test"

func TestAccOrgField(t *testing.T) {
	t.Parallel()
	t.Run("basic organization field", func(t *testing.T) {
		t.Parallel()
		testId := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
		fullResourceName := fmt.Sprintf("tf_acc_%s", testId)
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title":   config.StringVariable(fullResourceName),
						"test_id": config.StringVariable(testId),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyOrganizationFieldResourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(fullResourceName),
						),
					},
				},
			},
		})
	})
	t.Run("organization field dropdown", func(t *testing.T) {
		t.Parallel()
		testId := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
		fullResourceName := fmt.Sprintf("tf_acc_%s", testId)
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title":   config.StringVariable(fullResourceName),
						"test_id": config.StringVariable(testId),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyOrganizationFieldResourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(fullResourceName),
						),
					},
				},
			},
		})
	})
	t.Run("organization field dropdown add options", func(t *testing.T) {
		t.Parallel()
		testId := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
		fullResourceName := fmt.Sprintf("tf_acc_%s", testId)
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title":   config.StringVariable(fullResourceName),
						"test_id": config.StringVariable(testId),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyOrganizationFieldResourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(fullResourceName),
						),
					},
				},
				{
					ConfigFile: config.TestNameFile("optionsAdded.tf"),
					ConfigVariables: config.Variables{
						"title":   config.StringVariable(fullResourceName),
						"test_id": config.StringVariable(testId),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyOrganizationFieldResourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(fullResourceName),
						),
					},
				},
			},
		})
	})
	t.Run("organization field dropdown missing", func(t *testing.T) {
		t.Parallel()
		testId := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
		fullResourceName := fmt.Sprintf("tf_acc_%s", testId)
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title":   config.StringVariable(fullResourceName),
						"test_id": config.StringVariable(testId),
					},
					ExpectError: regexp.MustCompile(`.*Could not create Organization Field, unexpected error: 422*`),
				},
			},
		})
	})

	t.Run("should disable field", func(t *testing.T) {
		t.Parallel()
		testId := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
		fullResourceName := fmt.Sprintf("tf_acc_%s", testId)
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title":   config.StringVariable(fullResourceName),
						"test_id": config.StringVariable(testId),
					},
				},
				{
					ConfigFile: config.TestNameFile("disabled.tf"),
					ConfigVariables: config.Variables{
						"title":   config.StringVariable(fullResourceName),
						"test_id": config.StringVariable(testId),
					},
				},
			},
		})
	})
}
