resource "zendesk_ticket_field" "test" {
  title             = "test_acc_custom_field_macro_${var.title}"
  type              = "tagger"
  agent_description = "test2"

  custom_field_options = [
    {
      name  = "Test update 2"
      value = "${var.title}_test_tag_macro"
    },
    {
      name  = "Test 2"
      value = "${var.title}_test_tag_macro_2"
    }
  ]
}

resource "zendesk_macro" "test" {
  depends_on = [zendesk_ticket_field.test]
  title      = var.title
  actions = [
    {
      field = "status"
      value = "open"
    },
    {
      field           = "custom_field",
      value           = "${var.title}_test_tag_macro_2"
      custom_field_id = zendesk_ticket_field.test.id
    }
  ]
}


variable "title" {
  type     = string
  nullable = false
}
