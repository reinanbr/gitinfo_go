package gitinfo

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// GitInfo contains information about a Git repository
type GitInfo struct {
	CommitHash string
	Branch     string
	Tag        string
	Author     string
	CommitDate time.Time
	Message    string
	IsDirty    bool
	RemoteURL  string
}

// GetGitInfo returns complete information about the current Git repository
func GetGitInfo() (*GitInfo, error) {
	return GetGitInfoFromPath(".")
}

// GetGitInfoFromPath returns information about the Git repository at the specified path
func GetGitInfoFromPath(path string) (*GitInfo, error) {
	info := &GitInfo{}

	// Check if it's a Git repository
	if !isGitRepository(path) {
		return nil, fmt.Errorf("not a valid Git repository: %s", path)
	}

	var err error

	// Get commit hash
	info.CommitHash, err = getCommitHash(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get commit hash: %w", err)
	}

	// Get current branch
	info.Branch, err = getCurrentBranch(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get current branch: %w", err)
	}

	// Get latest tag
	info.Tag, _ = getLatestTag(path) // Not a critical error

	// Get commit information
	info.Author, info.CommitDate, info.Message, err = getCommitInfo(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get commit information: %w", err)
	}

	// Check if there are uncommitted changes
	info.IsDirty, err = isWorkingTreeDirty(path)
	if err != nil {
		return nil, fmt.Errorf("failed to check working tree status: %w", err)
	}

	// Get remote URL
	info.RemoteURL, _ = getRemoteURL(path) // Not a critical error

	return info, nil
}

// GetCommitHash returns the hash of the current commit in the Git repository
func GetCommitHash() (string, error) {
	return getCommitHash(".")
}

// GetCurrentBranch returns the current branch of the Git repository
func GetCurrentBranch() (string, error) {
	return getCurrentBranch(".")
}

// IsWorkingTreeDirty checks if there are uncommitted changes
func IsWorkingTreeDirty() (bool, error) {
	return isWorkingTreeDirty(".")
}

// Helper functions

func isGitRepository(path string) bool {
	gitDir := filepath.Join(path, ".git")
	if _, err := os.Stat(gitDir); err == nil {
		return true
	}

	// Check if there's a .git file (worktree case)
	if file, err := os.Open(gitDir); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		if scanner.Scan() {
			line := scanner.Text()
			return strings.HasPrefix(line, "gitdir:")
		}
	}

	return false
}

func runGitCommand(path string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = path
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func getCommitHash(path string) (string, error) {
	return runGitCommand(path, "rev-parse", "HEAD")
}

func getCurrentBranch(path string) (string, error) {
	return runGitCommand(path, "rev-parse", "--abbrev-ref", "HEAD")
}

func getLatestTag(path string) (string, error) {
	return runGitCommand(path, "describe", "--tags", "--abbrev=0")
}

func getCommitInfo(path string) (author string, date time.Time, message string, err error) {
	// Get author
	author, err = runGitCommand(path, "log", "-1", "--pretty=format:%an")
	if err != nil {
		return
	}

	// Get commit date
	dateStr, err := runGitCommand(path, "log", "-1", "--pretty=format:%ci")
	if err != nil {
		return
	}

	date, err = time.Parse("2006-01-02 15:04:05 -0700", dateStr)
	if err != nil {
		return
	}

	// Get commit message
	message, err = runGitCommand(path, "log", "-1", "--pretty=format:%s")
	return
}

func isWorkingTreeDirty(path string) (bool, error) {
	output, err := runGitCommand(path, "status", "--porcelain")
	if err != nil {
		return false, err
	}
	return len(strings.TrimSpace(output)) > 0, nil
}

func getRemoteURL(path string) (string, error) {
	return runGitCommand(path, "remote", "get-url", "origin")
}
