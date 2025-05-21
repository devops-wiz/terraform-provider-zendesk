resource "zendesk_organization_field" "test" {
  key    = "test_key"
  name   = var.title
  type   = "text"
  active = true
}

variable "title" {
  type     = string
  nullable = false
}