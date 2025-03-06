output "service_url" {
  value       = google_cloud_run_service.llm_txt_server.status[0].url
  description = "URL of the deployed Cloud Run service"
}

