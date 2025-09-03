package gitinfo

import (
	"testing"
	"time"
)

func TestGetCommitHash(t *testing.T) {
	hash, err := GetCommitHash()
	if err != nil {
		t.Skipf("Skipping test - not a Git repository or error: %v", err)
	}
	if len(hash) != 40 {
		t.Errorf("Commit hash should be 40 characters, got: %d", len(hash))
	}
}

func TestGetCurrentBranch(t *testing.T) {
	branch, err := GetCurrentBranch()
	if err != nil {
		t.Skipf("Skipping test - not a Git repository or error: %v", err)
	}
	if branch == "" {
		t.Error("Branch name cannot be empty")
	}
}

func TestIsWorkingTreeDirty(t *testing.T) {
	_, err := IsWorkingTreeDirty()
	if err != nil {
		t.Skipf("Skipping test - not a Git repository or error: %v", err)
	}
	// We don't check the specific value as it may vary
}

func TestGetGitInfo(t *testing.T) {
	info, err := GetGitInfo()
	if err != nil {
		t.Skipf("Skipping test - not a Git repository or error: %v", err)
	}

	// Check if required fields are populated
	if info.CommitHash == "" {
		t.Error("CommitHash cannot be empty")
	}
	if info.Branch == "" {
		t.Error("Branch cannot be empty")
	}
	if info.Author == "" {
		t.Error("Author cannot be empty")
	}
	if info.Message == "" {
		t.Error("Message cannot be empty")
	}
	if info.CommitDate.IsZero() {
		t.Error("CommitDate cannot be zero")
	}
	if info.CommitDate.After(time.Now()) {
		t.Error("CommitDate cannot be in the future")
	}
}

func TestGetGitInfoFromPath(t *testing.T) {
	// Test with non-existent path
	_, err := GetGitInfoFromPath("/nonexistent/path")
	if err == nil {
		t.Error("Should return error for non-existent path")
	}

	// Test with temporary directory (not a Git repository)
	_, err = GetGitInfoFromPath("/tmp")
	if err == nil {
		t.Error("Should return error for directory that is not a Git repository")
	}
}
