package sdk

import (
	"log/slog"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

type ApplyErrataOption func(*applyErrataOptions)

type applyErrataOptions struct {
	skipSharedEntities bool
}

func WithSkipSharedEntities(skip bool) ApplyErrataOption {
	return func(o *applyErrataOptions) {
		o.skipSharedEntities = skip
	}
}

func ApplyErrata(spec *spec.Specification, opts ...ApplyErrataOption) (err error) {
	var options applyErrataOptions
	for _, opt := range opts {
		opt(&options)
	}
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
				applyErrataToCluster(spec, entity, errata.SDK, options)
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
			case *matter.DeviceType:
				applyErrataToDeviceType(entity, typeOverrides)
			}
		}
		if extraTypes != nil {
			addExtraTypes(extraTypes, entities)
			addedExtraEntities = true
		}
	}
	if addedExtraEntities {
		spec.ResolveDataTypeReferences()
		spec.BuildDataTypeReferences()
		spec.BuildClusterReferences()
	}
	spec.ResolveConformances()
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
		if override.Response != "" {
			st.Response = types.ParseDataType(override.Response, types.DataTypeRankScalar)
			// Since ApplyErrata runs after type resolution, we must manually resolve the response name
			// to its command entity so that IDL and other templates can reference it properly.
			if cluster, ok := st.Parent().(*matter.Cluster); ok {
				var desiredDirection matter.Interface
				switch st.Direction {
				case matter.InterfaceServer:
					desiredDirection = matter.InterfaceClient
				case matter.InterfaceClient:
					desiredDirection = matter.InterfaceServer
				}
				for _, cmd := range cluster.Commands {
					if cmd.Direction == desiredDirection && cmd.Name == override.Response {
						st.Response.Entity = cmd
						break
					}
				}
				if st.Response.Entity == nil {
					slog.Warn("Failed to resolve overridden response", slog.String("command", st.Name), slog.String("response", override.Response))
				}
			}
		}
		applyErrataToFields(st.Fields, override)
		for _, f := range override.ExtraFields {
			var found bool
			for _, field := range st.Fields {
				if field.Name == f.Name {
					found = true
					break
				}
			}
			if !found {
				field := matter.NewField(nil, st, types.EntityTypeCommandField)
				field.Name = f.Name
				if f.Type != "" {
					var rank types.DataTypeRank = types.DataTypeRankScalar
					if f.List {
						rank = types.DataTypeRankList
					}
					field.Type = types.ParseDataType(f.Type, rank)
				}
				applyErrataToField(field, f)
				st.Fields = append(st.Fields, field)
			}
		}
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
		switch override.FabricScoping {
		case "none":
			ev.Access.FabricScoping = matter.FabricScopingUnscoped
		case "fabric-scoped":
			ev.Access.FabricScoping = matter.FabricScopingScoped
		}
		switch override.FabricSensitivity {
		case "none":
			ev.Access.FabricSensitivity = matter.FabricSensitivityInsensitive
		case "fabric-sensitive":
			ev.Access.FabricSensitivity = matter.FabricSensitivitySensitive
		}
		applyErrataToFields(ev.Fields, override)
		for _, f := range override.ExtraFields {
			var found bool
			for _, field := range ev.Fields {
				if field.Name == f.Name {
					found = true
					break
				}
			}
			if !found {
				field := matter.NewField(nil, ev, types.EntityTypeEventField)
				field.Name = f.Name
				if f.Type != "" {
					var rank types.DataTypeRank = types.DataTypeRankScalar
					if f.List {
						rank = types.DataTypeRankList
					}
					field.Type = types.ParseDataType(f.Type, rank)
				}
				applyErrataToField(field, f)
				ev.Fields = append(ev.Fields, field)
			}
		}
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
