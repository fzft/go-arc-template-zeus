#!/bin/bash

function validate_env {
  local envvar=$1

  # Get the value of the environment variable and convert it to uppercase
  local value=$(echo "${!envvar:-dev}" | tr 'A-Z' 'a-z')

  case "$value" in
    dev|test|prod|staging)
      # valid values, do nothing
      echo "$value"
      ;;
    *)
      echo "Unknown environment: ${value}"
      exit 1
      ;;
  esac
}