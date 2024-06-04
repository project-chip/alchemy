package dm

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"path/filepath"
	"slices"
	"sync"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/pipeline"
	"github.com/hasty/alchemy/matter"
	"github.com/iancoleman/orderedmap"
)

type Renderer struct {
	sdkRoot string

	clusters     []*matter.Cluster
	clustersLock sync.Mutex
}

func NewRenderer(sdkRoot string) *Renderer {
	return &Renderer{sdkRoot: sdkRoot}
}

func (p *Renderer) Name() string {
	return "Saving data model"
}

func (p *Renderer) Type() pipeline.ProcessorType {
	return pipeline.ProcessorTypeIndividual
}

func (p *Renderer) Process(cxt context.Context, input *pipeline.Data[*ascii.Doc], index int32, total int32) (outputs []*pipeline.Data[string], extra []*pipeline.Data[*ascii.Doc], err error) {
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
		s, err = renderAppCluster(cxt, doc, appClusters)
		if err != nil {
			err = fmt.Errorf("failed rendering app clusters %s: %w", doc.Path, err)
			return
		}
		p.clustersLock.Lock()
		p.clusters = append(p.clusters, appClusters...)
		p.clustersLock.Unlock()
		outputs = append(outputs, &pipeline.Data[string]{Path: getAppClusterPath(p.sdkRoot, doc.Path), Content: s})
	}
	if len(deviceTypes) > 0 {
		var s string
		s, err = renderDeviceType(cxt, doc, deviceTypes)
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

func (p *Renderer) GenerateClusterIDsJson() (*pipeline.Data[string], error) {
	p.clustersLock.Lock()
	defer p.clustersLock.Unlock()
	slices.SortFunc(p.clusters, func(a *matter.Cluster, b *matter.Cluster) int {
		return a.ID.Compare(b.ID)
	})
	o := orderedmap.New()
	for _, c := range p.clusters {
		if c.ID.Valid() {
			o.Set(c.ID.IntString(), c.Name)
		}
	}
	b, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		err = fmt.Errorf("error marshaling cluster ID json: %w", err)
		return nil, err
	}
	path := filepath.Join(p.sdkRoot, "/data_model/clusters/cluster_ids.json")
	return pipeline.NewData(path, string(b)), nil
}
