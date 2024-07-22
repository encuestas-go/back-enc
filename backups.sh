#!/bin/bash

# Check if the SQL file is provided as an argument
if [ -z "$1" ]; then
  echo "Usage: $0 <sql_file>"
  exit 1
fi

# Set the SQL file in the docker-compose.yml
sed -i "s/SQL_FILE: .*/SQL_FILE: \"$1\"/" docker-compose.yml

# Restart the Docker container
docker-compose down
docker-compose up -d

# Show logs to verify
docker-compose logs -f db