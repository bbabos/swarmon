#!/usr/bin/env bash

branch=$(git rev-parse --abbrev-ref HEAD)

# build
cd configs/docker/alertmanager && docker build --no-cache -t babobene/alertmanager:"$branch" .
cd ../dockerd-exporter && docker build --no-cache -t babobene/dockerd-exporter:"$branch" .
cd ../grafana && docker build --no-cache -t babobene/grafana:"$branch" .
cd ../node-exporter && docker build --no-cache -t babobene/node-exporter:"$branch" .
cd ../prometheus && docker build --no-cache -t babobene/prometheus:"$branch" .

# push
docker push babobene/alertmanager:"$branch"
docker push babobene/dockerd-exporter:"$branch"
docker push babobene/grafana:"$branch"
docker push babobene/node-exporter:"$branch"
docker push babobene/prometheus:"$branch"
