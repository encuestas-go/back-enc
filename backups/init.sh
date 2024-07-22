#!/bin/bash
set -e

# Define el archivo SQL a restaurar usando la variable de entorno SQL_FILE
SQL_FILE=${SQL_FILE:-"backup.sql"}

# Verifica si el archivo SQL existe
if [ -f "/docker-entrypoint-initdb.d/$SQL_FILE" ]; then
  echo "Restaurando desde $SQL_FILE..."
  mysql -uroot -proot ENCUESTA < "/docker-entrypoint-initdb.d/$SQL_FILE"
else
  echo "Archivo SQL $SQL_FILE no encontrado!"
  exit 1
fi