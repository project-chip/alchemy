//go:build github

package cmd

import (
	"github.com/project-chip/alchemy/cmd/action"
	"github.com/sethvargo/go-githubactions"
)

func init() {
	rootCmd.AddCommand(action.Disco)
	rootCmd.AddCommand(action.ZAP)
	defaultCommand = "disco"
}

func handleError(err error) {
	githubactions.Fatalf("failed disco ball action: %v\n", err)
}
