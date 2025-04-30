package action

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/cmd/action/github"
	"github.com/project-chip/alchemy/cmd/action/github/templates"
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/config"
	"github.com/project-chip/alchemy/disco"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/sethvargo/go-githubactions"
)

type Disco struct {
	WriteComment bool `default:"false" hidden:"" help:"Write comment directly"`
}

func (c *Disco) Run(cc *cli.Context) (err error) {

	action := githubactions.New()

	action.Infof("Alchemy %s", config.Version())

	githubContext, err := githubactions.Context()
	if err != nil {
		return fmt.Errorf("failed on getting GitHub context: %w", err)

	}
	pr, err := github.ReadPullRequest(cc, githubContext, action)
	if err != nil {
		action.Errorf("failed on reading pull request: %s", err.Error())
		return fmt.Errorf("failed on reading pull request: %w", err)
	}
	if pr == nil {
		return nil
	}
	var changedFiles []string
	changedFiles, err = github.GetPRChangedFiles(cc, githubContext, action, pr)
	if err != nil {
		return fmt.Errorf("failed on getting pull request changes: %w", err)
	}
	if len(changedFiles) == 0 {
		action.Infof("No changes found\n")
		action.SetOutput("disco_status", "noop")
		return nil
	}

	var changedDocs []string
	for _, path := range changedFiles {
		if filepath.Ext(path) == ".adoc" {
			changedDocs = append(changedDocs, path)
		}
	}

	if len(changedDocs) == 0 {
		action.Infof("No changed asciidoc files found\n")
		action.SetOutput("disco_status", "noop")
		return nil
	}

	pipelineOptions := pipeline.ProcessingOptions{NoProgress: true}

	var out bytes.Buffer
	writer := files.NewPatcher[string]("Generating patch file...", &out)

	err = disco.Pipeline(cc, githubContext.Workspace, changedDocs, pipelineOptions, disco.DefaultOptions, nil, writer)
	if err != nil {
		return fmt.Errorf("failed disco-balling: %v", err)
	}

	var t *raymond.Template
	if out.Len() > 0 {
		slog.Info("Setting disco_status to patched")

		action.SetOutput("disco_status", "patched")
		err = os.WriteFile("disco.patch", out.Bytes(), os.ModeAppend|0644)
		if err != nil {
			return fmt.Errorf("failed saving patch: %v", err)
		}

		t, err = templates.LoadDiscoPatchedTemplate()
		if err != nil {
			err = fmt.Errorf("error loading disco patched template: %w", err)
			return
		}

	} else {
		slog.Info("Setting disco_status to unpatched")
		action.SetOutput("disco_status", "unpatched")

		t, err = templates.LoadDiscoUnpatchedTemplate()
		if err != nil {
			err = fmt.Errorf("error loading disco unpatched template: %w", err)
			return
		}
	}
	var comment string
	comment, err = t.Exec(map[string]any{})
	if err != nil {
		err = fmt.Errorf("error rendering disco comment template: %w", err)
		return
	}
	if c.WriteComment {
		slog.Info("Writing comment", "comment", comment)
		err = github.WriteComment(cc, githubContext, action, pr, "disco-ball", comment)
		if err != nil {
			return
		}
	} else {
		slog.Info("Setting comment", "comment", comment)
		action.SetOutput("comment", comment)
	}
	return nil
}
