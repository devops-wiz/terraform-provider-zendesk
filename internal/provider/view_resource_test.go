package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"regexp"
	"strconv"
	"testing"

	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const dummyViewResourceName = "zendesk_view.test"

func TestAccView(t *testing.T) {

	fullResourceName := fmt.Sprintf(
		"test_acc_%s",
		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum),
	)

	t.Run("basic_view", func(t *testing.T) {
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
			},
		})
	})

	t.Run("basic_view_with_any", func(t *testing.T) {
		var view zendesk.View
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestStepFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					Check: resource.ComposeTestCheckFunc(
						testAccCheckViewResourceExists(dummyViewResourceName, &view, t),
					),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyViewResourceName,
							tfjsonpath.New("title"),
							knownvalue.StringExact(fullResourceName),
						),
						statecheck.ExpectKnownValue(
							dummyViewResourceName,
							tfjsonpath.New("conditions").AtMapKey("any").AtSliceIndex(1).AtMapKey("value"),
							knownvalue.StringExact("task"),
						),
					},
				},
				{
					ConfigFile: config.TestStepFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(dummyViewResourceName, plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyViewResourceName,
							tfjsonpath.New("conditions").AtMapKey("any").AtSliceIndex(1).AtMapKey("value"),
							knownvalue.StringExact("question"),
						),
					},
				},
			},
		})
	})

	t.Run("should fail invalid columns", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile:  config.TestNameFile("main.tf"),
					ExpectError: regexp.MustCompile(".*Error: Invalid Attribute Value Match*"),
				},
			},
		})
	})

}

func testAccCheckViewResourceExists(resourceName string, view *zendesk.View, t *testing.T) resource.TestCheckFunc {
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

		resp, err := client.GetView(ctx, convertedId)

		if err != nil {
			return err
		}

		*view = resp

		return nil

	}
}
