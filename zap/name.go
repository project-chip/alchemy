package zap

import (
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func CleanName(name string) string {
	if !strings.Contains(name, " ") {
		return name
	}
	return strcase.ToCamel(name)
}

func ClusterName(path string, errata *errata.ZAP, entities []types.Entity) string {

	if errata.TemplatePath != "" {
		return errata.TemplatePath
	}

	path = filepath.Base(path)
	name := strings.TrimSuffix(path, filepath.Ext(path))

	var suffix string
	for _, m := range entities {
		switch m.(type) {
		case *matter.Cluster, matter.ClusterGroup:
			suffix = "Cluster"
		}
	}
	if !strings.HasSuffix(name, suffix) {
		name += " " + suffix
	}
	return strcase.ToKebab(name)
}

func DeviceTypeName(deviceType *matter.DeviceType) string {
	name := matter.Case(deviceType.Name)
	return "MA-" + strings.ToLower(name)
}
