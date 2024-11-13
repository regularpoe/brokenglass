package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Scanning from: %s\n", currentDir)
	err = filepath.Walk(currentDir, checkGitRepo)
	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}
}

func checkGitRepo(path string, info os.FileInfo, err error) error {
	if err != nil {
		return nil
	}

	if !info.IsDir() {
		return nil
	}

	if info.Name() == ".git" {
		return filepath.SkipDir
	}

	gitDir := filepath.Join(path, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		return nil
	}

	cmd := exec.Command("git", "-C", path, "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error checking git status in %s: %v\n", path, err)
		return nil
	}

	if len(output) > 0 {
		fmt.Printf("\n📦 Repository: %s\n", path)

		cmdBranch := exec.Command("git", "-C", path, "branch", "--show-current")
		branch, err := cmdBranch.Output()
		if err == nil {
			fmt.Printf("🌿 Branch: %s", string(branch))
		}

		changes := strings.Split(string(output), "\n")
		staged := 0
		unstaged := 0
		untracked := 0

		for _, change := range changes {
			if len(change) < 2 {
				continue
			}
			switch change[0] {
			case 'M':
				staged++
			case ' ':
				if change[1] == 'M' {
					unstaged++
				}
			case '?':
				untracked++
			}
		}

		if staged > 0 {
			fmt.Printf("📝 Staged changes: %d\n", staged)
		}
		if unstaged > 0 {
			fmt.Printf("⚠️  Unstaged changes: %d\n", unstaged)
		}
		if untracked > 0 {
			fmt.Printf("❓ Untracked files: %d\n", untracked)
		}

		fmt.Println("\nChanges:")
		for _, change := range changes {
			if len(change) < 2 {
				continue
			}
			fmt.Printf("%s\n", change)
		}
		fmt.Println(strings.Repeat("-", 50))
	}

	return nil
}

