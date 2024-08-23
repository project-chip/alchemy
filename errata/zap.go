package errata

import (
	"github.com/project-chip/alchemy/matter"
)

type ZAP struct {
	SuppressAttributePermissions bool              `yaml:"suppress-attribute-permissions,omitempty"`
	ClusterDefinePrefix          string            `yaml:"cluster-define-prefix,omitempty"`
	SuppressClusterDefinePrefix  bool              `yaml:"suppress-cluster-define-prefix,omitempty"`
	DefineOverrides              map[string]string `yaml:"override-defines,omitempty"`

	WritePrivilegeAsRole bool                `yaml:"write-privilege-as-role,omitempty"`
	SeparateStructs      map[string]struct{} `yaml:"separate-structs,omitempty"`

	TemplatePath string `yaml:"template-path,omitempty"`

	ClusterSplit map[string]string `yaml:"cluster-split,omitempty"`

	Domain matter.Domain `yaml:"domain,omitempty"`
}

func GetZAP(path string) *ZAP {
	e := GetErrata(path)
	return &e.ZAP
}
