#!/usr/bin/env bash

branch=$(git rev-parse --abbrev-ref HEAD)

webhook="http://webhook.url.com"
domain="localhost"
adminuser="admin"
adminpw="admin"
traefik_user="admin"
traefik_pw='$$apr1$$EmruEHQ6$$/vaexUtWlwpuI9c24ki7a1' # echo $(htpasswd -nb user password) | sed -e s/\\$/\\$\\$/g
slack_url=$webhook
slack_user="Alertmanager"
stackname="swarmon"
traefik_port="80"
schema="http"
metric_port="9323"
gwbridge="172.18.0.1" # docker run --rm --net host alpine ip -o addr show docker_gwbridge
cgroup="# - /cgroup:/sys/fs/cgroup:ro"

cd ../config/docker/
cat docker-compose.yml |
    sed "s/{{.Tag}}/$branch/g" |
    sed "s/{{.Domain}}/$domain/g" |
    sed "s/{{.AdminUser.Name}}/$adminuser/g" |
    sed "s/{{.AdminUser.Password}}/$adminpw/g" |
    sed "s/{{.Docker.StackName}}/$stackname/g" |
    sed "s/{{.Schema}}/$schema/g" |
    sed "s@{{.Slack.Webhook}}@$slack_url@g" |
    sed "s/{{.Slack.AlertUser}}/$slack_user/g" |
    sed "s/{{.Traefik.Port}}/$traefik_port/g" |
    sed "s/{{.Docker.MetricPort}}/$metric_port/g" |
    sed "s/{{.Docker.GwBridgeIP}}/$gwbridge/g" |
    sed "s@{{.Cgroup}}@$cgroup@g" |
    sed "s/{{.Traefik.BAUser}}/$traefik_user/g" |
    sed "s@{{.Traefik.BAPassword}}@$traefik_pw@g" >tmp-docker-compose.yml

docker stack deploy -c tmp-docker-compose.yml $stackname
rm tmp-docker-compose.yml

exit 0
