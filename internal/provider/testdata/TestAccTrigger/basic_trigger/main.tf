resource "zendesk_ticket_field" "test" {
  title             = "custom_field_trigger_${var.test_id}"
  type              = "tagger"
  agent_description = "test_${var.test_id}"



  custom_field_options = [
    {
      name  = "Test update 2"
      value = "test_tag_trigger_${var.test_id}"
    },
    {
      name  = "Test 2"
      value = "test_tag_trigger_2_${var.test_id}"
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
      value           = "test_tag_trigger_${var.test_id}"
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

  position = zendesk_trigger.test2.position + 1
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
      value           = "test_tag_trigger_${var.test_id}"
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
  name = "${var.title}_category"
}


variable "title" {
  type     = string
  nullable = false
}

variable "test_id" {
  type     = string
  nullable = false
}
