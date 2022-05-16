#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}" -verbose up

echo "start the app"
exec "$@"
