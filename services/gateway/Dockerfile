FROM artmexbet/go-dep AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux go build -a -o router ./services/gateway/cmd/app/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/router .
COPY ./services/gateway/config.yml .
CMD ["./router"]
