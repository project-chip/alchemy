package disco

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/project-chip/alchemy/cmd/action/github"
	"github.com/project-chip/alchemy/config"
	"github.com/project-chip/alchemy/disco"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/sethvargo/go-githubactions"
	"github.com/spf13/cobra"
)

func Ball(cmd *cobra.Command, args []string) error {
	cxt := context.Background()

	action := githubactions.New()

	action.Infof("Alchemy %s", config.Version())

	githubContext, err := githubactions.Context()
	if err != nil {
		return fmt.Errorf("failed on getting GitHub context: %w", err)

	}
	pr, err := github.ReadPullRequest(cxt, githubContext, action)
	if err != nil {
		return fmt.Errorf("failed on reading pull request: %w", err)
	}
	if pr == nil {
		return nil
	}
	var changedFiles []string
	changedFiles, err = github.GetPRChangedFiles(cxt, githubContext, action, pr)
	if err != nil {
		return fmt.Errorf("failed on getting pull request changes: %w", err)
	}
	if len(changedFiles) == 0 {
		action.Infof("No changes found\n")
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
		return nil
	}

	pipelineOptions := pipeline.Options{NoProgress: true}

	var out bytes.Buffer
	writer := files.NewPatcher[string]("Generating patch file...", &out)

	err = disco.Pipeline(cxt, ".", changedDocs, pipelineOptions, nil, nil, writer)
	if err != nil {
		return fmt.Errorf("failed disco-balling: %v", err)
	}

	if out.Len() > 0 {
		action.SetOutput("disco_status", "patched")
		err = os.WriteFile("disco.patch", out.Bytes(), os.ModeAppend|0644)
		if err != nil {
			return fmt.Errorf("failed saving patch: %v", err)
		}
	} else {
		action.SetOutput("disco_status", "unpatched")
	}
	return nil
}
