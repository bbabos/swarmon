#!/usr/bin/env bash

domain=$(< stackconfig.json jq -r '.Domain')
schema=$(< stackconfig.json jq -r '.Schema')
traefik_port=$(< stackconfig.json jq -r '.Traefik.Port')
grafana_sub=$(< stackconfig.json jq -r '.Traefik.GrafanaSubDomain')
alertm_sub=$(< stackconfig.json jq -r '.Traefik.AlertmanagerSubDomain')
prom_sub=$(< stackconfig.json jq -r '.Traefik.PrometheusSubDomain')
ba_user=$(< stackconfig.json jq -r '.Traefik.BAUser')
ba_pass=$(< stackconfig.json jq -r '.Traefik.BAPassword')
slack_webhook=$(< stackconfig.json jq -r '.Slack.Webhook')
slack_user=$(< stackconfig.json jq -r '.Slack.AlertUser')
slack_channel=$(< stackconfig.json jq -r '.Slack.Channel')
grafana_domain=$schema://$grafana_sub.$domain:$traefik_port
alertm_domain=$schema://$alertm_sub.$domain:$traefik_port
prom_domain=$schema://$prom_sub.$domain:$traefik_port

function testSiteAccess {
    if [[ "$2" == BA ]]; then
        curl -u $ba_user:$ba_pass -H Host:$1.$domain $schema://$domain:$traefik_port > /dev/null 2>&1
    else
        curl -H Host:$1.$domain $schema://$domain:$traefik_port > /dev/null 2>&1
    fi
    if [[ "$?" != 0 ]]; then
        echo "TEST FAILED  > testSiteAccess with subdomain: $1"
    else
        echo "TEST SUCCEED > testSiteAccess with subdomain: $1"
    fi
}

function testSlackIntegration {
    curl -X POST -H 'Content-type: application/json' --data '{"text":"swarmon integration test","channel":"'$slack_channel'","username":"'"$slack_user"'"}' $slack_webhook > /dev/null 2>&1
    if [[ "$?" != 0 ]]; then
        echo "TEST FAILED  > testSlackIntegration"
    else
        echo "TEST SUCCEED > testSlackIntegration"
    fi
}

function testDockerServices {
    services=$(docker service ls | awk '{print $2":"$4}')

    for service in $services; do
        if [ $service != "NAME:REPLICAS" ]; then
            servicename=$(echo $service | cut -d ':' -f 1)
            echo $service | grep '0/*' > /dev/null 2>&1
            if [[ "$?" != 0 ]]; then
                echo "TEST SUCCEED > testDockerServices on service: $servicename"
            else
                echo "TEST FAILED  > testDockerServices on service: $servicename"
            fi
        fi
    done
}
 
# Test site access
testSiteAccess $prom_sub BA
testSiteAccess $alertm_sub BA
testSiteAccess $grafana_sub

# Docker services check
testDockerServices

# Slack integration test
testSlackIntegration
