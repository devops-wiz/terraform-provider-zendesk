package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"regexp"
	"strconv"
	"testing"

	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const dummyTicketFieldResourceName = "zendesk_ticket_field.test"

func TestAccTicketField(t *testing.T) {
	fullResourceName := fmt.Sprintf("test_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	t.Run("basic ticket field", func(t *testing.T) {
		var ticketField zendesk.TicketField
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					Check: resource.ComposeTestCheckFunc(
						testAccCheckTicketFieldResourceExists(dummyTicketFieldResourceName, &ticketField, t),
						testAccCheckTicketFieldAttributes(&ticketField),
					),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyTicketFieldResourceName,
							tfjsonpath.New("custom_field_options"),
							knownvalue.ListSizeExact(2),
						),
						statecheck.ExpectKnownValue(
							dummyTicketFieldResourceName,
							tfjsonpath.New("title"),
							knownvalue.StringExact(fullResourceName),
						),
						statecheck.ExpectKnownValue(
							dummyTicketFieldResourceName,
							tfjsonpath.New("required"),
							knownvalue.Bool(false),
						),
					},
				},
			},
		})
	})

	t.Run("visible in portal and required in portal invalid", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ExpectError: regexp.MustCompile(`.*Invalid attribute combination.*`),
				},
			},
		})
	})

	t.Run("editable in portal invalid", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ExpectError: regexp.MustCompile(`.*Invalid attribute combination.*`),
				},
			},
		})
	})

	t.Run("should fail invalid field", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ExpectError: regexp.MustCompile(`.*Attribute "custom_field_options" must be specified when "type" is specified*`),
				},
			},
		})
	})

	t.Run("should destroy then create when type changed", func(t *testing.T) {
		var ticketField zendesk.TicketField
		type1 := "tagger"
		type2 := "multiselect"

		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title":     config.StringVariable(fullResourceName),
						"fieldType": config.StringVariable(type1),
					},
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(dummyTicketFieldResourceName, "title", fullResourceName),
						testAccCheckTicketFieldResourceExists(dummyTicketFieldResourceName, &ticketField, t),
						testAccCheckTicketFieldAttributes(&ticketField),
					),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyTicketFieldResourceName,
							tfjsonpath.New("custom_field_options"),
							knownvalue.ListSizeExact(2),
						),
					},
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectNonEmptyPlan(),
							plancheck.ExpectResourceAction(dummyTicketFieldResourceName, plancheck.ResourceActionCreate),
						},
					},
				},
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title":     config.StringVariable(fullResourceName),
						"fieldType": config.StringVariable(type2),
					},
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(dummyTicketFieldResourceName, "title", fullResourceName),
						testAccCheckTicketFieldResourceExists(dummyTicketFieldResourceName, &ticketField, t),
						testAccCheckTicketFieldAttributes(&ticketField),
					),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyTicketFieldResourceName,
							tfjsonpath.New("custom_field_options"),
							knownvalue.ListSizeExact(2),
						),
					},
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectNonEmptyPlan(),
							plancheck.ExpectResourceAction(dummyTicketFieldResourceName, plancheck.ResourceActionDestroyBeforeCreate),
						},
					},
				},
			},
		})
	})

	t.Run("should import system field", func(t *testing.T) {
		subjectField := "zendesk_ticket_field.subject"
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile:    config.TestNameFile("import.tf"),
					ResourceName:  subjectField,
					ImportState:   true,
					ImportStateId: strconv.FormatInt(1900002750725, 10),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							subjectField,
							tfjsonpath.New("title"),
							knownvalue.StringExact("Subject"),
						),
					},
				},
			},
		})
	})

	t.Run("should import custom field", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile:    config.TestNameFile("import.tf"),
					ResourceName:  "zendesk_ticket_field.affected_dynatrace_component",
					ImportState:   true,
					ImportStateId: strconv.FormatInt(16590742000407, 10),
				},
			},
		})
	})

	t.Run("change description", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
				},
				{
					ConfigFile: config.TestNameFile("changed.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
				},
			},
		})
	})

	t.Run("change required option", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
				},
				{
					ConfigFile: config.TestNameFile("changed.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
				},
			},
		})
	})

}

func testAccCheckTicketFieldResourceExists(resourceName string, ticketField *zendesk.TicketField, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ticket field ID is not set")
		}

		client := getZdTestClient()
		ctx := getTestContext(t)

		convertedId, err := strconv.ParseInt(rs.Primary.ID, 10, 64)

		if err != nil {

			return fmt.Errorf("error converting")
		}

		tflog.SetField(ctx, "test_id", rs.Primary.ID)

		resp, err := client.GetTicketField(ctx, convertedId)

		if err != nil {
			return err
		}

		*ticketField = resp

		return nil

	}
}

func testAccCheckTicketFieldAttributes(ticketField *zendesk.TicketField) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if ticketField.Active != true {
			return fmt.Errorf("ticket field is not active")
		}
		return nil
	}
}
