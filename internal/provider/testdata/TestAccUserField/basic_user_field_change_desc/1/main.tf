resource "zendesk_user_field" "test" {
  key         = "test_key_${var.title}"
  name        = var.title
  type        = "dropdown"
  description = "test"
  custom_field_options = [{
    name  = "Test value"
    value = "test_value_tag_${var.title}"
  }]
}

variable "title" {
  type     = string
  nullable = false
}