FROM golang:1.22-bookworm as builder

WORKDIR /app

COPY go.mod go.sum ./


COPY vendor/ ./vendor/

COPY . .

RUN go build -mod=vendor -o api-gateway .

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/api-gateway ./app

COPY --from=builder /app/.env ./app


EXPOSE 8080

CMD ["./api-gateway"]