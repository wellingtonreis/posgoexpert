FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go build -o ./bin/cli ./cmd/main.go

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/bin/cli /app/cli
RUN chmod +x /app/cli
CMD ["/app/cli"]