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

echo "PASS: Workflow file, trigger, and login steps verified"
exit 0
