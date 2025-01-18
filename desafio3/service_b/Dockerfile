FROM golang:1.23.0 AS builder

WORKDIR /service_b

RUN GOBIN=/usr/local/bin go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy
RUN go build -o ./bin/service_b ./cmd/api/main.go

FROM alpine:latest

RUN apk --no-cache add bash libc6-compat

WORKDIR /service_b

COPY --from=builder /service_b/bin/service_b .
COPY --from=builder /usr/local/bin/air /usr/local/bin/air

EXPOSE 9000

CMD ["air"]
