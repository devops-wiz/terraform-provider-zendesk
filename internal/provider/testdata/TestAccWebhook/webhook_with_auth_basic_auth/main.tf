resource "zendesk_webhook" "test" {
  name           = var.name
  endpoint       = "https://example.com/api"
  http_method    = "GET"
  request_format = "json"

  authentication = {
    add_position = "header"
    credentials = {
      username = "real_username"
      password = "test_password"
    }
    type = "basic_auth"
  }
}

variable "name" {
  type     = string
  nullable = false
}
