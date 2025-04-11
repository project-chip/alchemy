package python

import (
	"embed"
	"log/slog"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/handlebars"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

//go:embed templates
var templateFiles embed.FS

var template pipeline.Once[*raymond.Template]

func (sp *PythonTestRenderer) loadTemplate(spec *spec.Specification) (*raymond.Template, error) {
	t, err := template.Do(func() (*raymond.Template, error) {

		ov := handlebars.NewOverlay(sp.templateRoot, templateFiles, "templates")
		err := ov.Flush()
		if err != nil {
			slog.Error("Error flushing embedded templates", slog.Any("error", err))
		}
		t, err := handlebars.LoadTemplate("{{> python_test}}", ov)
		if err != nil {
			return nil, err
		}

		handlebars.RegisterCommonHelpers(t)

		registerSpecHelpers(t, spec)

		return t, nil
	})
	if err != nil {
		return nil, err
	}
	return t.Clone(), nil
}
