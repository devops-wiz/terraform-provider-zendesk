package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"testing"
)

var dummyTicketFormResourceName = "zendesk_ticket_form.test"

func TestAccTicketForm(t *testing.T) {
	t.Parallel()
	t.Run("basic ticket form", func(t *testing.T) {
		t.Parallel()
		fullResourceName := fmt.Sprintf("tf_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
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
							dummyTicketFormResourceName,
							tfjsonpath.New("form_name"),
							knownvalue.StringExact(fullResourceName)),
						statecheck.ExpectKnownValue(
							dummyTicketFormResourceName,
							tfjsonpath.New("end_user_display_name"),
							knownvalue.Null(),
						),
					},
				},
			},
		})
	})

	t.Run("basic ticket form change display name", func(t *testing.T) {
		t.Parallel()
		fullResourceName := fmt.Sprintf("tf_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
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
							dummyTicketFormResourceName,
							tfjsonpath.New("form_name"),
							knownvalue.StringExact(fullResourceName)),
						statecheck.ExpectKnownValue(
							dummyTicketFormResourceName,
							tfjsonpath.New("end_user_display_name"),
							knownvalue.NotNull(),
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
							dummyTicketFormResourceName,
							tfjsonpath.New("form_name"),
							knownvalue.StringExact(fullResourceName)),
						statecheck.ExpectKnownValue(
							dummyTicketFormResourceName,
							tfjsonpath.New("end_user_display_name"),
							knownvalue.Null(),
						),
					},
				},
			},
		})
	})

	t.Run("basic ticket form change ticket fields", func(t *testing.T) {
		t.Parallel()
		fullResourceName := fmt.Sprintf("tf_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
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
							dummyTicketFormResourceName,
							tfjsonpath.New("form_name"),
							knownvalue.StringExact(fullResourceName)),
					},
				},
				{
					ConfigFile: config.TestStepFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyTicketFormResourceName,
							tfjsonpath.New("form_name"),
							knownvalue.StringExact(fullResourceName)),
					},
				},
			},
		})
	})

	t.Run("ticket form with agent conditions", func(t *testing.T) {
		t.Parallel()
		fullResourceName := fmt.Sprintf("tf_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
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
							dummyTicketFormResourceName,
							tfjsonpath.New("form_name"),
							knownvalue.StringExact(fullResourceName)),
					},
				},
			},
		})
	})

	t.Run("ticket form with agent and end user conditions", func(t *testing.T) {
		t.Parallel()
		fullResourceName := fmt.Sprintf("tf_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
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
							dummyTicketFormResourceName,
							tfjsonpath.New("form_name"),
							knownvalue.StringExact(fullResourceName)),
					},
				},
			},
		})
	})

	t.Run("ticket form with agent and end user conditions changed", func(t *testing.T) {
		t.Parallel()
		fullResourceName := fmt.Sprintf("tf_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
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
							dummyTicketFormResourceName,
							tfjsonpath.New("form_name"),
							knownvalue.StringExact(fullResourceName)),
					},
				},
				{
					ConfigFile: config.TestStepFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyTicketFormResourceName,
							tfjsonpath.New("form_name"),
							knownvalue.StringExact(fullResourceName)),
					},
				},
			},
		})
	})
}
