#!/usr/bin/env bash

helm uninstall k8ssandra-release

## Cluster environment bootstrapping

kubectl config view --flatten --minify > ../../.kubeconfig

echo "Bootstrapping initial cluster environment complete, kubeconfig temporarily exported to ../../.kubeconfig"

echo "Deploying k8ssandra..."

helm repo add k8ssandra-stable https://helm.k8ssandra.io/stable

helm repo update

helm install k8ssandra-release k8ssandra-stable/k8ssandra -f k8ssandra.values.yaml

echo "Deploying k8ssandra complete."

echo "Deploying k8ssandra-app-user secret..."

kubectl create secret generic k8ssandra-api-service-user --from-literal=username='admin' --from-literal=password='admin123'

echo "Application username:"

echo $(kubectl get secret k8ssandra-api-service-user -o=jsonpath='{.data.username}' | base64 --decode)

echo "Application password:"

echo $(kubectl get secret k8ssandra-api-service-user -o=jsonpath='{.data.password}' | base64 --decode)

echo "Deploying k8ssandra-app secret complete."