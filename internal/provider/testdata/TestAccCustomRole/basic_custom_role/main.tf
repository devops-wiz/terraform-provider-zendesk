resource "zendesk_custom_role" "test" {
  name      = var.title
  role_type = 0
  configuration = {
    manage_automations = true
  }
}

variable "title" {
  type     = string
  nullable = true
}