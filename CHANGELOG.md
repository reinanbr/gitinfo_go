# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial release of GitInfo Go library
- `GetGitInfo()` function to get complete Git repository information
- `GetCommitHash()` function to get current commit hash
- `GetCurrentBranch()` function to get current branch name
- `IsWorkingTreeDirty()` function to check for uncommitted changes
- `GetGitInfoFromPath(path)` function to get Git info from specific path
- Support for getting commit author, date, and message
- Support for getting latest tag and remote URL
- Comprehensive test suite
- Example usage in `example/main.go`
- Full documentation in README.md

### Changed
- N/A

### Deprecated
- N/A

### Removed
- N/A

### Fixed
- N/A

### Security
- N/A
