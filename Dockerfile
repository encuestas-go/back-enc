FROM mysql:5.7

# Copiar el directorio de backups al contenedor
COPY backups /docker-entrypoint-initdb.d

# Otorgar permisos de ejecución al script init.sh
RUN chmod +x /docker-entrypoint-initdb.d/init.sh