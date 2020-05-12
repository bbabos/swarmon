FROM prom/node-exporter:v0.18.1

ENV NODE_ID=none

USER root

COPY conf /etc/node-exporter/

ENTRYPOINT  [ "/etc/node-exporter/entrypoint.sh"]
CMD [ "/bin/node_exporter" ]
