package zap

import (
	"github.com/hasty/alchemy/matter"
	"github.com/iancoleman/strcase"
)

func ZAPDeviceTypeName(deviceType *matter.DeviceType) string {
	name := strcase.ToKebab(deviceType.Name)
	return "MA-" + name
}
