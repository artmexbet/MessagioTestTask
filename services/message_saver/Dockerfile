FROM artmexbet/go-dep AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux go build -a -o message_saver ./services/message_saver/cmd/app/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/message_saver .
COPY ./services/message_saver/config.yml .
CMD ["./message_saver"]
