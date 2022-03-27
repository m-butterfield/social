resource "google_app_engine_application" "default" {
  project     = var.project
  location_id = "us-central"
}

resource "google_cloud_tasks_queue" "publish_post" {
  name     = "social-publish-post"
  location = var.default_region
}

resource "google_project_iam_member" "social_task_creator" {
  project = var.project
  role    = "roles/cloudtasks.enqueuer"
  member  = "serviceAccount:${google_service_account.social_cloud_run.email}"
}
