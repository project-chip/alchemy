package render

import (
	"context"
	"path/filepath"

	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

type ProvisionalPatcher struct {
	spec    *spec.Specification
	sdkRoot string
}

func NewProvisionalPatcher(sdkRoot string, spec *spec.Specification) *ProvisionalPatcher {
	return &ProvisionalPatcher{sdkRoot: sdkRoot, spec: spec}
}

func (p ProvisionalPatcher) Name() string {
	return "Patching files with provisional clusters and device types"
}

func (p ProvisionalPatcher) Process(cxt context.Context, inputs []*pipeline.Data[struct{}]) (outputs []*pipeline.Data[[]byte], err error) {

	files := make([]string, 0, len(inputs))
	for _, input := range inputs {
		files = append(files, filepath.Base(input.Path))
	}

	var path string
	var value []byte

	path, value, err = patchLintBytes(p.sdkRoot, files)
	if err != nil {
		return
	}
	outputs = append(outputs, pipeline.NewData(path, value))

	path, value, err = patchTestsYamlBytes(p.sdkRoot, files)
	if err != nil {
		return
	}
	outputs = append(outputs, pipeline.NewData(path, value))
	return
}
