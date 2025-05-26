#!/bin/bash

function get_valid_migrations() {
	migrations=( $( ls ./migrations -1p | grep / | sed 's/^\(.*\)\//\1/') )
}

function is_valid_migration() {
	local migration_name="$1"
	get_valid_migrations
	for item in "${migrations[@]}"; do
		if [[ "$item" == "$migration_name" ]]; then
			return 0
		fi
	done
	return 1
}

function usage() {
		echo "Usage: $0 <migration_type> <migration_name>"
		echo "Example: $0 init add_user_table"
		get_valid_migrations
		echo "Available migrations: ${migrations[*]}"
		exit 1
}

migration_type="$1"
migration_name="$2"

if [ -z "$migration_name" ]; then
	usage
fi

if ! is_valid_migration "$migration_type"; then
	echo "Invalid migration type: $migration_type"
	exit 1
fi

migration_dir="./migrations/$migration_type"
goose create "$migration_name" sql -dir "$migration_dir"

