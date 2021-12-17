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
