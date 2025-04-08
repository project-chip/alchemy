package errata

import (
	"github.com/goccy/go-yaml"
	"github.com/project-chip/alchemy/matter"
)

type ZAP struct {
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

	Domain matter.Domain `yaml:"domain,omitempty"`

	DeviceTypeNames map[string]string `yaml:"device-type-names,omitempty"`

	TypeNames         map[string]string `yaml:"type-names,omitempty"`
	ForceIncludeTypes []string          `yaml:"force-include-types,omitempty"`

	Types      map[string]*ZAPType `yaml:"types,omitempty"`
	ExtraTypes map[string]*ZAPType `yaml:"extra-types,omitempty"`
}

type ZAPType struct {
	Type         string      `yaml:"type,omitempty"`
	OverrideName string      `yaml:"override-name,omitempty"`
	Fields       []*ZAPField `yaml:"fields,omitempty"`
	Priority     string      `yaml:"priority,omitempty"`
	Description  string      `yaml:"description,omitempty"`
}

type ZAPField struct {
	Name      string `yaml:"name,omitempty"`
	Type      string `yaml:"type,omitempty"`
	Bit       string `yaml:"bit,omitempty"`
	Value     string `yaml:"value,omitempty"`
	List      bool   `yaml:"list,omitempty"`
	MaxLength int64  `yaml:"max-length,omitempty"`

	OverrideName string `yaml:"override-name,omitempty"`
	OverrideType string `yaml:"override-type,omitempty"`
}

func GetZAP(path string) *ZAP {
	e := GetErrata(path)
	return &e.ZAP
}

func (zap *ZAP) TypeName(typeName string) string {
	if zap == nil || (zap.TypeNames == nil && len(zap.Types) == 0) {
		return typeName
	}
	t, ok := zap.Types[typeName]
	if ok && t.OverrideName != "" {
		return t.OverrideName
	}

	tn, ok := zap.TypeNames[typeName]
	if ok {
		return tn
	}
	return typeName
}

func (zap *ZAP) TypeDescription(typeName string, defaultDescription string) string {
	if zap == nil || len(zap.Types) == 0 {
		return defaultDescription
	}
	t, ok := zap.Types[typeName]
	if ok && t.Description != "" {
		return t.Description
	}
	return defaultDescription
}

func (zap *ZAP) FieldName(typeName string, fieldName string) string {
	if len(zap.Types) == 0 {
		return fieldName
	}
	t, ok := zap.Types[typeName]
	if !ok {
		return fieldName
	}
	for _, f := range t.Fields {
		if f.Name == fieldName {
			if f.OverrideName != "" {
				return f.OverrideName
			}
			break
		}
	}
	return fieldName
}

func (zap *ZAP) FieldTypeName(typeName string, fieldName string, defaultTypeName string) string {
	if len(zap.Types) == 0 {
		return defaultTypeName
	}
	t, ok := zap.Types[typeName]
	if !ok {
		return defaultTypeName
	}
	for _, f := range t.Fields {
		if f.Name == fieldName {
			if f.OverrideType != "" {
				return f.OverrideType
			}
			break
		}
	}
	return defaultTypeName
}

func (zap *ZAP) EventPriority(eventName string, defaultPriority string) string {
	if len(zap.Types) == 0 {
		return defaultPriority
	}
	t, ok := zap.Types[eventName]
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
