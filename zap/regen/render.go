package regen

import (
	"context"
	"embed"
	"log/slog"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mailgun/raymond/v2"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/handlebars"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
	"github.com/project-chip/alchemy/zap"
)

//go:embed templates
var templateFiles embed.FS

type IdlRenderer struct {
	spec *spec.Specification

	commonAttributes matter.FieldSet
}

func NewIdlRenderer(spec *spec.Specification) (IdlRenderer, error) {
	renderer := IdlRenderer{spec: spec}
	renderer.commonAttributes = matter.FieldSet{
		&matter.Field{
			Name:        "GeneratedCommandList",
			ID:          matter.NewNumber(65528),
			Type:        types.NewDataType(types.BaseDataTypeCommandID, true),
			Access:      matter.Access{Read: matter.PrivilegeView},
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
		&matter.Field{
			Name:        "AcceptedCommandList",
			ID:          matter.NewNumber(65529),
			Type:        types.NewDataType(types.BaseDataTypeCommandID, true),
			Access:      matter.Access{Read: matter.PrivilegeView},
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
		&matter.Field{
			Name:        "AttributeList",
			ID:          matter.NewNumber(65531),
			Type:        types.NewDataType(types.BaseDataTypeAttributeID, true),
			Access:      matter.Access{Read: matter.PrivilegeView},
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
		&matter.Field{
			Name:        "FeatureMap",
			ID:          matter.NewNumber(65532),
			Type:        types.NewDataType(types.BaseDataTypeMap32, false),
			Access:      matter.Access{Read: matter.PrivilegeView},
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
		&matter.Field{
			Name:        "ClusterRevision",
			ID:          matter.NewNumber(65533),
			Type:        types.NewDataType(types.BaseDataTypeUInt16, false),
			Access:      matter.Access{Read: matter.PrivilegeView},
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
	}
	return renderer, nil
}

func (p IdlRenderer) Name() string {
	return "Writing Matter files"
}

func (p IdlRenderer) Process(cxt context.Context, input *pipeline.Data[*zap.File], index int32, total int32) (outputs []*pipeline.Data[string], extras []*pipeline.Data[*zap.File], err error) {

	dir := filepath.Dir(input.Path)
	base := filepath.Base(input.Path)
	extension := filepath.Ext(base)
	file := strings.TrimSuffix(base, extension)
	path := filepath.Join(dir, file+".matter")

	slog.Info("converting zap path", "path", input.Path, "matter", path)

	var t *raymond.Template
	t, err = p.loadTemplate(p.spec)
	if err != nil {
		return
	}

	var endpoints []Endpoint
	clusters := make(map[*matter.Cluster]*ClusterInfo)

	for _, endpoint := range input.Content.Endpoints {
		if endpoint.EndpointTypeIndex > len(input.Content.EndpointTypes)-1 {
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

	spec.TraverseEntities(p.spec, func(parentCluster *matter.Cluster, parent, entity types.Entity) parse.SearchShould {
		if parentCluster == nil {
			_, ok := globalEntities[entity]
			if ok {
				globalEntities[entity] = true
			}
			return parse.SearchShouldContinue
		}
		ce, ok := clusterEntities[parentCluster]
		if !ok {
			return parse.SearchShouldSkip
		}
		if !entityShouldBeIncluded(entity) || !entityShouldBeIncluded(parent) {
			return parse.SearchShouldSkip
		}
		_, isDirectChildOfCluster := parent.(*matter.Cluster)
		switch entity := entity.(type) {
		case *matter.Namespace:
			field, ok := parent.(*matter.Field)
			if ok {
				entity = entity.Clone()
				entity.Name = field.Name
				globalEntities[entity] = true
			}
			return parse.SearchShouldContinue
		case *matter.Bitmap, *matter.Enum, *matter.Struct:
			if isDirectChildOfCluster && !forceIncludeEntity(p.spec, parentCluster, entity) {
				// We won't include these entities if they're only referenced by the cluster itself, not any of its attributes, commands or events
				return parse.SearchShouldSkip
			}
		}
		ce[entity] = struct{}{}
		globalEntities[entity] = false
		return parse.SearchShouldContinue
	})

	for entity, isGlobal := range globalEntities {
		if !isGlobal {
			continue
		}
		switch entity := entity.(type) {
		case *matter.Bitmap:
			globalBitmaps = append(globalBitmaps, entity)
		case *matter.Enum:
			globalEnums = append(globalEnums, entity)
		case *matter.Struct:
			globalStructs = append(globalStructs, entity)
		case *matter.Namespace:
			ns := matter.NewEnum(entity.Source(), entity.Parent())
			ns.Name = entity.Name + "Tag"
			ns.Type = types.NewDataType(types.BaseDataTypeEnum8, false)
			for _, tag := range entity.SemanticTags {
				nst := matter.NewEnumValue(tag.Source(), ns)
				nst.Name = tag.Name
				nst.Value = tag.ID
				ns.Values = append(ns.Values, nst)
			}
			globalEnums = append(globalEnums, ns)
		}
	}

	for _, clusterInfo := range clusterList {
		ce, ok := clusterEntities[clusterInfo.Cluster]
		if !ok {
			continue
		}
		for e := range ce {
			isGlobal := globalEntities[e]
			if isGlobal {
				continue
			}
			switch e := e.(type) {
			case *matter.Bitmap:
				clusterInfo.ReferencedBitmaps = append(clusterInfo.ReferencedBitmaps, e)
			case *matter.Enum:
				clusterInfo.ReferencedEnums = append(clusterInfo.ReferencedEnums, e)
			case *matter.Struct:
				clusterInfo.ReferencedStructs = append(clusterInfo.ReferencedStructs, e)
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
	if errata == nil || errata.SDK.ExtraTypes == nil {
		return false
	}
	switch e := e.(type) {
	case *matter.Bitmap:
		if _, ok := errata.SDK.ExtraTypes.Bitmaps[e.Name]; ok {
			return true
		}
	case *matter.Enum:
		if _, ok := errata.SDK.ExtraTypes.Enums[e.Name]; ok {
			return true
		}
	case *matter.Struct:
		if _, ok := errata.SDK.ExtraTypes.Structs[e.Name]; ok {
			return true
		}
	}
	return false
}
