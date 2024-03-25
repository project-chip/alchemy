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

		/*client := github.NewClient(nil).WithAuthToken(token)

		owner, repo := githubContext.Repo()
		githubactions.Infof("Fetching PR from: %s/%s\n", owner, repo)
		fmt.Printf("Fetching PR from: %s/%s\n", owner, repo)
		var files []*github.CommitFile
		files, _, err = client.PullRequests.ListFiles(cxt, owner, repo, *prInfo.Number, nil)
		if err != nil {
			return nil, err
		}
		githubactions.Infof("changed file count: %d\n", len(files))
		fmt.Printf("changed file count: %d\n", len(files))
		changedFiles := make(map[string]string)
		for _, file := range files {
			githubactions.Infof("changed file: %v\n", file)

			fmt.Printf("changed file: %v\n", file)
			if *file.Status == "deleted" {
				continue
			}
			var patch strings.Builder
			if file.PreviousFilename != nil {
				patch.WriteString(fmt.Sprintf("--- %s\n", file.GetPreviousFilename()))
			} else {
				patch.WriteString(fmt.Sprintf("--- %s\n", file.GetFilename()))

			}
			patch.WriteString(fmt.Sprintf("+++ %s\n", file.GetFilename()))
			patch.WriteString(*file.Patch)

			githubactions.Infof("mangled patch: %v\n", patch.String())
			changedFiles[*file.Filename] = patch.String()
		}
		return changedFiles, nil*/

	default:
		action.Infof("unsupported event name: %s", name)
	}
	return nil, nil
}
