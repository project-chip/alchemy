package errata

import (
	"github.com/goccy/go-yaml"
	"github.com/project-chip/alchemy/matter"
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

	WritePrivilegeAsRole bool `yaml:"write-privilege-as-role,omitempty"`

	SeparateStructs UniqueStringList `yaml:"separate-structs,omitempty"`
	SeparateBitmaps UniqueStringList `yaml:"separate-bitmaps,omitempty"`
	SeparateEnums   UniqueStringList `yaml:"separate-enums,omitempty"`

	SharedBitmaps UniqueStringList `yaml:"shared-bitmaps,omitempty"`
	SharedEnums   UniqueStringList `yaml:"shared-enums,omitempty"`
	SharedStructs UniqueStringList `yaml:"shared-structs,omitempty"`

	TemplatePath string `yaml:"template-path,omitempty"`

	ClusterSplit map[string]string `yaml:"cluster-split,omitempty"`
	ClusterSkip  []string          `yaml:"cluster-skip,omitempty"`

	TypeNames map[string]string `yaml:"type-names,omitempty"`

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

func (sdk *SDK) HasSdkPatch() bool {
	if sdk == nil {
		return false
	}
	if sdk.Types != nil {
		return true
	}
	if len(sdk.TypeNames) > 0 {
		return true
	}
	if sdk.ExtraTypes != nil {
		return true
	}
	if len(sdk.SharedBitmaps) > 0 {
		return true
	}
	if len(sdk.SharedEnums) > 0 {
		return true
	}
	if len(sdk.SharedStructs) > 0 {
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

	Quality string `yaml:"quality,omitempty"`
	Access  string `yaml:"access,omitempty"`
}

type SDKTypeCollection map[string]*SDKType

func (zap *SDK) OverrideDeviceTypeName(deviceType *matter.DeviceType, defaultName string) string {
	if zap.Types == nil {
		return defaultName
	}
	if zap.Types.DeviceTypes != nil {
		if override, ok := zap.Types.DeviceTypes[deviceType.Name]; ok && override.OverrideName != "" {
			return override.OverrideName
		}
	}
	return defaultName
}

func (zap *SDK) OverrideDeviceType(deviceType *matter.DeviceType, defaultTypeName string) string {
	if zap.Types == nil {
		return defaultTypeName
	}
	if zap.Types.DeviceTypes != nil {
		if override, ok := zap.Types.DeviceTypes[deviceType.Name]; ok && override.OverrideType != "" {
			return override.OverrideType
		}
	}
	return defaultTypeName
}

type UniqueStringList map[string]struct{}

func (i UniqueStringList) MarshalYAML() ([]byte, error) {
	structs := make([]string, 0, len(i))
	for s := range i {
		structs = append(structs, s)
	}
	return yaml.Marshal(structs)
}

func (i *UniqueStringList) UnmarshalYAML(b []byte) error {
	*i = make(UniqueStringList)
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
