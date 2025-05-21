resource "zendesk_macro" "test_macro" {
  depends_on = [zendesk_ticket_field.test2]
  title      = "Test Macro TF"
  actions = [
    {
      field = "status"
      value = "open"
    },
    {
      field           = "custom_field",
      value           = "test_tag_2"
      custom_field_id = "1234567890"
    }
  ]
  restriction = {
    type = "Group"
    ids  = [1900001029865]
  }
}
