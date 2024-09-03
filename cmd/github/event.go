package github

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/go-github/v60/github"
	"github.com/sethvargo/go-githubactions"
)

func readPullRequest(cxt context.Context, githubContext *githubactions.GitHubContext, action *githubactions.Action) (*github.PullRequest, error) {
	name := action.Getenv("GITHUB_EVENT_NAME")
	path := action.Getenv("GITHUB_EVENT_PATH")
	if len(path) == 0 || len(name) == 0 {
		return nil, fmt.Errorf("missing environment variables")
	}
	eb, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	switch name {
	case "pull_request":
		var prInfo github.PullRequest
		err = json.Unmarshal(eb, &prInfo)
		if err != nil {
			return nil, err
		}
		return &prInfo, nil

	default:
		action.Infof("unsupported event name: %s", name)
	}
	return nil, nil
}
