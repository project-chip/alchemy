package python

import (
	"log/slog"

	"github.com/project-chip/alchemy/matter/types"
)

func toPythonType(baseDataType types.BaseDataType) string {
	switch baseDataType {
	case types.BaseDataTypeBoolean:
		return "bool"
	case types.BaseDataTypeUInt8,
		types.BaseDataTypeUInt16,
		types.BaseDataTypeUInt24,
		types.BaseDataTypeUInt32,
		types.BaseDataTypeUInt40,
		types.BaseDataTypeUInt48,
		types.BaseDataTypeUInt56,
		types.BaseDataTypeUInt64,
		types.BaseDataTypeInt8,
		types.BaseDataTypeInt16,
		types.BaseDataTypeInt24,
		types.BaseDataTypeInt32,
		types.BaseDataTypeInt40,
		types.BaseDataTypeInt48,
		types.BaseDataTypeInt56,
		types.BaseDataTypeInt64:
		return "int"
	case types.BaseDataTypeString:
		return "str"
	default:
		slog.Info("Unimplemented base type Python conversion", slog.String("baseDataType", baseDataType.String()))
		return "unknown"
	}
}
