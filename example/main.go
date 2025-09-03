package main

import (
	"fmt"
	"log"

	gitinfo "github.com/reinanbr/gitinfo_go"
)

func main() {
	fmt.Println("=== GitInfo Example ===")

	// Get complete repository information
	info, err := gitinfo.GetGitInfo()
	if err != nil {
		log.Printf("Error getting Git information: %v\n", err)
		return
	}

	fmt.Printf("Commit Hash: %s\n", info.CommitHash[:8]+"...") // Show only first 8 characters
	fmt.Printf("Branch: %s\n", info.Branch)
	fmt.Printf("Author: %s\n", info.Author)
	fmt.Printf("Message: %s\n", info.Message)
	fmt.Printf("Date: %s\n", info.CommitDate.Format("2006-01-02 15:04:05"))

	if info.Tag != "" {
		fmt.Printf("Latest Tag: %s\n", info.Tag)
	} else {
		fmt.Println("Latest Tag: (no tags found)")
	}

	if info.RemoteURL != "" {
		fmt.Printf("Remote URL: %s\n", info.RemoteURL)
	} else {
		fmt.Println("Remote URL: (no remote configured)")
	}

	if info.IsDirty {
		fmt.Println("Working Tree: 🔄 has uncommitted changes")
	} else {
		fmt.Println("Working Tree: ✅ clean")
	}

	fmt.Println("\n=== Testing individual functions ===")

	// Test individual functions
	if hash, err := gitinfo.GetCommitHash(); err == nil {
		fmt.Printf("Commit Hash (individual): %s\n", hash[:8]+"...")
	}

	if branch, err := gitinfo.GetCurrentBranch(); err == nil {
		fmt.Printf("Branch (individual): %s\n", branch)
	}

	if isDirty, err := gitinfo.IsWorkingTreeDirty(); err == nil {
		fmt.Printf("Is Dirty (individual): %t\n", isDirty)
	}
}
