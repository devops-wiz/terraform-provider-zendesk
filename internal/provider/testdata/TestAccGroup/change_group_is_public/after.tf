resource "zendesk_group" "test" {
  name      = var.title
  is_public = true
}

variable "title" {
  type     = string
  nullable = false
}