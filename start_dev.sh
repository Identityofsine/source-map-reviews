#!/bin/bash

source ./bash/docker.sh
source ./bash/yaml.sh

# run the docker compose command
docker compose -f docker-compose.yaml -f docker-compose.dev.yaml up

if [ "$1" == "--kill" ] || [ "$1" == "-k" ]; then
	echo "Killing Docker"
	if [ "$DOCKER_CONTEXT" == "default" ]; then
		sudo systemctl stop docker
	fi
	if [ "$DOCKER_CONTEXT" == "docker-desktop" ]; then
		if [ "$OS" == "mac" ]; then
			osascript -e 'quit app "Docker"'
		elif [ "$OS" == "linux" ]; then
			systemctl --user stop docker-desktop
		fi
	fi
fi
