package templates

import (
	"embed"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/handlebars"
	"github.com/project-chip/alchemy/internal/pipeline"
)

//go:embed disco
var discoTemplateFiles embed.FS

var discoPatchedTemplate pipeline.Once[*raymond.Template]

func LoadDiscoPatchedTemplate() (*raymond.Template, error) {
	t, err := discoPatchedTemplate.Do(func() (*raymond.Template, error) {

		t, err := handlebars.LoadTemplate("{{> disco/patched}}", discoTemplateFiles)
		if err != nil {
			return nil, err
		}

		handlebars.RegisterCommonHelpers(t)

		return t, nil
	})
	if err != nil {
		return nil, err
	}
	return t.Clone(), nil
}

var discoUnpatchedTemplate pipeline.Once[*raymond.Template]

func LoadDiscoUnpatchedTemplate() (*raymond.Template, error) {
	t, err := discoUnpatchedTemplate.Do(func() (*raymond.Template, error) {

		t, err := handlebars.LoadTemplate("{{> disco/unpatched}}", discoTemplateFiles)
		if err != nil {
			return nil, err
		}

		handlebars.RegisterCommonHelpers(t)

		return t, nil
	})
	if err != nil {
		return nil, err
	}
	return t.Clone(), nil
}
