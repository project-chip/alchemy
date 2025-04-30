package action

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/project-chip/alchemy/cmd/action/github"
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/config"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/provisional"
	"github.com/sethvargo/go-githubactions"
)

type Provisional struct {
}

func (c *Provisional) Run(cc *cli.Context) (err error) {

	action := githubactions.New()

	action.Infof("Alchemy %s", config.Version())

	var workingDir string
	workingDir, err = os.Getwd()
	if err != nil {
		return fmt.Errorf("failed on getting working directory: %w", err)
	}
	action.Infof("Working directory: %s", workingDir)

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
		return fmt.Errorf("empty pull request")
	}

	pr, err = github.GetPR(cc, githubContext, action, pr)
	if err != nil {
		return fmt.Errorf("failed on getting pull request: %w", err)
	}

	var changedFiles []string
	changedFiles, err = github.GetPRChangedFiles(cc, githubContext, action, pr)
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

	pipelineOptions := pipeline.ProcessingOptions{NoProgress: true}

	slog.Info("pull request", "pr", pr)
	slog.Info("pull request", "head", pr.GetHead())
	slog.Info("pull request", "base", pr.GetBase())

	var baseRoot, headRoot string
	baseRoot, err = os.MkdirTemp("", "alchemy.base")
	if err != nil {
		return fmt.Errorf("failed on getting temp base dir: %w", err)
	}

	baseRoot, err = github.Checkout(cc, githubContext, action, pr, pr.GetBase().GetRef(), baseRoot)
	if err != nil {
		return fmt.Errorf("failed checking out base: %w", err)
	}

	headRoot = "."

	_ = pipelineOptions
	_ = headRoot

	var violations *provisional.Violations
	violations, err = provisional.Pipeline(cc, baseRoot, headRoot, changedDocs, pipelineOptions, nil, nil)
	if err != nil {
		return fmt.Errorf("failed checking provisional status: %v", err)
	}

	sha256.Sum256([]byte(""))

	owner, repo := githubContext.Repo()

	var comment string
	if len(violations.Set) > 0 {
		action.SetOutput("provisional_status", "violations")
		var sb strings.Builder
		sb.WriteString("<img align=\"left\" width=\"75\" height=\"75\" src=\"https://raw.githubusercontent.com/project-chip/alchemy/refs/heads/main/alchemy.svg\">\n")
		sb.WriteString(`<b><a href="https://github.com/project-chip/alchemy">Alchemy</a> found provisionality violations in this pull request</b>`)
		sb.WriteString("\n\n")
		sb.WriteString("---\n\n")

		vmap := make(map[string][]provisional.Violation)
		var paths []string
		for _, v := range violations.Set {
			_, ok := vmap[v.Path]
			if !ok {
				paths = append(paths, v.Path)
			}
			vmap[v.Path] = append(vmap[v.Path], v)
		}

		slices.Sort(paths)
		sb.WriteString("<table>\n")
		for _, path := range paths {
			vs := vmap[path]
			slices.SortFunc(vs, func(a provisional.Violation, b provisional.Violation) int {
				return a.Line - b.Line
			})
			sb.WriteString("<tr><td colspan=\"4\"><b>")
			sb.WriteString(path)
			sb.WriteString("</b></td></tr>\n")
			sb.WriteString("<tr><td>Name</td><td>Type</td><td>Source</td><td>Violation</td></tr>\n")
			for _, v := range vs {
				sb.WriteString("<tr>")
				sb.WriteString(fmt.Sprintf("<td>%s</td>", matter.EntityName(v.Entity)))
				sb.WriteString(fmt.Sprintf("<td>%s</td>", v.Entity.EntityType().String()))
				pathHash := sha256.Sum256([]byte(path))
				link := fmt.Sprintf("https://github.com/%s/%s/pull/%d/files#diff-%sR%d", owner, repo, pr.GetNumber(), hex.EncodeToString(pathHash[:]), v.Line)
				sb.WriteString(fmt.Sprintf("<td><a href=\"%s\">Line %d</a> </td>", link, v.Line))
				sb.WriteString(fmt.Sprintf("<td>%s</td>", v.Type.String()))
				sb.WriteString("</tr>\n")
			}
		}
		sb.WriteString("</table>\n\n")
		sb.WriteString("> [!CAUTION]\n")
		sb.WriteString("> These issues must be resolved before this PR can be merged.\n\n")

		comment = sb.String()
	} else {
		action.SetOutput("provisional_status", "no_violations")
		comment = "✨ Provisional violations have been resolved."
	}

	err = github.WriteComment(cc, githubContext, action, pr, "provisional-violations", comment)
	if err != nil {
		return
	}
	return nil
}
