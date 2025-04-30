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

//go:embed provisional
var provisionalTemplateFiles embed.FS

var provisionalViolationsTemplate pipeline.Once[*raymond.Template]

func LoadProvisionalViolationsTemplate() (*raymond.Template, error) {
	t, err := provisionalViolationsTemplate.Do(func() (*raymond.Template, error) {

		t, err := handlebars.LoadTemplate("{{> provisional/violations}}", provisionalTemplateFiles)
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

var provisionalNoViolationsTemplate pipeline.Once[*raymond.Template]

func LoadProvisionalNoViolationsTemplate() (*raymond.Template, error) {
	t, err := provisionalNoViolationsTemplate.Do(func() (*raymond.Template, error) {

		t, err := handlebars.LoadTemplate("{{> provisional/no_violations}}", provisionalTemplateFiles)
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
