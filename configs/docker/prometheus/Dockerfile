FROM prom/prometheus:v2.20.1

COPY conf /etc/prometheus/

CMD [ "--config.file=/etc/prometheus/prometheus.yml", \
    "--storage.tsdb.path=/prometheus" ]
