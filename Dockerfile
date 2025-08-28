FROM golang:1.24.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o subscription-service ./cmd/service/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/subscription-service .

COPY --from=builder /app/migrations/ ./migrations