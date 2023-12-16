package ascii

import (
	"fmt"
	"log/slog"

	"github.com/hasty/alchemy/matter"
)

type modelIndex map[string]map[matter.Model]struct{}

func (mi modelIndex) addModel(name string, model matter.Model) {
	m, ok := mi[name]
	if !ok {
		m = make(map[matter.Model]struct{})
		mi[name] = m
	}
	m[model] = struct{}{}
}

func BuildSpec(docs []*Doc) (spec *matter.Spec, err error) {

	buildTree(docs)

	spec = &matter.Spec{

		Clusters:    make(map[uint64]*matter.Cluster),
		DeviceTypes: make(map[uint64]*matter.DeviceType),
		ClusterRefs: make(map[matter.Model]map[*matter.Cluster]struct{}),
		DocRefs:     make(map[matter.Model]string),

		Bitmaps: make(map[string]*matter.Bitmap),
		Enums:   make(map[string]*matter.Enum),
		Structs: make(map[string]*matter.Struct),
	}
	modelIndex := make(modelIndex)

	for _, d := range docs {
		slog.Info("building spec", "path", d.Path)
		var models []matter.Model
		models, err = d.ToModel()
		if err != nil {
			slog.Warn("error building models", "doc", d.Path, "error", err)
			continue
		}
		for _, m := range models {
			switch m := m.(type) {
			case *matter.Cluster:
				spec.Clusters[m.ID.Value()] = m
				for _, en := range m.Bitmaps {
					_, ok := spec.Bitmaps[en.Name]
					if ok {
						slog.Debug("multiple bitmaps with same name", "name", en.Name)
					} else {
						spec.Bitmaps[en.Name] = en
					}
					spec.DocRefs[en] = d.Path
					modelIndex.addModel(en.Name, en)
				}
				for _, en := range m.Enums {
					_, ok := spec.Enums[en.Name]
					if ok {
						slog.Debug("multiple enums with same name", "name", en.Name)
					} else {
						spec.Enums[en.Name] = en
					}
					spec.DocRefs[en] = d.Path
					modelIndex.addModel(en.Name, en)
				}
				for _, en := range m.Structs {
					_, ok := spec.Structs[en.Name]
					if ok {
						slog.Debug("multiple structs with same name", "name", en.Name)
					} else {
						spec.Structs[en.Name] = en
					}
					spec.DocRefs[en] = d.Path
					modelIndex.addModel(en.Name, en)
				}
			case *matter.DeviceType:
				spec.DeviceTypes[m.ID.Value()] = m
			case *matter.Bitmap:
				_, ok := spec.Bitmaps[m.Name]
				if ok {
					slog.Debug("multiple bitmaps with same name", "name", m.Name)
				} else {

					spec.Bitmaps[m.Name] = m
				}
				modelIndex.addModel(m.Name, m)
			case *matter.Enum:
				_, ok := spec.Enums[m.Name]
				if ok {
					slog.Debug("multiple enums with same name", "name", m.Name)
				} else {
					spec.Enums[m.Name] = m
				}
				modelIndex.addModel(m.Name, m)
			case *matter.Struct:
				_, ok := spec.Structs[m.Name]
				if ok {
					slog.Debug("multiple structs with same name", "name", m.Name)
				} else {
					spec.Structs[m.Name] = m
				}
				modelIndex.addModel(m.Name, m)
			default:
				slog.Warn("unknown model type", "type", fmt.Sprintf("%T", m))
			}
			spec.DocRefs[m] = d.Path
		}
	}

	resolveDataTypeReferences(spec, modelIndex)

	for _, dt := range spec.DeviceTypes {
		for _, cr := range dt.ClusterRequirements {
			if c, ok := spec.Clusters[cr.ID.Value()]; ok {
				cr.Cluster = c
			}
		}
		for _, er := range dt.ElementRequirements {
			if c, ok := spec.Clusters[er.ID.Value()]; ok {
				er.Cluster = c
			}
		}
	}
	return
}

func resolveDataTypeReferences(spec *matter.Spec, mi modelIndex) {
	for _, s := range spec.Structs {
		for _, f := range s.Fields {
			resolveDataType(spec, mi, nil, f, f.Type)
		}
	}
	for _, cluster := range spec.Clusters {
		for _, a := range cluster.Attributes {
			if a.Type == nil {
				continue
			}
			resolveDataType(spec, mi, cluster, a, a.Type)
		}
		for _, s := range cluster.Structs {
			for _, f := range s.Fields {
				resolveDataType(spec, mi, cluster, f, f.Type)
			}
		}
		for _, s := range cluster.Events {
			for _, f := range s.Fields {
				resolveDataType(spec, mi, cluster, f, f.Type)
			}
		}
		for _, s := range cluster.Commands {
			for _, f := range s.Fields {
				resolveDataType(spec, mi, cluster, f, f.Type)
			}
		}
	}
}

func resolveDataType(spec *matter.Spec, mi modelIndex, cluster *matter.Cluster, field *matter.Field, dataType *matter.DataType) {
	if dataType == nil {
		slog.Warn("missing type on field", "name", field.Name)
		return
	}
	switch dataType.BaseType {
	case matter.BaseDataTypeList:
		resolveDataType(spec, mi, cluster, field, dataType.EntryType)
	case matter.BaseDataTypeCustom:
		if dataType.Model == nil {
			models := mi[dataType.Name]
			if len(models) == 0 {
				slog.Warn("unknown custom data type", "cluster", clusterName(cluster), "field", field.Name, "type", dataType.Name)
			} else if len(models) == 1 {
				for m := range models {
					dataType.Model = m
					break
				}
			} else {
				slog.Warn("ambiguous data type", "cluster", clusterName(cluster), "field", field.Name, "type", dataType.Name)
				for m := range models {
					slog.Warn("ambiguous data type", "model", m)
				}
			}
		}
		if cluster == nil || dataType.Model == nil {
			return
		}
		cm, ok := spec.ClusterRefs[dataType.Model]
		if !ok {
			cm = make(map[*matter.Cluster]struct{})
			spec.ClusterRefs[dataType.Model] = cm
		}
		cm[cluster] = struct{}{}
		slog.Debug("setting cluster", "name", cluster.Name, "type", dataType.Name)
		s, ok := dataType.Model.(*matter.Struct)
		if !ok {
			return
		}
		for _, f := range s.Fields {
			resolveDataType(spec, mi, cluster, f, f.Type)
		}
	}
}

func clusterName(cluster *matter.Cluster) string {
	if cluster != nil {
		return cluster.Name
	}
	return "none"
}
