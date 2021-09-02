FROM prom/prometheus:v2.29.1

COPY conf /etc/prometheus/

CMD [ "--config.file=/etc/prometheus/prometheus.yml", \
    "--storage.tsdb.path=/prometheus" ]
