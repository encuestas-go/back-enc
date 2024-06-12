FROM ubuntu:latest
LABEL authors="christian.hernandez"

ENTRYPOINT ["top", "-b"]