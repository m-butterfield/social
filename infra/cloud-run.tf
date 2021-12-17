resource "google_cloud_run_service" "social" {
  name     = "social"
  location = var.default_region

  template {
    spec {
      containers {
        image = "gcr.io/mattbutterfield/social"
        ports {
          container_port = 8000
        }
        env {
          name  = "GIN_MODE"
          value = "release"
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_service.social.location
  project  = google_cloud_run_service.social.project
  service  = google_cloud_run_service.social.name

  policy_data = data.google_iam_policy.noauth.policy_data
}
