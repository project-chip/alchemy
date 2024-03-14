package dm

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hasty/alchemy/matter"
)

type Renderer struct {
	sdkRoot string
}

func NewRenderer(sdkRoot string) *Renderer {
	return &Renderer{sdkRoot: sdkRoot}
}

func (p Renderer) Name() string {
	return "Saving data model"
}

func (p Renderer) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (p Renderer) Process(cxt context.Context, input *pipeline.Data[*ascii.Doc], index int32, total int32) (outputs []*pipeline.Data[string], extra []*pipeline.Data[*ascii.Doc], err error) {
	doc := input.Content
	entites, err := doc.Entities()
	if err != nil {
		slog.ErrorContext(cxt, "error converting doc to entities", "doc", doc.Path, "error", err)
		err = nil
		return
	}
	var appClusters []*matter.Cluster
	var deviceTypes []*matter.DeviceType
	for _, e := range entites {
		switch e := e.(type) {
		case *matter.Cluster:
			appClusters = append(appClusters, e)
		case *matter.DeviceType:
			deviceTypes = append(deviceTypes, e)
		}
	}
	if len(appClusters) > 0 {
		var s string
		s, err = renderAppCluster(cxt, appClusters)
		if err != nil {
			err = fmt.Errorf("failed rendering app clusters %s: %w", doc.Path, err)
			return
		}
		outputs = append(outputs, &pipeline.Data[string]{Path: getAppClusterPath(p.sdkRoot, doc.Path), Content: s})
	}
	if len(deviceTypes) > 0 {
		var s string
		s, err = renderDeviceType(cxt, deviceTypes)
		if err != nil {
			err = fmt.Errorf("failed rendering device types %s: %w", doc.Path, err)
			return
		}
		outputs = append(outputs, &pipeline.Data[string]{Path: getDeviceTypePath(p.sdkRoot, doc.Path), Content: s})
	}
	for _, o := range outputs {
		o.Content, err = patchLicense(o.Content, o.Path)
		if err != nil {
			err = fmt.Errorf("error patching license for %s: %w", o.Path, err)
			return
		}
	}
	return
}
