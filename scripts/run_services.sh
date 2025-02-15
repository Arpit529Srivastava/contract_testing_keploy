#!/bin/bash

# Define services and their respective directories
services=(
    "user-service"
    "order-service"
    "payment-service"
    "notification-service"
)

# Function to run each service
start_service() {
    service=$1
    echo "🚀 Starting $service..."
    (cd $service && go run main.go) &
}

# Run all services concurrently
for service in "${services[@]}"; do
    start_service $service
done

# Wait for all background processes
wait
