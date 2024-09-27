#!/usr/bin/env bash

set -e pipefail

TOKEN_NAME="deployer-token"
DOCKER_IMAGE_REGISTRY_URL=registry.registry.svc.cluster.local:5000

function build_env() {
  IP=$(minikube -p $1 ip)
  SERVER="https://$IP:8443"
  CA=$(kubectl --context $1 get secret/$TOKEN_NAME -o jsonpath='{.data.ca\.crt}')
  TOKEN=$(kubectl --context $1 get secret/$TOKEN_NAME -o jsonpath='{.data.token}' | base64 --decode)

  echo "DOCKER_IMAGE_REGISTRY_URL=$DOCKER_IMAGE_REGISTRY_URL
KUBE_SERVER=$SERVER
KUBE_CA=$CA
KUBE_TOKEN=$TOKEN" >".env.$1"
}

declare -a envs=("dev" "stg" "prod")

for env in "${envs[@]}"; do
  build_env $env
done
