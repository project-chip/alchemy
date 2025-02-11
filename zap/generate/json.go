package generate

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path"
	"slices"

	"github.com/iancoleman/orderedmap"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/tidwall/pretty"
)

var prettyOptions = pretty.Options{
	Width:    80,
	Prefix:   "",
	Indent:   "    ",
	SortKeys: false,
}

func (p *ProvisionalPatcher) patchZapJSONFile(sdkRoot string, file string, files []string) (zclJSONPath string, zclJSONBytes []byte, err error) {
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

	p.patchAttributeAccessInterfaceAttributes(o)

	zclJSONBytes, err = json.MarshalIndent(o, "", "    ")
	if err != nil {
		err = fmt.Errorf("error marshaling %s: %w", zclJSONPath, err)
		return
	}
	zclJSONBytes = pretty.PrettyOptions(zclJSONBytes, &prettyOptions)
	return
}

func (p *ProvisionalPatcher) patchAttributeAccessInterfaceAttributes(o *orderedmap.OrderedMap) {
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
		cluster, ok := p.spec.ClustersByName[clusterName]
		if !ok {
			slog.Warn("Unknown cluster name in attributeAccessInterfaceAttributes", "name", clusterName)
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
