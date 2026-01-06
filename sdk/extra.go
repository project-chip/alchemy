package sdk

import (
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

func addExtraTypes(extraTypes *errata.SDKTypes, entities []types.Entity) {
	if extraTypes == nil {
		return
	}
	var extraEntities []types.Entity
	for name, eb := range extraTypes.Bitmaps {
		bm := matter.NewBitmap(nil, nil)
		bm.Name = name
		typeName := eb.Type
		// This is a workaround for the errata file using the ZAP type name for bitmaps
		switch typeName {
		case "bitmap8":
			typeName = "map8"
		case "bitmap16":
			typeName = "map16"
		case "bitmap32":
			typeName = "map32"
		case "bitmap64":
			typeName = "map64"
		}
		bm.Type = types.ParseDataType(typeName, false)
		bm.Description = eb.Description
		for _, ef := range eb.Fields {
			b := matter.NewBitmapBit(nil, ef.Bit, ef.Name, "", nil)
			bm.Bits = append(bm.Bits, b)
		}
		extraEntities = append(extraEntities, bm)
	}
	for name, ee := range extraTypes.Enums {
		e := matter.NewEnum(nil, nil)
		e.Name = name
		e.Type = types.ParseDataType(ee.Type, false)
		e.Description = ee.Description
		for _, ef := range ee.Fields {
			ev := matter.NewEnumValue(nil, e)
			ev.Name = ef.Name
			ev.Value = matter.ParseNumber(ef.Value)
			e.Values = append(e.Values, ev)
		}
		extraEntities = append(extraEntities, e)
	}
	for name, es := range extraTypes.Structs {
		s := matter.NewStruct(nil, nil)
		s.Name = name
		s.Description = es.Description
		for i, ef := range es.Fields {
			f := matter.NewField(nil, s, types.EntityTypeStructField)
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
					addExtraEntity(cl, e)
				}
			case *matter.Cluster:
				addExtraEntity(v, e)
			}
		}
	}
}

func addExtraEntity(cluster *matter.Cluster, e types.Entity) {
	switch e := e.(type) {
	case *matter.Bitmap:
		for _, bm := range cluster.Bitmaps {
			if bm.Name == e.Name {
				return
			}
		}
		e.SetParent(cluster)
		cluster.AddBitmaps(e)
	case *matter.Enum:
		for _, en := range cluster.Enums {
			if en.Name == e.Name {
				return
			}
		}
		e.SetParent(cluster)
		cluster.AddEnums(e)
	case *matter.Struct:
		for _, s := range cluster.Structs {
			if s.Name == e.Name {
				return
			}
		}
		e.SetParent(cluster)
		cluster.AddStructs(e)
	}
}
