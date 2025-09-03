# GitInfo Go Library

[![CI](https://github.com/reinanbr/gitinfo_go/workflows/CI/badge.svg)](https://github.com/reinanbr/gitinfo_go/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/reinanbr/gitinfo_go)](https://goreportcard.com/report/github.com/reinanbr/gitinfo_go)
[![GoDoc](https://godoc.org/github.com/reinanbr/gitinfo_go?status.svg)](https://godoc.org/github.com/reinanbr/gitinfo_go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Go library for retrieving Git repository information.

## Features

- Get complete Git repository information
- Retrieve commit hash, branch, author, date, and message
- Check for uncommitted changes (dirty working tree)
- Get latest tag and remote URL
- Support for custom repository paths
- Comprehensive error handling
- Full test coverage
- **Command-line interface (CLI) included**

## Installation

### Library
```bash
go get github.com/reinanbr/gitinfo_go
```

### CLI Tool
```bash
# Clone and build
git clone https://github.com/reinanbr/gitinfo_go.git
cd gitinfo_go
make build-cli

# Or install to system PATH
make install-cli
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/reinanbr/gitinfo_go"
)

func main() {
    info, err := gitinfo.GetGitInfo()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Commit: %s\n", info.CommitHash[:8])
    fmt.Printf("Branch: %s\n", info.Branch)
    fmt.Printf("Author: %s\n", info.Author)
    fmt.Printf("Message: %s\n", info.Message)
    fmt.Printf("Dirty: %t\n", info.IsDirty)
}
```

## Usage

### Basic Usage

```go
// Get complete repository information
info, err := gitinfo.GetGitInfo()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Commit: %s\n", info.CommitHash)
fmt.Printf("Branch: %s\n", info.Branch)
fmt.Printf("Author: %s\n", info.Author)
fmt.Printf("Message: %s\n", info.Message)
fmt.Printf("Date: %s\n", info.CommitDate.Format("2006-01-02 15:04:05"))
fmt.Printf("Tag: %s\n", info.Tag)
fmt.Printf("Remote URL: %s\n", info.RemoteURL)
fmt.Printf("Working tree dirty: %t\n", info.IsDirty)
```

### Individual Functions

```go
// Get only commit hash
hash, err := gitinfo.GetCommitHash()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Commit hash:", hash)

// Get only current branch
branch, err := gitinfo.GetCurrentBranch()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Current branch:", branch)

// Check for uncommitted changes
isDirty, err := gitinfo.IsWorkingTreeDirty()
if err != nil {
    log.Fatal(err)
}
fmt.Println("Working tree dirty:", isDirty)
```

### Specify Repository Path

```go
// Get information from specific repository
info, err := gitinfo.GetGitInfoFromPath("/path/to/repo")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Repository info: %+v\n", info)
```

## GitInfo Structure

The `GitInfo` struct contains the following information:

```go
type GitInfo struct {
    CommitHash   string    // Current commit hash
    Branch       string    // Current branch name
    Tag          string    // Latest tag (if exists)
    Author       string    // Last commit author
    CommitDate   time.Time // Last commit date
    Message      string    // Last commit message
    IsDirty      bool      // Whether there are uncommitted changes
    RemoteURL    string    // Origin remote URL (if exists)
}
```

## Requirements

- Go 1.19 or higher
- Git installed on the system
- Valid Git repository

## Development

### Running Tests

```bash
go test -v ./...
```

### Running Example

```bash
cd example
go run main.go
```

### Using Makefile

```bash
make test          # Run tests
make test-coverage # Run tests with coverage
make fmt           # Format code
make vet           # Run go vet
make check         # Run all checks
make run           # Run example
```

## CLI Tool

The package includes a command-line interface for quick Git information retrieval:

```bash
# Basic usage
gitinfo-cli

# JSON output
gitinfo-cli -json

# Specific information
gitinfo-cli -hash    # Show only commit hash
gitinfo-cli -branch  # Show only branch name
gitinfo-cli -dirty   # Check if working tree is dirty
```

See [CLI.md](CLI.md) for complete CLI documentation.

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for a list of changes and releases.
