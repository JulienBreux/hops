#!/bin/bash

# Test: Workflow file exists
if [ ! -f ".github/workflows/publish.yml" ]; then
    echo "FAIL: .github/workflows/publish.yml does not exist"
    exit 1
fi

# Test: Workflow trigger is correct
if ! grep -q "tags:" ".github/workflows/publish.yml" || ! grep -q "v\*" ".github/workflows/publish.yml"; then
    echo "FAIL: Trigger events for tag pushes (v*) not found"
    exit 1
fi

# Test: Login to Docker Hub exists
if ! grep -q "Login to Docker Hub" ".github/workflows/publish.yml"; then
    echo "FAIL: Login to Docker Hub step not found"
    exit 1
fi

# Test: Login to GHCR exists
if ! grep -q "Login to GitHub Container Registry" ".github/workflows/publish.yml"; then
    echo "FAIL: Login to GitHub Container Registry step not found"
    exit 1
fi

# Test: Docker metadata exists
if ! grep -q "docker/metadata-action" ".github/workflows/publish.yml"; then
    echo "FAIL: Docker metadata action not found"
    exit 1
fi

# Test: Build and push exists
if ! grep -q "docker/build-push-action" ".github/workflows/publish.yml"; then
    echo "FAIL: Build and push action not found"
    exit 1
fi

# Test: Platforms configured
if ! grep -q "linux/amd64,linux/arm64,linux/arm/v7" ".github/workflows/publish.yml"; then
    echo "FAIL: Target platforms not correctly configured"
    exit 1
fi

echo "PASS: Workflow file, trigger, login, and build steps verified"
exit 0
