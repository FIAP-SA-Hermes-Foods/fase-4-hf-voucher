resource "kubernetes_service" "hf-voucher-service" { 
    metadata { 
        name = "hf-voucher-service"
        namespace = "dev"
    }

    spec { 
        type = "LoadBalancer"
        selector = { 
            app = "hf-voucher-go-app"
        }
        port { 
            protocol = "TCP"
            name = "hf-voucher-http-port"
            port = 8080
            target_port = 8080
        }
        port { 
            protocol = "TCP"
            name = "hf-voucher-rpc-port"
            port = 8070
            target_port = 8070
        }
    }
}
