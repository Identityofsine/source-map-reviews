#!/bin/bash

# Script to test database migrations by dropping and recreating the database
# This will help test the new index migrations

echo "🧹 Cleaning up existing containers and volumes..."

# Stop and remove containers
docker compose -f docker-compose.yaml -f docker-compose.dev-serve.yaml down

# Remove the database volume to start fresh
docker volume rm backend_de_archive-dev-serve 2>/dev/null || true

echo "🔨 Building and starting containers with fresh database..."

# Start with fresh database
docker compose -f docker-compose.yaml -f docker-compose.dev-serve.yaml up --build -d

echo "⏳ Waiting for database to be ready..."
sleep 10

echo "📊 Checking migration status..."
docker compose -f docker-compose.yaml -f docker-compose.dev-serve.yaml logs api | tail -20

echo "🔍 Checking database connection..."
docker compose -f docker-compose.yaml -f docker-compose.dev-serve.yaml exec db psql -U docker -d app -c "\l"

echo "📋 Listing all indexes to verify they were created..."
docker compose -f docker-compose.yaml -f docker-compose.dev-serve.yaml exec db psql -U docker -d app -c "
SELECT 
    schemaname,
    tablename,
    indexname,
    indexdef
FROM pg_indexes 
WHERE schemaname = 'public' 
    AND indexname LIKE 'idx_%'
ORDER BY tablename, indexname;
"

echo "✅ Test complete! Check the output above for any errors." 