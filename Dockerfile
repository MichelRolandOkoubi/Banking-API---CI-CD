# Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY src/ .
RUN go build -o banking-api main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/banking-api .
EXPOSE 8080
CMD ["./banking-api"]