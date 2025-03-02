package python

import (
	"context"
	"log/slog"
	"path/filepath"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/testplan"
)

type GeneratorOption func(g *PythonTestGenerator)
type PythonTestGenerator struct {
	sdkRoot      string
	templateRoot string
	overwrite    bool

	spec       *spec.Specification
	picsLabels map[string]string
}

func TemplateRoot(templateRoot string) func(*PythonTestGenerator) {
	return func(g *PythonTestGenerator) {
		g.templateRoot = templateRoot
	}
}

func Overwrite(overwrite bool) func(*PythonTestGenerator) {
	return func(g *PythonTestGenerator) {
		g.overwrite = overwrite
	}
}

func NewPythonTestGenerator(spec *spec.Specification, sdkRoot string, picsLabels map[string]string, options ...GeneratorOption) *PythonTestGenerator {
	ptg := &PythonTestGenerator{spec: spec, sdkRoot: sdkRoot, picsLabels: picsLabels}
	for _, o := range options {
		o(ptg)
	}
	return ptg
}

func (sp PythonTestGenerator) Name() string {
	return "Generating test plans"
}

func (sp *PythonTestGenerator) Process(cxt context.Context, input *pipeline.Data[*testplan.Test], index int32, total int32) (outputs []*pipeline.Data[string], extras []*pipeline.Data[*testplan.Test], err error) {

	outPath := sp.getPath(input.Path)
	slog.Info("generating", "in", input.Path, "out", outPath)

	test := input.Content

	if test.Config.Cluster == "" {
		return
	}

	var t *raymond.Template
	t, err = sp.loadTemplate(sp.spec)
	if err != nil {
		return
	}
	variables := make(map[string]struct{})
	t.RegisterHelper("variable", variableHelper(variables))
	t.RegisterHelper("value", valueHelper(variables))
	tc := map[string]any{
		"test": test,
	}
	var out string
	out, err = t.Exec(tc)
	if err != nil {
		return
	}
	outputs = append(outputs, pipeline.NewData(outPath, out))
	return
}

func (sp *PythonTestGenerator) getPath(path string) string {

	path = getTestName(path)
	path += ".py"
	return filepath.Join(sp.sdkRoot, "src/python_testing", path)
}

func getTestName(path string) string {
	path = filepath.Base(path)
	path = text.TrimCaseInsensitivePrefix(path, "test_")
	path = text.TrimCaseInsensitiveSuffix(path, ".yaml")
	return path
}
