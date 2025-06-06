package provider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/JacobPotter/go-zendesk/zendesk"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"log"
	"regexp"
	"strings"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const resourceType = "zendesk_ticket_field"

var rName = fmt.Sprintf("%s.test", resourceType)

var baseTestField = zendesk.TicketField{
	Type: "text",
}

func init() {
	resource.AddTestSweepers(resourceType, &resource.Sweeper{
		Name: resourceType,
		F: func(_ string) error {
			client, err := getZdTestClient()
			if err != nil {
				return err
			}

			fields, _, err := client.GetTicketFieldsOBP(context.Background(), &zendesk.OBPOptions{
				PageOptions: zendesk.PageOptions{
					PerPage: 100,
				},
			})

			if err != nil {
				return err
			}

			for _, field := range fields {
				if strings.HasPrefix(field.Title, "tf_acc_") {
					err = client.DeleteTicketField(context.Background(), field.ID)
					if err != nil {
						return err
					}
					log.Printf("Deleted ticket field %s", field.Title)
				}
			}

			return nil
		},
	})
}

func TestAccTicketField_basic(t *testing.T) {
	t.Parallel()

	t.Run("basic ticket field", func(t *testing.T) {
		fullResourceName := fmt.Sprintf("tf_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							rName,
							tfjsonpath.New("custom_field_options"),
							knownvalue.ListSizeExact(2),
						),
						statecheck.ExpectKnownValue(
							rName,
							tfjsonpath.New("title"),
							knownvalue.StringExact(fullResourceName),
						),
						statecheck.ExpectKnownValue(
							rName,
							tfjsonpath.New("required"),
							knownvalue.Bool(false),
						),
					},
				},
			},
		})
	})

	t.Run("visible in portal and required in portal invalid", func(t *testing.T) {
		fullResourceName := fmt.Sprintf("tf_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ExpectError: regexp.MustCompile(`.*Invalid attribute combination.*`),
				},
			},
		})
	})

	t.Run("editable in portal invalid", func(t *testing.T) {
		fullResourceName := fmt.Sprintf("tf_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ExpectError: regexp.MustCompile(`.*Invalid attribute combination.*`),
				},
			},
		})
	})

	t.Run("should fail invalid field", func(t *testing.T) {
		fullResourceName := fmt.Sprintf("tf_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
					ExpectError: regexp.MustCompile(`.*Attribute "custom_field_options" must be specified when "type" is specified*`),
				},
			},
		})
	})

	t.Run("should destroy then create when type changed", func(t *testing.T) {
		fullResourceName := fmt.Sprintf("tf_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

		t.Parallel()
		type1 := "tagger"
		type2 := "multiselect"

		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title":     config.StringVariable(fullResourceName),
						"fieldType": config.StringVariable(type1),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							rName,
							tfjsonpath.New("custom_field_options"),
							knownvalue.ListSizeExact(2),
						),
					},
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectNonEmptyPlan(),
							plancheck.ExpectResourceAction(rName, plancheck.ResourceActionCreate),
						},
					},
				},
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title":     config.StringVariable(fullResourceName),
						"fieldType": config.StringVariable(type2),
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							rName,
							tfjsonpath.New("custom_field_options"),
							knownvalue.ListSizeExact(2),
						),
					},
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectNonEmptyPlan(),
							plancheck.ExpectResourceAction(rName, plancheck.ResourceActionDestroyBeforeCreate),
						},
					},
				},
			},
		})
	})

	t.Run("change description", func(t *testing.T) {
		fullResourceName := fmt.Sprintf("tf_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
				},
				{
					ConfigFile: config.TestNameFile("changed.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
				},
			},
		})
	})

	t.Run("change required option", func(t *testing.T) {
		fullResourceName := fmt.Sprintf("tf_acc_%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

		t.Parallel()
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ConfigFile: config.TestNameFile("main.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
				},
				{
					ConfigFile: config.TestNameFile("changed.tf"),
					ConfigVariables: config.Variables{
						"title": config.StringVariable(fullResourceName),
					},
				},
			},
		})
	})

}

func TestAccTicketField_update(t *testing.T) {
	t.Parallel()

	t.Run("ticket field basic", func(t *testing.T) {
		resourceName := acctest.RandomWithPrefix("tf_acc_ticket_field")
		t.Parallel()
		testField := baseTestField
		testField.Title = resourceName
		testFieldChanged := testField
		testFieldChanged.Description = "changed"
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccTicketField(t, testField),
				},
				{
					Config: testAccTicketField(t, testFieldChanged),
				},
				{
					ImportState:     true,
					ImportStateKind: resource.ImportBlockWithID,
					ResourceName:    rName,
				},
			},
		})
	})

}

func TestAccTicketField_expectError(t *testing.T) {
	t.Parallel()

	t.Run("ticket field invalid attribute combo", func(t *testing.T) {
		t.Parallel()
		resourceName := acctest.RandomWithPrefix("tf_acc_ticket_field")
		randomTag := acctest.RandomWithPrefix("tf_acc_tag")
		testField := baseTestField
		testField.Type = "tagger"
		testField.CustomFieldOptions = []zendesk.CustomFieldOption{
			{
				Name:  "test",
				Value: randomTag,
			},
		}
		testField.Title = resourceName
		testField.VisibleInPortal = false
		testField.EditableInPortal = true
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{

				{
					Config:      testAccTicketField(t, testField),
					ExpectError: regexp.MustCompile(`.*Invalid attribute combination.*`),
				},
			},
		})
	})
	t.Run("ticket field change editable in portal", func(t *testing.T) {
		t.Parallel()
		resourceName := acctest.RandomWithPrefix("tf_acc_ticket_field")
		randomTag := acctest.RandomWithPrefix("tf_acc_tag")
		testField := baseTestField
		testField.Type = "tagger"
		testField.CustomFieldOptions = []zendesk.CustomFieldOption{
			{
				Name:  "test",
				Value: randomTag,
			},
		}
		testField.Title = resourceName
		testField.VisibleInPortal = true
		testField.EditableInPortal = true
		testFieldChanged := testField
		testFieldChanged.VisibleInPortal = false
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccTicketField(t, testField),
				},
				{
					Config:      testAccTicketField(t, testFieldChanged),
					ExpectError: regexp.MustCompile(`.*Invalid attribute combination.*`),
				},
				{
					Config: testAccTicketField(t, testField),
				},
			},
		})
	})
}

func testAccTicketField(t *testing.T, field zendesk.TicketField) string {
	t.Helper()

	if field.Type == "" || field.Title == "" {
		t.Fatal("type and title must be set")
	}

	tmpl, err := template.New(ticketFieldTmpl).ParseFiles(ticketFieldTmplPath)
	if err != nil {
		t.Fatal(err)
	}

	var tfFile bytes.Buffer

	err = tmpl.Execute(&tfFile, field)

	if err != nil {
		t.Fatal(err)
	}

	return tfFile.String()
}
