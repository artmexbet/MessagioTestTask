#FROM build_image AS builder
FROM golang:1.22-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux go build -a -o message_saver ./services/message_saver/cmd/app/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/message_saver .
COPY ./services/message_saver/config.yml .
CMD ["./message_saver"]
