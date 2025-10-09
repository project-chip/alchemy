package action

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/cmd/action/github"
	"github.com/project-chip/alchemy/cmd/action/github/templates"
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/config"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/provisional"
	"github.com/sethvargo/go-githubactions"
)

type MergeGuard struct {
	WriteComment bool `default:"false" hidden:"" help:"Write comment directly"`
}

func (c *MergeGuard) Run(cc *cli.Context) (err error) {

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
	action.Infof("Workspace: %s", githubContext.Workspace)

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

	headRoot = githubContext.Workspace

	_ = pipelineOptions
	_ = headRoot

	var out bytes.Buffer
	writer := files.NewPatcher[string]("Generating patch file...", &out)

	specs, err := spec.LoadSpecSet(cc, baseRoot, headRoot, changedDocs, pipelineOptions, nil)
	if err != nil {
		return fmt.Errorf("failed to load specs: %v", err)
	}

	var violations map[string][]provisional.Violation
	violations, err = provisional.ProcessSpecs(cc, &specs, pipelineOptions, writer)
	if err != nil {
		return fmt.Errorf("failed checking provisional status: %v", err)
	}

	owner, repo := githubContext.Repo()

	var comment string
	if len(violations) > 0 {
		action.SetOutput("provisional_status", "violations")

		err = os.WriteFile("provisional.patch", out.Bytes(), os.ModeAppend|0644)
		if err != nil {
			return fmt.Errorf("failed saving provisional patch: %v", err)
		}

		var t *raymond.Template
		t, err = templates.LoadProvisionalViolationsTemplate()
		if err != nil {
			err = fmt.Errorf("error loading violation template: %w", err)
			return
		}

		vc := templates.ViolationComment{}

		var paths []string
		for path := range violations {
			paths = append(paths, path)
		}

		slices.Sort(paths)

		for _, path := range paths {
			vs := violations[path]
			slices.SortFunc(vs, func(a provisional.Violation, b provisional.Violation) int {
				return a.Line - b.Line
			})
			vf := templates.ViolationFile{Path: path}
			for _, v := range vs {
				vv := templates.Violation{}
				vv.EntityName = matter.EntityName(v.Entity)
				vv.EntityType = entityTypeName(v.Entity)

				parent := v.Entity.Parent()
				for {
					if parent == nil {
						break
					}
					vv.EntityName = matter.EntityName(parent) + "." + vv.EntityName
					parent = parent.Parent()
				}

				pathHash := sha256.Sum256([]byte(path))
				vv.SourceLink = fmt.Sprintf("https://github.com/%s/%s/pull/%d/files#diff-%sR%d", owner, repo, pr.GetNumber(), hex.EncodeToString(pathHash[:]), v.Line)
				vv.SourceLine = v.Line
				if v.Type.Has(provisional.ViolationTypeNonProvisional) {
					vv.Violations = append(vv.Violations, "Not marked Provisional")
				}
				if v.Type.Has(provisional.ViolationTypeNotIfDefd) {
					vv.Violations = append(vv.Violations, "Not in in-progress ifdef")
				}
				vf.Violations = append(vf.Violations, vv)
			}
			vc.Files = append(vc.Files, vf)
		}

		tc := map[string]any{
			"comment": vc,
		}
		comment, err = t.Exec(tc)
		if err != nil {
			return
		}
	} else {
		action.SetOutput("provisional_status", "no_violations")

		var t *raymond.Template
		t, err = templates.LoadProvisionalNoViolationsTemplate()
		if err != nil {
			err = fmt.Errorf("error loading no violation template: %w", err)
			return
		}
		comment, err = t.Exec(map[string]any{})
		if err != nil {
			return
		}
	}

	if c.WriteComment {
		err = github.WriteComment(cc, githubContext, action, pr, "disco-ball", comment)
		if err != nil {
			return
		}
	} else {
		action.SetOutput("comment", comment)
	}
	return nil
}

func entityTypeName(e types.Entity) string {
	switch e := e.(type) {
	case *matter.Struct:
		return "Struct"
	case *matter.Feature:
		return "Feature"
	case *matter.Field:
		switch e.EntityType() {
		case types.EntityTypeAttribute:
			return "Attribute"
		case types.EntityTypeEventField:
			return "Event Field"
		case types.EntityTypeCommandField:
			return "Command Field"
		case types.EntityTypeStructField:
			return "Struct Field"
		default:
			return "Field"
		}
	case *matter.Bitmap:
		return "Bitmap"
	case *matter.Enum:
		return "Enum"
	case *matter.EnumValue:
		return "Enum Value"
	case matter.Bit:
		return "Bit"
	default:
		return e.EntityType().String()
	}
}
