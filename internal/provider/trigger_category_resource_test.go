package provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

const dummyTriggerCatgegoryResourceName = "zendesk_trigger_category.test"

func TestAccTriggerCategoryBasic(t *testing.T) {

	var triggerCategory zendesk.TriggerCategory

	fullResourceName := fmt.Sprintf("test_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTriggerCatConfig(fullResourceName),
				Check:  resource.ComposeTestCheckFunc(testAccCheckTriggerCategoryResourceExists(dummyTriggerCatgegoryResourceName, &triggerCategory, t)),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dummyTriggerCatgegoryResourceName, tfjsonpath.New("name"), knownvalue.StringExact(fullResourceName)),
				},
			},
		},
	})
}

func testAccCheckTriggerCategoryResourceExists(resourceName string, triggerCat *zendesk.TriggerCategory, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]

		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Ticket field ID is not set")
		}

		client := getZdTestClient()
		ctx := getTestContext(t)

		convertedId, err := strconv.ParseInt(rs.Primary.ID, 10, 64)

		if err != nil {

			return fmt.Errorf("Error converting")
		}

		tflog.SetField(ctx, "test_id", rs.Primary.ID)

		resp, err := client.GetTriggerCategory(ctx, convertedId)

		if err != nil {
			return err
		}

		*triggerCat = resp

		return nil

	}
}

func testAccTriggerCatConfig(name string) string {
	return fmt.Sprintf(`
resource "zendesk_trigger_category" "test" {
	name = "%s"
}
`,
		name,
	)
}
