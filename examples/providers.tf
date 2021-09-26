terraform {
  required_providers {
    airbyte = {
      version = "~> 0.0.1"
      source  = "jamakase.com/custom/airbyte"
    }
  }
  required_version = "~> 1.0.4"
}

provider "airbyte" {
  host     = "http://localhost:8001/api"
}
