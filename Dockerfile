FROM alpine:3.10.2

MAINTAINER gsxhnd <gsxhnd@gmail.com>

WORKDIR /opt/code
COPY . /opt/code
RUN make
