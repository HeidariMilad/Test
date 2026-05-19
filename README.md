# Easy SOCKS5 Go Proxy for GitHub Codespaces

A containerized SOCKS5 proxy server written in Go, optimized for automatic execution inside GitHub Codespaces. It binds to port `443` inside the container and is configured to be publicly accessible from the internet.

## Features

- **Automatic Startup**: Starts automatically in the background when the Codespace container initializes.
- **Port 443 Access**: Configured to run on port `443` inside the container. Using Linux capabilities (`setcap`), the process can bind to this privileged port without running as root.
- **Codespaces Port Forwarding**: Automatically forwards port `443` with `public` visibility.
- **Optional Credentials**: Supports optional username/password authentication for securing your proxy.

## Project Structure

```
├── .devcontainer/
│   ├── devcontainer.json   # Codespace lifecycle, ports, & extensions config
│   └── Dockerfile          # Prepares Go environment with cap-binding tools
├── main.go                 # Go SOCKS5 server implementation
├── go.mod                  # Go module definition
├── start.sh                # Helper script to start/restart the server
├── stop.sh                 # Helper script to stop the server
└── status.sh               # Helper script to view process state and logs
```

## Configuration

The server reads configuration from environment variables. You can define these in your GitHub Codespaces settings/secrets or export them inside the terminal:

| Variable | Description | Default |
| :--- | :--- | :--- |
| `PROXY_PORT` | The port the proxy server binds to inside the container. | `443` |
| `PROXY_USERNAME` | Optional username for proxy authentication. | *(None)* |
| `PROXY_PASSWORD` | Optional password for proxy authentication. | *(None)* |

> [!IMPORTANT]
> Since the port is set to public visibility by default, it is **highly recommended** to set `PROXY_USERNAME` and `PROXY_PASSWORD` in your Codespaces environment secrets to prevent unauthorized access to your proxy.

## Control Scripts

You can manage the proxy server manually using the provided shell scripts in the workspace root:

- **Check Status & Logs**:
  ```bash
  ./status.sh
  ```
- **Stop Server**:
  ```bash
  ./stop.sh
  ```
- **Start/Restart Server**:
  ```bash
  ./start.sh
  ```

Logs are written to `socks5.log` in the root of the project workspace.

## How to Connect from Your Local Machine

Because GitHub Codespaces forwards ports over TLS/HTTPS by default for public web access, standard SOCKS5 clients (which speak raw TCP) cannot connect directly to the raw HTTPS URL (`https://<codespace>-443.app.github.dev`).

To route raw TCP traffic through the proxy, use one of the following methods:

### Method 1: GitHub CLI Port Forwarding (Recommended)
You can use the GitHub CLI (`gh`) to establish a direct local TCP bridge to your Codespace port:

1. Install the GitHub CLI on your local machine.
2. Authenticate using `gh auth login`.
3. Start the forwarder:
   ```bash
   gh codespace ports forward 443:443 -c <YOUR_CODESPACE_NAME>
   ```
4. Now, configure your local applications or browser proxy settings to use:
   - **Type**: SOCKS5
   - **Host**: `127.0.0.1`
   - **Port**: `443`
   - **Credentials**: (If configured)

### Method 2: Configure in VS Code Desktop
If you open the Codespace in VS Code Desktop:
1. VS Code automatically manages local port forwarding.
2. Go to the **Ports** tab in the panel.
3. You will see port `443` forwarded to `localhost:443` locally.
4. You can use `socks5://127.0.0.1:443` directly.
