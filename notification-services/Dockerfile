FROM golang:1.23.6
WORKDIR /app/notification-services
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8084
CMD ["go" "run" "./notification-services/main.go"]
