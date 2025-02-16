#!/bin/bash

# Get the parent directory of the script
BASE_DIR="$(cd "$(dirname "$0")"/.. && pwd)"

# Define service directories (relative to BASE_DIR)
services=("user-services" "notification-services" "order-services" "payment-services")

# Run each service in the background
for service in "${services[@]}"; do
    SERVICE_PATH="$BASE_DIR/$service"
    
    if [ -d "$SERVICE_PATH" ]; then
        echo "Starting $service in $SERVICE_PATH..."
        (cd "$SERVICE_PATH" && go run main.go &)
    else
        echo "Error: Directory $SERVICE_PATH does not exist!"
    fi
done

# Wait for all background processes to complete
wait
