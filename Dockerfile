
# ---------- Build stage ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod only (go.sum may not exist)
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

# ---------- Runtime stage ----------
FROM alpine:latest

WORKDIR /app

# Copy binary
COPY --from=builder /app/server .

# Copy templates and static files
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./server"]

