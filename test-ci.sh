#!/bin/bash

# CI Simulation Script
# This script simulates what GitHub Actions CI will do

set -e

echo "🚀 Simulating GitHub Actions CI workflow..."
echo

echo "📦 Step 1: Download dependencies"
go mod download
go mod tidy
echo "✅ Dependencies downloaded"
echo

echo "🔧 Step 2: Setup Git (simulation)"
# In CI, this will be: git config --global user.name "CI Test"
echo "✅ Git configured"
echo

echo "🧪 Step 3: Run tests with race detection"
go test -v -race -coverprofile=coverage.out ./...
echo "✅ Tests passed"
echo

echo "🔍 Step 4: Run go vet"
go vet ./...
echo "✅ Go vet passed"
echo

echo "📝 Step 5: Check formatting"
if [ "$(gofmt -s -l . | wc -l)" -gt 0 ]; then
    echo "❌ The following files need to be formatted:"
    gofmt -s -l .
    exit 1
else
    echo "✅ All files are properly formatted"
fi
echo

echo "🏗️  Step 6: Build example"
cd example
go build -v ./...
cd ..
echo "✅ Example built successfully"
echo

echo "🛠️  Step 7: Build CLI"
cd cmd/gitinfo-cli
go build -v ./...
cd ../..
echo "✅ CLI built successfully"
echo

echo "🎉 All CI checks passed! The workflow should work on GitHub Actions."
echo

echo "📊 Coverage report generated: coverage.out"
if command -v go &> /dev/null && go version | grep -q "go1\.[2-9][0-9]\|go1\.[2-9]"; then
    echo "📈 Coverage summary:"
    go tool cover -func=coverage.out | tail -1
fi
