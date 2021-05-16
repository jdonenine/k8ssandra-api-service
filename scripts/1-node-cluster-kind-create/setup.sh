#!/usr/bin/env bash

## Cluster environment bootstrapping

echo "Bootstrapping initial cluster environment..."

kind delete cluster --name k8ssandra-dev

kind create cluster --image "kindest/node:v1.18.15"  --config kind.config.yaml

kubectl config use-context kind-k8ssandra-dev

kubectl config view --flatten --minify > ../../.kubeconfig

echo "Bootstrapping initial cluster environment complete, kubeconfig temporarily exported to ../../.kubeconfig"

echo "Deploying k8ssandra..."

helm repo add k8ssandra-stable https://helm.k8ssandra.io/stable

helm repo add traefik https://helm.traefik.io/traefik

helm repo update

helm install traefik traefik/traefik -f traefik.values.yaml

helm install k8ssandra k8ssandra-stable/k8ssandra -f k8ssandra.values.yaml

echo "Deploying k8ssandra complete."

echo "Deploying k8ssandra-app-user secret..."

kubectl create secret generic k8ssandra-api-service-user --from-literal=username='admin' --from-literal=password='admin123'

echo "Application username:"

echo $(kubectl get secret k8ssandra-api-service-user -o=jsonpath='{.data.username}' | base64 --decode)

echo "Application password:"

echo $(kubectl get secret k8ssandra-api-service-user -o=jsonpath='{.data.password}' | base64 --decode)

echo "Deploying k8ssandra-app secret complete."