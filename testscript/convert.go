package testscript

import (
	"context"

	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/testplan"
)

type TestScriptConverter struct {
	spec       *spec.Specification
	sdkRoot    string
	picsLabels map[string]string
}

func NewTestScriptConverter(spec *spec.Specification, sdkRoot string, picsLabels map[string]string) *TestScriptConverter {
	return &TestScriptConverter{spec: spec, sdkRoot: sdkRoot, picsLabels: picsLabels}
}

func (sp TestScriptConverter) Name() string {
	return "Converting test plan to test script"
}

func (sp *TestScriptConverter) Process(cxt context.Context, input *pipeline.Data[*testplan.Test], index int32, total int32) (outputs []*pipeline.Data[*Test], extras []*pipeline.Data[*testplan.Test], err error) {
	return
}
