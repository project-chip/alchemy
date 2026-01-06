package zap

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type Configurator struct {
	Spec    *spec.Specification
	Docs    []*asciidoc.Document
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

	ExternalEntities map[types.Entity]struct{}
}

func (c *Configurator) IsEmpty() bool {
	if len(c.Clusters) > 0 {
		return false
	}
	if len(c.Features) > 0 {
		return false
	}
	if len(c.Bitmaps) > 0 {
		return false
	}
	if len(c.Enums) > 0 {
		return false
	}
	if len(c.Structs) > 0 {
		return false
	}
	return true
}

func NewConfigurator(spec *spec.Specification, docs []*asciidoc.Document, entities []types.Entity, outPath string, errata *errata.SDK, global bool) (*Configurator, error) {
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

		ExternalEntities: make(map[types.Entity]struct{}),
	}
	if len(docs) > 0 {
		//c.Domain = matter.DomainNames[docs[0].Domain]
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
	for _, name := range c.Errata.ClusterSkip {
		if name == v.Name {
			return
		}
	}
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

	if v.ID.Valid() {
		c.ClusterIDs = append(c.ClusterIDs, v.ID.HexString())
	}

	// Special case for status code and mode tag enums, which typically do not get referenced
	for _, e := range v.Enums {
		if strings.EqualFold(e.Name, "StatusCode") || strings.EqualFold(e.Name, "StatusCodeEnum") || strings.EqualFold(e.Name, "ModeTag") {
			c.Enums[e] = append(c.Enums[e], v.ID)
		}
		if c.Errata != nil && c.Errata.SeparateEnums != nil {
			if _, ok := c.Errata.SeparateEnums[e.Name]; ok {
				c.Enums[e] = append(c.Enums[e], v.ID)
			}
		}
	}
	if c.Errata != nil {
		if len(c.Errata.SeparateBitmaps) > 0 {
			for _, b := range v.Bitmaps {
				if _, ok := c.Errata.SeparateBitmaps[b.Name]; ok {
					c.Bitmaps[b] = append(c.Bitmaps[b], v.ID)
				}
			}

		}
	}
	c.Clusters[v] = false
}

func (c *Configurator) addTypes(parentEntity types.Entity, fs matter.FieldSet) {
	for _, f := range fs {
		if f.Type == nil {
			continue
		}
		if conformance.IsZigbee(f.Conformance) || IsDisallowed(f, f.Conformance) {
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
		if c.typeBelongsToOtherCluster(entity, parentEntity) {
			return
		}
	}
	c.addEntityType(parentEntity, entity)
}

func (c *Configurator) typeBelongsToOtherCluster(entity types.Entity, parentEntity types.Entity) bool {
	typeParent := entity.Parent()
	for {
		if typeParent == nil { // This is a global type, and doesn't belong to any cluster
			return true
		}
		if typeParent == parentEntity {
			return false
		}
		c.ExternalEntities[typeParent] = struct{}{}
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
	refs.Range(func(cluster *matter.Cluster, value struct{}) bool {
		clusterIDs = append(clusterIDs, cluster.ID)
		return true
	})
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
