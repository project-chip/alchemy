//go:build github

package cmd

import (
	"github.com/hasty/alchemy/cmd/github"
	"github.com/sethvargo/go-githubactions"
)

func init() {
	rootCmd.AddCommand(github.Command)
	defaultCommand = "github"
}

func handleError(err error) {
	githubactions.Fatalf("failed disco ball action: %v\n", err)
}
