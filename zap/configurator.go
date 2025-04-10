package zap

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type Configurator struct {
	Spec    *spec.Specification
	Docs    []*spec.Doc
	Domain  string
	OutPath string

	Features []*matter.Number
	Bitmaps  map[*matter.Bitmap][]*matter.Number
	Enums    map[*matter.Enum][]*matter.Number
	Clusters map[*matter.Cluster]bool
	Structs  map[*matter.Struct][]*matter.Number

	ClusterIDs []string
	Errata     *errata.SDK
	Global     bool
}

func NewConfigurator(spec *spec.Specification, docs []*spec.Doc, entities []types.Entity, outPath string, errata *errata.SDK, global bool) (*Configurator, error) {
	c := &Configurator{
		Spec:    spec,
		Docs:    docs,
		OutPath: outPath,

		Bitmaps:  make(map[*matter.Bitmap][]*matter.Number),
		Enums:    make(map[*matter.Enum][]*matter.Number),
		Clusters: make(map[*matter.Cluster]bool),
		Structs:  make(map[*matter.Struct][]*matter.Number),

		Errata: errata,
		Global: global,
	}
	if len(docs) > 0 {
		c.Domain = matter.DomainNames[docs[0].Domain]
	}

	for _, m := range entities {
		switch v := m.(type) {
		case *matter.ClusterGroup:
			for _, cl := range v.Clusters {
				c.addCluster(v, cl)
			}
			for _, bm := range v.Bitmaps {
				c.addEntityType(v, bm)
			}
			for _, e := range v.Enums {
				c.addEntityType(v, e)
			}
			for _, s := range v.Structs {
				c.addEntityType(v, s)
			}
		case *matter.Cluster:
			c.addCluster(v, v)
		case *matter.Bitmap, *matter.Enum, *matter.Struct:
			c.addEntityType(nil, v)
		}
	}
	c.addExtraTypes(errata, entities)
	return c, nil
}

func (c *Configurator) addCluster(parentEntity types.Entity, v *matter.Cluster) {
	c.addTypes(parentEntity, v.Attributes)
	if v.Features != nil {
		c.addEntityType(parentEntity, v.Features)
	}
	for _, cmd := range v.Commands {
		c.addTypes(parentEntity, cmd.Fields)
	}
	for _, e := range v.Events {
		c.addTypes(parentEntity, e.Fields)
	}

	c.addForcedTypes(v, parentEntity)

	if v.ID.Valid() {
		c.ClusterIDs = append(c.ClusterIDs, v.ID.HexString())
	}

	// Special case for status code and mode tag enums, which typically do not get referenced
	for _, e := range v.Enums {
		if strings.EqualFold(e.Name, "StatusCode") || strings.EqualFold(e.Name, "StatusCodeEnum") || strings.EqualFold(e.Name, "ModeTag") {
			c.Enums[e] = append(c.Enums[e], v.ID)
		}
	}
	for _, name := range c.Errata.ClusterSkip {
		if name == v.Name {
			return
		}
	}
	c.Clusters[v] = false
}

func (c *Configurator) addForcedTypes(cluster *matter.Cluster, parentEntity types.Entity) {
	if c.Errata != nil && len(c.Errata.ForceIncludeTypes) > 0 {
		for _, bm := range cluster.Bitmaps {
			for _, force := range c.Errata.ForceIncludeTypes {
				if strings.EqualFold(bm.Name, force) {
					c.addEntityType(parentEntity, bm)
					break
				}
			}
		}
		for _, en := range cluster.Enums {
			for _, force := range c.Errata.ForceIncludeTypes {
				if strings.EqualFold(en.Name, force) {
					c.addEntityType(parentEntity, en)
					break
				}
			}
		}
		for _, s := range cluster.Structs {
			for _, force := range c.Errata.ForceIncludeTypes {
				if strings.EqualFold(s.Name, force) {
					c.addEntityType(parentEntity, s)
					break
				}
			}
		}
	}
}

func (c *Configurator) addTypes(parentEntity types.Entity, fs matter.FieldSet) {
	for _, f := range fs {
		if f.Type == nil {
			continue
		}
		if conformance.IsZigbee(fs, f.Conformance) || conformance.IsDisallowed(f.Conformance) {
			continue
		}
		c.addType(parentEntity, f.Type)
	}
}

func (c *Configurator) addType(parentEntity types.Entity, dt *types.DataType) {
	if dt == nil {
		return
	}

	if dt.IsArray() {
		c.addType(parentEntity, dt.EntryType)
		return
	}

	entity := dt.Entity
	if entity == nil {
		slog.Debug("skipping data type with no entity", "name", dt.Name)
		return
	}
	if parentEntity != nil {
		if typeBelongsToOtherCluster(entity, parentEntity) {
			slog.Debug("skipping data type for different entity", "name", dt.Name, "parent", entity, "context", parentEntity)
			return
		}
	}
	c.addEntityType(parentEntity, entity)
}

func typeBelongsToOtherCluster(entity types.Entity, parentEntity types.Entity) bool {
	typeParent := entity.Parent()
	for {
		if typeParent == nil { // This is a global type, and doesn't belong to any cluster
			return true
		}
		if typeParent == parentEntity {
			return false
		}
		typeParent = typeParent.Parent()
	}
}

func (c *Configurator) addEntityType(parentEntity types.Entity, entity types.Entity) {

	switch entity := entity.(type) {
	case *matter.Features:
		c.Features = c.getClusterCodes(entity)
	case *matter.Bitmap:
		c.Bitmaps[entity] = c.getClusterCodes(entity)
	case *matter.Enum:
		c.Enums[entity] = c.getClusterCodes(entity)
	case *matter.Struct:
		c.Structs[entity] = c.getClusterCodes(entity)
		if !c.Global {
			c.addTypes(parentEntity, entity.Fields)
		}
	}
}

func (c *Configurator) getClusterCodes(entity types.Entity) (clusterIDs []*matter.Number) {
	if c.Global {
		// If these are global objects, we don't care what their associated cluster IDs are
		return []*matter.Number{matter.InvalidID}
	}
	refs, ok := c.Spec.ClusterRefs.Get(entity)
	if !ok {
		slog.Warn("unknown cluster ref when searching for cluster codes", c.DocLogs(), matter.LogEntity("entity", entity))
		return
	}
	for ref := range refs {
		clusterIDs = append(clusterIDs, ref.ID)
	}
	matter.SortNumbers(clusterIDs)
	return
}

func (c *Configurator) DocLogs() slog.Attr {
	var paths []any
	for _, d := range c.Docs {
		paths = append(paths, slog.String("path", d.Path.Relative))
	}
	return slog.Group("paths", paths...)
}
