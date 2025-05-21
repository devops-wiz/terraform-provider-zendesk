resource "zendesk_organization_field" "test" {
  type = "dropdown"
  name = var.title
  key  = "test_dropdown_2"
  custom_field_options = [
    {
      name  = "Test"
      value = "${var.title}_test_tag_original"
    },
    {
      name  = "Test 2"
      value = "${var.title}_test_tag2_changed"
    }
  ]
}

variable "title" {
  type     = string
  nullable = false
}