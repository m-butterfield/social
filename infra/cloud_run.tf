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
          name  = "WORKER_BASE_URL"
          value = "${google_cloud_run_service.social-worker.status[0].url}/"
        }
        env {
          name  = "TASK_SERVICE_ACCOUNT_EMAIL"
          value = google_service_account.social_cloud_run.email
        }
        env {
          name = "DB_SOCKET"
          value_from {
            secret_key_ref {
              name = google_secret_manager_secret.social_db_socket.secret_id
              key  = "latest"
            }
          }
        }
        volume_mounts {
          mount_path = "/secrets"
          name       = "secrets"
        }
      }
      volumes {
        name = "secrets"
        secret {
          secret_name = google_secret_manager_secret.social_uploader_service_account.secret_id
          items {
            key  = "latest"
            path = "uploadercreds.json"
          }
        }
      }
      service_account_name = google_service_account.social_cloud_run.email
    }
    metadata {
      annotations = {
        "run.googleapis.com/cloudsql-instances"    = google_sql_database_instance.mattbutterfield.connection_name
        "autoscaling.knative.dev/maxScale"         = "100"
        "client.knative.dev/user-image"            = "gcr.io/mattbutterfield/social"
        "run.googleapis.com/client-name"           = "gcloud"
        "run.googleapis.com/client-version"        = "378.0.0"
        "run.googleapis.com/execution-environment" = "gen1"
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

resource "google_cloud_run_domain_mapping" "social" {
  location = var.default_region
  name     = "social.mattbutterfield.com"

  metadata {
    namespace = var.project
  }

  spec {
    route_name = google_cloud_run_service.social.name
  }
}

resource "google_cloud_run_service" "social-worker" {
  name     = "social-worker"
  location = var.default_region

  template {
    spec {
      containers {
        image = "gcr.io/mattbutterfield/social-worker"
        ports {
          container_port = 8001
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
              key  = "latest"
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
        "client.knative.dev/user-image"         = "gcr.io/mattbutterfield/social-worker"
        "run.googleapis.com/client-name"        = "gcloud"
        "run.googleapis.com/client-version"     = "378.0.0"
      }
    }
  }
}

resource "google_project_iam_member" "social_cloud_run_invoker" {
  project = var.project
  role    = "roles/run.invoker"
  member  = "serviceAccount:${google_service_account.social_cloud_run.email}"
}
