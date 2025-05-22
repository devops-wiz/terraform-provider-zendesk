resource "zendesk_organization_field" "test" {
  type = "dropdown"
  name = var.title
  key  = "test_dropdown_${var.test_id}"

}

variable "title" {
  type     = string
  nullable = false
}

variable "test_id" {
  type     = string
  nullable = false
}