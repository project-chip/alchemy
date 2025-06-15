package templates

import (
	"fmt"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/handlebars"
	"github.com/project-chip/alchemy/internal/pipeline"
)

var templates = make(map[string]*pipeline.Once[*raymond.Template])

func LoadTemplate(path string) (*raymond.Template, error) {
	te, ok := templates[path]
	if !ok {
		te = &pipeline.Once[*raymond.Template]{}
		templates[path] = te
	}
	t, err := te.Do(func() (*raymond.Template, error) {

		t, err := handlebars.LoadTemplate(fmt.Sprintf("{{> %s}}", path), discoTemplateFiles)
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
