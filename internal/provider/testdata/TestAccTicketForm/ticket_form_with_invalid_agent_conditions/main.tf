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
    // Subject
    1900002750725,
    // Description
    1900002750745,
    // Status
    1900002750765,
    // Group
    1900002750825,
    // Assignee
    1900002750845,
    // Ticket Status (custom)
    9193431385367,
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

