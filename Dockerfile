# Start with a Go base image
FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o go-eshop-console .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/go-eshop-console .

COPY .env .env

EXPOSE 8080

CMD ["./go-eshop-console"]
