package render

import (
	"embed"
	"log/slog"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/handlebars"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/conformance"
)

//go:embed templates
var templateFiles embed.FS

var template pipeline.Once[*raymond.Template]

type templateContext struct {
	ReferenceStore conformance.ReferenceStore
}

func (sp *Renderer) loadTemplate() (*raymond.Template, error) {
	t, err := template.Do(func() (*raymond.Template, error) {

		ov := handlebars.NewOverlay(sp.templateRoot, templateFiles, "templates")
		err := ov.Flush()
		if err != nil {
			slog.Error("Error flushing embedded templates", slog.Any("error", err))
		}
		t, err := handlebars.LoadTemplate("{{> test_plan}}", ov)
		if err != nil {
			return nil, err
		}

		handlebars.RegisterCommonHelpers(t)

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
		return t, nil
	})
	if err != nil {
		return nil, err
	}
	return t.Clone(), nil
}
