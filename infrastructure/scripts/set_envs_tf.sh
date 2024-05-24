#!/bin/bash
echo "{
    \"kube_reg_server\": \"$KUBE_REG_SERVER\",
    \"kube_reg_username\": \"$KUBE_REG_USERNAME\",
    \"kube_reg_password\": \"$KUBE_REG_PASSWORD\", 
    \"kube_reg_email\": \"$KUBE_REG_EMAIL\",
    \"image_tag\": \"$IMAGE_API_HTTP_URL\",
    \"image_api_http_url\": \"$IMAGE_API_HTTP_URL\",
    \"image_api_rpc_url\": \"$IMAGE_API_HTTP_RPC\"
}"
