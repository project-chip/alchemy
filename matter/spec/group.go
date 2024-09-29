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
	default:
		slog.Warn("Unexpected type adding entity to spec", log.Type("type", entity))
	}
}

func (si *Specification) addEntityByName(name string, entity types.Entity, cluster *matter.Cluster) {
	m, ok := si.entities[name]
	if !ok {
		m = make(map[types.Entity]*matter.Cluster)
		si.entities[name] = m
	}
	existing, ok := m[entity]
	if ok {
		slog.Warn("Registering same entity twice", "cluster", cluster.Name, "name", name, "address", fmt.Sprintf("%p", entity), "existing", fmt.Sprintf("%p", existing))
	}
	m[entity] = cluster
}
