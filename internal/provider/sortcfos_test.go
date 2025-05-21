package provider

import (
	"github.com/devops-wiz/terraform-provider-zendesk/internal/models"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"testing"
)

func TestSortCustomFieldOptionsRun(t *testing.T) {
	t.Parallel()

	testCfoListInput, err := getTestCfoList(testCfosModelInput)
	if err != nil {
		t.Fatal(err)
	}

	testCfoListExpected, err := getTestCfoList(testCfosModelExpected)
	if err != nil {
		t.Fatal(err)
	}

	testCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"valid": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{testCfoListInput}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(testCfoListExpected),
			},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := function.RunResponse{
				Result: function.NewResultData(types.ListUnknown(types.ObjectType{AttrTypes: models.CustomFieldOptionResourceBase{}.AttributeTypes()})),
			}

			options := SortCustomFieldOptions{}
			options.Run(t.Context(), testCase.request, &got)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestSortCustomFieldOptionsValid(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					locals {
					  cfos=[
						{name:"T",value:"T"},
						{name:"H",value:"H"},
						{name:"D",value:"D"},
						{name:"L",value:"L"}
					  ]
					}
					
					output "test" {
					  value = provider::zendesk::sort_custom_field_options(local.cfos)
					}
`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownOutputValue("test", knownvalue.ListExact([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"name": knownvalue.StringExact("D"),
							}),
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"name": knownvalue.StringExact("H"),
							}),
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"name": knownvalue.StringExact("L"),
							}),
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"name": knownvalue.StringExact("T"),
							}),
						})),
					},
				},
			},
		},
	})
}
