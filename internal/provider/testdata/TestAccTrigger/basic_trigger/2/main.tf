resource "zendesk_ticket_field" "test" {
  title             = "test_acc_custom_field_trigger_${var.title}"
  type              = "tagger"
  agent_description = "test2"



  custom_field_options = [
    {
      name  = "Test update 2"
      value = "test_tag_trigger"
    },
    {
      name  = "Test 2"
      value = "test_tag_trigger_2"
    }
  ]
}


resource "zendesk_trigger" "test" {
  depends_on = [zendesk_trigger_category.test, zendesk_ticket_field.test]
  title      = var.title
  actions = [
    {
      field = "status"
      value = "open"
    },
    {
      field           = "custom_field",
      value           = "test_tag_trigger"
      custom_field_id = zendesk_ticket_field.test.id
    }
  ]
  conditions = {
    any = [
      {
        field    = "status",
        operator = "is",
        value    = "open"
      },
      {
        field    = "via_id"
        operator = "includes"
        values   = ["4", "75"]
      }
  ] }
  category_id = zendesk_trigger_category.test.id

  position = 2
}

resource "zendesk_trigger" "test2" {
  depends_on = [zendesk_trigger_category.test, zendesk_ticket_field.test]
  title      = var.title
  actions = [
    {
      field = "status"
      value = "open"
    },
    {
      field           = "custom_field",
      value           = "test_tag_trigger"
      custom_field_id = zendesk_ticket_field.test.id
    }
  ]
  conditions = {
    any = [
      {
        field    = "status",
        operator = "is",
        value    = "open"
      }
  ] }
  category_id = zendesk_trigger_category.test.id

  position = 1
}

resource "zendesk_trigger_category" "test" {
  name = "test_acc_${var.title}_category"
}


variable "title" {
  type     = string
  nullable = false
}
