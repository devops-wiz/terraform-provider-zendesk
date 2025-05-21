resource "zendesk_macro" "test" {
  title = var.title
  actions = [
    {
      field = "status"
      value = "open"
    },
    {
      field = "ticket_form_id"
      value = zendesk_ticket_form.test.id
    }
  ]
}

resource "zendesk_ticket_form" "test" {
  form_name = "${var.title}_form"

  ticket_field_ids = [
    37446012469780,
    37446012469908,
    37446026474900,
    37446026477460,
    37446012470548,
    37446026475284
  ]



}


variable "title" {
  type     = string
  nullable = false
}
