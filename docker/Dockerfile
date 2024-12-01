# Stage 1: Build
FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR "/app"

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o binary cmd/main.go

# Stage 2: Compression
FROM alpine:latest AS compressor

RUN apk add --no-cache upx
COPY --from=builder /app/binary /binary
RUN upx --lzma /binary

# Stage 3: Final Image
FROM alpine:latest

COPY --from=compressor /binary /binary
EXPOSE 8080

ENTRYPOINT ["/binary"]