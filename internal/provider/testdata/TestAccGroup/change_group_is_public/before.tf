resource "zendesk_group" "test" {
  name      = var.title
  is_public = false
}

variable "title" {
  type     = string
  nullable = false
}