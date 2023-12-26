package zap

import (
	"log/slog"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/conformance"
)

type Configurator struct {
	spec *matter.Spec
	doc  *ascii.Doc

	Bitmaps  map[*matter.Bitmap]bool
	Enums    map[*matter.Enum]bool
	Clusters map[*matter.Cluster]bool
	Structs  map[*matter.Struct]bool

	ClusterIDs []string
}

func NewConfigurator(spec *matter.Spec, doc *ascii.Doc, models []matter.Model) (*Configurator, error) {
	c := &Configurator{
		spec:     spec,
		doc:      doc,
		Bitmaps:  make(map[*matter.Bitmap]bool),
		Enums:    make(map[*matter.Enum]bool),
		Clusters: make(map[*matter.Cluster]bool),
		Structs:  make(map[*matter.Struct]bool),
	}
	for _, m := range models {
		switch v := m.(type) {
		case *matter.Cluster:

			c.addTypes(v.Attributes)
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
	}
	return c, nil
}

func (c *Configurator) addTypes(fs matter.FieldSet) {
	for _, f := range fs {
		if f.Type == nil {
			continue
		}
		if conformance.IsZigbee(fs, f.Conformance) {
			continue
		}
		c.addType(f.Type)
	}
}

func (c *Configurator) addType(dt *matter.DataType) {
	if dt == nil {
		return
	}

	if dt.IsArray() {
		c.addType(dt.EntryType)
		return
	}

	model := dt.Model
	if model == nil {
		slog.Warn("skipping data type with no model", "name", dt.Name)
		return
	}
	path := c.spec.DocRefs[model]
	if path != c.doc.Path {
		// This model came from a different document, and will thus end up in its xml file, so should not be repeated here
		slog.Warn("skipping data type from another document", "name", dt.Name)
		return
	}

	switch model := dt.Model.(type) {
	case *matter.Bitmap:
		c.Bitmaps[model] = false
	case *matter.Enum:
		c.Enums[model] = false
	case *matter.Struct:
		c.Structs[model] = false
		c.addTypes(model.Fields)
	}
}
