FROM golang:1.23.6

WORKDIR /app/payment-services

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8082

CMD ["go", "run", "./payment-services/main.go"]
