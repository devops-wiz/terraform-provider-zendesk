resource "zendesk_organization_field" "test" {
  key  = "test_key"
  name = var.title
  type = "text"
}

variable "title" {
  type     = string
  nullable = false
}