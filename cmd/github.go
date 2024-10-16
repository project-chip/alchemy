//go:build github

package cmd

import (
	"github.com/project-chip/alchemy/cmd/github"
	"github.com/sethvargo/go-githubactions"
)

func init() {
	rootCmd.AddCommand(github.Disco)
	rootCmd.AddCommand(github.ZAP)
	defaultCommand = "disco"
}

func handleError(err error) {
	githubactions.Fatalf("failed disco ball action: %v\n", err)
}
