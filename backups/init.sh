#!/bin/bash
set -e

# Define el archivo SQL que se va a ejecutar
SQL_FILE=${SQL_FILE:-"backup.sql"}

# Verifica si el archivo SQL existe
if [ -f "/docker-entrypoint-initdb.d/$SQL_FILE" ]; then
  echo "Esperando a que MySQL se inicie..."

  # Esperar a que MySQL esté listo
  while ! mysqladmin ping -h"localhost" --silent; do
    sleep 1
  done

  echo "Restaurando desde $SQL_FILE..."
  mysql -uroot -proot ENCUESTA < "/docker-entrypoint-initdb.d/$SQL_FILE"
else
  echo "Archivo SQL $SQL_FILE no encontrado!"
  exit 1
fi