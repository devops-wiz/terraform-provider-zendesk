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

const dummySLAPolicyResourceName = "zendesk_sla_policy.test"

func TestAccSlaPolicy(t *testing.T) {
	t.Parallel()
	fullResourceName := fmt.Sprintf("test_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	t.Run("basic sla resource", func(t *testing.T) {
		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				testAccPreCheck(t)
			},
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummySLAPolicyResourceName,
							tfjsonpath.New("title"),
							knownvalue.StringExact(fullResourceName),
						),
					},
				},
			},
		})
	})

	t.Run("sla resource with settings", func(t *testing.T) {
		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				testAccPreCheck(t)
			},
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummySLAPolicyResourceName,
							tfjsonpath.New("metrics_settings").
								AtMapKey("first_reply_time").
								AtMapKey("fulfill_on_agent_internal_note"),
							knownvalue.Bool(true),
						),
					},
				},
			},
		})
	})

	t.Run("sla resource with position", func(t *testing.T) {
		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				testAccPreCheck(t)
			},
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestStepFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummySLAPolicyResourceName,
							tfjsonpath.New("metrics_settings").
								AtMapKey("first_reply_time").
								AtMapKey("fulfill_on_agent_internal_note"),
							knownvalue.Bool(true),
						),
					},
				},
				{
					ConfigFile: config.TestStepFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummySLAPolicyResourceName,
							tfjsonpath.New("metrics_settings").
								AtMapKey("first_reply_time").
								AtMapKey("fulfill_on_agent_internal_note"),
							knownvalue.Bool(true),
						),
					},
				},
			},
		})
	})
}
