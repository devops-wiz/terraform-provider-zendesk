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

var dummyScheduleResourceName = "zendesk_schedule.test"

func TestAccSchedule(t *testing.T) {
	t.Parallel()
	fullResourceName := fmt.Sprintf("tf_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	t.Run("basic schedule one day", func(t *testing.T) {
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
							dummyScheduleResourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(fullResourceName),
						),
						statecheck.ExpectKnownValue(
							dummyScheduleResourceName,
							tfjsonpath.New("intervals").AtMapKey("sunday").AtMapKey("start_time"),
							knownvalue.Int64Exact(5),
						),
					},
				},
			},
		})
	})

	t.Run("basic schedule all days", func(t *testing.T) {
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
							dummyScheduleResourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(fullResourceName),
						),
						statecheck.ExpectKnownValue(
							dummyScheduleResourceName,
							tfjsonpath.New("intervals").AtMapKey("sunday").AtMapKey("start_time"),
							knownvalue.Int64Exact(5),
						),
					},
				},
			},
		})
	})

	t.Run("update schedule", func(t *testing.T) {
		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestStepFile("main.tf"),
					ConfigVariables: map[string]config.Variable{
						"title": config.StringVariable(fullResourceName),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyScheduleResourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(fullResourceName),
						),
						statecheck.ExpectKnownValue(
							dummyScheduleResourceName,
							tfjsonpath.New("intervals").AtMapKey("tuesday").AtMapKey("start_time"),
							knownvalue.Int64Exact(5),
						),
					},
				},
				{
					ConfigFile: config.TestStepFile("main.tf"),
					ConfigVariables: map[string]config.Variable{
						"title": config.StringVariable(fullResourceName),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							dummyScheduleResourceName,
							tfjsonpath.New("name"),
							knownvalue.StringExact(fullResourceName),
						),
						statecheck.ExpectKnownValue(
							dummyScheduleResourceName,
							tfjsonpath.New("intervals").AtMapKey("tuesday").AtMapKey("start_time"),
							knownvalue.Int64Exact(4),
						),
					},
				},
			},
		})
	})
}
