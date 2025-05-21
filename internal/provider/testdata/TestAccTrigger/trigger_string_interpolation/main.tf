


locals {
  fields = {
    org_field_prefix = "organization.custom_fields"
  }
}

resource "zendesk_trigger" "test" {
  depends_on = [zendesk_trigger_category.test]
  title      = var.title
  actions = [
    {
      field = "status"
      value = "open"
    },
  ]
  conditions = {
    any = [
      {
        field    = "${local.fields.org_field_prefix}.account_risk_level",
        operator = "includes",
        value    = "blah"
      }
  ] }
  category_id = zendesk_trigger_category.test.id
}

resource "zendesk_trigger_category" "test" {
  name = "test_acc_${var.title}_category"
}


variable "title" {
  type     = string
  nullable = false
}
