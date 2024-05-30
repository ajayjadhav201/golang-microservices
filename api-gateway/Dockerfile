# Stage 1: Build
FROM golang:1.22 AS builder

# set working directory
WORKDIR /app
# 
COPY go.mod ./
# 
# Set GIN_MODE to release
ENV GIN_MODE=release
#
RUN go mod download
# 
COPY . .
# 
RUN CGO_ENABLED=0 go build -o /api-gateway
# 
# 
# Stage 2: Run
FROM gcr.io/distroless/base-debian12 as runner
COPY --from=builder /api-gateway /api-gateway
EXPOSE 8080
CMD ["/api-gateway"]