#!/bin/bash

BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
COMMIT_HASH=$(git rev-parse --short HEAD)
BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD)

function insert_yaml() {
  export COMMIT_HASH=$(git rev-parse --short HEAD)
  export BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD)
	yq eval ".commit = strenv(COMMIT_HASH) | .branch = strenv(BRANCH_NAME)" -i ./config/server.yaml
}

# check if yq is installed
if ! command -v yq &> /dev/null
then
		echo "yq could not be found, please install it first."
		exit
fi

echo "Setting build date to $BUILD_DATE, $COMMIT_HASH, and $BRANCH_NAME"

# check if command ran successfully
if ! insert_yaml; then
	echo "Failed to insert commit hash and branch name into server.yaml"
	exit 1
fi
if [ $? -ne 0 ]; then
	echo "Failed to set build date"
	exit 1
fi
