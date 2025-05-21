resource "zendesk_organization_field" "test" {
  type = "dropdown"
  name = var.title
  key  = "test_dropdown_${var.test_id}"
  custom_field_options = [
    {
      name  = "Test"
      value = "test_tag_original_${var.test_id}"
    },
    {
      name  = "Test 2"
      value = "test_tag_added_${var.test_id}"
    }
  ]
}

variable "title" {
  type     = string
  nullable = false
}

variable "test_id" {
  type     = string
  nullable = false
}