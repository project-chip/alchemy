package generate

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"slices"

	"github.com/iancoleman/orderedmap"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/tidwall/pretty"
)

var prettyOptions = pretty.Options{
	Width:    80,
	Prefix:   "",
	Indent:   "    ",
	SortKeys: false,
}

type ZclPatcher struct {
	spec                *spec.Specification
	sdkRoot             string
	provisionalZclFiles pipeline.Paths
}

func NewZclPatcher(sdkRoot string, spec *spec.Specification, provisionalZclFiles pipeline.Paths) *ZclPatcher {
	return &ZclPatcher{sdkRoot: sdkRoot, spec: spec, provisionalZclFiles: provisionalZclFiles}
}

func (p ZclPatcher) Name() string {
	return "Patching ZCL JSON files with clusters"
}

func (p ZclPatcher) Process(cxt context.Context, inputs []*pipeline.Data[*spec.Doc]) (outputs []*pipeline.Data[[]byte], err error) {

	files := make([]string, 0, p.provisionalZclFiles.Size())
	p.provisionalZclFiles.Range(func(key string, value *pipeline.Data[struct{}]) bool {
		files = append(files, filepath.Base(value.Path))
		return true
	})

	clusterMap := make(map[string]*matter.Cluster)
	for _, input := range inputs {
		doc := input.Content
		var entities []types.Entity
		entities, err = doc.Entities()
		if err != nil {
			return
		}
		for _, e := range entities {
			switch e := e.(type) {
			case *matter.Cluster:
				clusterMap[e.Name] = e
			}
		}
	}

	var path string
	var value []byte
	path, value, err = p.patchZapJSONFile(p.sdkRoot, "src/app/zap-templates/zcl/zcl.json", files, clusterMap)
	if err != nil {
		return
	}
	outputs = append(outputs, pipeline.NewData(path, value))

	path, value, err = p.patchZapJSONFile(p.sdkRoot, "src/app/zap-templates/zcl/zcl-with-test-extensions.json", files, clusterMap)
	if err != nil {
		return
	}
	outputs = append(outputs, pipeline.NewData(path, value))
	return
}

func (p *ZclPatcher) patchZapJSONFile(sdkRoot string, file string, files []string, clusterMap map[string]*matter.Cluster) (zclJSONPath string, zclJSONBytes []byte, err error) {
	zclJSONPath = path.Join(sdkRoot, file)
	zclJSONBytes, err = os.ReadFile(zclJSONPath)
	if err != nil {
		return
	}

	o := orderedmap.New()
	err = json.Unmarshal(zclJSONBytes, &o)
	if err != nil {
		return
	}
	val, ok := o.Get("xmlFile")
	if !ok {
		err = fmt.Errorf("missing xmlFile element in %s", zclJSONPath)
		return
	}
	is, ok := val.([]any)
	if !ok {
		err = fmt.Errorf("xmlFile element in %s is not array", zclJSONPath)
		return
	}
	xmls := make([]string, 0, len(is)+len(files))
	fileMap := make(map[string]struct{})
	for _, file := range files {
		fileMap[file] = struct{}{}
	}
	for _, i := range is {
		if s, ok := i.(string); ok {
			xmls = append(xmls, s)
			delete(fileMap, s)
		}
	}

	xmls = mergeLines(xmls, fileMap, 2)

	xmls = slices.Compact(xmls)
	o.Set("xmlFile", xmls)

	p.patchAttributeAccessInterfaceAttributes(o, clusterMap)

	zclJSONBytes, err = json.MarshalIndent(o, "", "    ")
	if err != nil {
		err = fmt.Errorf("error marshaling %s: %w", zclJSONPath, err)
		return
	}
	zclJSONBytes = pretty.PrettyOptions(zclJSONBytes, &prettyOptions)
	return
}

func (p *ZclPatcher) patchAttributeAccessInterfaceAttributes(o *orderedmap.OrderedMap, clusterMap map[string]*matter.Cluster) {
	val, ok := o.Get("attributeAccessInterfaceAttributes")
	if !ok {
		return
	}
	om, ok := val.(orderedmap.OrderedMap)
	if !ok {
		slog.Warn("attributeAccessInterfaceAttributes has unexpected type", log.Type("type", val))
		return
	}
	for _, clusterName := range om.Keys() {
		iav, ok := om.Get(clusterName)
		if !ok {
			slog.Warn("attributeAccessInterfaceAttributes has missing cluster name", slog.String("clusterName", clusterName))
			continue
		}
		ia, ok := iav.([]any)
		if !ok {
			slog.Warn("attributeAccessInterfaceAttributes has unexpected entry type", log.Type("type", iav))
			continue
		}
		cluster, ok := clusterMap[clusterName]
		if !ok {
			slog.Debug("Unknown cluster name in attributeAccessInterfaceAttributes", "name", clusterName)
			continue
		}
		oa := make([]string, 0, len(ia))
		for _, a := range ia {
			as, ok := a.(string)
			if !ok {
				slog.Warn("attributeAccessInterfaceAttributes has unexpected entry type", slog.String("clusterName", clusterName), log.Type("type", a))
			}
			found := false
			for _, ca := range cluster.Attributes {
				if ca.Name == as {
					if conformance.IsZigbee(cluster, ca.Conformance) || conformance.IsDisallowed(ca.Conformance) {
						break
					}

					if matter.NonGlobalIDInvalidForEntity(ca.ID, types.EntityTypeAttribute) {
						break
					}
					found = true
					break
				}
			}
			if found || isGlobalAttributeName(as) {
				oa = append(oa, as)
			}
		}
		om.Set(clusterName, oa)
	}
}

func isGlobalAttributeName(attributeName string) bool {
	switch attributeName {
	case "ClusterRevision", "FeatureMap", "AttributeList", "EventList", "AcceptedCommandList", "GeneratedCommandList":
		return true
	}
	return false
}
