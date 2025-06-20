#!/bin/sh

set -e

echo "Running DB Migration"
/app/migrate -path /app/migration_file -database "$DB_CONNECTION" -verbose up

echo "Starting Server"
exec "$@"