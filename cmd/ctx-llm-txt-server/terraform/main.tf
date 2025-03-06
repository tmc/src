provider "google" {
  project = var.project_id
  region  = var.region
}

resource "google_cloud_run_service" "llm_txt_server" {
  name     = "ctx-llm-txt-server"
  location = var.region

  template {
    spec {
      containers {
        image = "gcr.io/${var.project_id}/ctx-llm-txt-server:latest"
        
        resources {
          limits = {
            cpu    = "1000m"
            memory = "256Mi"
          }
        }
        
        env {
          name  = "DOMAIN"
          value = "llm.txt.tmc.dev"
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

resource "google_cloud_run_service_iam_member" "public_access" {
  service  = google_cloud_run_service.llm_txt_server.name
  location = google_cloud_run_service.llm_txt_server.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}

