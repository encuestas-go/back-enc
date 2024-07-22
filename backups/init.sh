#!/bin/bash
set -e

# Define el archivo SQL que se va a ejecutar
SQL_FILE=${SQL_FILE:-"backup.sql"}

# Función para verificar si MySQL está listo
mysql_ready() {
    mysqladmin ping -h"127.0.0.1" -uroot -proot --silent
}

# Verifica si el archivo SQL existe
if [ -f "/docker-entrypoint-initdb.d/$SQL_FILE" ]; then
    echo "Esperando a que MySQL se inicie..."

    # Esperar a que MySQL esté listo
    while !(mysql_ready)
    do
        echo "Esperando a que MySQL se inicie..."
        sleep 2
    done

    echo "Restaurando desde $SQL_FILE..."
    mysql -h127.0.0.1 -uroot -proot ENCUESTA < "/docker-entrypoint-initdb.d/$SQL_FILE"
else
    echo "Archivo SQL $SQL_FILE no encontrado!"
    exit 1
fi