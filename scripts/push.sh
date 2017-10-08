#!/bin/bash

APP_VERSION=0.0.1
REGISTRY_NAME=

# aws ecr get-login --region us-east-2

docker tag go-tweet-processor-web:latest ${REGISTRY_NAME}/go-tweet-processor-web:${APP_VERSION}

docker push ${REGISTRY_NAME}/go-tweet-processor-web:${APP_VERSION}