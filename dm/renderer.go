package dm

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"sync"

	"github.com/iancoleman/orderedmap"
	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/internal/files"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type Renderer struct {
	dmRoot string

	clusters     []*matter.Cluster
	clustersLock sync.Mutex
}

func NewRenderer(dmRoot string) *Renderer {
	return &Renderer{dmRoot: dmRoot}
}

func (p *Renderer) Name() string {
	return "Saving data model"
}

func (p *Renderer) Process(cxt context.Context, input *pipeline.Data[*spec.Doc], index int32, total int32) (outputs []*pipeline.Data[string], extra []*pipeline.Data[*spec.Doc], err error) {
	doc := input.Content
	entites, err := doc.Entities()
	if err != nil {
		slog.ErrorContext(cxt, "error converting doc to entities", "doc", doc.Path, "error", err)
		err = nil
		return
	}
	var appClusters []types.Entity
	var deviceTypes []*matter.DeviceType
	for _, e := range entites {
		switch e := e.(type) {
		case *matter.ClusterGroup, *matter.Cluster:
			appClusters = append(appClusters, e)
		case *matter.DeviceType:
			deviceTypes = append(deviceTypes, e)
		}
	}

	if len(appClusters) == 1 {
		var s string
		s, _, err = p.renderAppCluster(doc, appClusters[0])
		if err != nil {
			err = fmt.Errorf("failed rendering app clusters %s: %w", doc.Path, err)
			return
		}
		outputs = append(outputs, &pipeline.Data[string]{Path: getAppClusterPath(p.dmRoot, doc.Path, ""), Content: s})
	} else if len(appClusters) > 1 {
		for _, e := range appClusters {
			var s string
			var clusterName string
			s, clusterName, err = p.renderAppCluster(doc, e)

			if err != nil {
				err = fmt.Errorf("failed rendering app clusters %s: %w", doc.Path, err)
				return
			}
			if !text.HasCaseInsensitiveSuffix(clusterName, " Cluster") {
				clusterName += " Cluster"
			}
			clusterName = strcase.ToCamel(clusterName)
			outputs = append(outputs, &pipeline.Data[string]{Path: getAppClusterPath(p.dmRoot, doc.Path, clusterName), Content: s})
		}
	}

	if len(deviceTypes) == 1 {
		var s string
		s, err = renderDeviceType(doc, deviceTypes[0])
		if err != nil {
			err = fmt.Errorf("failed rendering device type %s: %w", doc.Path, err)
			return
		}
		outputs = append(outputs, &pipeline.Data[string]{Path: getDeviceTypePath(p.dmRoot, doc.Path, ""), Content: s})
	} else if len(deviceTypes) > 1 {
		for _, dt := range deviceTypes {
			var s string
			s, err = renderDeviceType(doc, dt)
			if err != nil {
				err = fmt.Errorf("failed rendering device types %s: %w", doc.Path, err)
				return
			}
			outputs = append(outputs, &pipeline.Data[string]{Path: getDeviceTypePath(p.dmRoot, doc.Path, strcase.ToCamel(dt.Name)), Content: s})

		}
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

	clusters := make(map[uint64]string)

	path := filepath.Join(p.dmRoot, "/clusters/cluster_ids.json")

	exists, err := files.Exists(path)
	if err != nil {
		return nil, err
	}
	if exists {
		var clusterListBytes []byte
		clusterListBytes, err = os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		var clusterList map[string]any
		err = json.Unmarshal(clusterListBytes, &clusterList)
		if err != nil {
			return nil, err
		}
		for id, name := range clusterList {
			mid := matter.ParseNumber(id)
			if mid.Valid() {
				clusters[mid.Value()] = name.(string)
			}
		}
	}

	p.clustersLock.Lock()
	defer p.clustersLock.Unlock()
	for _, c := range p.clusters {
		if c.ID.Valid() {
			clusters[c.ID.Value()] = c.Name
		}
	}

	var clusterIDs []uint64
	for id := range clusters {
		clusterIDs = append(clusterIDs, id)
	}

	slices.Sort(clusterIDs)
	o := orderedmap.New()
	for _, cid := range clusterIDs {
		name := clusters[cid]
		id := strconv.FormatUint(cid, 10)
		o.Set(id, name)

	}
	b, err := json.MarshalIndent(o, "", "    ")
	if err != nil {
		err = fmt.Errorf("error marshaling cluster ID json: %w", err)
		return nil, err
	}
	return pipeline.NewData(path, string(b)), nil
}
