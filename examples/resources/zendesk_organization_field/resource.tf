resource "zendesk_organization_field" "test" {
  type = "dropdown"
  name = var.title
  key  = "test_dropdown"
  custom_field_options = [{
    name  = "Test"
    value = "${var.title}_test_tag"
  }]
}

variable "title" {
  type     = string
  nullable = false
}