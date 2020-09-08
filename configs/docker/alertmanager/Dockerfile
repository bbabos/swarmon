FROM prom/alertmanager:v0.21.0

COPY conf /etc/alertmanager/

ENTRYPOINT  [ "/etc/alertmanager/entrypoint.sh" ]
CMD [ "--config.file=/etc/alertmanager/alertmanager.yml", \
    "--storage.path=/alertmanager" ]
