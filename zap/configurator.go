package zap

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type Configurator struct {
	Spec *spec.Specification
	Doc  *spec.Doc

	Features []*matter.Number
	Bitmaps  map[*matter.Bitmap][]*matter.Number
	Enums    map[*matter.Enum][]*matter.Number
	Clusters map[*matter.Cluster]bool
	Structs  map[*matter.Struct][]*matter.Number

	ClusterIDs []string
}

func NewConfigurator(spec *spec.Specification, doc *spec.Doc, entities []types.Entity) (*Configurator, error) {
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
				if hasAtomicAttributes(cl) {
					c.Spec.ClusterRefs.Add(cl, atomicRequest)
					c.Spec.ClusterRefs.Add(cl, atomicResponse)
					cl.Commands = append(cl.Commands, atomicRequest)
					cl.Commands = append(cl.Commands, atomicResponse)
				}
			}
		case *matter.Cluster:
			if hasAtomicAttributes(v) {
				c.Spec.ClusterRefs.Add(v, atomicRequest)
				c.Spec.ClusterRefs.Add(v, atomicResponse)
				v.Commands = append(v.Commands, atomicRequest)
				v.Commands = append(v.Commands, atomicResponse)
			}
		}

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
	for _, cmd := range v.Commands {
		c.addTypes(cmd.Fields)
	}
	for _, e := range v.Events {
		c.addTypes(e.Fields)
	}

	if v.ID.Valid() {
		c.ClusterIDs = append(c.ClusterIDs, v.ID.HexString())
	}

	// Special case for status code enums, which typically do not get referenced
	for _, e := range v.Enums {
		if strings.EqualFold(e.Name, "StatusCode") || strings.EqualFold(e.Name, "StatusCodeEnum") {
			c.Enums[e] = append(c.Enums[e], v.ID)
		}
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
		slog.Warn("unknown cluster ref when searching for cluster codes", slog.String("path", c.Doc.Path), matter.LogEntity(entity))
		return
	}
	for ref := range refs {
		clusterIDs = append(clusterIDs, ref.ID)
	}
	matter.SortNumbers(clusterIDs)
	return
}
