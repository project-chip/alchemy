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

	WritePrivilegeAsRole bool            `yaml:"write-privilege-as-role,omitempty"`
	SeparateStructs      SeparateStructs `yaml:"separate-structs,omitempty"`

	TemplatePath string `yaml:"template-path,omitempty"`

	ClusterSplit map[string]string `yaml:"cluster-split,omitempty"`
	ClusterSkip  []string          `yaml:"cluster-skip,omitempty"`

	Domain matter.Domain `yaml:"domain,omitempty"`

	TypeNames map[string]string `yaml:"type-names,omitempty"`
}

func GetZAP(path string) *ZAP {
	e := GetErrata(path)
	return &e.ZAP
}

func (zap *ZAP) TypeName(typeName string) string {
	if zap == nil || zap.TypeNames == nil {
		return typeName
	}
	tn, ok := zap.TypeNames[typeName]
	if ok {
		return tn
	}
	return typeName
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
