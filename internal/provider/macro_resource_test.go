package provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

const dummyMacroResourceName = "zendesk_macro.test"

func TestAccMacro(t *testing.T) {
	var macro zendesk.Macro
	fullResourceName := fmt.Sprintf("test_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	t.Run("basic macro", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ResourceName: dummyMacroResourceName,
					ConfigFile:   config.TestStepFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					Check: resource.ComposeTestCheckFunc(
						testAccCheckMacroResourceExists(dummyMacroResourceName, &macro, t),
					), ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectNonEmptyPlan(),
							plancheck.ExpectResourceAction(dummyMacroResourceName, plancheck.ResourceActionCreate),
						},
					}, ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyMacroResourceName,
							tfjsonpath.New("title"),
							knownvalue.StringExact(fullResourceName),
						),
						statecheck.ExpectKnownValue(
							dummyMacroResourceName,
							tfjsonpath.New("actions"),
							knownvalue.ListSizeExact(2),
						),
					},
				},
				{
					ResourceName: dummyMacroResourceName,
					ConfigFile:   config.TestStepFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					Check: resource.ComposeTestCheckFunc(
						testAccCheckMacroResourceExists(dummyMacroResourceName, &macro, t),
					),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectNonEmptyPlan(),
							plancheck.ExpectResourceAction(dummyMacroResourceName, plancheck.ResourceActionUpdate),
						},
					}, ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyMacroResourceName,
							tfjsonpath.New("actions"),
							knownvalue.ListExact([]knownvalue.Check{
								knownvalue.MapExact(
									map[string]knownvalue.Check{
										"field":                knownvalue.StringExact("status"),
										"value":                knownvalue.StringExact("open"),
										"notification_subject": knownvalue.Null(),
										"slack_workspace":      knownvalue.Null(),
										"slack_channel":        knownvalue.Null(),
										"slack_title":          knownvalue.Null(),
										"target":               knownvalue.Null(),
										"content_type":         knownvalue.Null(),
										"custom_field_id":      knownvalue.Null(),
									},
								),
								knownvalue.MapExact(
									map[string]knownvalue.Check{
										"field":                knownvalue.StringExact("custom_field"),
										"value":                knownvalue.StringExact(fmt.Sprintf("%s_test_tag_macro_2", fullResourceName)),
										"notification_subject": knownvalue.Null(),
										"slack_workspace":      knownvalue.Null(),
										"slack_channel":        knownvalue.Null(),
										"slack_title":          knownvalue.Null(),
										"target":               knownvalue.Null(),
										"content_type":         knownvalue.Null(),
										"custom_field_id":      knownvalue.NotNull(),
									},
								),
							},
							),
						),
						statecheck.ExpectKnownValue(
							dummyMacroResourceName,
							tfjsonpath.New("title"),
							knownvalue.StringExact(fullResourceName),
						),
						statecheck.ExpectKnownValue(
							dummyMacroResourceName,
							tfjsonpath.New("actions"),
							knownvalue.ListSizeExact(2),
						),
					},
				},
			},
		})
	})

	t.Run("side conversation", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ResourceName: dummyMacroResourceName,
					ConfigFile:   config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					Check: resource.ComposeTestCheckFunc(
						testAccCheckMacroResourceExists(dummyMacroResourceName, &macro, t),
					), ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectNonEmptyPlan(),
							plancheck.ExpectResourceAction(dummyMacroResourceName, plancheck.ResourceActionCreate),
						},
					}, ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyMacroResourceName,
							tfjsonpath.New("title"),
							knownvalue.StringExact(fullResourceName),
						),
						statecheck.ExpectKnownValue(
							dummyMacroResourceName,
							tfjsonpath.New("actions"),
							knownvalue.ListSizeExact(3),
						),
					},
				},
			},
		})
	})

}

func testAccCheckMacroResourceExists(resourceName string, macro *zendesk.Macro, t *testing.T) resource.TestCheckFunc {
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

		resp, err := client.GetMacro(ctx, convertedId)

		if err != nil {
			return err
		}

		*macro = resp

		return nil

	}
}
