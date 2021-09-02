FROM alpine:3.14
WORKDIR /etc/dockerd-exporter/

RUN apk add socat

COPY conf .

ENTRYPOINT [ "./entrypoint.sh" ]
