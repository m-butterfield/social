terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "3.5.0"
    }
  }
}

provider "google" {
  credentials = file("/var/terraform/mattbutterfield.json")

  project = "mattbutterfield"
  region  = "us-central1"
}

terraform {
  backend "gcs" {
    bucket  = "social-tf-state-prod"
    prefix  = "terraform/state"
  }
}
