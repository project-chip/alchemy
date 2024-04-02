package zap

import (
	"path/filepath"
	"strings"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
	"github.com/iancoleman/strcase"
)

func CleanName(name string) string {
	if !strings.Contains(name, " ") {
		return name
	}
	return strcase.ToCamel(name)
}

func ClusterName(path string, errata *Errata, entities []types.Entity) string {

	if errata.TemplatePath != "" {
		return errata.TemplatePath
	}

	path = filepath.Base(path)
	name := strings.TrimSuffix(path, filepath.Ext(path))

	var suffix string
	for _, m := range entities {
		switch m.(type) {
		case *matter.Cluster:
			suffix = "Cluster"
		}
	}
	if !strings.HasSuffix(name, suffix) {
		name += " " + suffix
	}
	return strcase.ToKebab(name)
}

func DeviceTypeName(deviceType *matter.DeviceType) string {
	name := strcase.ToKebab(deviceType.Name)
	return "MA-" + name
}

func getZapPath(sdkRoot string, name string) string {
	newPath := filepath.Join(sdkRoot, "src/app/zap-templates/zcl/data-model/chip", name+".xml")
	return newPath
}

func ClusterPath(sdkRoot string, path string, entities []types.Entity) string {
	newFile := filepath.Base(path)
	errata, ok := Erratas[newFile]
	if !ok {
		errata = DefaultErrata
	}

	newFile = ClusterName(path, errata, entities)
	newFile = strcase.ToKebab(newFile)
	return getZapPath(sdkRoot, newFile)
}
