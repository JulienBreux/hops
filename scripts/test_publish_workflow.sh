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

echo "PASS: Workflow file and trigger verified"
exit 0
