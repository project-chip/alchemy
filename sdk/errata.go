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
	for path, errata := range spec.Errata.All() {
		typeOverrides := errata.SDK.Types
		typeNames := errata.SDK.TypeNames
		if typeOverrides == nil && len(typeNames) == 0 {
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

func applyErrataToField(entity *matter.Field, override *errata.SDKType) {
	if override.OverrideName != "" {
		entity.Name = override.OverrideName
	}
	if override.OverrideType != "" {
		entity.Type = types.ParseDataType(override.OverrideType, false)
	}
	if override.Conformance != "" {
		entity.Conformance = conformance.ParseConformance(override.Conformance)
	}
	if override.Constraint != "" {
		entity.Constraint = constraint.ParseString(override.Constraint)
	}
	if override.Fallback != "" {
		entity.Fallback = constraint.ParseLimit(override.Fallback)
	}
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
