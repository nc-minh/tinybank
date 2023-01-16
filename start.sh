#!/bin/sh

set -e

echo "run db migrations"
source /app/app.env
/app/migrate -path /app/db/migrations -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"