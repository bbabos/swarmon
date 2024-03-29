#!/usr/bin/env bash

branch=$(git rev-parse --abbrev-ref HEAD)

domain="localhost"
adminuser="admin"
adminpw="admin"
traefik_user="admin"
traefik_pw='$$apr1$$EmruEHQ6$$/vaexUtWlwpuI9c24ki7a1' # echo $(htpasswd -nb user password) | sed -e s/\\$/\\$\\$/g
slack_url="http://webhook.url.com"
slack_user="Alert"
slack_channel="test"
stackname="swarmon"
traefik_port="80"
schema="http"
metric_port="9323"
gwbridge="172.26.0.1" # docker run --rm --net host alpine ip -o addr show docker_gwbridge
cgroup_path="/cgroup"
cgroup_enabled="#"
hostname_path="~/hostname"
prom_domain="prometheus"
grafana_domain="grafana"
alert_domain="alerts"

cd configs/docker/
sed <docker-compose.yml "s/{{.Docker.Tag}}/$branch/g" |
    sed "s/{{.Domain}}/$domain/g" |
    sed "s/{{.AdminUser.Name}}/$adminuser/g" |
    sed "s/{{.AdminUser.Password}}/$adminpw/g" |
    sed "s/{{.Docker.StackName}}/$stackname/g" |
    sed "s/{{.Schema}}/$schema/g" |
    sed "s@{{.Slack.Webhook}}@$slack_url@g" |
    sed "s/{{.Slack.AlertUser}}/$slack_user/g" |
    sed "s/{{.Slack.Channel}}/$slack_channel/g" |
    sed "s/{{.Traefik.Port}}/$traefik_port/g" |
    sed "s/{{.Docker.MetricPort}}/$metric_port/g" |
    sed "s/{{.Docker.GwBridgeIP}}/$gwbridge/g" |
    sed "s@{{.Cgroup.Path}}@$cgroup_path@g" |
    sed "s@{{.Cgroup.Enabled}}@$cgroup_enabled@g" |
    sed "s@{{.HostNamePath}}@$hostname_path@g" |
    sed "s/{{.Traefik.BAUser}}/$traefik_user/g" |
    sed "s/{{.Traefik.PrometheusSubDomain}}/$prom_domain/g" |
    sed "s/{{.Traefik.GrafanaSubDomain}}/$grafana_domain/g" |
    sed "s/{{.Traefik.AlertmanagerSubDomain}}/$alert_domain/g" |
    sed "s@{{.Traefik.BAPassword}}@$traefik_pw@g" >tmp-docker-compose.yml

docker stack deploy -c tmp-docker-compose.yml "$stackname"
rm tmp-docker-compose.yml
