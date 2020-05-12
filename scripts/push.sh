#!/usr/bin/env bash

branch=$(git rev-parse --abbrev-ref HEAD)

docker push babobene/alertmanager:$branch
docker push babobene/dockerd-exporter:$branch
docker push babobene/grafana:$branch
docker push babobene/node-exporter:$branch
docker push babobene/prometheus:$branch

exit 0
