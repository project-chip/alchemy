package generate

import (
	"context"

	"github.com/hasty/alchemy/internal/pipeline"
)

type ProvisionalPatcher struct {
	sdkRoot string
}

func NewProvisionalPatcher(sdkRoot string) *ProvisionalPatcher {
	return &ProvisionalPatcher{sdkRoot: sdkRoot}
}

func (p ProvisionalPatcher) Name() string {
	return "Patching files with provisional clusters and device types"
}

func (p ProvisionalPatcher) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeCollective
}

func (p ProvisionalPatcher) Process(cxt context.Context, inputs []*pipeline.Data[struct{}]) (outputs []*pipeline.Data[[]byte], err error) {

	files := make([]string, 0, len(inputs))
	for _, input := range inputs {
		files = append(files, input.Path)
	}

	var path string
	var value []byte
	path, value, err = patchZapJsonFile(p.sdkRoot, "src/app/zap-templates/zcl/zcl.json", files)
	if err != nil {
		return
	}
	outputs = append(outputs, pipeline.NewData[[]byte](path, value))
	path, value, err = patchZapJsonFile(p.sdkRoot, "src/app/zap-templates/zcl/zcl-with-test-extensions.json", files)
	if err != nil {
		return
	}
	outputs = append(outputs, pipeline.NewData[[]byte](path, value))

	path, value, err = patchLintBytes(p.sdkRoot, files)
	if err != nil {
		return
	}
	outputs = append(outputs, pipeline.NewData[[]byte](path, value))

	path, value, err = patchTestsYamlBytes(p.sdkRoot, files)
	if err != nil {
		return
	}
	outputs = append(outputs, pipeline.NewData[[]byte](path, value))
	return
}
