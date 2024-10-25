package testplan

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/handlebars"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/conformance"
)

//go:embed templates/*.handlebars
var templateFiles embed.FS

var template pipeline.Once[*raymond.Template]

type templateContext struct {
	ReferenceStore conformance.ReferenceStore
}

func (sp *Generator) loadTemplate() (*raymond.Template, error) {
	t, err := template.Do(func() (*raymond.Template, error) {

		ov := handlebars.NewOverlay(sp.templateRoot, templateFiles)

		files, err := ov.ReadDir("templates")
		if err != nil {
			return nil, err
		}
		t := raymond.MustParse("{{> test_plan}}")
		t.RegisterHelper("raw", handlebars.RawHelper)
		t.RegisterHelper("ifSet", handlebars.IfSetHelper)
		t.RegisterHelper("ifEqual", handlebars.IfEqualHelper)
		t.RegisterHelper("quote", handlebars.QuoteHelper)
		t.RegisterHelper("add", handlebars.AddHelper)
		t.RegisterHelper("brace", handlebars.BraceHelper)
		t.RegisterHelper("format", handlebars.FormatHelper)

		t.RegisterHelper("conformance", conformanceHelper)
		t.RegisterHelper("picsConformance", picsConformanceHelper)
		t.RegisterHelper("constraint", constraintHelper)
		t.RegisterHelper("id", idHelper)
		t.RegisterHelper("shortId", shortIdHelper)
		t.RegisterHelper("entityId", entityIdentifierHelper)
		t.RegisterHelper("entityIdPadded", entityIdentifierPaddedHelper)
		t.RegisterHelper("entityIdPadding", entityIdentifierPaddingHelper)
		t.RegisterHelper("dataType", dataTypeHelper)
		t.RegisterHelper("dataTypeArticle", dataTypeArticleHelper)
		t.RegisterHelper("ifHasQuality", ifHasQualityHelper)
		t.RegisterHelper("ifIsConstraint", ifIsConstraintHelper)
		t.RegisterHelper("limit", limitHelper)
		t.RegisterHelper("ifDataTypeIsArray", ifDataTypeIsArrayHelper)
		t.RegisterHelper("minLimit", minimumHelper)
		t.RegisterHelper("maxLimit", maximumHelper)

		for _, file := range files {
			if file.IsDir() {
				continue
			}
			name := file.Name()
			val, err := fs.ReadFile(ov, filepath.Join("templates/", name))
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
