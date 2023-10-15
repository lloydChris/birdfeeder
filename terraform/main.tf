terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "5.1.0"
    }
  }
}

provider "google" {
  project = "noted-amphora-402018"
  region  = "us-central1"
  zone    = "us-central1-c"
}

resource "google_project" "bird_feeder" {
  name       = "birdfeeder"
  project_id = "birdfeeder1"
}

resource "google_project_service" "project" {
  project = "birdfeeder1"
  service = "vision.googleapis.com"

  timeouts {
    create = "30m"
    update = "40m"
  }

}

resource "google_billing_project_info" "default" {
  project         = google_project.bird_feeder.project_id
  billing_account = "01DCB7-C51474-48DE9C"
}
