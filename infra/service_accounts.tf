resource "google_service_account" "social_cloud_run" {
  account_id = "social-cloud-run"
}

resource "google_project_iam_member" "social_cloud_run_owner" {
  project = var.project
  role    = "roles/owner"
  member  = "serviceAccount:${google_service_account.social_cloud_run.email}"
}
