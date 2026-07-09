FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY . .

RUN go build -o /subs-api ./cmd/api

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /subs-api /usr/local/bin/subs-api

EXPOSE 8080

CMD ["subs-api"]
