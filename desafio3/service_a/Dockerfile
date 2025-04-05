FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/service_a ./cmd/api/main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder /app/bin/service_a .

ENV PORT=8000
EXPOSE 8000

CMD ["/app/service_a"]
