# Technology Stack

## Core Language
- **Go**: Used as the primary language for building the lightweight, high-performance backend services necessary for Hops.

## Backend Framework
- **Echo / Gin (Go)**: A lightweight and fast routing framework to expose the HTTP/gRPC health-checking APIs efficiently.

## Frontend Framework
- **Svelte**: Used for the Interactive SPA. Svelte compiles to highly optimized vanilla JavaScript, perfect for building a lean, fast, and responsive real-time dashboard.

## Data Storage & Configuration
- **In-Memory / Config-only**: To ensure absolute simplicity and a zero-friction setup, Hops will run in a stateless manner. Configuration will be managed entirely via YAML files or Environment variables, storing runtime state in-memory without requiring external databases.

## Infrastructure & Deployment
- **Docker**: Used to containerize the application for consistent deployment across environments.
- **Alpine Linux**: The base image for the final production container, ensuring a minimal and secure footprint.
- **GitHub Actions**: Automates the CI/CD pipeline, building and publishing multi-architecture Docker images (`linux/amd64`, `linux/arm64`, `linux/arm/v7`) to multiple registries.