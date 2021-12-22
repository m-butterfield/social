variable "default_region" {
  type = string
}

variable "project" {
  type = string
}

variable "db_password" {
  type      = string
  sensitive = true
}

variable "db_socket" {
  type      = string
  sensitive = true
}

variable "social_uploader_service_account" {
  type      = string
  sensitive = true
}
