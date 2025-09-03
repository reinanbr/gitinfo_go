#!/bin/bash

# GitInfo CLI Test Script
# This script demonstrates various CLI functionalities

echo "=== GitInfo CLI Test Suite ==="
echo

# Build the CLI tool
echo "🔨 Building CLI tool..."
cd cmd/gitinfo-cli
go build -o gitinfo-cli main.go
if [ $? -ne 0 ]; then
    echo "❌ Failed to build CLI tool"
    exit 1
fi
echo "✅ CLI tool built successfully"
echo

# Test help
echo "📖 Testing help command:"
./gitinfo-cli -help
echo

# Create a temporary Git repository for testing
echo "🏗️  Creating temporary Git repository for testing..."
TEMP_REPO=$(mktemp -d)
cd "$TEMP_REPO"
git init > /dev/null 2>&1
git config user.name "Test User"
git config user.email "test@example.com"
echo "# Test Repository" > README.md
git add README.md
git commit -m "Initial commit" > /dev/null 2>&1

echo "✅ Test repository created at: $TEMP_REPO"
echo

# Return to CLI directory
cd - > /dev/null

# Test various CLI functions
echo "🧪 Testing CLI functions:"
echo

echo "1. Basic Git info (human readable):"
./gitinfo-cli -path "$TEMP_REPO"
echo

echo "2. JSON output:"
./gitinfo-cli -path "$TEMP_REPO" -json
echo

echo "3. Verbose output:"
./gitinfo-cli -path "$TEMP_REPO" -verbose
echo

echo "4. Commit hash only:"
./gitinfo-cli -path "$TEMP_REPO" -hash
echo

echo "5. Branch name only:"
./gitinfo-cli -path "$TEMP_REPO" -branch
echo

echo "6. Working tree status:"
./gitinfo-cli -path "$TEMP_REPO" -dirty
echo "Exit code: $?"
echo

# Test with dirty working tree
echo "7. Testing dirty working tree:"
echo "Modified file" >> "$TEMP_REPO/README.md"
./gitinfo-cli -path "$TEMP_REPO" -dirty
echo "Exit code: $?"
echo

echo "8. Testing error case (non-Git directory):"
./gitinfo-cli -path "/tmp" 2>&1 || echo "Exit code: $?"
echo

# Clean up
echo "🧹 Cleaning up..."
rm -rf "$TEMP_REPO"
echo "✅ Test completed successfully"

echo
echo "💡 CLI tool is ready to use!"
echo "   Build: go build -o gitinfo-cli main.go"
echo "   Usage: ./gitinfo-cli -help"
