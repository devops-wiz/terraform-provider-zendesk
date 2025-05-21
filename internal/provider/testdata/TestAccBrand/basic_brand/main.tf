resource "zendesk_brand" "test" {
  name      = var.title
  subdomain = var.subdomain
}


variable "title" {
  type     = string
  nullable = false
}

variable "subdomain" {
  type     = string
  nullable = false
}