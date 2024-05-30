FROM golang:1.18 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /user-service

FROM gcr.io/distroless/base-debian10
COPY --from=builder /user-service /user-service
EXPOSE 8081
CMD ["/user-service"]