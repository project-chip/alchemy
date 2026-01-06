package sdk

import (
	"log/slog"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func ApplyErrata(spec *spec.Specification) (err error) {
	var addedExtraEntities bool
	for path, errata := range spec.Errata.All() {
		if !errata.SDK.HasSdkPatch() {
			continue
		}
		typeOverrides := errata.SDK.Types
		typeNames := errata.SDK.TypeNames
		extraTypes := errata.SDK.ExtraTypes
		doc, ok := spec.Docs[path]
		if !ok {
			slog.Warn("Errata refers to unknown document", "path", path)
			continue
		}
		entities := spec.EntitiesForDocument(doc)
		for _, entity := range entities {
			switch entity := entity.(type) {
			case *matter.Cluster:
				applyErrataToCluster(spec, entity, errata.SDK)
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
