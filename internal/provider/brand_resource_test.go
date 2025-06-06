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

var dummyBrandResourceName = "zendesk_brand.test"

func TestAccBrand(t *testing.T) {
	t.Parallel()
	fullResourceName := fmt.Sprintf("tf_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	testSubDomain := fmt.Sprintf("testsubdomain%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	t.Run("basic_brand", func(t *testing.T) {
		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: map[string]config.Variable{
						"title":     config.StringVariable(fullResourceName),
						"subdomain": config.StringVariable(testSubDomain),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyBrandResourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(fullResourceName),
						),
					},
				},
			},
		})
	})
}
