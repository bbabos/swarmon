FROM prom/node-exporter:v1.0.0

ENV NODE_ID=none

USER root

COPY conf /etc/node-exporter/

ENTRYPOINT  [ "/etc/node-exporter/entrypoint.sh"]
CMD [ "/bin/node_exporter" ]
