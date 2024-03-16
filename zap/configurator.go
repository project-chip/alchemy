package zap

import (
	"log/slog"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
	"github.com/hasty/alchemy/matter/types"
)

type Configurator struct {
	Spec *matter.Spec
	Doc  *ascii.Doc

	Bitmaps  map[*matter.Bitmap][]*matter.Number
	Enums    map[*matter.Enum][]*matter.Number
	Clusters map[*matter.Cluster]bool
	Structs  map[*matter.Struct][]*matter.Number

	ClusterIDs []string
}

func NewConfigurator(spec *matter.Spec, doc *ascii.Doc, entities []types.Entity) (*Configurator, error) {
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
		case *matter.Cluster:

			c.addTypes(v.Attributes)
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
		case *matter.Bitmap, *matter.Enum, *matter.Struct:
			c.addEntityType(v)
		}
	}
	return c, nil
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
		slog.Warn("unknown cluster ref", "val", entity)
		return
	}
	for ref := range refs {
		clusterIDs = append(clusterIDs, ref.ID)
	}
	matter.SortNumbers(clusterIDs)
	return
}
