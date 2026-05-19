# Stage 1: Build the static Go binary
FROM golang:1.22-bookworm AS builder

WORKDIR /app

# Copy dependency manifests
COPY go.mod ./
# Download dependencies
RUN go mod download

# Copy source code
COPY main.go ./

# Compile statically
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o easy-socks5 main.go

# Stage 2: Final devcontainer image
FROM mcr.microsoft.com/devcontainers/base:debian

# Install libcap2-bin for capability management
RUN apt-get update && apt-get install -y libcap2-bin && rm -rf /var/lib/apt/lists/*

# Copy the pre-built binary
COPY --from=builder /app/easy-socks5 /usr/local/bin/easy-socks5

# Set capabilities so the binary can bind to port 443 without root/sudo
RUN setcap 'cap_net_bind_service=+ep' /usr/local/bin/easy-socks5
