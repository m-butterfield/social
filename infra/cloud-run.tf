resource "google_cloud_run_service" "social" {
  name     = "social"
  location = var.default_region

  template {
    spec {
      containers {
        image = "gcr.io/mattbutterfield/social"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
