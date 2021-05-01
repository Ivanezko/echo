#!/bin/bash
set -e

kubectl --kubeconfig="k8s-stalker-prod-kubeconfig.yaml" apply k8s-prod.yaml

echo "======= Deploy to DO k8s cluster ======="
echo "======= Build container"
docker build -t echo:k8sprod .
echo "======= Tag container"
docker tag echo:k8sprod registry.digitalocean.com/stalkerwebber/echo
echo "======= Push to DO repository"
docker push registry.digitalocean.com/stalkerwebber/echo
echo "======= Restart k8s deployment"
kubectl --kubeconfig="k8s-stalker-prod-kubeconfig.yaml" rollout restart deployment/echo
