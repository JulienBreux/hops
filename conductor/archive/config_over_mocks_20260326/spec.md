# Specification: Use the configuration instead of mocks and handle empty configuration

## Overview
This track focuses on replacing the mock service data currently used by the Go backend with a real configuration loading mechanism. It will also introduce an improved user experience for the Svelte frontend when no services are configured.

## Objectives
- Replace hardcoded mock data in the backend with a configuration loading system.
- Support loading configurations from both a YAML file and environment variables.
- Implement an empty state in the frontend to guide users when no services are defined.

## Functional Requirements
- **Backend - Configuration Loading**:
  - The backend MUST load its configuration from a YAML file or environment variables.
  - The default configuration file path SHOULD be specified via the `CONFIG_PATH` environment variable.
  - The configuration MUST support a hierarchical structure: a list of services, each with its own list of checks.
- **Backend - Data Serving**:
  - The `/api/health` endpoint MUST return the real configuration data instead of mocks.
- **Frontend - Empty State**:
  - The SPA MUST detect when the backend returns an empty list of services.
  - In the empty state, the SPA MUST display an illustration and a "Quick Start Guide" to help users configure their first service.
  - The guide SHOULD provide a simple example of the configuration format.

## Non-Functional Requirements
- **Reliability**: If the configuration file is malformed or missing (and no env vars are provided), the backend SHOULD start but return an empty service list rather than crashing.
- **Maintainability**: Use standard libraries for YAML parsing and structured logging.

## Acceptance Criteria
- Pushing a valid `config.yaml` to the location specified by `CONFIG_PATH` results in the services being correctly displayed in the dashboard.
- Removing all service definitions from the configuration results in the dashboard displaying the "No configured service" illustration and Quick Start Guide.
- The application correctly merges or prioritizes configuration from environment variables if both are present.

## Out of Scope
- Dynamic reloading of configuration without restarting the service (can be a future track).
- Authentication or access control for the configuration.
