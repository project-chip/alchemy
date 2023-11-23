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

	WriteRoleAsPrivilege bool
	SeparateStructs      []string
}

type ZapSettings struct {
	Erratas map[string]*ZapErrata
}
