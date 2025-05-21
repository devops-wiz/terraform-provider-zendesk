terraform {
  required_providers {
    zendesk = {
      source  = "app.terraform.io/dynatrace-business-systems/zendesk"
      version = "0.0.0"
    }
  }
}

locals {
  cfos = [
    { name : "T", value : "T" },
    { name : "H", value : "H" },
    { name : "D", value : "D" },
    { name : "L", value : "L" }
  ]
}

output "test" {
  value = provider::zendesk::sort_custom_field_options(local.cfos)
}