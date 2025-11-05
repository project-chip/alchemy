package regen

import (
	"context"
	"embed"
	"log/slog"
	"path/filepath"
	"slices"
	"strings"

	"github.com/mailgun/raymond/v2"
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
	clusters := make(map[*matter.Cluster]struct{})

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
			cluster, ok := p.spec.ClustersByID[uint64(clusterRef.Code)]
			if !ok {
				//err = fmt.Errorf("unrecognized cluster id in %s: %d", input.Path, clusterId.Code)
				return
			}
			clusters[cluster] = struct{}{}
			switch clusterRef.Side {
			case "server":
				ep.Servers = append(ep.Servers, cluster)
			case "client":
				ep.Clients = append(ep.Clients, cluster)
			}
		}
		endpoints = append(endpoints, ep)
	}

	clusterList := make([]*matter.Cluster, 0, len(clusters))
	clusterEntities := make(map[*matter.Cluster][]types.Entity)
	globalEntities := make(map[types.Entity]bool)
	for cluster := range clusters {
		clusterList = append(clusterList, cluster)
		clusterEntities[cluster] = []types.Entity{}
	}

	slices.SortFunc(clusterList, func(a *matter.Cluster, b *matter.Cluster) int {
		return a.ID.Compare(b.ID)
	})

	var globalEnums []*matter.Enum
	var globalStructs []*matter.Struct
	var globalBitmaps []*matter.Bitmap

	spec.IterateOverDataTypes(p.spec, func(cluster *matter.Cluster, parent, entity types.Entity) {
		if cluster == nil {
			_, ok := globalEntities[entity]
			if ok {
				globalEntities[entity] = true
			}
			return
		}
		ce, ok := clusterEntities[cluster]
		if !ok {
			return
		}
		switch entity := entity.(type) {
		case *matter.Namespace:
			field, ok := parent.(*matter.Field)
			if ok {
				entity = entity.Clone()
				entity.Name = field.Name
				globalEntities[entity] = true
			}
			return
		}
		clusterEntities[cluster] = append(ce, entity)
		globalEntities[entity] = false

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

		sp.registerIdlHelpers(t, spec)

		return t, nil
	})
	if err != nil {
		return nil, err
	}
	return t.Clone(), nil
}
