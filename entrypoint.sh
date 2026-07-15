#!/bin/sh
docker-entrypoint.sh postgres &

env | grep '^POSTGRES_' | sed 's/^POSTGRES_/PG/' > /app/.env

until pg_isready -q; do sleep 1; done
exec /app/program