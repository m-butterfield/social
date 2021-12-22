resource "google_secret_manager_secret" "social_db_socket" {
  secret_id = "social-db-socket"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "social_db_socket_v1" {
  secret      = google_secret_manager_secret.social_db_socket.name
  secret_data = var.db_socket
}

resource "google_secret_manager_secret_iam_member" "cloud_run_social_db_socket" {
  project   = var.project
  secret_id = google_secret_manager_secret.social_db_socket.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.social_cloud_run.email}"
}

resource "google_secret_manager_secret" "social_uploader_service_account" {
  secret_id = "social-uploader-service-account"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "social_uploader_service_account_v1" {
  secret      = google_secret_manager_secret.social_uploader_service_account.name
  secret_data = var.social_uploader_service_account
}

resource "google_secret_manager_secret_iam_member" "cloud_run_uploader_service_account" {
  project   = var.project
  secret_id = google_secret_manager_secret.social_uploader_service_account.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.social_cloud_run.email}"
}

