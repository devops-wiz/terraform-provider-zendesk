resource "zendesk_ticket_field" "test" {
  title              = "${var.title} qa test"
  type               = "tagger"
  agent_description  = "test2"
  visible_in_portal  = true
  editable_in_portal = true

  custom_field_options = [
    {
      name  = "Test update 2"
      value = "${var.title}_test_tag_1"
    },
    {
      name  = "Test 2"
      value = "${var.title}_test_tag_2"
    }
  ]
}

resource "zendesk_ticket_field" "test2" {
  title              = "${var.title} qa test2"
  type               = "tagger"
  agent_description  = "test22"
  visible_in_portal  = true
  editable_in_portal = true

  custom_field_options = [
    {
      name  = "Test update 2"
      value = "${var.title}_test_tag_12"
    },
    {
      name  = "Test 2"
      value = "${var.title}_test_tag_22"
    },
    {
      name  = "Test 3"
      value = "${var.title}_test_tag_32"
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

  agent_conditions = {
    (zendesk_ticket_field.test2.id) = {

      field_value_map = {
        (zendesk_ticket_field.test2.custom_field_options[2].value) = {
          child_field_conditions = [
            {
              id          = zendesk_ticket_field.test.id
              is_required = true
              required_on_statuses = {
                type     = "SOME_STATUSES",
                statuses = ["pending", "hold", "solved"]
              }
            }
          ]
        }
      }
    }
  }

  end_user_conditions = {
    (zendesk_ticket_field.test2.id) = {

      field_value_map = {
        (zendesk_ticket_field.test2.custom_field_options[0].value) = {
          child_field_conditions = [
            {
              id          = zendesk_ticket_field.test.id
              is_required = true
            }
          ]
        }
      }
    }
  }

}

variable "title" {
  type     = string
  nullable = false
}

