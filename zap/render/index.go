package render

import (
	"context"
	"path/filepath"

	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter/spec"
)

type IndexFilesPatcher struct {
	spec    *spec.Specification
	sdkRoot string
}

func NewIndexFilesPatcher(sdkRoot string, spec *spec.Specification) *IndexFilesPatcher {
	return &IndexFilesPatcher{sdkRoot: sdkRoot, spec: spec}
}

func (ifp IndexFilesPatcher) Name() string {
	return "Patching index files with clusters and device types"
}

func (ifp IndexFilesPatcher) Process(cxt context.Context, inputs []*pipeline.Data[string]) (outputs []*pipeline.Data[[]byte], err error) {

	files := make([]string, 0, len(inputs))
	for _, input := range inputs {
		files = append(files, filepath.Base(input.Path))
	}

	var path string
	var value []byte

	path, value, err = patchLintBytes(ifp.sdkRoot, files)
	if err != nil {
		return
	}
	outputs = append(outputs, pipeline.NewData(path, value))

	path, value, err = patchTestsYamlBytes(ifp.sdkRoot, files)
	if err != nil {
		return
	}
	outputs = append(outputs, pipeline.NewData(path, value))
	return
}
