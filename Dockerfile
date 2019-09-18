FROM alpine:3.10.2

MAINTAINER gsxhnd <gsxhnd@gmail.com>

WORKDIR /opt/code
ADD gin-demo /opt/code
EXPOSE 8080
