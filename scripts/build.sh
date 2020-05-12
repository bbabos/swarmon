#!/usr/bin/env bash

branch=$(git rev-parse --abbrev-ref HEAD)

cd ../config/docker/
cd alertmanager/
docker build --no-cache -t babobene/alertmanager:$branch .
cd ../dockerd-exporter/
docker build --no-cache -t babobene/dockerd-exporter:$branch .
cd ../grafana/
docker build --no-cache -t babobene/grafana:$branch .
cd ../node-exporter/
docker build --no-cache -t babobene/node-exporter:$branch .
cd ../prometheus/
docker build --no-cache -t babobene/prometheus:$branch .

exit 0
