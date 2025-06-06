resource "zendesk_ticket_field" "test" {
  title             = "tf_acc_custom_field_automation_${var.test_id}"
  type              = "tagger"
  agent_description = "test2"

  custom_field_options = [
    {
      name  = "Test update 2"
      value = "test_tag_automation_${var.test_id}"
    },
    {
      name  = "Test 2"
      value = "test_tag_automation_2_${var.test_id}"
    }
  ]
}

resource "zendesk_ticket_form" "test" {
  form_name = var.title

  ticket_field_ids = [
    37446012469780,
    37446012469908,
    37446026474900,
    37446026477460,
    37446012470548,
    37446026475284,
    zendesk_ticket_field.test.id
  ]



}


resource "zendesk_automation" "test" {
  depends_on = [zendesk_ticket_field.test]
  title      = var.title
  actions = [
    {
      field = "status"
      value = "open"
    },
    {
      field           = "custom_field",
      value           = "test_tag_automation_${var.test_id}"
      custom_field_id = zendesk_ticket_field.test.id
    }
  ]
  conditions = {
    all = [
      {
        field    = "status",
        operator = "is",
        value    = "open"
      },
      {
        field    = "NEW"
        operator = "is"
        value    = "2"
      },
      {
        field    = "ticket_form_id"
        operator = "is_not"
        value    = zendesk_ticket_form.test.id
      },
      {
        field           = "custom_field",
        value           = "test_tag_automation_${var.test_id}"
        operator        = "is_not"
        custom_field_id = zendesk_ticket_field.test.id
      }
  ] }

  position = 2
}

variable "title" {
  type     = string
  nullable = false
}

variable "test_id" {
  type     = string
  nullable = false
}


