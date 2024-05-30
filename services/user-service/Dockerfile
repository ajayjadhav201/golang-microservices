# Stage 1: Build
FROM golang:1.22 AS builder

# set working directory
WORKDIR /app
# 
COPY go.mod ./
#
RUN go mod download
# 
COPY . .
# 
RUN CGO_ENABLED=0 go build -o /user-service
# 
# 
# Stage 2: Run
FROM gcr.io/distroless/base-debian12 as runner
COPY --from=builder /user-service /user-service
EXPOSE 8081
CMD ["/user-service"]