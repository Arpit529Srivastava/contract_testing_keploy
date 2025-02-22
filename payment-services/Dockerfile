FROM golang:1.23-alpine

# Set environment variables
ENV GO111MODULE=off 
ENV CGO_ENABLED=0 
ENV GOOS=linux 
ENV GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o payment-service .

# Expose the port the service runs on
EXPOSE 8082

# Command to run the executable
CMD ["./payment-service"]