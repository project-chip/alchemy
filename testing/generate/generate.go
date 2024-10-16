package generate

import (
	"context"
	"log/slog"
	"path/filepath"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/testing/parse"
)

type PythonTestGenerator struct {
	sdkRoot   string
	overwrite bool

	spec   *spec.Specification
	labels map[string]string
}

func NewPythonTestGenerator(spec *spec.Specification, sdkRoot string, overwrite bool, labels map[string]string) *PythonTestGenerator {
	return &PythonTestGenerator{spec: spec, sdkRoot: sdkRoot, overwrite: overwrite, labels: labels}
}

func (sp PythonTestGenerator) Name() string {
	return "Generating test plans"
}

func (sp PythonTestGenerator) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (sp *PythonTestGenerator) Process(cxt context.Context, input *pipeline.Data[*parse.Test], index int32, total int32) (outputs []*pipeline.Data[string], extras []*pipeline.Data[*parse.Test], err error) {

	outPath := sp.getPath(input.Path)
	slog.Info("generating", "in", input.Path, "out", outPath)

	var test *test
	test, err = sp.convert(input.Content, input.Path)
	if err != nil {
		slog.WarnContext(cxt, "Error converting test to model", slog.String("path", input.Path), slog.Any("error", err))
		err = nil
		return
	}

	if test.Config.Cluster == "" {
		return
	}

	var t *raymond.Template
	t, err = loadTemplate(sp.spec)
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
