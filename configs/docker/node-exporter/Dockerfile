FROM prom/node-exporter:v1.2.2

ENV NODE_ID=none

USER root

COPY conf /etc/node-exporter/

ENTRYPOINT  [ "/etc/node-exporter/entrypoint.sh"]
CMD [ "/bin/node_exporter" ]
