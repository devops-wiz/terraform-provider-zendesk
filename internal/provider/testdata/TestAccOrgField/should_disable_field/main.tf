resource "zendesk_organization_field" "test" {
  key    = "test_key_${var.test_id}"
  name   = var.title
  type   = "text"
  active = true
}

variable "title" {
  type     = string
  nullable = false
}

variable "test_id" {
  type     = string
  nullable = false
}