package spec

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

type DocGroup struct {
	Root string
	referenceIndex

	Docs []*Doc
}

func NewDocGroup(root string) *DocGroup {
	return &DocGroup{
		Root:           root,
		referenceIndex: newReferenceIndex(),
	}
}

func setSpec(d *Doc, si *Specification, docGroup *DocGroup) {
	if d.group != nil {
		if d.group.Root != docGroup.Root && d.Path.Base() != "matter-defines.adoc" {
			slog.Warn("multiple doc group roots", "path", d.Path.String(), "root", d.group.Root, "newRoot", docGroup.Root)
		}
		return
	}
	docGroup.Docs = append(docGroup.Docs, d)
	d.spec = si
	d.group = docGroup
	for _, c := range d.children {
		setSpec(c, si, docGroup)
	}
}

func (dg *DocGroup) Anchors(id string) []*Anchor {
	return dg.anchors[id]
}

func (dg *DocGroup) CrossReferences(id string) []*CrossReference {
	return dg.crossReferences[id]
}

func (si *Specification) addEntity(entity types.Entity, cluster *matter.Cluster) {
	switch entity := entity.(type) {
	case *matter.Bitmap:
		si.addEntityByName(entity.Name, entity, cluster)
	case *matter.Enum:
		si.addEntityByName(entity.Name, entity, cluster)
	case *matter.Struct:
		si.addEntityByName(entity.Name, entity, cluster)
	case *matter.TypeDef:
		si.addEntityByName(entity.Name, entity, cluster)
	case *matter.Namespace:
		si.addEntityByName(entity.Name, entity, cluster)
	default:
		slog.Warn("Unexpected type adding entity to spec", log.Type("type", entity))
	}
}

func (si *Specification) addEntityByName(name string, entity types.Entity, cluster *matter.Cluster) {
	m, ok := si.entities[name]
	if !ok {
		m = make(map[types.Entity]map[*matter.Cluster]struct{})
		si.entities[name] = m
	}
	clusters, ok := m[entity]
	if !ok {
		clusters = make(map[*matter.Cluster]struct{})
		m[entity] = clusters
	}
	_, ok = clusters[cluster]
	if ok {
		slog.Debug("Registering same entity twice", "cluster", cluster.Name, "name", name, "address", fmt.Sprintf("%p", cluster))
		return
	}
	clusters[cluster] = struct{}{}
}
