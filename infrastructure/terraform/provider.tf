terraform { 
    required_providers { 
        mycloud = { 
            source = "hashicorp/kubernetes"
        } 
    }
}

provider "kubernetes" { 
    config_path = "~/.kube/config"
}
