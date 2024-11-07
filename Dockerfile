FROM golang:1.22.3-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o stresstester ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/stresstester .

ENTRYPOINT ["/app/stresstester"]
