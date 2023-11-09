FROM golang:1.19.2-alpine3.16 AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go build -o backend cmd/main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/backend .
COPY --from=builder /app/doc ./doc

ENTRYPOINT ["/app/backend"]