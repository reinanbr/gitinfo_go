# API Documentation

## Package gitinfo

Package gitinfo provides functions to retrieve Git repository information.

### Types

#### type GitInfo

```go
type GitInfo struct {
    CommitHash   string    // Current commit hash (40 characters)
    Branch       string    // Current branch name
    Tag          string    // Latest tag (empty if no tags)
    Author       string    // Author of the last commit
    CommitDate   time.Time // Date of the last commit
    Message      string    // Message of the last commit
    IsDirty      bool      // Whether working tree has uncommitted changes
    RemoteURL    string    // URL of origin remote (empty if no remote)
}
```

GitInfo contains comprehensive information about a Git repository.

### Functions

#### func GetGitInfo

```go
func GetGitInfo() (*GitInfo, error)
```

GetGitInfo returns complete information about the current Git repository. It's equivalent to calling `GetGitInfoFromPath(".")`.

**Returns:**
- `*GitInfo`: Repository information
- `error`: Error if not a valid Git repository or Git command fails

**Example:**
```go
info, err := gitinfo.GetGitInfo()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Current commit: %s\n", info.CommitHash)
```

#### func GetGitInfoFromPath

```go
func GetGitInfoFromPath(path string) (*GitInfo, error)
```

GetGitInfoFromPath returns information about the Git repository at the specified path.

**Parameters:**
- `path string`: Path to the Git repository

**Returns:**
- `*GitInfo`: Repository information
- `error`: Error if path is not a valid Git repository or Git command fails

**Example:**
```go
info, err := gitinfo.GetGitInfoFromPath("/path/to/repo")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Repository at %s: %+v\n", "/path/to/repo", info)
```

#### func GetCommitHash

```go
func GetCommitHash() (string, error)
```

GetCommitHash returns the hash of the current commit in the Git repository.

**Returns:**
- `string`: 40-character commit hash
- `error`: Error if not a valid Git repository or Git command fails

**Example:**
```go
hash, err := gitinfo.GetCommitHash()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Current commit: %s\n", hash)
```

#### func GetCurrentBranch

```go
func GetCurrentBranch() (string, error)
```

GetCurrentBranch returns the current branch of the Git repository.

**Returns:**
- `string`: Current branch name
- `error`: Error if not a valid Git repository or Git command fails

**Example:**
```go
branch, err := gitinfo.GetCurrentBranch()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Current branch: %s\n", branch)
```

#### func IsWorkingTreeDirty

```go
func IsWorkingTreeDirty() (bool, error)
```

IsWorkingTreeDirty checks if there are uncommitted changes in the working tree.

**Returns:**
- `bool`: True if there are uncommitted changes, false otherwise
- `error`: Error if not a valid Git repository or Git command fails

**Example:**
```go
isDirty, err := gitinfo.IsWorkingTreeDirty()
if err != nil {
    log.Fatal(err)
}
if isDirty {
    fmt.Println("Repository has uncommitted changes")
} else {
    fmt.Println("Repository is clean")
}
```

## Error Handling

All functions return meaningful error messages for common scenarios:

- **Not a Git repository**: When the path doesn't contain a valid Git repository
- **Git command failed**: When underlying Git commands fail (e.g., corrupted repository)
- **Parse errors**: When Git output cannot be parsed (rare)

## Requirements

- **Go**: Version 1.19 or higher
- **Git**: Must be installed and available in PATH
- **Repository**: Must be a valid Git repository (contains `.git` directory or file)

## Thread Safety

This library is thread-safe. All functions can be called concurrently from multiple goroutines.
