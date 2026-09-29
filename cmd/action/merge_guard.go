package action

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/cmd/action/github"
	"github.com/project-chip/alchemy/cmd/action/github/templates"
	"github.com/project-chip/alchemy/cmd/cli"
	"github.com/project-chip/alchemy/config"
	"github.com/project-chip/alchemy/errdiff"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/mle"
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

	slog.Info("pull request", slog.Int64("id", pr.GetID()))

	base := pr.GetBase()
	if base == nil {
		return fmt.Errorf("pull request missing base")
	}

	slog.Info("base", slog.String("sha", base.GetSHA()))

	head := pr.GetHead()
	if head == nil {
		return fmt.Errorf("pull request missing head")
	}

	slog.Info("head", slog.String("sha", head.GetSHA()))

	var baseRoot, headRoot string
	baseRoot, err = os.MkdirTemp("", "alchemy.base")
	if err != nil {
		return fmt.Errorf("failed on getting temp base dir: %w", err)
	}

	baseRoot, err = github.Checkout(cc, githubContext, action, pr, base.GetRef(), baseRoot)
	if err != nil {
		return fmt.Errorf("failed checking out base: %w", err)
	}

	headRoot = githubContext.Workspace

	var patch bytes.Buffer
	writer := files.NewPatcher[string]("Generating patch file...", &patch)

	specs, err := spec.LoadSpecPullRequest(cc, baseRoot, headRoot, pipelineOptions)
	if err != nil {
		return fmt.Errorf("failed to load specs: %v", err)
	}

	var vp map[string][]spec.Violation
	vp, err = provisional.ProcessSpecs(cc, &specs, pipelineOptions, writer)
	if err != nil {
		return fmt.Errorf("failed checking provisional status: %v", err)
	}

	var ve map[string][]spec.Violation = errdiff.ProcessComparison(&specs)

	var vm map[string][]spec.Violation
	vm, err = mle.Process(headRoot, specs.Head)
	if err != nil {
		return fmt.Errorf("failed checking Master List Enforcer status: %v", err)
	}

	var vev map[string][]spec.Violation = spec.ProcessEventConformanceComparison(&specs)

	violations := spec.MergeViolations(vp, ve, vm, vev)

	owner, repo := githubContext.Repo()

	var comment string
	var violationError error
	if len(violations) > 0 {
		action.SetOutput("merge_guard_status", "violations")
		violationError = errors.New("merge guard violations found")

		err = os.WriteFile("provisional.patch", patch.Bytes(), os.ModeAppend|0644)
		if err != nil {
			return fmt.Errorf("failed saving provisional patch: %v", err)
		}

		var t *raymond.Template
		t, err = templates.LoadMergeGuardViolationsTemplate()
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
			slices.SortFunc(vs, func(a spec.Violation, b spec.Violation) int {
				return a.Line - b.Line
			})

			var relPath string
			relPath, err = filepath.Rel(headRoot, path)
			if err != nil {
				err = fmt.Errorf("the relative path could not be determined. Path: %s, Root: %s", path, headRoot)
				return
			}

			vf := templates.ViolationFile{Path: relPath}
			for _, v := range vs {
				vv := templates.Violation{}

				vv.EntityName, vv.EntityType = getViolationEntity(v)

				pathHash := sha256.Sum256([]byte(relPath))
				vv.SourceLink = fmt.Sprintf("https://github.com/%s/%s/pull/%d/files#diff-%sR%d", owner, repo, pr.GetNumber(), hex.EncodeToString(pathHash[:]), v.Line)
				vv.SourceLine = v.Line
				if v.Type.Has(spec.ViolationTypeNonProvisional) {
					vv.Violations = append(vv.Violations, "Not marked Provisional")
				}
				if v.Type.Has(spec.ViolationTypeNotIfDefd) {
					vv.Violations = append(vv.Violations, "Not in in-progress ifdef")
				}
				if v.Type.Has(spec.ViolationNewParseError) {
					vv.Violations = append(vv.Violations, "New Parse Error introduced by this PR: "+v.Text)
				}
				if v.Type.Has(spec.ViolationMasterList) {
					vv.Violations = append(vv.Violations, "Incompatible with Master List: "+v.Text)
				}
				if v.Type.Has(spec.ViolationEventConformance) {
					vv.Violations = append(vv.Violations, "Invalid event conformance: "+v.Text)
				}
				vf.Violations = append(vf.Violations, vv)
			}
			vc.Files = append(vc.Files, vf)

			annotateViolations(action, path, vs)
		}

		tc := map[string]any{
			"comment": vc,
		}
		comment, err = t.Exec(tc)
		if err != nil {
			return
		}
	} else {
		action.SetOutput("merge_guard_status", "no_violations")

		var t *raymond.Template
		t, err = templates.LoadMergeGuardNoViolationsTemplate()
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
		err = github.WriteComment(cc, githubContext, action, pr, "merge-guard", comment)
		if err != nil {
			return
		}
	} else {
		action.SetOutput("comment", comment)
	}

	if serr := github.WriteSummary(cc, action, comment); serr != nil {
		slog.Error("failed to write summary", "error", serr)
	}
	if violationError != nil {
		return violationError
	}
	return
}

func getViolationEntity(v spec.Violation) (entityName, entityType string) {
	if v.Entity == nil {
		entityName = "-"
		entityType = "-"
		return
	}

	entityName = matter.EntityName(v.Entity)
	entityType = entityTypeName(v.Entity)

	parent := v.Entity.Parent()
	for {
		if parent == nil {
			break
		}
		entityName = matter.EntityName(parent) + "." + entityName
		parent = parent.Parent()
	}
	return
}

func annotateViolations(action *githubactions.Action, path string, violations []spec.Violation) {
	for _, v := range violations {

		entityName, _ := getViolationEntity(v)
		var annotation strings.Builder
		if v.Type.Has(spec.ViolationTypeNonProvisional) {
			fmt.Fprintf(&annotation, "• The entity \"%s\" is newly introduced by this PR, and needs to have provisional conformance.\n", entityName)
		}
		if v.Type.Has(spec.ViolationTypeNotIfDefd) {
			fmt.Fprintf(&annotation, "• The entity \"%s\" must be wrapped in an in-progress ifdef.\n", entityName)
		}
		if v.Type.Has(spec.ViolationNewParseError) {
			annotation.WriteString("• This PR introduces a new parse error.\n")
			annotation.WriteString(v.Text)
			annotation.WriteRune('\n')
		}
		if v.Type.Has(spec.ViolationMasterList) {
			annotation.WriteString("• This PR introduces an incompatibility with the Master List:\n")
			annotation.WriteString(v.Text)
			annotation.WriteRune('\n')
		}
		if v.Type.Has(spec.ViolationEventConformance) {
			annotation.WriteString("• This PR introduces an invalid event conformance:\n")
			annotation.WriteString(v.Text)
			annotation.WriteRune('\n')
		}
		github.Annotate(action, github.AnnotationLevelError, "Merge Guard Violation", annotation.String(), path, v.Line)
	}
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
