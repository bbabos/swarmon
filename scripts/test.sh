#!/usr/bin/env bash

# TODOs
# check domains (200 OK) + basic auths
# check docker services n/n availability
# test slack integration

domain=$(< ../stackconfig.json jq -r '.Domain')
schema=$(< ../stackconfig.json jq -r '.Schema')
traefik_port=$(< ../stackconfig.json jq -r '.Traefik.Port')
grafana_sub=$(< ../stackconfig.json jq -r '.Traefik.GrafanaSubDomain')
alertm_sub=$(< ../stackconfig.json jq -r '.Traefik.AlertmanagerSubDomain')
prom_sub=$(< ../stackconfig.json jq -r '.Traefik.PrometheusSubDomain')
ba_user=$(< ../stackconfig.json jq -r '.Traefik.BAUser')
ba_pass=$(< ../stackconfig.json jq -r '.Traefik.BAPassword')
grafana_domain=$schema://$grafana_sub.$domain:$traefik_port
alertm_domain=$schema://$alertm_sub.$domain:$traefik_port
prom_domain=$schema://$prom_sub.$domain:$traefik_port
slack_webhook=$(< ../stackconfig.json jq -r '.Slack.Webhook')
slack_user=$(< ../stackconfig.json jq -r '.Slack.AlertUser')
slack_channel=$(< ../stackconfig.json jq -r '.Slack.Channel')

function siteCheck {
    curl -u $ba_user:$ba_pass -H Host:$1.$domain $schema://$domain:$traefik_port > /dev/null 2>&1
    if [[ $? != 0 ]]; then
        echo "TEST FAILED > with subdomain: $1"
    else
        echo "TEST SUCCEED > with subdomain: $1"
    fi
}

function testSlackIntegration {
    curl -X POST -H 'Content-type: application/json' --data '{"text":"Integration test","channel":"'$slack_channel'","username":"'"$slakc_user"'"}' $slack_webhook
}

# Site checks with BA
siteCheck $prom_sub
siteCheck $alertm_sub

# Slack integration test
testSlackIntegration