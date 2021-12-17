resource "google_storage_bucket" "content" {
  name     = "social-content"
  location = "US"
}

resource "google_storage_bucket_access_control" "public_rule" {
  bucket = google_storage_bucket.content.name
  role   = "READER"
  entity = "allUsers"
}
