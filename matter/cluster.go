package matter

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

type ClusterGroup struct {
	entity
	Name     string     `json:"name,omitempty"`
	Clusters []*Cluster `json:"clusters"`
	AssociatedDataTypes
}

func NewClusterGroup(name string, source asciidoc.Element, clusters []*Cluster) *ClusterGroup {
	cg := &ClusterGroup{
		Name:     name,
		entity:   entity{source: source},
		Clusters: clusters,
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

type Cluster struct {
	entity
	ID          *Number         `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Revisions   []*Revision     `json:"revisions,omitempty"`
	Parent      *Cluster        `json:"-"`
	Conformance conformance.Set `json:"conformance,omitempty"`

	Hierarchy string `json:"hierarchy,omitempty"`
	Role      string `json:"role,omitempty"`
	Scope     string `json:"scope,omitempty"`
	PICS      string `json:"pics,omitempty"`

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

func (c *Cluster) Inherit(parent *Cluster) (linkedEntities []types.Entity, err error) {
	c.Parent = parent
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

	c.Attributes = c.Attributes.Inherit(parent.Attributes)

	for _, pbm := range parent.Bitmaps {
		var matching *Bitmap
		for _, b := range c.Bitmaps {
			if b.Name == pbm.Name {
				matching = b
				break
			}
		}
		if matching == nil {
			c.Bitmaps = append(c.Bitmaps, pbm)
			linkedEntities = append(linkedEntities, pbm)
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
			c.Enums = append(c.Enums, pe)
			linkedEntities = append(linkedEntities, pe)
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
			c.Structs = append(c.Structs, ps)
			linkedEntities = append(linkedEntities, ps)
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
			c.Events = append(c.Events, pe.Clone())
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
			c.Commands = append(c.Commands, pc.Clone())
			continue
		}
		matching.Inherit(pc)
	}

	return
}

func (c *Cluster) Identifier(name string) (types.Entity, bool) {
	if c == nil {
		return nil, false
	}
	var cr types.Entity
	var ok bool
	if c.Features != nil {
		cr, ok = c.Features.Identifier(name)
		if ok {
			return cr, ok
		}

	}
	stores := []conformance.IdentifierStore{c.Attributes, c.Commands, c.Events, c.Enums, c.Bitmaps, c.Structs}
	for _, s := range stores {
		cr, ok = s.Identifier(name)
		if ok {
			return cr, true
		}
	}
	return nil, false
}
