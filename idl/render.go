package idl

import (
	"context"
	"embed"
	"log/slog"
	"path/filepath"
	"slices"
	"strings"
	"unicode"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/handlebars"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/provisional"
)

//go:embed templates
var templateFiles embed.FS

type IdlRenderer struct {
	spec *spec.Specification

	commonAttributes matter.FieldSet

	SuppressEndpoints   bool
	SuppressProvisional string
	PerTrait            bool

	provisionalFilter ProvisionalFilter
}

func NewIdlRenderer(spec *spec.Specification) (IdlRenderer, error) {
	renderer := IdlRenderer{spec: spec}
	renderer.commonAttributes = matter.FieldSet{
		&matter.Field{
			Name:        "GeneratedCommandList",
			ID:          matter.NewNumber(65528),
			Type:        types.NewDataType(types.BaseDataTypeCommandID, types.DataTypeRankList),
			Access:      matter.Access{Read: matter.PrivilegeView},
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
		&matter.Field{
			Name:        "AcceptedCommandList",
			ID:          matter.NewNumber(65529),
			Type:        types.NewDataType(types.BaseDataTypeCommandID, types.DataTypeRankList),
			Access:      matter.Access{Read: matter.PrivilegeView},
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
		&matter.Field{
			Name:        "AttributeList",
			ID:          matter.NewNumber(65531),
			Type:        types.NewDataType(types.BaseDataTypeAttributeID, types.DataTypeRankList),
			Access:      matter.Access{Read: matter.PrivilegeView},
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
		&matter.Field{
			Name:        "FeatureMap",
			ID:          matter.NewNumber(65532),
			Type:        types.NewDataType(types.BaseDataTypeMap32, types.DataTypeRankScalar),
			Access:      matter.Access{Read: matter.PrivilegeView},
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
		&matter.Field{
			Name:        "ClusterRevision",
			ID:          matter.NewNumber(65533),
			Type:        types.NewDataType(types.BaseDataTypeUInt16, types.DataTypeRankScalar),
			Access:      matter.Access{Read: matter.PrivilegeView},
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
	}
	return renderer, nil
}

func (p IdlRenderer) Name() string {
	return "Writing Matter files"
}

func (p IdlRenderer) Process(cxt context.Context, input *pipeline.Data[*File], index int32, total int32) (outputs []*pipeline.Data[string], extras []*pipeline.Data[*File], err error) {

	dir := filepath.Dir(input.Path)
	base := filepath.Base(input.Path)
	extension := filepath.Ext(base)
	file := strings.TrimSuffix(base, extension)
	path := filepath.Join(dir, file+".matter")

	slog.Info("converting zap path", "path", input.Path, "matter", path)

	filter := ProvisionalFilter{
		Mode: p.SuppressProvisional,
	}
	if p.SuppressProvisional == "keep-existing" {
		elements, err := parseExistingMatterElements(path)
		if err != nil {
			slog.Warn("Failed to parse existing matter file for provisional elements", slog.String("path", path), slog.Any("err", err))
		} else {
			filter.ExistingElements = elements
		}
	}
	p.provisionalFilter = filter

	var endpoints []Endpoint
	clusters := make(map[*matter.Cluster]*ClusterInfo)

	for _, endpoint := range input.Content.Endpoints {
		if endpoint.EndpointTypeIndex < 0 || endpoint.EndpointTypeIndex >= len(input.Content.EndpointTypes) {
			continue
		}
		endpointType := input.Content.EndpointTypes[endpoint.EndpointTypeIndex]

		deviceType, ok := p.spec.DeviceTypesByID[uint64(endpointType.DeviceTypeCode)]
		if !ok {
			continue
		}
		ep := Endpoint{ID: endpoint.EndpointId, EndpointType: endpointType}
		ep.DeviceType = deviceType
		for _, clusterRef := range ep.Clusters {
			c, ok := p.spec.ClustersByID[uint64(clusterRef.Code)]
			if !ok {
				slog.Warn("Unrecognized cluster code", slog.String("path", input.Path), slog.Int("clusterCode", clusterRef.Code))
				continue
			}
			if !entityShouldBeIncluded(p.spec, p.provisionalFilter, c) {
				continue
			}
			ci := &ClusterInfo{Cluster: c}
			clusters[c] = ci
			switch clusterRef.Side {
			case "server":
				ep.Servers = append(ep.Servers, ci)
			case "client":
				ep.Clients = append(ep.Clients, ci)
			}
		}
		endpoints = append(endpoints, ep)
	}

	clusterList := make([]*ClusterInfo, 0, len(clusters))
	clusterEntities := make(map[*matter.Cluster]map[types.Entity]struct{})
	globalEntities := make(map[types.Entity]bool)
	for entity := range p.spec.GlobalObjects {
		globalEntities[entity] = false
	}
	for cluster, ci := range clusters {
		clusterList = append(clusterList, ci)
		clusterEntities[cluster] = make(map[types.Entity]struct{})
	}

	slices.SortFunc(clusterList, func(a *ClusterInfo, b *ClusterInfo) int {
		return a.Cluster.ID.Compare(b.Cluster.ID)
	})

	var globalEnums []*matter.Enum
	var globalStructs []*matter.Struct
	var globalBitmaps []*matter.Bitmap
	namespaces := make(map[string]*matter.Namespace)

	spec.TraverseEntities(p.spec, func(parentCluster *matter.Cluster, parent, entity types.Entity) parse.SearchShould {

		ce, ok := clusterEntities[parentCluster]
		if !ok {
			return parse.SearchShouldSkip
		}

		if !entityShouldBeIncluded(p.spec, p.provisionalFilter, entity) || !entityShouldBeIncluded(p.spec, p.provisionalFilter, parent) {
			return parse.SearchShouldSkip
		}
		_, isDirectChildOfCluster := parent.(*matter.Cluster)
		switch entity := entity.(type) {
		case *matter.Namespace:
			field, ok := parent.(*matter.Field)
			if ok {
				_, existing := namespaces[field.Name]
				if !existing {
					namespaces[field.Name] = entity
				}
			}
			return parse.SearchShouldContinue
		case *matter.Bitmap, *matter.Enum, *matter.Struct:
			if isDirectChildOfCluster && !forceIncludeEntity(p.spec, parentCluster, entity) {
				// We won't include these entities if they're only referenced by the cluster itself, not any of its attributes, commands or events
				return parse.SearchShouldSkip
			}
		}
		if _, isGlobal := p.spec.GlobalObjects[entity]; isGlobal {
			globalEntities[entity] = true
			ce[entity] = struct{}{}
			return parse.SearchShouldContinue
		}
		ce[entity] = struct{}{}
		globalEntities[entity] = false
		return parse.SearchShouldContinue
	})

	for c, ce := range clusterEntities {
		doc, ok := p.spec.DocRefs[c]
		if !ok {
			continue
		}
		errata := p.spec.Errata.Get(doc.Path.Relative)
		if errata == nil || errata.SDK.Types == nil {
			continue
		}
		for name, entry := range errata.SDK.Types.Enums {
			if entry.ForceLocal {
				for entity := range globalEntities {
					if en, ok := entity.(*matter.Enum); ok && en.Name == name {
						exists := false
						for _, existing := range c.Enums {
							if existing.Name == name {
								exists = true
								break
							}
						}
						if !exists {
							ce[entity] = struct{}{}
							c.Enums = append(c.Enums, en)
						}
					}
				}
			}
		}
		for name, entry := range errata.SDK.Types.Bitmaps {
			if entry.ForceLocal {
				for entity := range globalEntities {
					if bm, ok := entity.(*matter.Bitmap); ok && bm.Name == name {
						exists := false
						for _, existing := range c.Bitmaps {
							if existing.Name == name {
								exists = true
								break
							}
						}
						if !exists {
							ce[entity] = struct{}{}
							c.Bitmaps = append(c.Bitmaps, bm)
						}
					}
				}
			}
		}
		for name, entry := range errata.SDK.Types.Structs {
			if entry.ForceLocal {
				for entity := range globalEntities {
					if st, ok := entity.(*matter.Struct); ok && st.Name == name {
						exists := false
						for _, existing := range c.Structs {
							if existing.Name == name {
								exists = true
								break
							}
						}
						if !exists {
							ce[entity] = struct{}{}
							c.Structs = append(c.Structs, st)
						}
					}
				}
			}
		}
	}

	for entity, isGlobal := range globalEntities {
		if !isGlobal {
			continue
		}
		if provisional.Check(p.spec, entity, entity) == provisional.StateUnreferenced {
			continue
		}
		switch entity := entity.(type) {
		case *matter.Bitmap:
			globalBitmaps = append(globalBitmaps, entity)
		case *matter.Enum:
			globalEnums = append(globalEnums, entity)
		case *matter.Struct:
			globalStructs = append(globalStructs, entity)

		default:
			slog.Warn("Unexpected entity in global entities", matter.LogEntity("entity", entity))
		}
	}

	for _, ns := range p.spec.Namespaces {
		name := ns.Name
		if _, existing := namespaces[name]; !existing {
			namespaces[name] = ns
		}
		doc, ok := p.spec.DocRefs[ns]
		if ok {
			errata := p.spec.Errata.Get(doc.Path.Relative)
			if errata != nil && errata.SDK.Types != nil {
				if entry, ok := errata.SDK.Types.Enums[name+"Tag"]; ok && entry.OverrideName != "" {
					namespaces[entry.OverrideName] = ns
				}
			}
		}
	}

	for fieldName, ns := range namespaces {
		en := matter.NewEnum(ns.Source(), ns.Parent())
		if strings.HasSuffix(fieldName, "Tag") {
			en.Name = fieldName
		} else {
			en.Name = fieldName + "Tag"
		}
		doc, ok := p.spec.DocRefs[ns]
		if ok {
			errata := p.spec.Errata.Get(doc.Path.Relative)
			if errata != nil && errata.SDK.Types != nil {
				if entry, ok := errata.SDK.Types.Enums[en.Name]; ok && entry.OverrideName != "" {
					en.Name = entry.OverrideName
				}
			}
		}
		en.Type = types.NewDataType(types.BaseDataTypeEnum8, types.DataTypeRankScalar)
		for _, tag := range ns.SemanticTags {
			nst := matter.NewEnumValue(tag.Source(), en)
			nst.Name = tag.Name
			nst.Value = tag.ID
			en.Values = append(en.Values, nst)
		}
		globalEnums = append(globalEnums, en)
		for c, ce := range clusterEntities {
			doc, ok := p.spec.DocRefs[c]
			if !ok {
				continue
			}
			errata := p.spec.Errata.Get(doc.Path.Relative)
			if errata == nil || errata.SDK.Types == nil {
				continue
			}
			if entry, ok := errata.SDK.Types.Enums[en.Name]; ok && entry.Keep {
				exists := false
				for _, existing := range c.Enums {
					if existing.Name == en.Name {
						exists = true
						break
					}
				}
				if !exists {
					ce[en] = struct{}{}
					c.Enums = append(c.Enums, en)
				}
			}
		}
	}

	slices.SortFunc(globalEnums, func(a, b *matter.Enum) int {
		return strings.Compare(a.Name, b.Name)
	})

	slices.SortFunc(globalBitmaps, func(a, b *matter.Bitmap) int {
		return strings.Compare(a.Name, b.Name)
	})

	slices.SortFunc(globalStructs, func(a, b *matter.Struct) int {
		return strings.Compare(a.Name, b.Name)
	})

	for _, clusterInfo := range clusterList {
		ce, ok := clusterEntities[clusterInfo.Cluster]
		if !ok {
			continue
		}

		var errata *errata.Errata
		if doc, ok := p.spec.DocRefs[clusterInfo.Cluster]; ok && p.spec.Errata != nil {
			errata = p.spec.Errata.Get(doc.Path.Relative)
		}

		ceNames := make(map[string]struct{}, len(ce))
		for e := range ce {
			ceNames[matter.EntityName(e)] = struct{}{}
		}
		for _, s := range clusterInfo.Cluster.Structs {
			if _, ok := ceNames[s.Name]; ok {
				if errata != nil && errata.SDK.Types != nil {
					if entry, ok := errata.SDK.Types.Structs[s.Name]; ok && entry.ForceGlobal {
						continue
					}
				}
				clusterInfo.ReferencedStructs = append(clusterInfo.ReferencedStructs, s)
			}
		}
		for _, en := range clusterInfo.Cluster.Enums {
			if _, ok := ceNames[en.Name]; ok {
				if errata != nil && errata.SDK.Types != nil {
					if entry, ok := errata.SDK.Types.Enums[en.Name]; ok && entry.ForceGlobal {
						continue
					}
				}
				clusterInfo.ReferencedEnums = append(clusterInfo.ReferencedEnums, en)
			}
		}
		for _, bm := range clusterInfo.Cluster.Bitmaps {
			if _, ok := ceNames[bm.Name]; ok {
				if errata != nil && errata.SDK.Types != nil {
					if entry, ok := errata.SDK.Types.Bitmaps[bm.Name]; ok && entry.ForceGlobal {
						continue
					}
				}
				clusterInfo.ReferencedBitmaps = append(clusterInfo.ReferencedBitmaps, bm)
			}
		}
	}

	tc := map[string]any{
		"bitmaps":   globalBitmaps,
		"enums":     globalEnums,
		"structs":   globalStructs,
		"clusters":  clusterList,
		"endpoints": endpoints,
	}
	if p.SuppressEndpoints {
		tc["endpoints"] = nil
	}

	var t *raymond.Template
	t, err = p.loadTemplate(p.spec)
	if err != nil {
		return
	}

	if p.PerTrait {
		for _, clusterInfo := range clusterList {
			c := clusterInfo.Cluster
			ce := clusterEntities[c]

			ceNames := make(map[string]struct{}, len(ce))
			for e := range ce {
				ceNames[matter.EntityName(e)] = struct{}{}
			}

			var clusterGlobalEnums []*matter.Enum
			var errata *errata.Errata
			if doc, ok := p.spec.DocRefs[c]; ok && p.spec.Errata != nil {
				errata = p.spec.Errata.Get(doc.Path.Relative)
			}

			for i := 0; i < len(c.Enums); {
				local := c.Enums[i]
				globalForced := false
				if errata != nil && errata.SDK.Types != nil {
					if entry, ok := errata.SDK.Types.Enums[local.Name]; ok && entry.ForceGlobal {
						globalForced = true
					}
				}
				if globalForced {
					clusterGlobalEnums = append(clusterGlobalEnums, local)
					c.Enums = append(c.Enums[:i], c.Enums[i+1:]...)
				} else {
					i++
				}
			}

			for _, en := range globalEnums {
				if _, ok := ceNames[en.Name]; ok {
					isLocal := false
					for _, local := range c.Enums {
						if local.Name == en.Name {
							isLocal = true
							break
						}
					}
					if !isLocal {
						clusterGlobalEnums = append(clusterGlobalEnums, en)
					}
				}
			}
			var clusterGlobalBitmaps []*matter.Bitmap
			for i := 0; i < len(c.Bitmaps); {
				local := c.Bitmaps[i]
				globalForced := false
				if errata != nil && errata.SDK.Types != nil {
					if entry, ok := errata.SDK.Types.Bitmaps[local.Name]; ok && entry.ForceGlobal {
						globalForced = true
					}
				}
				if globalForced {
					clusterGlobalBitmaps = append(clusterGlobalBitmaps, local)
					c.Bitmaps = append(c.Bitmaps[:i], c.Bitmaps[i+1:]...)
				} else {
					i++
				}
			}
			for _, bm := range globalBitmaps {
				if _, ok := ceNames[bm.Name]; ok {
					isLocal := false
					for _, local := range c.Bitmaps {
						if local.Name == bm.Name {
							isLocal = true
							break
						}
					}
					if !isLocal {
						clusterGlobalBitmaps = append(clusterGlobalBitmaps, bm)
					}
				}
			}
			var clusterGlobalStructs []*matter.Struct
			for i := 0; i < len(c.Structs); {
				local := c.Structs[i]
				globalForced := false
				if errata != nil && errata.SDK.Types != nil {
					if entry, ok := errata.SDK.Types.Structs[local.Name]; ok && entry.ForceGlobal {
						globalForced = true
					}
				}
				if globalForced {
					clusterGlobalStructs = append(clusterGlobalStructs, local)
					c.Structs = append(c.Structs[:i], c.Structs[i+1:]...)
				} else {
					i++
				}
			}
			for _, s := range globalStructs {
				if _, ok := ceNames[s.Name]; ok {
					isLocal := false
					for _, local := range c.Structs {
						if local.Name == s.Name {
							isLocal = true
							break
						}
					}
					if !isLocal {
						clusterGlobalStructs = append(clusterGlobalStructs, s)
					}
				}
			}

			clusterListSingle := []*ClusterInfo{clusterInfo}

			clusterTc := map[string]any{
				"bitmaps":   clusterGlobalBitmaps,
				"enums":     clusterGlobalEnums,
				"structs":   clusterGlobalStructs,
				"clusters":  clusterListSingle,
				"endpoints": nil,
			}
			var clusterOut string
			clusterOut, err = t.Exec(clusterTc)
			if err != nil {
				slog.Error("error rendering matter template for cluster", slog.String("name", c.Name), slog.Any("err", err))
				return
			}
			clusterFileName := getClusterFileName(c.Name)
			clusterPath := filepath.Join(dir, clusterFileName)
			outputs = append(outputs, pipeline.NewData(clusterPath, clusterOut))
		}
		return
	}

	var out string
	out, err = t.Exec(tc)
	if err != nil {
		slog.Error("error rendering matter template", slog.Any("err", err))
		return
	}
	outputs = append(outputs, pipeline.NewData(path, out))
	return
}

var template pipeline.Once[*raymond.Template]

func (sp *IdlRenderer) loadTemplate(spec *spec.Specification) (*raymond.Template, error) {
	t, err := template.Do(func() (*raymond.Template, error) {

		ov := handlebars.NewOverlay("", templateFiles, "templates")
		err := ov.Flush()
		if err != nil {
			slog.Error("Error flushing embedded templates", slog.Any("error", err))
		}
		t, err := handlebars.LoadTemplate("{{> matter}}", ov)
		if err != nil {
			return nil, err
		}

		handlebars.RegisterCommonHelpers(t)

		return t, nil
	})
	if err != nil {
		return nil, err
	}
	t = t.Clone()
	sp.registerIdlHelpers(t, spec)
	return t, nil
}

func forceIncludeEntity(spec *spec.Specification, cluster *matter.Cluster, e types.Entity) bool {
	if e, ok := e.(*matter.Enum); ok {
		if strings.EqualFold(e.Name, "StatusCode") || strings.EqualFold(e.Name, "StatusCodeEnum") || strings.EqualFold(e.Name, "ModeTag") {
			return true
		}
	}
	doc, ok := spec.DocRefs[cluster]
	if !ok {
		return false
	}
	errata := spec.Errata.Get(doc.Path.Relative)
	if errata == nil || (errata.SDK.ExtraTypes == nil && errata.SDK.Types == nil) {
		return false
	}
	switch e := e.(type) {
	case *matter.Bitmap:
		if errata.SDK.ExtraTypes != nil {
			if _, ok := errata.SDK.ExtraTypes.Bitmaps[e.Name]; ok {
				return true
			}
		}
		if errata.SDK.Types != nil {
			if entry, ok := errata.SDK.Types.Bitmaps[e.Name]; ok && entry.Keep {
				return true
			}
		}
	case *matter.Enum:
		if errata.SDK.ExtraTypes != nil {
			if _, ok := errata.SDK.ExtraTypes.Enums[e.Name]; ok {
				return true
			}
		}
		if errata.SDK.Types != nil {
			if entry, ok := errata.SDK.Types.Enums[e.Name]; ok && entry.Keep {
				return true
			}
		}
	case *matter.Struct:
		if errata.SDK.ExtraTypes != nil {
			if _, ok := errata.SDK.ExtraTypes.Structs[e.Name]; ok {
				return true
			}
		}
		if errata.SDK.Types != nil {
			if entry, ok := errata.SDK.Types.Structs[e.Name]; ok && entry.Keep {
				return true
			}
		}
	}
	return false
}

func getClusterFileName(clusterName string) string {
	name := caseify(clusterName, false, true)
	name = idlNameToLowerSnakeCase(name)
	return name + ".matter"
}

func idlNameToLowerSnakeCase(text string) string {
	replaceTerms := map[string]string{
		"BLE":   "Ble",
		"IDs":   "Ids",
		"IPv4":  "Ipv4",
		"IPv6":  "Ipv6",
		"iOS":   "Ios",
		"Int8U": "Int8u",
		"KWh":   "Kwh",
		"KVAh":  "Kvah",
	}

	for old, new := range replaceTerms {
		text = strings.ReplaceAll(text, old, new)
	}

	runes := []rune(text)
	if len(runes) == 0 {
		return ""
	}

	var splits []string
	var current strings.Builder
	current.WriteRune(runes[0])

	for i := 1; i < len(runes); i++ {
		shouldSplit := false
		r := runes[i]
		prev := runes[i-1]

		if unicode.IsUpper(r) {
			if !unicode.IsUpper(prev) {
				shouldSplit = true
			} else if i+1 < len(runes) && unicode.IsLower(runes[i+1]) {
				shouldSplit = true
			}
		}

		if shouldSplit {
			splits = append(splits, current.String())
			current.Reset()
		}
		current.WriteRune(r)
	}
	splits = append(splits, current.String())

	for i, s := range splits {
		splits[i] = strings.ToLower(s)
	}

	return strings.Join(splits, "_")
}


