FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# Install libcap2-bin so we can set capabilities on the compiled binary
RUN apt-get update && apt-get install -y libcap2-bin && rm -rf /var/lib/apt/lists/*
