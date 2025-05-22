package provider

import (
	"fmt"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"

	"strconv"
	"testing"

	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const dummyTriggerResourceName = "zendesk_trigger.test"

func TestAccTrigger(t *testing.T) {
	t.Parallel()
	t.Run("basic trigger", func(t *testing.T) {
		t.Parallel()
		testId := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
		fullResourceName := fmt.Sprintf("test_acc_%s", testId)
		var trigger zendesk.Trigger
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
						resource.TestCheckResourceAttr(dummyTriggerResourceName, "title", fullResourceName),
						testAccCheckTriggerResourceExists(dummyTriggerResourceName, &trigger, t),
					),
				},
			},
		})
	})

	t.Run("trigger notification user", func(t *testing.T) {
		t.Parallel()
		testId := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
		fullResourceName := fmt.Sprintf("test_acc_%s", testId)
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
						resource.TestCheckResourceAttr(dummyTriggerResourceName, "title", fullResourceName),
					),
				},
			},
		})
	})

	t.Run("auto reply", func(t *testing.T) {
		t.Parallel()
		testId := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
		fullResourceName := fmt.Sprintf("test_acc_%s", testId)
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
						resource.TestCheckResourceAttr(dummyTriggerResourceName, "title", fullResourceName),
					),
				},
			},
		})
	})

}

func testAccCheckTriggerResourceExists(resourceName string, trigger *zendesk.Trigger, t *testing.T) resource.TestCheckFunc {
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

		resp, err := client.GetTrigger(ctx, convertedId)

		if err != nil {
			return err
		}

		*trigger = resp

		return nil

	}
}

func TestTriggerSchema(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	schemaRequest := fwresource.SchemaRequest{}
	schemaResponse := &fwresource.SchemaResponse{}

	NewTriggerResource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	// Validate the schema
	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}

}
