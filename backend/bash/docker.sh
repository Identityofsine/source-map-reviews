#!/bin/bash
# Start the development environment

function systemctl_exists() {
	if [ "$(command -v systemctl)" ]; then
		return 0
	else
		return 1
	fi
}

function systemctl_running() {
	if [ "$(systemctl is-active $1)" == "active" ]; then
		return 0
	else
		return 1
	fi
}

function get_os() {
	if [ "$(uname)" == "Darwin" ]; then
		echo "mac"
	elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
		echo "linux"
	else
		echo "unknown"
	fi
}

function docker_running() {
	if [ "$(docker ps -q)" ]; then
		return 0
	else
		return 1
	fi
}

function get_docker_context() {
	if [ "$(docker context ls | grep -c 'default \*')" -eq 1 ]; then
		echo "default"
	elif [ "$(docker context ls | grep -c 'docker-linux \*')" -eq 1 ]; then
		echo "docker-desktop"
	elif [ "$(docker context ls | grep -c '\*   Docker Desktop')" -eq 1 ]; then
		echo "docker-desktop"
	else
		echo "unknown"
	fi
}

OS=$(get_os)
DOCKER_CONTEXT=$(get_docker_context)

if docker_running; then
	echo "Docker is already running"
else
	echo "Starting Docker"
	if [ "$(get_docker_context)" == "default" ]; then
		echo "Default context found"
		sudo systemctl start docker
	fi 
	if [ "$DOCKER_CONTEXT" == "docker-desktop" ]; then
		echo "Docker Desktop context found"
		if [ "$OS" == "mac" ]; then
			open /Applications/Docker.app
			while ! docker_running; do
				sleep 1
				echo "waiting for docker to start"
			done
		elif [ "$OS" == "linux" ]; then
			if systemctl_exists; then
				echo "systemctl exists"
			else 
				echo "systemctl does not exist... exiting"
				exit 1
			fi
			if docker_running; then 
				echo "Docker Desktop is already running"
			else
				echo "Starting Docker Desktop"	
				systemctl --user start docker-desktop
				#wait for docker to start
				while ! docker_running; do
					sleep 1
					echo "waiting for docker to start"
				done
			fi

		fi
	fi
fi
