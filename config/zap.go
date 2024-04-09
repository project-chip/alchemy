package config

import "github.com/hasty/alchemy/matter"

type ZapErrata struct {
	TopOrder      []matter.Section
	ClusterOrder  []matter.Section
	DataTypeOrder []matter.Section

	SuppressAttributePermissions bool
	ClusterDefinePrefix          string
	SuppressClusterDefinePrefix  bool
	DefineOverrides              map[string]string

	WritePrivilegeAsRole bool
	SeparateStructs      map[string]struct{}

	TemplatePath string

	ClusterSplit map[string]string

	Domain matter.Domain
}

type ZapSettings struct {
	Erratas map[string]*ZapErrata
}
