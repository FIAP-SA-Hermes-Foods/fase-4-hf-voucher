resource "kubernetes_horizontal_pod_autoscaler" "hf-voucher-hpa" {
  metadata {
    name = "hf-voucher-hpa"
    namespace = "dev"
  }

  spec {
    min_replicas = 1
    max_replicas = 2

    scale_target_ref {
      kind = "Deployment"
      name = "hf-voucher-deployment"
    }

    metric { 
        type = "Resource"
        resource { 
            name = "cpu"
            target { 
                type = "Utilization"
                average_utilization = 70
            }
        }
    }
  }
}
