#!/bin/bash

BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
COMMIT_HASH=$(git rev-parse --short HEAD)
BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD)

# check if yq is installed
if ! command -v yq &> /dev/null
then
		echo "yq could not be found, please install it first."
		exit
fi


yq eval ".commit = strenv(COMMIT_HASH) | .branch = strenv(BRANCH_NAME)" -i ./config/server.yaml
