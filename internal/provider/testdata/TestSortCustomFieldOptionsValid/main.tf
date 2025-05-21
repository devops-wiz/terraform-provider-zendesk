terraform {
  required_providers {
    zendesk = {
      source  = "devops-wiz/zendesk"
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