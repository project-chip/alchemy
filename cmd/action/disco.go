package action

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

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

type EnforcementLevel string

const (
	// EnforcementMandatory applies to new files with suggestions or existing discoballed files with suggestions.
	EnforcementMandatory EnforcementLevel = "discoball-mandatory"
	EnforcementOptional  EnforcementLevel = "discoball-optional"
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
	var changedFiles map[string]github.FileStatus
	changedFiles, err = github.GetPRChangedFilesWithStatus(cc, githubContext, action, pr)
	if err != nil {
		return fmt.Errorf("failed on getting pull request changes: %w", err)
	}
	if len(changedFiles) == 0 {
		action.Infof("No changes found\n")
		action.SetOutput("disco_status", "noop")
		return nil
	}

	var changedDocs []string
	fileEnforcementLevel := make(map[string]EnforcementLevel)
	for path, status := range changedFiles {
		if filepath.Ext(path) == ".adoc" {
			changedDocs = append(changedDocs, path)
			if status == github.FileStatusAdded {
				fileEnforcementLevel[path] = EnforcementMandatory
				continue
			}
			fullPath := filepath.Join(githubContext.Workspace, path)
			b, err := os.ReadFile(fullPath)
			if err != nil {
				slog.Warn("failed to read original file to check for discoballed marker", "path", fullPath, "error", err)
				fileEnforcementLevel[path] = EnforcementOptional
				continue
			}
			if strings.Contains(string(b), "\n:alchemy-discoballed:") || strings.HasPrefix(string(b), ":alchemy-discoballed:") {
				fileEnforcementLevel[path] = EnforcementMandatory
			} else {
				fileEnforcementLevel[path] = EnforcementOptional
			}
		}
	}

	if len(changedDocs) == 0 {
		action.Infof("No changed asciidoc files found\n")
		action.SetOutput("disco_status", "noop")
		return nil
	}

	pipelineOptions := pipeline.ProcessingOptions{NoProgress: true}

	var outMandatory bytes.Buffer
	var outOptional bytes.Buffer
	writer := files.NewPatcher[string]("Generating patch file...", &outMandatory)
	writer.GetWriter = func(path string) io.Writer {
		if fileEnforcementLevel[path] == EnforcementMandatory {
			return &outMandatory
		}
		return &outOptional
	}
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

	hasMandatory := outMandatory.Len() > 0
	hasOptional := outOptional.Len() > 0

	if hasMandatory {
		action.SetOutput("has_violations", "true")
		var violations []string
		for _, path := range writer.ModifiedFiles {
			if fileEnforcementLevel[path] == EnforcementMandatory {
				status := changedFiles[path]
				var msg string
				if status == github.FileStatusAdded {
					msg = fmt.Sprintf("a new file is added to PR and it has alchemy discoball suggestions: %s", path)
				} else {
					msg = fmt.Sprintf("a file that already had :alchemy-discoballed: has fixes suggested: %s", path)
				}
				violations = append(violations, msg)
			}
		}
		for _, v := range violations {
			action.Errorf("%s", v)
		}
		action.SetOutput("violation_reason", fmt.Sprintf("Found %d files with violations. See logs for details.", len(violations)))
	} else {
		action.SetOutput("has_violations", "false")
	}

	if hasMandatory {
		slog.Info("Setting mandatory patch outputs")
		err = os.WriteFile("disco-mandatory.patch", outMandatory.Bytes(), 0644)
		if err != nil {
			return fmt.Errorf("failed saving mandatory patch: %v", err)
		}
		action.SetOutput("mandatory_patch_name", "disco_mandatory_patch")
		action.SetOutput("mandatory_patch_path", "disco-mandatory.patch")
	}

	if hasOptional {
		slog.Info("Setting optional patch outputs")
		err = os.WriteFile("disco-optional.patch", outOptional.Bytes(), 0644)
		if err != nil {
			return fmt.Errorf("failed saving optional patch: %v", err)
		}
		action.SetOutput("optional_patch_name", "disco_optional_patch")
		action.SetOutput("optional_patch_path", "disco-optional.patch")
	}

	if hasMandatory || hasOptional {
		action.SetOutput("template_name", "disco/combined")
	} else {
		slog.Info("Setting disco_status to unpatched")
		action.SetOutput("disco_status", "unpatched")
		action.SetOutput("template_name", "disco/unpatched")
	}

	return nil
}
