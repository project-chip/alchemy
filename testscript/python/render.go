package python

import (
	"context"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/testscript"
)

type PythonTestRenderer struct {
	sdkRoot      string
	templateRoot string
	overwrite    bool

	spec       *spec.Specification
	picsLabels map[string]string
}

func NewPythonTestRenderer(spec *spec.Specification, sdkRoot string, picsLabels map[string]string, options ...GeneratorOption) *PythonTestRenderer {
	ptg := &PythonTestRenderer{spec: spec, sdkRoot: sdkRoot, picsLabels: picsLabels}
	for _, o := range options {
		o(ptg)
	}
	return ptg
}

func (sp PythonTestRenderer) Name() string {
	return "Rendering test scripts"
}

func (sp *PythonTestRenderer) Process(cxt context.Context, input *pipeline.Data[*testscript.Test], index int32, total int32) (outputs []*pipeline.Data[string], extras []*pipeline.Data[*testscript.Test], err error) {

	test := input.Content

	if test.Cluster == nil {
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
	outputs = append(outputs, pipeline.NewData(input.Path, out))
	return
}
