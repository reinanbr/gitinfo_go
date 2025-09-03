package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	gitinfo "github.com/reinanbr/gitinfo_go"
)

var (
	pathFlag    = flag.String("path", ".", "Path to Git repository")
	jsonFlag    = flag.Bool("json", false, "Output in JSON format")
	verboseFlag = flag.Bool("verbose", false, "Verbose output")
	hashOnly    = flag.Bool("hash", false, "Show only commit hash")
	branchOnly  = flag.Bool("branch", false, "Show only branch name")
	dirtyOnly   = flag.Bool("dirty", false, "Check only if working tree is dirty")
	helpFlag    = flag.Bool("help", false, "Show help")
)

func main() {
	flag.Parse()

	if *helpFlag {
		showHelp()
		return
	}

	// Resolve absolute path
	absPath, err := filepath.Abs(*pathFlag)
	if err != nil {
		log.Fatalf("Error resolving path: %v", err)
	}

	if *verboseFlag {
		fmt.Printf("Checking Git repository at: %s\n", absPath)
	}

	// Handle single-value flags
	if *hashOnly {
		if *pathFlag == "." {
			hash, err := gitinfo.GetCommitHash()
			if err != nil {
				log.Fatalf("Error getting commit hash: %v", err)
			}
			fmt.Println(hash)
		} else {
			info, err := gitinfo.GetGitInfoFromPath(absPath)
			if err != nil {
				log.Fatalf("Error getting Git info: %v", err)
			}
			fmt.Println(info.CommitHash)
		}
		return
	}

	if *branchOnly {
		if *pathFlag == "." {
			branch, err := gitinfo.GetCurrentBranch()
			if err != nil {
				log.Fatalf("Error getting branch: %v", err)
			}
			fmt.Println(branch)
		} else {
			info, err := gitinfo.GetGitInfoFromPath(absPath)
			if err != nil {
				log.Fatalf("Error getting Git info: %v", err)
			}
			fmt.Println(info.Branch)
		}
		return
	}

	if *dirtyOnly {
		if *pathFlag == "." {
			isDirty, err := gitinfo.IsWorkingTreeDirty()
			if err != nil {
				log.Fatalf("Error checking working tree: %v", err)
			}
			if isDirty {
				fmt.Println("dirty")
				os.Exit(1)
			} else {
				fmt.Println("clean")
				os.Exit(0)
			}
		} else {
			info, err := gitinfo.GetGitInfoFromPath(absPath)
			if err != nil {
				log.Fatalf("Error getting Git info: %v", err)
			}
			if info.IsDirty {
				fmt.Println("dirty")
				os.Exit(1)
			} else {
				fmt.Println("clean")
				os.Exit(0)
			}
		}
		return
	}

	// Get complete Git information
	var info *gitinfo.GitInfo
	if *pathFlag == "." {
		info, err = gitinfo.GetGitInfo()
	} else {
		info, err = gitinfo.GetGitInfoFromPath(absPath)
	}

	if err != nil {
		log.Fatalf("Error getting Git information: %v", err)
	}

	// Output results
	if *jsonFlag {
		outputJSON(info)
	} else {
		outputHuman(info, *verboseFlag)
	}
}

func showHelp() {
	fmt.Printf(`GitInfo CLI - Git Repository Information Tool

Usage: %s [flags]

Flags:
  -path string
        Path to Git repository (default ".")
  -json
        Output in JSON format
  -verbose
        Verbose output with additional details
  -hash
        Show only commit hash
  -branch
        Show only branch name
  -dirty
        Check only if working tree is dirty (exit code 1 if dirty, 0 if clean)
  -help
        Show this help message

Examples:
  %s                           # Show Git info for current directory
  %s -path /path/to/repo       # Show Git info for specific repository
  %s -json                     # Output in JSON format
  %s -hash                     # Show only commit hash
  %s -branch                   # Show only branch name
  %s -dirty                    # Check if working tree is dirty
  %s -verbose                  # Show detailed information

Exit Codes:
  0 - Success
  1 - Error or working tree is dirty (when using -dirty flag)
`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}

func outputJSON(info *gitinfo.GitInfo) {
	fmt.Printf(`{
  "commit_hash": "%s",
  "branch": "%s",
  "tag": "%s",
  "author": "%s",
  "commit_date": "%s",
  "message": "%s",
  "is_dirty": %t,
  "remote_url": "%s"
}
`, info.CommitHash, info.Branch, info.Tag, info.Author,
		info.CommitDate.Format("2006-01-02T15:04:05Z07:00"),
		info.Message, info.IsDirty, info.RemoteURL)
}

func outputHuman(info *gitinfo.GitInfo, verbose bool) {
	if verbose {
		fmt.Println("=== Git Repository Information ===")
	}

	fmt.Printf("Commit Hash: %s\n", info.CommitHash[:8]+"...")
	fmt.Printf("Branch: %s\n", info.Branch)
	fmt.Printf("Author: %s\n", info.Author)
	fmt.Printf("Date: %s\n", info.CommitDate.Format("2006-01-02 15:04:05"))
	fmt.Printf("Message: %s\n", info.Message)

	if info.Tag != "" {
		fmt.Printf("Latest Tag: %s\n", info.Tag)
	} else if verbose {
		fmt.Println("Latest Tag: (no tags found)")
	}

	if info.RemoteURL != "" {
		fmt.Printf("Remote URL: %s\n", info.RemoteURL)
	} else if verbose {
		fmt.Println("Remote URL: (no remote configured)")
	}

	if info.IsDirty {
		fmt.Println("Working Tree: 🔄 has uncommitted changes")
	} else {
		fmt.Println("Working Tree: ✅ clean")
	}

	if verbose {
		fmt.Printf("Full Commit Hash: %s\n", info.CommitHash)
	}
}
