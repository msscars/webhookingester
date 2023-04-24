#! /bin/bash

docker build -t moscars-webhookingester-api:latest ./src/api
docker build -t moscars-webhookingester-publisher:latest ./src/publisher

minikube image load moscars-webhookingester-api:latest
minikube image load moscars-webhookingester-publisher:latest