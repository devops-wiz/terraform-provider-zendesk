package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/JacobPotter/go-zendesk/zendesk"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const dummyWebhookResourceName = "zendesk_webhook.test"

func TestAccWebhook(t *testing.T) {
	fullResourceName := fmt.Sprintf("test_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	t.Run("basic webhook", func(t *testing.T) {
		var webhook zendesk.Webhook
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"name": config.StringVariable(fullResourceName),
					},
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(dummyWebhookResourceName, "name", fullResourceName),
						testAccCheckWebhookResourceExists(dummyWebhookResourceName, &webhook, t),
					),
				},
			},
		})
	})

	t.Run("webhook with auth basic auth", func(t *testing.T) {
		var webhook zendesk.Webhook
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"name": config.StringVariable(fullResourceName),
					},
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(dummyWebhookResourceName, "name", fullResourceName),
						testAccCheckWebhookResourceExists(dummyWebhookResourceName, &webhook, t),
					),
				},
			},
		})
	})

}

func testAccCheckWebhookResourceExists(resourceName string, webhook *zendesk.Webhook, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("webhook ID is not set")
		}

		client := getZdTestClient()
		ctx := getTestContext(t)

		tflog.SetField(ctx, "test_id", rs.Primary.ID)

		resp, err := client.GetWebhook(ctx, rs.Primary.ID)

		if err != nil {
			return err
		}

		*webhook = resp

		return nil

	}
}

func TestWebhookResourceSchema(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	schemaRequest := fwresource.SchemaRequest{}
	schemaResponse := &fwresource.SchemaResponse{}

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
