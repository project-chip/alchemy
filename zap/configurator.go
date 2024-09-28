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
	Doc     *spec.Doc
	OutPath string

	Features []*matter.Number
	Bitmaps  map[*matter.Bitmap][]*matter.Number
	Enums    map[*matter.Enum][]*matter.Number
	Clusters map[*matter.Cluster]bool
	Structs  map[*matter.Struct][]*matter.Number

	ClusterIDs []string
}

func NewConfigurator(spec *spec.Specification, doc *spec.Doc, entities []types.Entity, outPath string, errata *errata.ZAP) (*Configurator, error) {
	c := &Configurator{
		Spec:    spec,
		Doc:     doc,
		OutPath: outPath,

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
				c.addCluster(v, cl, errata)
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
			c.addCluster(v, v, errata)
		case *matter.Bitmap, *matter.Enum, *matter.Struct:
			c.addEntityType(nil, v)
		}
	}
	return c, nil
}

func (c *Configurator) addCluster(parentEntity types.Entity, v *matter.Cluster, errata *errata.ZAP) {
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

	// Special case for status code enums, which typically do not get referenced
	for _, e := range v.Enums {
		if strings.EqualFold(e.Name, "StatusCode") || strings.EqualFold(e.Name, "StatusCodeEnum") {
			c.Enums[e] = append(c.Enums[e], v.ID)
		}
	}
	for _, name := range errata.ClusterSkip {
		if name == v.Name {
			return
		}
	}
	c.Clusters[v] = false
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
			slog.Warn("skipping data type for different entity", "name", dt.Name, "parent", entity, "context", parentEntity)
			return
		}
	}
	c.addEntityType(parentEntity, entity)
}

func typeBelongsToOtherCluster(entity types.Entity, parentEntity types.Entity) bool {
	var typeParent types.Entity
	switch entity := entity.(type) {
	case *matter.Bitmap:
		typeParent = entity.ParentEntity
	case *matter.Enum:
		typeParent = entity.ParentEntity
	case *matter.Struct:
		typeParent = entity.ParentEntity
	}
	if typeParent == nil { // This is a global type, and doesn't belong to any cluster
		return true
	}
	return typeParent != parentEntity
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
		c.addTypes(parentEntity, entity.Fields)
	}
}

func (c *Configurator) getClusterCodes(entity types.Entity) (clusterIDs []*matter.Number) {
	refs, ok := c.Spec.ClusterRefs.Get(entity)
	if !ok {
		slog.Warn("unknown cluster ref when searching for cluster codes", slog.String("path", c.Doc.Path.String()), matter.LogEntity(entity))
		return
	}
	for ref := range refs {
		clusterIDs = append(clusterIDs, ref.ID)
	}
	matter.SortNumbers(clusterIDs)
	return
}
