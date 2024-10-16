package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v60/github"
	"github.com/sethvargo/go-githubactions"
)

func GetPRChangedFiles(cxt context.Context, githubContext *githubactions.GitHubContext, action *githubactions.Action, pr *github.PullRequest) (changedFiles []string, err error) {

	token := action.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("missing github token")
	}
	client := github.NewClient(nil).WithAuthToken(token)

	owner, repo := githubContext.Repo()

	action.Infof("Fetching PR from: %s/%s\n", owner, repo)

	var files []*github.CommitFile
	files, _, err = client.PullRequests.ListFiles(cxt, owner, repo, *pr.Number, nil)
	if err != nil {
		err = fmt.Errorf("failed listing files in PR: %w", err)
		return
	}
	for _, file := range files {
		if *file.Status == "deleted" {
			continue
		}
		changedFiles = append(changedFiles, *file.Filename)
	}
	return
}
