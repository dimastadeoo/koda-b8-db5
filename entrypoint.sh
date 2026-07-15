#!/bin/sh
docker-entrypoint.sh postgres &

# Write .env to /app/.env and keep POSTGRES_ prefix
env | grep '^POSTGRES_' > /app/.env

until pg_isready -q; do sleep 1; done
exec /app/program