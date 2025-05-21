resource "zendesk_organization_field" "test" {
  type = "dropdown"
  name = var.title
  key  = "test_dropdown_2"
  custom_field_options = [
    {
      name  = "Test"
      value = "${var.title}_test_tag_original"
    }
  ]
}

variable "title" {
  type     = string
  nullable = false
}