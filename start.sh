#!/bin/sh

set -e

echo "Running DB Migration"
source /app/app.env
/app/migrate -path /app/migration_file -database "$DB_CONNECTION" -verbose up

echo "Starting Server"
exec "$@"