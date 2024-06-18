package spec

import (
	"log/slog"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
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
		if d.group.Root != docGroup.Root {
			slog.Warn("multiple doc group roots", "path", d.Path, "root", d.group.Root, "newRoot", docGroup.Root)
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

func (si *Specification) addEntity(name string, entity types.Entity, cluster *matter.Cluster) {
	m, ok := si.entities[name]
	if !ok {
		m = make(map[types.Entity]*matter.Cluster)
		si.entities[name] = m
	}
	m[entity] = cluster
}
