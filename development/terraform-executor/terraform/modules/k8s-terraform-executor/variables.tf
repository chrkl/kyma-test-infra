variable "terraform_executor_k8s_service_account" {
  type = object({
    name      = string
    namespace = string
  })
  description = "Details of terraform executor k8s service account."
}

variable "terraform_executor_gcp_service_account" {
  type = object({
    id         = string
    project_id = string
  })
  description = "Details of terraform executor gcp service account."
}