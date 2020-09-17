#!/bin/sh -e

socat -dd TCP-LISTEN:"$METRIC_PORT",fork TCP:"$GW_BRIDGE_IP":"$METRIC_PORT"

exec "$@"
