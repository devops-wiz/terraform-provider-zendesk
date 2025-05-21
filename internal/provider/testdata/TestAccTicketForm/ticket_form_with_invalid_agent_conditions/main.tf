resource "zendesk_ticket_field" "test" {
  title             = "qa test"
  type              = "tagger"
  agent_description = "test2"

  custom_field_options = [
    {
      name  = "Test update 2"
      value = "test_tag_1"
    },
    {
      name  = "Test 2"
      value = "test_tag_2"
    }
  ]
}

resource "zendesk_ticket_field" "test2" {
  title             = "qa test2"
  type              = "tagger"
  agent_description = "test22"

  custom_field_options = [
    {
      name  = "Test update 2"
      value = "test_tag_12"
    },
    {
      name  = "Test 2"
      value = "test_tag_22"
    }
  ]
}

resource "zendesk_ticket_form" "test" {
  form_name = var.title

  ticket_field_ids = [
    // Default ticket field IDs for form
    37446012469780,
    37446012469908,
    37446026474900,
    37446026477460,
    37446012470548,
    37446026475284,
    // Custom ticket field id
    zendesk_ticket_field.test.id,

    zendesk_ticket_field.test2.id

  ]

  agent_conditions = [{
    parent_field_id = zendesk_ticket_field.test2.id
    child_fields = [{
      id          = zendesk_ticket_field.test.id
      is_required = false
      required_on_statuses = {
        type = "ALL_STATUSES"
      }
    }]
    value = zendesk_ticket_field.test2.custom_field_options[0].value
  }]

}

variable "title" {
  type     = string
  nullable = false
}

