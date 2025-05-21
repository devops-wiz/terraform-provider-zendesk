resource "zendesk_user_field" "test" {
  key  = "test_key_${var.title}"
  name = var.title
  type = "text"
}

variable "title" {
  type     = string
  nullable = false
}