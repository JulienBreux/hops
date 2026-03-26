# ⚡ Hops

Hops is a lightweight, real-time health-checking application designed to verify connectivity between services using HTTP or gRPC protocols. It provides a visual dashboard to monitor the health of your distributed architecture with zero-friction setup.

## ✨ Features

- **Real-Time Monitoring**: Instant visibility into service health via a Svelte-based SPA.
- **Multi-Protocol Support**: Built-in checkers for both HTTP(S) and gRPC.
- **Simple Configuration**: Manage monitored services using straightforward YAML or environment variables.
- **Stateless & Lightweight**: Distributed as a single Docker container with no external database requirements.
- **Deployment Ready**: Optimized for Cloud Run and other serverless platforms.

## 🚀 Quick Start

### Prerequisites

- [Go](https://go.dev/doc/install) 1.26+
- [Node.js](https://nodejs.org/en/download/) 25+
- [Docker](https://docs.docker.com/get-docker/) (optional, for containerization)

### Local Development

1. **Clone the repository**:
   ```bash
   git clone https://github.com/julienbreux/hops.git
   cd hops
   ```

2. **Install dependencies and build**:
   ```bash
   make build
   ```

3. **Run the application**:
   ```bash
   ./hops
   ```
   The dashboard will be available at `http://localhost:8080`.

## ⚙️ Configuration

Hops is configured via a `config.yaml` file or environment variables.

### YAML Configuration

Create a `config.yaml` in the root directory:

```yaml
services:
  - name: API Gateway
    description: "Main entry point for external traffic"
    emoji: "🚀"
    checks:
      - name: health-check
        type: http
        endpoint: https://api.example.com/health
        method: GET
  - name: Auth Service
    description: "Internal authentication service"
    emoji: "🔒"
    checks:
      - name: grpc-health
        type: grpc
        endpoint: auth-service:50051
```

### Environment Variables

Any value in the YAML configuration can be overridden or provided via environment variables using the `${VAR_NAME}` syntax.

## 🐳 Docker

### Build and Run locally

Build the image:
```bash
make docker-build
```

Run the container:
```bash
make docker-run
```

### Deploy to Cloud Run

Hops is perfectly suited for Google Cloud Run. To deploy:

1. Build and push the image to Artifact Registry:
   ```bash
   docker tag hops:latest gcr.io/[PROJECT-ID]/hops
   docker push gcr.io/[PROJECT-ID]/hops
   ```

2. Deploy to Cloud Run:
   ```bash
   gcloud run deploy hops --image gcr.io/[PROJECT-ID]/hops --platform managed
   ```

## 🛠 Development

The following `make` targets are available:

- `make build`: Build Go and Svelte components.
- `make test`: Run Go and Svelte tests.
- `make lint`: Run linters for both backend and frontend.
- `make docker-build`: Create the Docker image.
- `make docker-run`: Run the image locally on port 8080.
