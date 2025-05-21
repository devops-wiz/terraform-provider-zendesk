resource "zendesk_brand" "test" {
  name      = var.title
  subdomain = "testsubdomain123"
}


variable "title" {
  type     = string
  nullable = false
}