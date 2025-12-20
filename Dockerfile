# ===== BUILD STAGE =====
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

# Cài git (cần cho go mod)
RUN apk add --no-cache git

# Copy go mod & sum trước để cache
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./cmd/api

# ===== RUN STAGE =====
FROM alpine:3.20

WORKDIR /app

# Cài cert để gọi HTTPS (Cloudinary)
RUN apk add --no-cache ca-certificates

# Copy binary
COPY --from=builder /app/app .

# Copy entrypoint
COPY docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
