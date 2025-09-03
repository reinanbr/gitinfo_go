# GitInfo CLI Tool

A command-line interface for the GitInfo Go library that allows you to quickly retrieve Git repository information from the terminal.

## Installation

### Option 1: Build from source
```bash
# Clone the repository
git clone https://github.com/reinanbr/gitinfo_go.git
cd gitinfo_go

# Build the CLI tool
make build-cli

# The binary will be available at: cmd/gitinfo-cli/gitinfo-cli
```

### Option 2: Install to system PATH
```bash
# Build and install to /usr/local/bin
make install-cli

# Now you can use it from anywhere
gitinfo-cli -help
```

## Usage

### Basic Commands

```bash
# Show Git info for current directory
gitinfo-cli

# Show Git info for specific repository
gitinfo-cli -path /path/to/repo

# Show help
gitinfo-cli -help
```

### Output Formats

```bash
# Human-readable output (default)
gitinfo-cli

# JSON output
gitinfo-cli -json

# Verbose output with additional details
gitinfo-cli -verbose
```

### Specific Information

```bash
# Show only commit hash
gitinfo-cli -hash

# Show only branch name
gitinfo-cli -branch

# Check if working tree is dirty
gitinfo-cli -dirty
```

## Examples

### Basic Repository Information
```bash
$ gitinfo-cli
Commit Hash: a1b2c3d4...
Branch: main
Author: John Doe
Date: 2025-09-03 15:30:45
Message: Add new feature
Working Tree: ✅ clean
```

### JSON Output for Scripts
```bash
$ gitinfo-cli -json
{
  "commit_hash": "a1b2c3d4e5f6789...",
  "branch": "main",
  "tag": "v1.2.3",
  "author": "John Doe",
  "commit_date": "2025-09-03T15:30:45Z",
  "message": "Add new feature",
  "is_dirty": false,
  "remote_url": "https://github.com/user/repo.git"
}
```

### Use in Shell Scripts
```bash
#!/bin/bash

# Get current branch
BRANCH=$(gitinfo-cli -branch)
echo "Current branch: $BRANCH"

# Check if working tree is dirty
if gitinfo-cli -dirty > /dev/null; then
    echo "Working tree is clean"
else
    echo "Working tree has uncommitted changes"
    exit 1
fi

# Get commit hash for tagging
HASH=$(gitinfo-cli -hash)
echo "Current commit: $HASH"
```

### CI/CD Integration
```bash
# In your CI/CD pipeline
BUILD_INFO=$(gitinfo-cli -json)
echo "Build info: $BUILD_INFO"

# Check for uncommitted changes before deployment
gitinfo-cli -dirty || {
    echo "Error: Cannot deploy with uncommitted changes"
    exit 1
}
```

## Exit Codes

- **0**: Success
- **1**: Error occurred or working tree is dirty (when using `-dirty` flag)

## Available Flags

| Flag | Description | Example |
|------|-------------|---------|
| `-path` | Path to Git repository | `-path /home/user/project` |
| `-json` | Output in JSON format | `-json` |
| `-verbose` | Verbose output with additional details | `-verbose` |
| `-hash` | Show only commit hash | `-hash` |
| `-branch` | Show only branch name | `-branch` |
| `-dirty` | Check if working tree is dirty | `-dirty` |
| `-help` | Show help message | `-help` |

## Testing

Run the CLI test suite:
```bash
# Run all CLI tests
make test-cli

# Or run the test script directly
./test-cli.sh
```

## Use Cases

1. **Development Workflow**: Quickly check repository status
2. **CI/CD Pipelines**: Validate repository state before builds
3. **Shell Scripts**: Integrate Git information into automation
4. **Deployment**: Ensure clean working tree before deployment
5. **Debugging**: Get comprehensive repository information for troubleshooting
