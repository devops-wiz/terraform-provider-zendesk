

terraform {
  required_providers {
    zendesk = {
      source  = "app.terraform.io/dynatrace-business-systems/zendesk"
      version = "0.2.0"
    }
  }
}

provider "zendesk" {
  # Configuration options 
}

  