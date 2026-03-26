# Implementation Plan: Use the configuration instead of mocks and handle empty configuration

## Phase 1: Backend Configuration Implementation
- [ ] Task: Define configuration data structures in Go
    - [ ] Create a `Config` struct that matches the hierarchical specification.
    - [ ] Add unit tests for the configuration struct's JSON/YAML marshaling.
- [ ] Task: Implement configuration loader
    - [ ] Create a function to load configuration from the file specified by `CONFIG_PATH`.
    - [ ] Implement logic to load or override configuration from environment variables.
    - [ ] Write unit tests for the configuration loader, covering successful loads and fallback to empty state.
- [ ] Task: Integrate configuration into health check API
    - [ ] Replace hardcoded mocks in the `/api/health` endpoint with data from the loaded configuration.
    - [ ] Update integration tests for the API to verify it returns the actual configured services.
- [ ] Task: Conductor - User Manual Verification 'Phase 1: Backend Configuration Implementation' (Protocol in workflow.md)

## Phase 2: Frontend Empty State and Guide
- [ ] Task: Implement empty state detection in Svelte
    - [ ] Update the dashboard store or component to identify when the service list is empty.
    - [ ] Add unit tests for the empty state detection logic.
- [ ] Task: Design and implement the Empty State UI
    - [ ] Create a new component for the empty state featuring an illustration.
    - [ ] Implement the "Quick Start Guide" with configuration examples and instructions.
    - [ ] Add visual tests or component tests for the new UI elements.
- [ ] Task: Conductor - User Manual Verification 'Phase 2: Frontend Empty State and Guide' (Protocol in workflow.md)
