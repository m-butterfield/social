terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.4.0"
    }
  }
}

provider "google" {
  credentials = file("/var/terraform/mattbutterfield.json")

  project = "mattbutterfield"
  region  = var.default_region
}

terraform {
  backend "gcs" {
    bucket = "social-tf-state-prod"
    prefix = "terraform/state"
  }
}
