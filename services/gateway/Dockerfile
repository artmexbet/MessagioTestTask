FROM golang:1.22-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux go build -a -o router ./services/gateway/cmd/app/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/router .
COPY ./services/gateway/config.yml .
CMD ["./router"]
