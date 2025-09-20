#!/usr/bin/env bash

# Usage: ./migrate.sh up|down

if [ $# -ne 1 ]; then
  echo "Usage: $0 up|down"
  exit 1
fi

action="$1"

if [[ "$action" != "up" && "$action" != "down" ]]; then
  echo "Error: argument must be 'up' or 'down'"
  exit 1
fi

# Run goose with the proper direction
goose postgres "postgres://postgres:postgres@localhost:5432/gator" "$action"
