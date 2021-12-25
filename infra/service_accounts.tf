resource "google_service_account" "social_cloud_run" {
  account_id = "social-cloud-run"
}

resource "google_service_account" "social_uploader" {
  account_id = "social-uploader"
}

resource "google_project_iam_member" "social_act_as_sa" {
  project = var.project
  role    = "roles/iam.serviceAccountUser"
  member  = "serviceAccount:${google_service_account.social_cloud_run.email}"
}
