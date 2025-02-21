FROM golang:1.23.6

WORKDIR /app/user-services

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080

# Set environment variables for PostgreSQL
ENV DB_USER=postgres
ENV DB_PASSWORD=postgres
ENV DB_NAME=userdb
ENV DB_HOST=postgres
ENV DB_PORT=5432
ENV DB_SSLMODE=disable

CMD ["go", "run", "./user-services/main.go"]
