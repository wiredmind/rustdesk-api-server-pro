#!/bin/sh
set -eu

umask 077
mkdir -p /app/data
cd /app/data

/app/rustdesk-api-server-pro sync

if [ ! -f .init.lock ] && [ -n "${ADMIN_USER:-}" ]; then
    if [ -z "${ADMIN_PASS_FILE:-}" ]; then
        echo "ADMIN_PASS_FILE is required when ADMIN_USER is set" >&2
        exit 1
    fi
    /app/rustdesk-api-server-pro user add "$ADMIN_USER" --password-file "$ADMIN_PASS_FILE" --admin
    touch .init.lock
fi

exec /app/rustdesk-api-server-pro start
