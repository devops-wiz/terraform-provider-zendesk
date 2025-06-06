resource "zendesk_ticket_field" "test" {
  title              = var.title
  type               = "tagger"
  agent_description  = "test2"
  visible_in_portal  = true
  required_in_portal = true

  custom_field_options = [
    {
      name  = "Test update 2"
      value = "test_tag"
    },
    {
      name  = "Test 2"
      value = "test_tag_2"
    }
  ]
}

variable "title" {
  type     = string
  nullable = false
}
