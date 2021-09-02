#!/usr/bin/env bash

domain=$(jq <stackconfig.json -r '.Domain')
schema=$(jq <stackconfig.json -r '.Schema')
traefik_port=$(jq <stackconfig.json -r '.Traefik.Port')
grafana_sub=$(jq <stackconfig.json -r '.Traefik.GrafanaSubDomain')
alertm_sub=$(jq <stackconfig.json -r '.Traefik.AlertmanagerSubDomain')
prom_sub=$(jq <stackconfig.json -r '.Traefik.PrometheusSubDomain')
ba_user=$(jq <stackconfig.json -r '.Traefik.BAUser')
ba_pass=$(jq <stackconfig.json -r '.Traefik.BAPassword')
slack_webhook=$(jq <stackconfig.json -r '.Slack.Webhook')
slack_user=$(jq <stackconfig.json -r '.Slack.AlertUser')
slack_channel=$(jq <stackconfig.json -r '.Slack.Channel')
stack_name=$(jq <stackconfig.json -r '.Docker.StackName')
metric_port=$(jq <stackconfig.json -r '.Docker.MetricPort')
grafana_domain=$schema://$grafana_sub.$domain:$traefik_port
alertm_domain=$schema://$alertm_sub.$domain:$traefik_port
prom_domain=$schema://$prom_sub.$domain:$traefik_port

function testSiteAccess {
    if [[ "$2" == BA ]]; then
        curl -u $ba_user:$ba_pass -H Host:$1.$domain $schema://$domain:$traefik_port >/dev/null 2>&1
    else
        curl -H Host:$1.$domain $schema://$domain:$traefik_port >/dev/null 2>&1
    fi
    if [[ "$?" != 0 ]]; then
        echo "TEST FAILED  > testSiteAccess with subdomain: $1"
    else
        echo "TEST SUCCEED > testSiteAccess with subdomain: $1"
    fi
}

function testSlackIntegration {
    curl -X POST -H 'Content-type: application/json' --data '{"text":"swarmon integration test","channel":"'$slack_channel'","username":"'"$slack_user"'"}' $slack_webhook >/dev/null 2>&1
    if [[ "$?" != 0 ]]; then
        echo "TEST FAILED  > testSlackIntegration"
    else
        echo "TEST SUCCEED > testSlackIntegration"
    fi
}

function testDockerServices {
    services=$(docker service ls | grep $stack_name | awk '{print $2":"$4}')

    for service in $services; do
        if [ $service != "NAME:REPLICAS" ]; then
            servicename=$(echo $service | cut -d ':' -f 1)
            echo $service | grep '0/*' >/dev/null 2>&1
            if [[ "$?" != 0 ]]; then
                echo "TEST SUCCEED > testDockerServices on service: $servicename"
            else
                echo "TEST FAILED  > testDockerServices on service: $servicename"
            fi
        fi
    done
}

function testDockerDaemon {
    docker info >/dev/null 2>&1
    if [[ "$?" != 0 ]]; then
        echo "TEST FAILED  > testDockerDaemon"
        exit 1
    else
        echo "TEST SUCCEED > testDockerDaemon"
    fi
}

function testStackCreation {
    docker stack ps $stack_name >/dev/null 2>&1
    if [[ "$?" != 0 ]]; then
        echo "TEST FAILED  > testStackCreation"
        exit 1
    else
        echo "TEST SUCCEED > testStackCreation"
    fi
}

function testDockerMetricPort {
    nc -z localhost $metric_port >/dev/null 2>&1
    if [[ "$?" != 0 ]]; then
        echo "TEST FAILED  > testDockerMetricPort"
    else
        echo "TEST SUCCEED > testDockerMetricPort"
    fi
}

# Test if Docker daemon is running
testDockerDaemon

# Check exposed Docker metrics port on localhost
testDockerMetricPort

# Test if Docker stack creation was succesfull
testStackCreation

# Test site access with or without BasicAuth
testSiteAccess $prom_sub BA
testSiteAccess $alertm_sub BA
testSiteAccess $grafana_sub

# Docker services check
testDockerServices

# Slack integration test
# testSlackIntegration
