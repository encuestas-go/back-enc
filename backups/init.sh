#!/bin/bash
set -e

# Define the SQL file to restore. This can be passed as an environment variable or hardcoded.
SQL_FILE=${SQL_FILE:-"backup.sql"}

# Check if the SQL file exists
if [ -f "/docker-entrypoint-initdb.d/$SQL_FILE" ]; then
  echo "Restoring $SQL_FILE..."
  mysql -uroot -proot ENCUESTA < "/docker-entrypoint-initdb.d/$SQL_FILE"
else
  echo "SQL file $SQL_FILE not found!"
  exit 1
fi