resource "zendesk_webhook" "test" {
  name           = var.name
  endpoint       = "https://example.com/status/200"
  http_method    = "GET"
  request_format = "json"
  status         = "active"
  custom_headers = {
    header-one = "value_one"
    header-two = "value_two"
  }
  authentication = {
    add_position = "header"
    data = {
      username = "real_username"
      password = var.password
    }
    type = "basic_auth"
  }
}

variable "name" {
  type     = string
  nullable = false
}

variable "password" {
  type      = string
  sensitive = true
}