# TODO add variables to some of this stuff

provider "google" {
  project = "albertlockett-test2"
  region  = "us-east1"
  zone    = "us-east1-b"
}

resource "google_service_account" "service_account" {
  account_id   = "gke-service-account"
  display_name = "GKE Service Account"
  description = "display name for gcp service account"
}

resource "google_container_cluster" "primary" {
  name     = "my-gke-cluster"
  location = "us-east1"

  # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true
  initial_node_count       = 1
}

resource "google_container_node_pool" "primary_preemptiable_nodes" {
  name       = "my-node-pool"
  location   = "us-east1"
  cluster    = google_container_cluster.primary.name
  node_count = 1

  node_config {
    preemptible  = true
    machine_type = "e2-medium"

    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    service_account = google_service_account.service_account.email
    oauth_scopes    = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }  
}