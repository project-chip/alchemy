package github

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/hasty/alchemy/disco"
	"github.com/hasty/alchemy/internal/files"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/sethvargo/go-githubactions"
)

func Action() error {
	cxt := context.Background()

	action := githubactions.New()

	githubContext, err := githubactions.Context()
	if err != nil {
		return fmt.Errorf("failed on getting GitHub context: %w", err)

	}
	pr, err := readPullRequest(cxt, githubContext, action)
	if err != nil {
		return fmt.Errorf("failed on reading pull request: %w", err)
	}
	if pr == nil {
		return nil
	}
	var changedFiles []string
	changedFiles, err = getPRChangedFiles(cxt, githubContext, action, pr)
	if err != nil {
		return fmt.Errorf("failed on getting pull request changes: %w", err)
	}
	if len(changedFiles) == 0 {
		action.Infof("No changes found\n")
		return nil
	}

	targeter := files.PathsTargeter(changedFiles...)

	pipelineOptions := pipeline.Options{NoProgress: true}

	var out bytes.Buffer
	writer := files.NewPatcher[string]("Generating patch file...", &out)

	err = disco.Pipeline(cxt, targeter, pipelineOptions, nil, nil, writer)
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
