package matter

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

type ClusterClassification struct {
	Hierarchy string  `json:"hierarchy,omitempty"`
	Role      string  `json:"role,omitempty"`
	Scope     string  `json:"scope,omitempty"`
	PICS      string  `json:"pics,omitempty"`
	Quality   Quality `json:"quality,omitempty"`
}

type ClusterGroup struct {
	entity
	Name string `json:"name,omitempty"`

	ClusterClassification

	Clusters []*Cluster `json:"clusters"`
	AssociatedDataTypes
}

func NewClusterGroup(name string, source asciidoc.Element, clusters []*Cluster) *ClusterGroup {
	cg := &ClusterGroup{
		Name:     name,
		entity:   entity{source: source},
		Clusters: clusters,
	}
	for _, cluster := range clusters {
		cluster.parentEntity = cg
	}
	cg.AssociatedDataTypes.parentEntity = cg
	return cg
}

func (c ClusterGroup) EntityType() types.EntityType {
	return types.EntityTypeClusterGroup
}

func (c ClusterGroup) Explode() []*Cluster {
	return c.Clusters
}

func (c *ClusterGroup) Equals(e types.Entity) bool {
	oc, ok := e.(*ClusterGroup)
	if !ok {
		return false
	}
	return c.Name == oc.Name
}

type Cluster struct {
	entity
	ID            *Number         `json:"id,omitempty"`
	Name          string          `json:"name,omitempty"`
	Description   string          `json:"description,omitempty"`
	Revisions     Revisions       `json:"revisions,omitempty"`
	ParentCluster *Cluster        `json:"-"`
	Conformance   conformance.Set `json:"conformance,omitempty"`
	Domain        string          `json:"domain,omitempty"`

	ClusterClassification

	Features *Features `json:"features,omitempty"`
	AssociatedDataTypes
	Attributes FieldSet   `json:"attributes,omitempty"`
	Events     EventSet   `json:"events,omitempty"`
	Commands   CommandSet `json:"commands,omitempty"`
}

func NewCluster(source asciidoc.Element) *Cluster {
	c := &Cluster{
		entity: entity{source: source},
	}
	c.AssociatedDataTypes.parentEntity = c
	return c
}

func (c *Cluster) EntityType() types.EntityType {
	return types.EntityTypeCluster
}

func (c *Cluster) Equals(e types.Entity) bool {
	oc, ok := e.(*Cluster)
	if !ok {
		return false
	}
	if c.ID.Valid() && oc.ID.Valid() {
		return c.ID.Equals(oc.ID)
	}
	return c.Name == oc.Name
}

func (c *Cluster) Inherit(parent *Cluster) (linkedEntities []types.Entity, err error) {
	c.ParentCluster = parent
	if parent.Features != nil {
		if c.Features == nil || len(c.Features.Bits) == 0 {
			c.Features = parent.Features.Clone()
		} else {
			err = c.Features.Inherit(&parent.Features.Bitmap)
			if err != nil {
				return
			}
		}
	}

	if len(c.Description) == 0 {
		c.Description = parent.Description
	}

	c.Attributes = c.Attributes.Inherit(c, parent.Attributes)

	for _, pbm := range parent.Bitmaps {
		var matching *Bitmap
		for _, b := range c.Bitmaps {
			if b.Name == pbm.Name {
				matching = b
				break
			}
		}
		if matching == nil {
			matching = pbm.Clone()
			c.Bitmaps = append(c.Bitmaps, matching)
			matching.parent = c
			linkedEntities = append(linkedEntities, matching)
			continue
		}
		err = matching.Inherit(pbm)
		if err != nil {
			return
		}
	}

	for _, pe := range parent.Enums {
		var matching *Enum
		for _, en := range c.Enums {
			if en.Name == pe.Name {
				matching = en
				break
			}
		}
		if matching == nil {
			matching = pe.Clone()
			c.Enums = append(c.Enums, matching)
			matching.parent = c
			linkedEntities = append(linkedEntities, matching)
			continue
		}
		err = matching.Inherit(pe)
		if err != nil {
			return
		}
	}

	for _, ps := range parent.Structs {
		var matching *Struct
		for _, s := range c.Structs {
			if s.Name == ps.Name {
				matching = s
				break
			}
		}
		if matching == nil {
			matching = ps.Clone()
			c.Structs = append(c.Structs, matching)
			matching.parent = c
			linkedEntities = append(linkedEntities, matching)
			continue
		}
		matching.Inherit(ps)
	}

	for _, pe := range parent.Events {
		var matching *Event
		for _, e := range c.Events {
			if e.ID.Equals(pe.ID) {
				matching = e
				break
			}
		}
		if matching == nil {
			matching = pe.Clone()
			matching.parent = c
			c.Events = append(c.Events, matching)
			continue
		}
		matching.Inherit(pe)
	}

	for _, pc := range parent.Commands {
		var matching *Command
		for _, c := range c.Commands {
			if c.ID.Equals(pc.ID) {
				matching = c
				break
			}
		}
		if matching == nil {
			matching = pc.Clone()
			matching.parent = c
			c.Commands = append(c.Commands, matching)
			continue
		}
		matching.Inherit(pc)
	}

	for _, pc := range parent.Constants {
		var matching *Constant
		for _, c := range c.Constants {
			if c.Name == pc.Name {
				matching = c
				break
			}
		}
		if matching == nil {
			matching = pc.Clone()
			matching.parent = c
			c.Constants = append(c.Constants, matching)
			continue
		}
		matching.Inherit(pc)
	}

	return
}

func (c *Cluster) Cluster() *Cluster {
	return c
}

func (c *Cluster) IterateDataTypes() EntityIterator {
	return iterateOverEntities(c)
}

func (c *Cluster) TraverseDataTypes(callback EntityCallback) {
	traverseEntities(c, callback)
}

func findCluster(entity types.Entity) *Cluster {
	switch e := entity.(type) {
	case *Cluster:
		return e
	case *ClusterGroup:
		return e.Clusters[0]
	case *Struct:
		return findCluster(e.parent)
	case *Field:
		return findCluster(e.parent)
	case *Event:
		return findCluster(e.parent)
	case *Command:
		return findCluster(e.parent)
	case *Bitmap:
		return findCluster(e.parent)
	case *Enum:
		return findCluster(e.parent)
	}
	return nil
}
