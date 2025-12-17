#!/bin/sh
set -e

echo "⏳ Waiting for database..."
# basic wait; bạn có thể nâng cấp healthcheck sau
sleep 3

echo "📦 Running migrations..."
./migrate

echo "🌱 Running seed..."
./seed || true

echo "🚀 Starting API..."
exec ./api
