provider "google" {
  project = var.project_id
  region  = var.region
}

resource "google_cloud_run_service" "airules_server" {
  name     = "ctx-airules-server"
  location = var.region

  template {
    spec {
      containers {
        image = "gcr.io/${var.project_id}/ctx-airules-server:latest"
        
        resources {
          limits = {
            cpu    = "1000m"
            memory = "256Mi"
          }
        }
        
        env {
          name  = "OUTPUT_DIR"
          value = "/app/outputs"
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
  service  = google_cloud_run_service.airules_server.name
  location = google_cloud_run_service.airules_server.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}

