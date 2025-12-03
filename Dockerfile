# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Run stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /go/bin/goose /bin/goose
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]
