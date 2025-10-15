package templates

import (
	"embed"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/handlebars"
	"github.com/project-chip/alchemy/internal/pipeline"
)

type Violation struct {
	EntityName string
	EntityType string
	SourceLink string
	SourceLine int
	Violations []string
}

type ViolationFile struct {
	Path       string
	Violations []Violation
}

type ViolationComment struct {
	Files []ViolationFile
}

//go:embed merge_guard
var mergeGuardTemplateFiles embed.FS

var mergeGuardViolationsTemplate pipeline.Once[*raymond.Template]

func LoadMergeGuardViolationsTemplate() (*raymond.Template, error) {
	t, err := mergeGuardViolationsTemplate.Do(func() (*raymond.Template, error) {

		t, err := handlebars.LoadTemplate("{{> merge_guard/violations}}", mergeGuardTemplateFiles)
		if err != nil {
			return nil, err
		}

		handlebars.RegisterCommonHelpers(t)

		//registerSpecHelpers(t, spec)

		return t, nil
	})
	if err != nil {
		return nil, err
	}
	return t.Clone(), nil
}

var mergeGuardNoViolationsTemplate pipeline.Once[*raymond.Template]

func LoadMergeGuardNoViolationsTemplate() (*raymond.Template, error) {
	t, err := mergeGuardNoViolationsTemplate.Do(func() (*raymond.Template, error) {

		t, err := handlebars.LoadTemplate("{{> merge_guard/no_violations}}", mergeGuardTemplateFiles)
		if err != nil {
			return nil, err
		}

		handlebars.RegisterCommonHelpers(t)

		//registerSpecHelpers(t, spec)

		return t, nil
	})
	if err != nil {
		return nil, err
	}
	return t.Clone(), nil
}
