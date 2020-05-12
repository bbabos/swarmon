#!/bin/sh -e

cd /etc/alertmanager

cat alertmanager.yml |
    sed "s@#url#@api_url: $SLACK_URL@g" |
    sed "s@#user#@username: $SLACK_USER@g" |
    sed "s@#title_link#@title_link: $SCHEMA://alerts.$DOMAIN@g" >tmp_alertmanager.yml

mv tmp_alertmanager.yml alertmanager.yml

set -- /bin/alertmanager "$@"

exec "$@"
