# Implementation Plan: Add Cloud Run Deployment Button

## Phase 1: Preparation & Research [checkpoint: e2492cb]
- [x] Task: Research official 'Deploy to Google Cloud' button specifications [eb56bbe]
    - [x] Locate the official SVG for the button
    - [x] Determine the deployment link URL structure (deploy.cloud.google.com)
    - [x] Identify parameters for repository, branch, and environment variables
- [x] Task: Conductor - User Manual Verification 'Phase 1: Preparation & Research' (Protocol in workflow.md) [e2492cb]

## Phase 2: Implementation
- [x] Task: Integrate button into README.md [04d3c66]
    - [x] Locate the appropriate insertion point at the top of README.md
    - [x] Add the Markdown for the button image and link
    - [x] Verify the link points to the `main` branch of `https://github.com/julienbreux/hops`
- [ ] Task: Configure deployment parameters
    - [ ] Add `app.json` or equivalent if required for Cloud Run button (or use URL params)
    - [ ] Ensure `julienbreux/hops:latest` image is referenced
    - [ ] Add placeholders for `HOPS_CONFIG_PATH`
- [ ] Task: Conductor - User Manual Verification 'Phase 2: Implementation' (Protocol in workflow.md)

## Phase 3: Final Verification
- [ ] Task: Verify deployment functionality
    - [ ] Click the button in the rendered README.md
    - [ ] Confirm redirection to Google Cloud Console
    - [ ] Validate that parameters (repo, branch) are correctly pre-filled
- [ ] Task: Conductor - User Manual Verification 'Phase 3: Final Verification' (Protocol in workflow.md)