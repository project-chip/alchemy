package sdk

import (
	"fmt"
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
				err = applyErrataToCluster(spec, entity, errata.SDK, options)
				if err != nil {
					return err
				}
			case *matter.Bitmap:
				applyErrataToBitmap(entity, typeNames, typeOverrides)
			case *matter.Enum:
				applyErrataToEnum(entity, typeNames, typeOverrides)
			case *matter.Struct:
				err = applyErrataToStruct(entity, typeNames, typeOverrides)
				if err != nil {
					return err
				}
			case *matter.Command:
				err = applyErrataToCommand(entity, typeNames, typeOverrides)
				if err != nil {
					return err
				}
			case *matter.Event:
				err = applyErrataToEvent(entity, typeNames, typeOverrides)
				if err != nil {
					return err
				}
			case *matter.DeviceType:
				applyErrataToDeviceType(entity, typeOverrides)
			}
		}
		if extraTypes != nil {
			err = addExtraTypes(extraTypes, entities)
			if err != nil {
				return err
			}

		}
	}
	for entity, doc := range spec.GlobalObjects {
		errata := spec.Errata.Get(doc.Path.Relative)
		if errata == nil || !errata.SDK.HasSdkPatch() {
			continue
		}
		typeOverrides := errata.SDK.Types
		typeNames := errata.SDK.TypeNames
		switch entity := entity.(type) {
		case *matter.Bitmap:
			applyErrataToBitmap(entity, typeNames, typeOverrides)
		case *matter.Enum:
			applyErrataToEnum(entity, typeNames, typeOverrides)
		case *matter.Struct:
			err = applyErrataToStruct(entity, typeNames, typeOverrides)
			if err != nil {
				return err
			}
		case *matter.Command:
			err = applyErrataToCommand(entity, typeNames, typeOverrides)
			if err != nil {
				return err
			}
		case *matter.Event:
			err = applyErrataToEvent(entity, typeNames, typeOverrides)
			if err != nil {
				return err
			}
		}
	}
	spec.ResolveDataTypeReferences()
	spec.BuildDataTypeReferences()
	spec.BuildClusterReferences()
	spec.ResolveConformances()
	return
}

func applyErrataToCommand(st *matter.Command, typeNames map[string]string, typeOverrides *errata.SDKTypes) error {
	if typeOverrides != nil {
		override, ok := typeOverrides.Commands[st.Name]
		if !ok {
			return nil
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
		err := applyErrataToFields(st.Fields, override)
		if err != nil {
			return err
		}
		err = injectExtraFields(st, &st.Fields, types.EntityTypeCommandField, override.ExtraFields)
		if err != nil {
			return err
		}
	}
	st.Name = applyTypeName(typeNames, st.Name)
	return nil
}

func applyErrataToEvent(ev *matter.Event, typeNames map[string]string, typeOverrides *errata.SDKTypes) error {
	if typeOverrides != nil {
		override, ok := typeOverrides.Events[ev.Name]
		if !ok {
			return nil
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
		err := applyErrataToFields(ev.Fields, override)
		if err != nil {
			return err
		}
		err = injectExtraFields(ev, &ev.Fields, types.EntityTypeEventField, override.ExtraFields)
		if err != nil {
			return err
		}
	}
	ev.Name = applyTypeName(typeNames, ev.Name)
	return nil
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

func overrideAccess(override *errata.SDKType, entityType types.EntityType, defaultAccess matter.Access) (matter.Access, error) {
	if override.Access == "" {
		return defaultAccess, nil
	}
	switch override.Access {
	case "none":
		return matter.Access{}, nil
	default:
		access, parsed := spec.ParseAccess(override.Access, entityType)
		if parsed {
			return access, nil
		}
		return defaultAccess, fmt.Errorf("failed to parse access string %q for entity type %s", override.Access, entityType)
	}
}
