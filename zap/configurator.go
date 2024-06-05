package zap

import (
	"log/slog"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/spec"
	"github.com/hasty/alchemy/matter/types"
)

type Configurator struct {
	Spec *matter.Spec
	Doc  *spec.Doc

	Features []*matter.Number
	Bitmaps  map[*matter.Bitmap][]*matter.Number
	Enums    map[*matter.Enum][]*matter.Number
	Clusters map[*matter.Cluster]bool
	Structs  map[*matter.Struct][]*matter.Number

	ClusterIDs []string
}

func NewConfigurator(spec *matter.Spec, doc *spec.Doc, entities []types.Entity) (*Configurator, error) {
	c := &Configurator{
		Spec:     spec,
		Doc:      doc,
		Bitmaps:  make(map[*matter.Bitmap][]*matter.Number),
		Enums:    make(map[*matter.Enum][]*matter.Number),
		Clusters: make(map[*matter.Cluster]bool),
		Structs:  make(map[*matter.Struct][]*matter.Number),
	}
	for _, m := range entities {
		switch v := m.(type) {
		case *matter.ClusterGroup:
			for _, cl := range v.Clusters {
				c.addCluster(cl)
			}
		case *matter.Cluster:
			c.addCluster(v)
		case *matter.Bitmap, *matter.Enum, *matter.Struct:
			c.addEntityType(v)
		}
	}
	return c, nil
}

func (c *Configurator) addCluster(v *matter.Cluster) {
	c.addTypes(v.Attributes)
	if v.Features != nil {
		c.addEntityType(v.Features)
	}
	for _, s := range v.Bitmaps {
		c.addEntityType(s)
	}
	for _, s := range v.Enums {
		c.addEntityType(s)
	}
	for _, s := range v.Structs {
		c.addEntityType(s)
	}
	for _, cmd := range v.Commands {
		c.addTypes(cmd.Fields)
	}
	for _, e := range v.Events {
		c.addTypes(e.Fields)
	}

	if v.ID.Valid() {
		c.ClusterIDs = append(c.ClusterIDs, v.ID.HexString())
	}
	c.Clusters[v] = false
}

func (c *Configurator) addTypes(fs matter.FieldSet) {
	for _, f := range fs {
		if f.Type == nil {
			continue
		}
		if conformance.IsZigbee(fs, f.Conformance) || conformance.IsDisallowed(f.Conformance) {
			continue
		}
		c.addType(f.Type)
	}
}

func (c *Configurator) addType(dt *types.DataType) {
	if dt == nil {
		return
	}

	if dt.IsArray() {
		c.addType(dt.EntryType)
		return
	}

	entity := dt.Entity
	if entity == nil {
		slog.Debug("skipping data type with no entity", "name", dt.Name)
		return
	}
	path := c.Spec.DocRefs[entity]
	if path != c.Doc.Path {
		// This entity came from a different document, and will thus end up in its xml file, so should not be repeated here

		slog.Debug("skipping data type from another document", "name", dt.Name, "path", c.Doc.Path)
		return
	}
	c.addEntityType(entity)
}

func (c *Configurator) addEntityType(entity types.Entity) {

	switch entity := entity.(type) {
	case *matter.Features:
		c.Features = c.getClusterCodes(entity)
	case *matter.Bitmap:
		c.Bitmaps[entity] = c.getClusterCodes(entity)
	case *matter.Enum:
		c.Enums[entity] = c.getClusterCodes(entity)
	case *matter.Struct:
		c.Structs[entity] = c.getClusterCodes(entity)
		c.addTypes(entity.Fields)
	}
}

func (c *Configurator) getClusterCodes(entity types.Entity) (clusterIDs []*matter.Number) {
	refs, ok := c.Spec.ClusterRefs[entity]
	if !ok {
		slog.Warn("unknown cluster ref when searching for cluster codes", "val", entity, "path", c.Doc.Path)
		return
	}
	for ref := range refs {
		clusterIDs = append(clusterIDs, ref.ID)
	}
	matter.SortNumbers(clusterIDs)
	return
}
