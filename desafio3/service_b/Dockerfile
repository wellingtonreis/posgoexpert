FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/serviceb ./cmd/api/main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder /app/bin/serviceb .

ENV PORT=8080
EXPOSE 8080

CMD ["/app/serviceb"]
