package provider

import (
	"fmt"
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

const dummyAutomationResourceName = "zendesk_automation.test"

func TestAccAutomation(t *testing.T) {
	t.Parallel()
	var automation zendesk.Automation
	testId := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	fullResourceName := fmt.Sprintf("test_acc_%s", testId)

	t.Run("basic automation", func(t *testing.T) {
		t.Parallel()
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
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(dummyAutomationResourceName, "title", fullResourceName),
						testAccCheckAutomationResourceExists(dummyAutomationResourceName, &automation, t),
					),
				},
			},
		},
		)
	})

	t.Run("should fail empty field", func(t *testing.T) {
		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile:  config.TestNameFile("main.tf"),
					ExpectError: regexp.MustCompile(".*field should not be empty*"),
				},
			},
		},
		)
	})

	t.Run("should fail invalid all condition", func(t *testing.T) {
		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile:  config.TestNameFile("main.tf"),
					ExpectError: regexp.MustCompile(".*Inappropriate value for attribute \"conditions\": attribute \"all\" is required.*"),
				},
			},
		},
		)
	})

	t.Run("should fail invalid condition field", func(t *testing.T) {
		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile:  config.TestNameFile("main.tf"),
					ExpectError: regexp.MustCompile(".*Error: Missing condition from \"All\" Conditions*"),
				},
			},
		},
		)
	})

}

func testAccCheckAutomationResourceExists(resourceName string, automation *zendesk.Automation, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ticket field ID is not set")
		}

		client := getZdTestClient()
		ctx := t.Context()

		convertedId, err := strconv.ParseInt(rs.Primary.ID, 10, 64)

		if err != nil {

			return fmt.Errorf("error converting")
		}

		tflog.SetField(ctx, "test_id", rs.Primary.ID)

		resp, err := client.GetAutomation(ctx, convertedId)

		if err != nil {
			return err
		}

		*automation = resp

		return nil

	}
}
