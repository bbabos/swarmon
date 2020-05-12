FROM prom/prometheus:v2.17.0

COPY conf /etc/prometheus/

CMD [ "--config.file=/etc/prometheus/prometheus.yml", \
    "--storage.tsdb.path=/prometheus" ]
