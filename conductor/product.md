# Initial Concept

⚡Hops

Hops is an application designed to verify connectivity between various services using HTTP or gRPC protocols. It serves as a health-checking tool to ensure that inter-service communication is functional across a distributed architecture in real-time.
To maintain ease of use, the configuration is managed through straightforward YAML files or environment variables.
This aligns with current multi-service functionality goals, such as defining job specifications and parameters within a simplified framework.
For every service monitored by Hops, the following attributes must be defined:
An emoji for quick visual identification.
A name and detailed description of the service.
Specific routes to test, including HTTP verbs or gRPC methods.

The user interface consists of a single-page application (SPA) that provides a clear representation of each service's status in real time.
This vision includes a mapping mechanism for services and jobs to provide a comprehensive overview of the system's health.

The Hops project is available in a Docker image format to be usable by service providers like Cloud Run. Deploying Hops on Cloud Run is really easy, with a one-click deployment button available in the repository.

## Target Audience
- **Backend Developers**: Primary users leveraging Hops to test local microservices connectivity during development.

## Core Features
- **Interactive SPA**: A real-time, visual Single-Page Application offering a clear representation of service health and system status.
- **Topology Mapping**: A mapping mechanism for services and jobs to provide a comprehensive system health overview.
- **Emoji-based Visual Identification**: Quick, intuitive visual indicators for monitored services.

## Primary Goal
- **Simplicity**: The overarching goal is zero-friction setup. Configuration is managed entirely through straightforward YAML files or environment variables, making it the easiest tool to deploy.

## Deployment Environments
- **Serverless / Cloud Run**: Optimized for stateless execution, allowing incredibly easy deployment on platforms like Cloud Run.
- **Docker / Local**: Distributed as a lightweight standalone container image, perfect for local development or straightforward server setups.
- **Automated Multi-Arch Releases**: Distributed as a lightweight multi-architecture Docker image (`linux/amd64`, `linux/arm64`, `linux/arm/v7`) automatically published to Docker Hub and GitHub Container Registry upon every release.