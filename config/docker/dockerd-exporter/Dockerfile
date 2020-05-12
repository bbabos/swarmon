FROM alpine:3.11

RUN apk add socat

ENV GW_BRIDGE_IP="172.18.0.1:9323"
ENV METRIC_PORT="9323"

COPY conf /etc/dockerd-exporter/

ENTRYPOINT [ "/etc/dockerd-exporter/entrypoint.sh" ]
