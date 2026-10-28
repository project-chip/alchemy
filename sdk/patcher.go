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

type Patcher struct {
	options applyErrataOptions
}

func NewPatcher(opts ...ApplyErrataOption) *Patcher {
	p := &Patcher{}
	for _, opt := range opts {
		opt(&p.options)
	}
	return p
}

type patch struct {
	spec *spec.Specification

	globalEnums   map[string]*matter.Enum
	globalBitmaps map[string]*matter.Bitmap
	globalStructs map[string]*matter.Struct
}

func (p *patch) patchGlobalObjects() (err error) {
	for entity, doc := range p.spec.GlobalObjects {
		errata := p.spec.Errata.Get(doc.Path.Relative)
		if errata == nil || !errata.SDK.HasSdkPatch() {
			continue
		}
		typeOverrides := errata.SDK.Types
		typeNames := errata.SDK.TypeNames
		switch entity := entity.(type) {
		case *matter.Bitmap:
			p.applyErrataToBitmap(entity, typeNames, typeOverrides)
			p.globalBitmaps[entity.Name] = entity
		case *matter.Enum:
			p.applyErrataToEnum(entity, typeNames, typeOverrides)
			p.globalEnums[entity.Name] = entity
		case *matter.Struct:
			err = p.applyErrataToStruct(entity, typeNames, typeOverrides)
			if err != nil {
				return err
			}
			p.globalStructs[entity.Name] = entity
		case *matter.Command:
			err = p.applyErrataToCommand(entity, typeNames, typeOverrides)
			if err != nil {
				return err
			}
		case *matter.Event:
			err = p.applyErrataToEvent(entity, typeNames, typeOverrides)
			if err != nil {
				return err
			}
		}
	}
	return
}

func (patcher *Patcher) BeforeParse(spec *spec.Specification) (err error) {
	return
}

func (patcher *Patcher) AfterParse(spec *spec.Specification) (err error) {
	return
}

func (patcher *Patcher) BeforeDataModelResolution(spec *spec.Specification) (err error) {
	if spec.Errata == nil {
		return
	}
	p := &patch{
		spec:          spec,
		globalEnums:   make(map[string]*matter.Enum),
		globalBitmaps: make(map[string]*matter.Bitmap),
		globalStructs: make(map[string]*matter.Struct),
	}
	resolveAtomicOperations(spec)
	err = p.patchGlobalObjects()
	if err != nil {
		return err
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
				err = p.applyErrataToCluster(spec, entity, errata.SDK, patcher.options)
				if err != nil {
					return err
				}
			case *matter.Bitmap:
				p.applyErrataToBitmap(entity, typeNames, typeOverrides)
			case *matter.Enum:
				p.applyErrataToEnum(entity, typeNames, typeOverrides)
			case *matter.Struct:
				err = p.applyErrataToStruct(entity, typeNames, typeOverrides)
				if err != nil {
					return err
				}
			case *matter.Command:
				err = p.applyErrataToCommand(entity, typeNames, typeOverrides)
				if err != nil {
					return err
				}
			case *matter.Event:
				err = p.applyErrataToEvent(entity, typeNames, typeOverrides)
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

	patchMiscellaneousErrata(spec)
	return
}

func (patcher *Patcher) AfterDataModelResolution(spec *spec.Specification) (err error) {
	patchContentLaunchCluster(spec)
	return
}

func (p *patch) applyErrataToCommand(st *matter.Command, typeNames map[string]string, typeOverrides *errata.SDKTypes) error {
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
		}
		err := p.applyErrataToFields(st.Fields, override)
		if err != nil {
			return err
		}
		err = p.injectExtraFields(st, &st.Fields, types.EntityTypeCommandField, override.ExtraFields)
		if err != nil {
			return err
		}
	}
	st.Name = applyTypeName(typeNames, st.Name)
	return nil
}

func (p *patch) applyErrataToEvent(ev *matter.Event, typeNames map[string]string, typeOverrides *errata.SDKTypes) error {
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
		err := p.applyErrataToFields(ev.Fields, override)
		if err != nil {
			return err
		}
		err = p.injectExtraFields(ev, &ev.Fields, types.EntityTypeEventField, override.ExtraFields)
		if err != nil {
			return err
		}
	}
	ev.Name = applyTypeName(typeNames, ev.Name)
	return nil
}

func applyErrataToDeviceType(deviceType *matter.DeviceType, typeOverrides *errata.SDKTypes) {
	if typeOverrides == nil || typeOverrides.DeviceTypes == nil {
		return
	}
	override, ok := typeOverrides.DeviceTypes[deviceType.Name]
	if !ok || override == nil {
		return
	}
	if override.OverrideProfileID != "" {
		deviceType.ProfileID = override.OverrideProfileID
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
