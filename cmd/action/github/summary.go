package github

import (
	"context"
	"fmt"
	"os"

	"github.com/sethvargo/go-githubactions"
)

func WriteSummary(cxt context.Context, action *githubactions.Action, result string) (err error) {
	summaryPath := action.Getenv("GITHUB_STEP_SUMMARY")
	if summaryPath == "" {
		// Not running inside a GitHub Action runner
		return
	}

	f, err := os.OpenFile(summaryPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open GITHUB_STEP_SUMMARY: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(result + "\n"); err != nil {
		return fmt.Errorf("failed to write to GITHUB_STEP_SUMMARY: %w", err)
	}
	return
}
