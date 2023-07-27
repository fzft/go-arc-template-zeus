#!/bin/bash

source ./scripts/base.sh

GET_URL=false

# Regex pattern for MySQL URL validation
URL_REGEX="^mysql:\/\/[^:]+:[^@]+@tcp\(([^:]+:[0-9]+)\)\/[^?]+\?.*$"

MYSQL_CONFIG_PREFIX="db.mysql"

# Loop over all arguments
while [[ $# -gt 0 ]]; do
  key="$1"

  case $key in
  --url)
    GET_URL=true
    shift # past argument
    ;;
  *)      # unknown option
    shift # past argument
    ;;
  esac
done

if [ "$GET_URL" = true ]; then
  ENV=$(validate_env "ENV")

  echo "Getting URL for ${ENV} environment"

  # Check if yq is installed, if not, exit
  if ! [ -x "$(command -v yq)" ]; then
    echo 'Error: yq is not installed. Install it and try again.' >&2
    exit 1
  fi

  # Get the absolute path of the currently executing script
  SCRIPTS_DIR="$(
    cd "$(dirname "$0")"
    pwd -P
  )"

  # Replace 'scripts' with 'config'
  CONFIG_DIR="${SCRIPTS_DIR/scripts/config}"

  CONFIG_FILE="${CONFIG_DIR}/config-${ENV}.yaml"

  # Check if file exists, if not, exit
  if [[ ! -f "$CONFIG_FILE" ]]; then
    echo "Error: $CONFIG_FILE does not exist." >&2
    exit 1
  fi

  # Assuming the yaml file has a structure like this:
  # mysql:
  #   host: "localhost"
  #   port: 3306
  #   password: "password"

  # Get host, port, password, user, dbname, and query from the environment or the config file
  HOST=${MYSQL_HOST:-$(yq e ".${MYSQL_CONFIG_PREFIX}.host" "${CONFIG_FILE}")}
  PORT=${MYSQL_PORT:-$(yq e ".${MYSQL_CONFIG_PREFIX}.port" "${CONFIG_FILE}")}
  PASSWORD=${MYSQL_PASSWORD:-$(yq e ".${MYSQL_CONFIG_PREFIX}.password" "${CONFIG_FILE}")}
  USER=${MYSQL_USER:-$(yq e ".${MYSQL_CONFIG_PREFIX}.user" "${CONFIG_FILE}")}
  DBNAME=${MYSQL_DB:-$(yq e ".${MYSQL_CONFIG_PREFIX}.db" "${CONFIG_FILE}")}
  QUERY=${MYSQL_QUERY:-$(yq e ".${MYSQL_CONFIG_PREFIX}.query" "${CONFIG_FILE}")}

  # Construct the URL
  URL="mysql://${USER}:${PASSWORD}@tcp(${HOST}:${PORT})/${DBNAME}?${QUERY}"

  # Validate the URL
  if [[ $URL =~ $URL_REGEX ]]; then
    echo ${URL}
  else
    echo "Error: Invalid URL" >&2
    exit 1
  fi
fi
