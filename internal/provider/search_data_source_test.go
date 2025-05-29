package provider

import (
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"regexp"
	"testing"
)

func TestAccSearch(t *testing.T) {
	t.Parallel()
	t.Run("user_search", func(t *testing.T) {
		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownOutputValue("test_user_email", knownvalue.NotNull()),
						statecheck.ExpectKnownValue(
							"data.zendesk_search.test",
							tfjsonpath.New("results").AtMapKey("users").AtSliceIndex(0).AtMapKey("email"),
							knownvalue.NotNull(),
						),
					},
				},
			},
		})
	})
	t.Run("org_search", func(t *testing.T) {
		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownOutputValue("test_org_id", knownvalue.NotNull()),
						statecheck.ExpectKnownValue(
							"data.zendesk_search.test",
							tfjsonpath.New("results").AtMapKey("organizations").AtSliceIndex(0).AtMapKey("id"),
							knownvalue.NotNull(),
						),
					},
				},
			},
		})
	})
	t.Run("other_search", func(t *testing.T) {
		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile:  config.TestNameFile("main.tf"),
					ExpectError: regexp.MustCompile(`.*Error: Unsupported result type*`),
				},
			},
		})
	})

}
