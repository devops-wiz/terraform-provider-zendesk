resource "zendesk_organization_field" "test" {
  type = "dropdown"
  name = var.title
  key  = "test_dropdown_2_${var.title}"

}

variable "title" {
  type     = string
  nullable = false
}