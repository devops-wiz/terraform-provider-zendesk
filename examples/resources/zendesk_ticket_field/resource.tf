resource "zendesk_ticket_field" "test2" {
  title = " Terraform test 3"
  type  = "tagger"

  custom_field_options = [
    {
      name  = "Test"
      value = "test_tag"
    }
  ]
}


