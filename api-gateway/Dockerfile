FROM golang:1.18 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /api-gateway

FROM gcr.io/distroless/base-debian10
COPY --from=builder /api-gateway /api-gateway
EXPOSE 8080
CMD ["/api-gateway"]