#!/bin/bash

# Install the speakeasy CLI
curl -fsSL https://raw.githubusercontent.com/speakeasy-api/speakeasy/main/install.sh | sh

# Setup samples directory
rmdir samples || true
mkdir samples

# Go module commands
go mod download
go mod tidy

# Generate starter usage sample with speakeasy
speakeasy generate usage -s openapi.1.0.0.json -l go -o samples/root.go