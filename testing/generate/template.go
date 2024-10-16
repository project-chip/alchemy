package generate

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

//go:embed templates/*.handlebars
var templateFiles embed.FS

var template pipeline.Once[*raymond.Template]

func loadTemplate(spec *spec.Specification) (*raymond.Template, error) {
	t, err := template.Do(func() (*raymond.Template, error) {
		files, err := templateFiles.ReadDir("templates")
		if err != nil {
			return nil, err
		}
		t := raymond.MustParse("{{> python_test}}")
		registerHelpers(t, spec)

		for _, file := range files {
			if file.IsDir() {
				continue
			}
			name := file.Name()
			val, err := templateFiles.ReadFile(filepath.Join("templates/", name))
			if err != nil {
				return nil, fmt.Errorf("error reading template file %s: %w", name, err)
			}
			p, err := raymond.Parse(string(val))
			if err != nil {
				return nil, fmt.Errorf("error parsing template file %s: %w", name, err)
			}
			t.RegisterPartialTemplate(strings.TrimSuffix(filepath.Base(name), filepath.Ext(name)), p)
		}
		return t, nil
	})
	if err != nil {
		return nil, err
	}
	return t.Clone(), nil
}
