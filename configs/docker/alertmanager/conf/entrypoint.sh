#!/bin/sh -e

cd /etc/alertmanager
< alertmanager.yml sed "s@#url#@api_url: $SLACK_URL@g" |
    sed "s@#user#@username: $SLACK_USER@g" |
    sed "s@#channel#@channel: $SLACK_CHANNEL@g" |
    sed "s@#title_link#@title_link: $SCHEMA://$SUB_DOMAIN.$DOMAIN@g" >tmp_alertmanager.yml
mv tmp_alertmanager.yml alertmanager.yml

set -- /bin/alertmanager "$@"
exec "$@"
