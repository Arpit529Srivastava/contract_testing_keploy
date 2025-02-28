FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum before downloading dependencies
COPY go.mod go.sum ./
RUN go mod tidy 

# Copy the entire project source
COPY . .

# Build the Go application
RUN go build -o order-services .

# Use a lightweight final image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/order-services .
EXPOSE 8081
CMD ["./order-services"]