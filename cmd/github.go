//go:build github

package cmd

import (
	"github.com/hasty/alchemy/cmd/github"
	"github.com/sethvargo/go-githubactions"
	"github.com/sirupsen/logrus"
)

func Execute() {
	logrus.SetLevel(logrus.ErrorLevel)
	err := github.Action()
	if err != nil {
		githubactions.Fatalf("failed disco ball action: %v\n", err)
	}
}
