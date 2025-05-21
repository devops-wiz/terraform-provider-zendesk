resource "zendesk_webhook" "test" {
  name           = var.name
  endpoint       = "https://example.com/status/200"
  http_method    = "GET"
  request_format = "json"
}

variable "name" {
  type     = string
  nullable = false
}
