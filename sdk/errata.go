package sdk

import (
	"log/slog"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func ApplyErrata(spec *spec.Specification) (err error) {
	var addedExtraEntities bool
	for path, errata := range spec.Errata.All() {
		typeOverrides := errata.SDK.Types
		typeNames := errata.SDK.TypeNames
		extraTypes := errata.SDK.ExtraTypes
		if typeOverrides == nil && len(typeNames) == 0 && extraTypes == nil {
			continue
		}
		doc, ok := spec.Docs[path]
		if !ok {
			slog.Warn("Errata refers to unknown document", "path", path)
			continue
		}
		entities := spec.EntitiesForDocument(doc)
		for _, entity := range entities {
			switch entity := entity.(type) {
			case *matter.Cluster:
				applyErrataToCluster(entity, typeNames, typeOverrides)
			case *matter.Bitmap:
				applyErrataToBitmap(entity, typeNames, typeOverrides)
			case *matter.Enum:
				applyErrataToEnum(entity, typeNames, typeOverrides)
			case *matter.Struct:
				applyErrataToStruct(entity, typeNames, typeOverrides)
			case *matter.Command:
				applyErrataToCommand(entity, typeNames, typeOverrides)
			case *matter.Event:
				applyErrataToEvent(entity, typeNames, typeOverrides)
			}
		}
		if extraTypes != nil {
			addExtraTypes(extraTypes, entities)
			addedExtraEntities = true
		}
	}
	if addedExtraEntities {
		spec.BuildDataTypeReferences()
		spec.BuildClusterReferences()
	}
	return
}

func applyErrataToCluster(cluster *matter.Cluster, typeNames map[string]string, typeOverrides *errata.SDKTypes) {
	if typeOverrides != nil {
		ac, ok := typeOverrides.Clusters[cluster.Name]
		if ok {
			if ac.Domain != "" {
				cluster.Domain = ac.Domain
			}
			if ac.Description != "" {
				cluster.Description = ac.Description
			}
		}
		for _, a := range cluster.Attributes {
			ao, ok := typeOverrides.Attributes[a.Name]
			if !ok {
				continue
			}
			applyErrataToField(a, ao)
		}
		if cluster.Features != nil {
			fc, ok := typeOverrides.Bitmaps["Features"]
			if ok {
				for _, feature := range cluster.Features.Bits {
					for _, f := range fc.Fields {
						if feature.Name() == f.Name {
							if f.OverrideName != "" {
								feature.SetName(f.OverrideName)
							}
							break
						}
					}
				}
			}
		}
	}

	for _, bm := range cluster.Bitmaps {
		applyErrataToBitmap(bm, typeNames, typeOverrides)
	}
	for _, en := range cluster.Enums {
		applyErrataToEnum(en, typeNames, typeOverrides)
	}
	for _, s := range cluster.Structs {
		applyErrataToStruct(s, typeNames, typeOverrides)
	}
	for _, cmd := range cluster.Commands {
		applyErrataToCommand(cmd, typeNames, typeOverrides)
	}
	for _, ev := range cluster.Events {
		applyErrataToEvent(ev, typeNames, typeOverrides)
	}
}

func applyErrataToBitmap(bitmap *matter.Bitmap, typeNames map[string]string, typeOverrides *errata.SDKTypes) {
	if typeOverrides != nil {
		override, ok := typeOverrides.Bitmaps[bitmap.Name]

		if ok {
			if override.OverrideName != "" {
				bitmap.Name = override.OverrideName
			}
			if override.OverrideType != "" {
				bitmap.Type = types.ParseDataType(override.OverrideType, false)
			}
			if len(override.Fields) == 0 {
				return
			}
			for _, f := range override.Fields {
				for _, b := range bitmap.Bits {
					if b.Name() == f.Name {
						if f.OverrideName != "" {
							b.SetName(f.OverrideName)
						}
						break
					}
				}
			}
		}
	}
	bitmap.Name = applyTypeName(typeNames, bitmap.Name)
}

func applyErrataToEnum(en *matter.Enum, typeNames map[string]string, typeOverrides *errata.SDKTypes) {
	if typeOverrides != nil {
		override, ok := typeOverrides.Enums[en.Name]
		if ok {
			if override.OverrideName != "" {
				en.Name = override.OverrideName
			}
			if override.OverrideType != "" {
				en.Type = types.ParseDataType(override.OverrideType, false)
			}
			if len(override.Fields) == 0 {
				return
			}
			for _, f := range override.Fields {
				for _, ev := range en.Values {
					if ev.Name == f.Name {
						if f.OverrideName != "" {
							ev.Name = f.OverrideName
						}
						if f.Conformance != "" {
							ev.Conformance = conformance.ParseConformance(f.Conformance)
						}
						break
					}
				}
			}
		}
	}
	en.Name = applyTypeName(typeNames, en.Name)
}

func applyErrataToStruct(st *matter.Struct, typeNames map[string]string, typeOverrides *errata.SDKTypes) {
	if typeOverrides != nil {
		ast, ok := typeOverrides.Structs[st.Name]
		if !ok {
			return
		}
		if ast.OverrideName != "" {
			st.Name = ast.OverrideName
		}
		applyErrataToFields(st.Fields, ast)
	}
	st.Name = applyTypeName(typeNames, st.Name)

}

func applyErrataToCommand(st *matter.Command, typeNames map[string]string, typeOverrides *errata.SDKTypes) {
	if typeOverrides != nil {
		override, ok := typeOverrides.Commands[st.Name]
		if !ok {
			return
		}
		if override.OverrideName != "" {
			st.Name = override.OverrideName
		}
		if override.Description != "" {
			st.Description = override.Description
		}
		if override.Conformance != "" {
			st.Conformance = conformance.ParseConformance(override.Conformance)
		}
		applyErrataToFields(st.Fields, override)
	}
	st.Name = applyTypeName(typeNames, st.Name)
}

func applyErrataToEvent(ev *matter.Event, typeNames map[string]string, typeOverrides *errata.SDKTypes) {
	if typeOverrides != nil {
		override, ok := typeOverrides.Events[ev.Name]
		if !ok {
			return
		}
		if override.OverrideName != "" {
			ev.Name = override.OverrideName
		}
		if override.Description != "" {
			ev.Description = override.Description
		}
		if override.Priority != "" {
			ev.Priority = override.Priority
		}
		if override.Conformance != "" {
			ev.Conformance = conformance.ParseConformance(override.Conformance)
		}
		applyErrataToFields(ev.Fields, override)
	}
	ev.Name = applyTypeName(typeNames, ev.Name)

}

func applyErrataToFields(fs matter.FieldSet, override *errata.SDKType) {
	if len(override.Fields) != 0 {
		for _, f := range override.Fields {
			for _, field := range fs {
				if field.Name == f.Name {
					applyErrataToField(field, f)
					break
				}
			}
		}
	}
}

func applyErrataToField(field *matter.Field, override *errata.SDKType) {
	if override.OverrideName != "" {
		field.Name = override.OverrideName
	}
	if override.OverrideType != "" {
		field.Type = types.ParseDataType(override.OverrideType, false)
	}
	if override.Conformance != "" {
		field.Conformance = conformance.ParseConformance(override.Conformance)
	}
	if override.Constraint != "" {
		field.Constraint = constraint.ParseString(override.Constraint)
	}
	if override.Fallback != "" {
		field.Fallback = constraint.ParseLimit(override.Fallback)
	}
	field.Quality = overrideQuality(override, field.Quality)
	field.Access = overrideAccess(override, field.EntityType(), field.Access)
}

func applyErrataToDeviceType(deviceType *matter.DeviceType, typeOverrides *errata.SDKTypes) {
	override, ok := typeOverrides.DeviceTypes[deviceType.Name]
	if !ok {
		return
	}
	if override.OverrideName != "" {
		deviceType.Name = override.OverrideName
	}
}

func applyTypeName(typeNames map[string]string, name string) string {
	if len(typeNames) == 0 {
		return name
	}
	if override, ok := typeNames[name]; ok {
		return override
	}
	return name
}

func overrideQuality(override *errata.SDKType, defaultQuality matter.Quality) matter.Quality {
	if override.Quality == "" {
		return defaultQuality
	}
	switch override.Quality {
	case "none":
		return matter.QualityNone
	default:
		return matter.ParseQuality(override.Quality)
	}
}

func overrideAccess(override *errata.SDKType, entityType types.EntityType, defaultAccess matter.Access) matter.Access {
	if override.Access == "" {
		return defaultAccess
	}
	switch override.Access {
	case "none":
		return matter.Access{}
	default:
		access, parsed := spec.ParseAccess(override.Access, entityType)
		if parsed {
			return access
		}
		return defaultAccess
	}
}

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
			b := matter.NewBitmapBit(nil, bm, ef.Bit, ef.Name, "", nil)
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
