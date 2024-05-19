resource "kubernetes_secret" "hfregcred" {
  metadata {
    name = "hfregcred"
    namespace = "dev"
  }

  type = "kubernetes.io/dockerconfigjson"

  data = {
    ".dockerconfigjson" = jsonencode({
      auths = {
        "${data.external.env_vars.result.kube_reg_server}" = {
          "username" = data.external.env_vars.result.kube_reg_username
          "password" = data.external.env_vars.result.kube_reg_password
          "email"    = data.external.env_vars.result.kube_reg_email
          "auth"     = base64encode("${data.external.env_vars.result.kube_reg_username}}:${data.external.env_vars.result.kube_reg_password}")
        }
      }
    })
  }
}
