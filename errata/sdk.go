package errata

import (
	"log/slog"
	"runtime/debug"

	"github.com/goccy/go-yaml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

type SDK struct {
	SkipFile                     bool                `yaml:"skip-file,omitempty"`
	SuppressAttributePermissions bool                `yaml:"suppress-attribute-permissions,omitempty"`
	ClusterDefinePrefix          string              `yaml:"cluster-define-prefix,omitempty"`
	SuppressClusterDefinePrefix  bool                `yaml:"suppress-cluster-define-prefix,omitempty"`
	DefineOverrides              map[string]string   `yaml:"override-defines,omitempty"`
	ClusterName                  string              `yaml:"cluster-name,omitempty"`
	ClusterAliases               map[string][]string `yaml:"cluster-aliases,omitempty"`
	ClusterListKeys              map[string]string   `yaml:"cluster-list-keys,omitempty"`

	WritePrivilegeAsRole bool            `yaml:"write-privilege-as-role,omitempty"`
	SeparateStructs      SeparateStructs `yaml:"separate-structs,omitempty"`

	TemplatePath string `yaml:"template-path,omitempty"`

	ClusterSplit map[string]string `yaml:"cluster-split,omitempty"`
	ClusterSkip  []string          `yaml:"cluster-skip,omitempty"`

	TypeNames         map[string]string `yaml:"type-names,omitempty"`
	ForceIncludeTypes []string          `yaml:"force-include-types,omitempty"`

	Types      *SDKTypes `yaml:"types,omitempty"`
	ExtraTypes *SDKTypes `yaml:"extra-types,omitempty"`
}

func (s *SDK) HasSpecPatch() bool {
	if s.ClusterName != "" {
		return true
	}
	if s.Types != nil {
		return true
	}
	if s.ExtraTypes != nil {
		return true
	}
	return false
}

type SDKTypes struct {
	Attributes  map[string]*SDKType `yaml:"attributes,omitempty"`
	Clusters    map[string]*SDKType `yaml:"clusters,omitempty"`
	Enums       map[string]*SDKType `yaml:"enums,omitempty"`
	Bitmaps     map[string]*SDKType `yaml:"bitmaps,omitempty"`
	Structs     map[string]*SDKType `yaml:"structs,omitempty"`
	Commands    map[string]*SDKType `yaml:"commands,omitempty"`
	Events      map[string]*SDKType `yaml:"events,omitempty"`
	DeviceTypes map[string]*SDKType `yaml:"device-types,omitempty"`
}

type SDKType struct {
	Type         string `yaml:"type,omitempty"`
	Name         string `yaml:"name,omitempty"`
	OverrideName string `yaml:"override-name,omitempty"`
	OverrideType string `yaml:"override-type,omitempty"`
	List         bool   `yaml:"list,omitempty"`

	Fields      []*SDKType `yaml:"fields,omitempty"`
	Domain      string     `yaml:"domain,omitempty"`
	Priority    string     `yaml:"priority,omitempty"`
	Description string     `yaml:"description,omitempty"`

	Bit   string `yaml:"bit,omitempty"`
	Value string `yaml:"value,omitempty"`

	Constraint  string `yaml:"constraint,omitempty"`
	Conformance string `yaml:"conformance,omitempty"`
	Fallback    string `yaml:"fallback,omitempty"`
}

type SDKTypeCollection map[string]*SDKType

func (z *SDK) getTypes(entityType types.EntityType) SDKTypeCollection {
	if z.Types == nil {
		return nil
	}
	switch entityType {
	case types.EntityTypeAttribute:
		return z.Types.Attributes
	case types.EntityTypeEnum:
		return z.Types.Enums
	case types.EntityTypeCluster:
		return z.Types.Clusters
	case types.EntityTypeBitmap:
		return z.Types.Bitmaps
	case types.EntityTypeStruct:
		return z.Types.Structs
	case types.EntityTypeCommand:
		return z.Types.Commands
	case types.EntityTypeEvent:
		return z.Types.Events
	case types.EntityTypeDeviceType:
		return z.Types.DeviceTypes
	default:
		slog.Warn("Unexpected entity type in ZAP errata types", slog.String("type", entityType.String()))
		debug.PrintStack()
	}
	return nil
}

func (ztc SDKTypeCollection) getType(typeName string) (*SDKType, bool) {
	if len(ztc) == 0 {
		return nil, false
	}
	t, ok := ztc[typeName]
	if !ok {
		return nil, false
	}
	return t, true
}

func (t *SDKType) getField(fieldName string) (f *SDKType, ok bool) {
	if t == nil || t.Fields == nil {
		return nil, false
	}
	for _, f := range t.Fields {
		if f.Name == fieldName {
			return f, true
		}
	}
	return nil, false
}

func (z *SDK) getType(entity types.Entity) (*SDKType, bool) {
	if z == nil {
		return nil, false
	}
	switch entity := entity.(type) {
	case *matter.Field:
		switch entity.EntityType() {
		case types.EntityTypeAttribute:
			return z.getTypes(types.EntityTypeAttribute).getType(entity.Name)
		case types.EntityTypeCommandField,
			types.EntityTypeStructField,
			types.EntityTypeEventField:
			if entity.Parent() == nil {
				return nil, false
			}
			if t, ok := z.getType(entity.Parent()); ok {
				return t.getField(entity.Name)
			}
		default:
			slog.Warn("Unexpected entity type in ZAP errata types", slog.String("type", entity.EntityType().String()))
		}
	case *matter.Enum:
		return z.getTypes(types.EntityTypeEnum).getType(entity.Name)
	case *matter.EnumValue:
		if entity.Parent() == nil {
			return nil, false
		}
		if e, ok := z.getType(entity.Parent()); ok {
			return e.getField(entity.Name)
		}
	case *matter.Bitmap:
		return z.getTypes(types.EntityTypeBitmap).getType(entity.Name)
	case *matter.BitmapBit:
		if entity.Parent() == nil {
			return nil, false
		}
		if e, ok := z.getType(entity.Parent()); ok {
			return e.getField(entity.Name())
		}
	case *matter.Feature:
		if f, ok := z.getTypes(types.EntityTypeBitmap).getType("Features"); ok {
			return f.getField(entity.Name())
		}
	case *matter.Struct:
		return z.getTypes(types.EntityTypeStruct).getType(entity.Name)
	case *matter.Cluster:
		return z.getTypes(types.EntityTypeCluster).getType(entity.Name)
	case *matter.Command:
		return z.getTypes(types.EntityTypeCommand).getType(entity.Name)
	case *matter.Event:
		return z.getTypes(types.EntityTypeEvent).getType(entity.Name)
	case *matter.DeviceType:
		return z.getTypes(types.EntityTypeDeviceType).getType(entity.Name)
	case nil:
		slog.Warn("Unexpected nil entity in ZAP errata types")
		debug.PrintStack()
	default:
		slog.Warn("Unexpected entity type in ZAP errata types", matter.LogEntity("entity", entity))
	}
	return nil, false
}

func (zap *SDK) OverrideName(entity types.Entity, defaultName string) string {
	if zap == nil || (zap.TypeNames == nil && zap.Types == nil) {
		return defaultName
	}
	t, ok := zap.getType(entity)
	if ok && t.OverrideName != "" {
		return t.OverrideName
	}

	tn, ok := zap.TypeNames[defaultName]
	if ok {
		return tn
	}
	return defaultName
}

func (zap *SDK) OverrideConformance(entity types.Entity) conformance.Conformance {
	if zap == nil || (zap.TypeNames == nil && zap.Types == nil) {
		return matter.EntityConformance(entity)
	}
	t, ok := zap.getType(entity)
	if ok && t.Conformance != "" {
		return conformance.ParseConformance(t.Conformance)
	}
	return matter.EntityConformance(entity)
}

func (zap *SDK) OverrideConstraint(entity types.Entity) constraint.Constraint {
	if zap == nil || (zap.TypeNames == nil && zap.Types == nil) {
		return matter.EntityConstraint(entity)
	}
	t, ok := zap.getType(entity)
	if ok && t.Constraint != "" {
		return constraint.ParseString(t.Constraint)
	}
	return matter.EntityConstraint(entity)
}

func (zap *SDK) OverrideFallback(entity types.Entity) constraint.Limit {
	if zap == nil || (zap.TypeNames == nil && zap.Types == nil) {
		return matter.EntityFallback(entity)
	}
	t, ok := zap.getType(entity)
	if ok && t.Fallback != "" {
		return constraint.ParseLimit(t.Fallback)
	}
	return matter.EntityFallback(entity)
}

func (zap *SDK) OverrideDomain(clusterName string, defaultDomain string) string {
	if zap == nil || zap.Types == nil {
		return defaultDomain
	}
	t, ok := zap.getTypes(types.EntityTypeCluster).getType(clusterName)
	if ok && t.Domain != "" {
		return t.Domain
	}
	return defaultDomain
}

func (zap *SDK) OverrideType(entity types.Entity, defaultTypeName string) string {
	if zap == nil || (zap.TypeNames == nil && zap.Types == nil) {
		return defaultTypeName
	}
	t, ok := zap.getType(entity)
	if ok && t.OverrideType != "" {
		return t.OverrideType
	}

	tn, ok := zap.TypeNames[defaultTypeName]
	if ok {
		return tn
	}
	return defaultTypeName
}

func (zap *SDK) OverrideDescription(entity types.Entity, defaultDescription string) string {
	if zap == nil || zap.Types == nil {
		return defaultDescription
	}
	t, ok := zap.getType(entity)
	if ok && t.Description != "" {
		return t.Description
	}
	return defaultDescription
}

func (zap *SDK) OverridePriority(entity types.Entity, defaultPriority string) string {
	if zap == nil || zap.Types == nil {
		return defaultPriority
	}
	t, ok := zap.getType(entity)
	if ok && t.Priority != "" {
		return t.Priority
	}
	return defaultPriority
}

type SeparateStructs map[string]struct{}

func (i SeparateStructs) MarshalYAML() ([]byte, error) {
	structs := make([]string, 0, len(i))
	for s := range i {
		structs = append(structs, s)
	}
	return yaml.Marshal(structs)
}

func (i *SeparateStructs) UnmarshalYAML(b []byte) error {
	*i = make(SeparateStructs)
	var structs []string
	err := yaml.Unmarshal(b, &structs)
	if err != nil {
		return err
	}
	for _, s := range structs {
		(*i)[s] = struct{}{}
	}
	return nil
}
