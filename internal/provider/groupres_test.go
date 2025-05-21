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

var dummyGroupResourceName = "zendesk_group.test"

func TestAccGroup(t *testing.T) {
	t.Parallel()

	t.Run("basic_group", func(t *testing.T) {
		fullResourceName := fmt.Sprintf("test_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: map[string]config.Variable{
						"title": config.StringVariable(fullResourceName),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyGroupResourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(fullResourceName),
						),
					},
				},
			},
		})
	})

	t.Run("change group is_public", func(t *testing.T) {
		fullResourceName := fmt.Sprintf("test_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("before.tf"),
					ConfigVariables: map[string]config.Variable{
						"title": config.StringVariable(fullResourceName),
					},
				}, {
					ConfigFile: config.TestNameFile("after.tf"),
					ConfigVariables: map[string]config.Variable{
						"title": config.StringVariable(fullResourceName),
					},
					ExpectError: regexp.MustCompile(`(\s\S|.)*Error: 422.*`),
				},
			},
		})
	})
}
