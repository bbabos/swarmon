FROM prom/prometheus:v2.19.3

COPY conf /etc/prometheus/

CMD [ "--config.file=/etc/prometheus/prometheus.yml", \
    "--storage.tsdb.path=/prometheus" ]
