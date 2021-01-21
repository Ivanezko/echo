#!/bin/bash
set -e

echo "======= Deploy to DO k8s cluster ======="
echo "======= Build container"
docker build -t echo:k8sprod .
echo "======= Tag container"
docker tag echo:k8sprod registry.digitalocean.com/ivanezko/echo
echo "======= Push to DO repository"
docker push registry.digitalocean.com/ivanezko/echo
echo "======= Restart k8s deployment"
kubectl rollout restart deployment/echo
