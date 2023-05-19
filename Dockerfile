# Build stage
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main server.go

# RUN stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 1234
CMD ["/app/main"]