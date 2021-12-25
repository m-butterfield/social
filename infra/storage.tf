resource "google_storage_bucket" "content" {
  name     = "social-content"
  location = "US"

  cors {
    origin          = ["*"]
    method          = ["GET", "HEAD", "PUT", "POST", "DELETE"]
    response_header = ["*"]
    max_age_seconds = 3600
  }
}

resource "google_project_iam_member" "social_uploader" {
  project = var.project
  role    = "roles/storage.objectCreator"
  member  = "serviceAccount:${google_service_account.social_uploader.email}"
}

resource "google_project_iam_member" "social_cloud_run_admin" {
  project = var.project
  role    = "roles/storage.objectAdmin"
  member  = "serviceAccount:${google_service_account.social_cloud_run.email}"
}
