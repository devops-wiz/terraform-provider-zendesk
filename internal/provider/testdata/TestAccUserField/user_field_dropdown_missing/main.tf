resource "zendesk_user_field" "test" {
  type = "dropdown"
  name = var.title
  key  = "test_dropdown_${var.title}"

}

variable "title" {
  type     = string
  nullable = false
}