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
        env {
          name = "DB_SOCKET"
          value_from {
            secret_key_ref {
              name = google_secret_manager_secret.social_db_socket.secret_id
              key  = "1"
            }
          }
        }
      }
      service_account_name = google_service_account.social_cloud_run.email
    }
    metadata {
      annotations = {
        "run.googleapis.com/cloudsql-instances" = google_sql_database_instance.mattbutterfield.connection_name
        "autoscaling.knative.dev/maxScale"      = "100"
        "client.knative.dev/user-image"         = "gcr.io/mattbutterfield/social"
        "run.googleapis.com/client-name"        = "gcloud"
        "run.googleapis.com/client-version"     = "367.0.0"
      }
    }
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
