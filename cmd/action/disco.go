package action

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/cmd/action/github"
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/config"
	"github.com/project-chip/alchemy/disco"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/sethvargo/go-githubactions"
)

type Disco struct {
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
	writer.Root = githubContext.Workspace

	parserOptions := spec.ParserOptions{
		Root: githubContext.Workspace,
	}

	attributes := []asciidoc.AttributeName{"in-progress"}

	err = disco.Pipeline(cc, parserOptions, changedDocs, pipelineOptions, disco.DefaultOptions, attributes, nil, writer)

	if err != nil {
		var message string
		switch err := err.(type) {
		case spec.Error:
			path, line := err.Origin()
			message = fmt.Sprintf("%s (%s:%d)", err.Error(), path, line)
		default:
			message = err.Error()
		}
		return fmt.Errorf("failed disco-balling: %s", message)
	}

	if out.Len() > 0 {
		slog.Info("Setting disco_status to patched")
		action.SetOutput("disco_status", "patched")

		err = os.WriteFile("disco.patch", out.Bytes(), os.ModeAppend|0644)
		if err != nil {
			return fmt.Errorf("failed saving patch: %v", err)
		}
		action.SetOutput("patch_name", "disco_patch")
		action.SetOutput("patch_path", "disco.patch")
		action.SetOutput("template_name", "disco/patched")
	} else {
		slog.Info("Setting disco_status to unpatched")
		action.SetOutput("disco_status", "unpatched")
		action.SetOutput("template_name", "disco/unpatched")
	}

	return nil
}
