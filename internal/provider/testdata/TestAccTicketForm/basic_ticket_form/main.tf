resource "zendesk_ticket_field" "test" {
  title             = "${var.title}_ticket_field"
  type              = "tagger"
  agent_description = "test2"

  custom_field_options = [
    {
      name  = "Test update 2"
      value = "test_tag_${var.title}"
    },
    {
      name  = "Test 2"
      value = "test_tag_${var.title}_2"
    }
  ]
}

resource "zendesk_ticket_form" "test" {
  form_name = var.title

  ticket_field_ids = [
    37446012469780,
    37446012469908,
    37446026474900,
    37446012470164,
    37446026477204,
    37446026477460,
    zendesk_ticket_field.test.id
  ]



}

variable "title" {
  type     = string
  nullable = false
}