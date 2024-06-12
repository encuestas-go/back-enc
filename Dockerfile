FROM mysql:5.7

RUN sed -i 's/chown -R mysql:mysql "$DATADIR"//' /usr/local/bin/docker-entrypoint.sh

RUN mkdir -p /var/lib/mysql && chown -R mysql:mysql /var/lib/mysql
