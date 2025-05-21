package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestAccDynamicContent(t *testing.T) {
	fullResourceName := fmt.Sprintf("test_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	t.Run("should create basic dynamic content", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: map[string]config.Variable{
						"title": config.StringVariable(fullResourceName),
					},
				},
			},
		})
	})

	t.Run("should be able to change content", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestStepFile("main.tf"),
					ConfigVariables: map[string]config.Variable{
						"title": config.StringVariable(fullResourceName),
					},
				},
				{
					ConfigFile: config.TestStepFile("main.tf"),
					ConfigVariables: map[string]config.Variable{
						"title": config.StringVariable(fullResourceName),
					},
				},
			},
		})
	})
}
