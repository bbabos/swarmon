FROM grafana/grafana:6.7.2

COPY datasources /etc/grafana/provisioning/datasources/
COPY dashboards.yml /etc/grafana/provisioning/dashboards/
COPY dashboards /etc/grafana/dashboards/

ENV GF_SECURITY_ADMIN_PASSWORD=admin
ENV GF_SECURITY_ADMIN_USER=admin
ENV GF_PATHS_PROVISIONING=/etc/grafana/provisioning/
# ENV GF_INSTALL_PLUGINS=grafana-image-renderer
ENV GF_INSTALL_IMAGE_RENDERER_PLUGIN=true
