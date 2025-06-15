package github

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v60/github"
	"github.com/sethvargo/go-githubactions"
	"github.com/walle/targz"
)

func GetPRChangedFiles(cxt context.Context, githubContext *githubactions.GitHubContext, action *githubactions.Action, pr *github.PullRequest) (changedFiles []string, err error) {

	token := action.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("missing github token")
	}
	client := github.NewClient(nil).WithAuthToken(token)

	owner, repo := githubContext.Repo()

	lo := github.ListOptions{
		PerPage: 100,
	}
	for {
		action.Infof("Fetching PR from: %s/%s; page %d\n", owner, repo, lo.Page)
		var files []*github.CommitFile
		var resp *github.Response
		files, resp, err = client.PullRequests.ListFiles(cxt, owner, repo, pr.GetNumber(), &lo)
		if err != nil {
			err = fmt.Errorf("failed listing files in PR: %w", err)
			return
		}
		for _, file := range files {
			if *file.Status == "deleted" {
				continue
			}
			slog.Info("changed file", "file", *file.Filename)
			changedFiles = append(changedFiles, *file.Filename)
		}
		if resp.NextPage == 0 {
			break
		}
		lo.Page = resp.NextPage
	}
	return
}

func GetPR(cxt context.Context, githubContext *githubactions.GitHubContext, action *githubactions.Action, pr *github.PullRequest) (pullRequest *github.PullRequest, err error) {

	token := action.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("missing github token")
	}
	client := github.NewClient(nil).WithAuthToken(token)

	owner, repo := githubContext.Repo()

	action.Infof("Reloading PR from: %s/%s\n", owner, repo)

	pullRequest, _, err = client.PullRequests.Get(cxt, owner, repo, pr.GetNumber())
	if err != nil {
		err = fmt.Errorf("failed listing files in PR: %w", err)
		return
	}
	return
}

func Checkout(cxt context.Context, githubContext *githubactions.GitHubContext, action *githubactions.Action, pr *github.PullRequest, ref string, outDir string) (tempDir string, err error) {
	token := action.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		err = fmt.Errorf("missing github token")
		return
	}
	client := github.NewClient(nil).WithAuthToken(token)

	owner, repo := githubContext.Repo()

	var link *url.URL
	var resp *github.Response
	link, resp, err = client.Repositories.GetArchiveLink(cxt, owner, repo, github.Tarball, &github.RepositoryContentGetOptions{Ref: ref}, 1)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusFound {
		err = fmt.Errorf("getArchiveLink returned status: %d, want %d", resp.StatusCode, http.StatusFound)
		return
	}
	resp.Body.Close()

	var req *http.Request
	req, err = client.NewRequest("GET", link.String(), nil)
	if err != nil {
		return
	}
	resp, err = client.BareDo(cxt, req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	pattern := fmt.Sprintf("%s-%s-*", owner, repo)

	var temp *os.File
	temp, err = os.CreateTemp("", pattern+".tar.gz")
	if err != nil {
		err = fmt.Errorf("failed creating temp file: %w", err)
		return
	}

	var written int64
	written, err = io.Copy(temp, resp.Body)
	if err != nil {
		err = fmt.Errorf("failed reading response: %w", err)
		return
	}
	slog.Info("Wrote bytes to temp file", slog.Int("count", int(written)), slog.String("path", temp.Name()))
	defer temp.Close()

	err = targz.Extract(temp.Name(), outDir)
	if err != nil {
		err = fmt.Errorf("failed extracting temp file: %w", err)
		return
	}
	slog.Info("Extracted repo to temp dir", slog.String("path", outDir))

	files, err := filepath.Glob(filepath.Join(outDir, pattern))
	if err != nil {
		err = fmt.Errorf("failed discovering files: %w", err)
		return
	}
	if len(files) == 1 {
		slog.Info("Returning repo dir", slog.String("path", files[0]))
		return files[0], nil
	}
	err = fmt.Errorf("failed to find extracted directory")
	return
}

func WriteComment(cxt context.Context, githubContext *githubactions.GitHubContext, action *githubactions.Action, pr *github.PullRequest, messageId string, comment string) (err error) {
	token := action.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		err = fmt.Errorf("missing github token")
		return
	}
	client := github.NewClient(nil).WithAuthToken(token)

	owner, repo := githubContext.Repo()

	var comments []*github.IssueComment
	var resp *github.Response
	comments, resp, err = client.Issues.ListComments(cxt, owner, repo, pr.GetNumber(), nil)
	if err != nil {
		err = fmt.Errorf("failed discovering files: %w", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("listComments returned status: %d, want %d", resp.StatusCode, http.StatusOK)
		return
	}

	messageIdComment := fmt.Sprintf("<!-- add-pr-comment:%s -->", messageId)
	slog.Info("fetched comments", "total", len(comments))
	var existingComment *github.IssueComment
	for _, c := range comments {
		slog.Info("comment", slog.Any("c", c))
		body := c.GetBody()
		if strings.HasPrefix(body, messageIdComment) {
			existingComment = c
			break
		}
	}

	comment = fmt.Sprintf("%s\n\n%s", messageIdComment, comment)
	ic := &github.IssueComment{Body: &comment}
	if existingComment != nil {
		_, _, err = client.Issues.EditComment(cxt, owner, repo, existingComment.GetID(), ic)
		if err != nil {
			err = fmt.Errorf("failed editing comment: %w", err)
			return
		}
	} else {
		_, _, err = client.Issues.CreateComment(cxt, owner, repo, pr.GetNumber(), ic)
		if err != nil {
			err = fmt.Errorf("failed creating comment: %w", err)
			return
		}
	}
	return
}
