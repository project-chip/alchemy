package zap

import (
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

func (c *Configurator) addExtraTypes(errata *errata.SDK, entities []types.Entity) {
	if errata.ExtraTypes == nil {
		return
	}
	var extraEntities []types.Entity
	for name, eb := range errata.ExtraTypes.Bitmaps {
		bm := matter.NewBitmap(nil, nil)
		bm.Name = name
		bm.Type = types.ParseDataType(eb.Type, false)
		bm.Description = eb.Description
		for _, ef := range eb.Fields {
			b := matter.NewBitmapBit(nil, ef.Bit, ef.Name, "", nil)
			bm.Bits = append(bm.Bits, b)
		}
		extraEntities = append(extraEntities, bm)
	}
	for name, ee := range errata.ExtraTypes.Enums {
		e := matter.NewEnum(nil, nil)
		e.Name = name
		e.Type = types.ParseDataType(ee.Type, false)
		e.Description = ee.Description
		for _, ef := range ee.Fields {
			ev := matter.NewEnumValue(nil, nil)
			ev.Name = ef.Name
			ev.Value = matter.ParseNumber(ef.Value)
			e.Values = append(e.Values, ev)
		}
		extraEntities = append(extraEntities, e)
	}
	for name, es := range errata.ExtraTypes.Structs {
		s := matter.NewStruct(nil, nil)
		s.Name = name
		s.Description = es.Description
		for i, ef := range es.Fields {
			f := matter.NewField(nil, nil, types.EntityTypeStructField)
			f.ID = matter.NewNumber(uint64(i))
			f.Name = ef.Name
			f.Type = types.ParseDataType(ef.Type, ef.List)
			if ef.Constraint != "" {
				f.Constraint = constraint.ParseString(ef.Constraint)
			}
			if ef.Conformance != "" {
				f.Conformance = conformance.ParseConformance(ef.Conformance)
			}
			f.Conformance = conformance.Set{&conformance.Mandatory{}}
			s.Fields = append(s.Fields, f)
		}
		extraEntities = append(extraEntities, s)
	}
	for _, e := range extraEntities {
		for _, m := range entities {
			switch v := m.(type) {
			case *matter.ClusterGroup:
				for _, cl := range v.Clusters {
					c.addExtraEntity(cl, e)
				}
			case *matter.Cluster:
				c.addExtraEntity(v, e)
			}
		}
	}
	c.Spec.BuildDataTypeReferences()
	c.Spec.BuildClusterReferences()
	for _, e := range extraEntities {
		for _, m := range entities {
			switch v := m.(type) {
			case *matter.ClusterGroup:
				for _, cl := range v.Clusters {
					c.addEntityType(cl, e)

				}
			case *matter.Cluster:
				c.addEntityType(v, e)
			}
		}
	}

}

func (c *Configurator) addExtraEntity(cluster *matter.Cluster, e types.Entity) {
	switch e := e.(type) {
	case *matter.Bitmap:
		cluster.AddBitmaps(e)
	case *matter.Enum:
		cluster.AddEnums(e)
	case *matter.Struct:
		cluster.AddStructs(e)
	}
}
