# Specification: Build core HTTP/gRPC polling engine and basic SPA dashboard

## Overview
This track focuses on implementing the initial foundational layer for Hops. It will build the backend Go engine capable of parsing the YAML configuration and polling HTTP/gRPC endpoints, alongside the initial frontend Svelte setup for the real-time Single-Page Application dashboard.

## Objectives
- Setup the Go backend utilizing Echo or Gin.
- Implement YAML configuration parsing and ENV variable fallback.
- Build the core polling engine for HTTP and gRPC.
- Setup the Svelte SPA frontend structure.
- Create a basic API endpoint to serve the real-time status to the SPA.