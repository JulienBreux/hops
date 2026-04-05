# Specification: Add Cloud Run Deployment Button

## Overview
This track aims to simplify the deployment of Hops to Google Cloud Run by adding a "Deploy to Google Cloud" button to the primary README.md file. This will provide a zero-friction path for users to try Hops in a cloud environment.

## Functional Requirements
- Add the official "Deploy to Google Cloud" button SVG/link to the `README.md`.
- The button must be placed at the top of the README, under the initial Hops description.
- The deployment link must target the `main` branch of the current repository.
- The deployment should be configured to use the official Hops Docker image (`julienbreux/hops:latest`).
- Provide placeholders or default configuration for key environment variables (e.g., `HOPS_CONFIG_PATH`).

## Non-Functional Requirements
- **Simplicity**: The deployment should require minimal user input once clicked.
- **Visual Integration**: The button should follow standard Google Cloud deployment button guidelines.

## Acceptance Criteria
- [ ] A "Deploy to Google Cloud" button is visible at the top of the `README.md`.
- [ ] Clicking the button takes the user to the Google Cloud Console's Cloud Run deployment page.
- [ ] The deployment page is pre-filled with the `main` branch of the Hops repository.
- [ ] The deployment is configured to use the `julienbreux/hops:latest` image.
- [ ] The deployment successfully starts a Hops instance on Cloud Run (verified manually).

## Out of Scope
- Configuring custom domain names during the automated deployment.
- Setting up complex VPC or networking configurations.
- Automating the creation of a Google Cloud Project if the user doesn't have one.