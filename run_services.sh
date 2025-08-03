#!/bin/bash

echo "Starting Product Management Services..."

# Function to check if a port is in use
check_port() {
    if lsof -Pi :$1 -sTCP:LISTEN -t >/dev/null ; then
        echo "Port $1 is already in use. Please stop the service using port $1 first."
        exit 1
    fi
}

# Check ports
check_port 8080
check_port 8081

echo "Starting HTTP API Service on port 8080..."
# Start HTTP API service in background
PORT=8080 go run main.go &
HTTP_API_PID=$!

echo "HTTP API Service started with PID: $HTTP_API_PID"

# Wait a moment for HTTP API to start
sleep 3

echo "Starting GraphQL Service on port 8081..."
# Start GraphQL service in background
cd graphql_service
GRAPHQL_PORT=8081 API_BASE_URL=http://localhost:8080 go run main.go &
GRAPHQL_PID=$!

echo "GraphQL Service started with PID: $GRAPHQL_PID"
cd ..

echo ""
echo "Services are running:"
echo "- HTTP API Service: http://localhost:8080"
echo "- GraphQL Service: http://localhost:8081"
echo "- GraphQL Playground: http://localhost:8081"
echo ""
echo "Press Ctrl+C to stop all services"

# Function to cleanup on exit
cleanup() {
    echo ""
    echo "Stopping services..."
    kill $HTTP_API_PID 2>/dev/null
    kill $GRAPHQL_PID 2>/dev/null
    echo "Services stopped."
    exit 0
}

# Set up signal handling
trap cleanup SIGINT SIGTERM

# Wait for background processes
wait 