FROM golang:1.23.0 AS builder

WORKDIR /service_a

RUN GOBIN=/usr/local/bin go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy
RUN go build -o ./bin/service_a ./cmd/api/main.go

FROM alpine:latest

RUN apk --no-cache add bash libc6-compat

WORKDIR /service_a

COPY --from=builder /service_a/bin/service_a .
COPY --from=builder /usr/local/bin/air /usr/local/bin/air

EXPOSE 8000

CMD ["air"]
