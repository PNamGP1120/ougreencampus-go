# ===== builder =====
FROM golang:1.25-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build 3 binaries
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/migrate ./cmd/migrate
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/seed ./cmd/seed

# ===== runtime =====
FROM alpine:3.20

WORKDIR /app
RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /out/api ./api
COPY --from=builder /out/migrate ./migrate
COPY --from=builder /out/seed ./seed

COPY entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

EXPOSE 8080
ENTRYPOINT ["./entrypoint.sh"]
