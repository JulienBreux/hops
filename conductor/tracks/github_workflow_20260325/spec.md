# Specification: Create the Github worflow to publish Docker image in Github and Docker hub

## Overview
This track focuses on automating the release process for Hops. It will introduce a GitHub Actions workflow that builds, tags, and publishes the project's Docker image to both Docker Hub and GitHub Container Registry (ghcr.io).

## Objectives
- Create a reusable and robust GitHub Actions workflow.
- Automate multi-architecture Docker image builds using Docker Buildx.
- Push images concurrently to Docker Hub and GitHub Container Registry.
- Enforce semantic versioning and commit SHA tagging for released images.

## Functional Requirements
- **Triggers**:
  - The workflow must trigger exclusively when a new Git tag is pushed (e.g., `refs/tags/v*`).
- **Target Architectures**:
  - The workflow must build the image for the following target platforms:
    - `linux/amd64` (Standard x86-64 servers)
    - `linux/arm64` (ARM-based servers like Graviton and Apple Silicon Macs)
    - `linux/arm/v7` (32-bit ARM architectures like older Raspberry Pis)
- **Tagging Strategy**:
  - Released images must be tagged using Semantic Versioning derived from the git tag. (e.g., tag `v1.2.3` creates Docker tags `1.2.3`, `1.2`, `1`, `latest`).
  - Released images must also be tagged with the exact Git commit SHA (`sha-<short-sha>`) for precise traceability.
- **Registries**:
  - Images must be pushed to Docker Hub (`docker.io`).
  - Images must be concurrently pushed to the GitHub Container Registry (`ghcr.io`).

## Acceptance Criteria
- Pushing a new version tag (e.g., `v1.0.0`) triggers a successful run of the GitHub Actions workflow.
- The workflow correctly configures QEMU and Docker Buildx to build multi-architecture images (`linux/amd64`, `linux/arm64`, `linux/arm/v7`).
- The built images are pushed and immediately visible on both Docker Hub and GitHub Container Registry with the correct semantic and SHA tags.
- Secrets for Docker Hub access (`DOCKERHUB_USERNAME`, `DOCKERHUB_TOKEN`) are correctly referenced in the workflow configuration.