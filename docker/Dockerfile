FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR "/app"

COPY ./ .

RUN go mod tidy

RUN go build -o binary cmd/main.go

FROM alpine:latest

RUN apk add --no-cache tzdata

COPY --from=builder /app .

EXPOSE 8080

ENTRYPOINT ["/binary"]