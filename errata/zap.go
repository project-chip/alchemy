package errata

import (
	"path/filepath"

	"github.com/project-chip/alchemy/matter"
)

type ZAP struct {
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

func GetZAP(path string) *ZAP {
	errata, ok := Erratas[filepath.Base(path)]
	if !ok {
		return &DefaultErrata.ZAP
	}
	return &errata.ZAP
}
