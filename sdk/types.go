package sdk

import (
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func ToUnderlyingType(dt types.BaseDataType) types.BaseDataType {
	switch dt {
	case types.BaseDataTypeEnum8:
		return types.BaseDataTypeUInt8
	case types.BaseDataTypeEnum16:
		return types.BaseDataTypeUInt16
	case types.BaseDataTypeInt40,
		types.BaseDataTypeInt48,
		types.BaseDataTypeInt56:
		return types.BaseDataTypeInt64
	case types.BaseDataTypeUInt40,
		types.BaseDataTypeUInt48,
		types.BaseDataTypeUInt56:
		return types.BaseDataTypeUInt64
	case types.BaseDataTypePower,
		types.BaseDataTypeApparentPower,
		types.BaseDataTypeReactivePower,
		types.BaseDataTypeAmperage,
		types.BaseDataTypeVoltage,
		types.BaseDataTypeEnergy,
		types.BaseDataTypeApparentEnergy,
		types.BaseDataTypeReactiveEnergy,
		types.BaseDataTypeMoney:
		return types.BaseDataTypeInt64
	case types.BaseDataTypeEpochMicroseconds,
		types.BaseDataTypeSystimeMicroseconds,
		types.BaseDataTypePosixMilliseconds,
		types.BaseDataTypeSystimeMilliseconds:
		return types.BaseDataTypeUInt64
	case types.BaseDataTypeEpochSeconds,
		types.BaseDataTypeElapsedSeconds:
		return types.BaseDataTypeUInt32
	case types.BaseDataTypeTemperature, types.BaseDataTypeTemperatureDifference:
		return types.BaseDataTypeInt16
	case types.BaseDataTypeSignedTemperature:
		return types.BaseDataTypeInt8
	case types.BaseDataTypeUnsignedTemperature:
		return types.BaseDataTypeUInt8
	case types.BaseDataTypePercent:
		return types.BaseDataTypeUInt8
	case types.BaseDataTypePercentHundredths:
		return types.BaseDataTypeUInt16
	default:
		return dt
	}
}

func FindBaseType(dataType *types.DataType) types.BaseDataType {
	switch entity := dataType.Entity.(type) {
	case *matter.Enum:
		return entity.BaseDataType()
	case *matter.Bitmap:
		return entity.BaseDataType()
	default:
		return dataType.BaseType
	}
}

func CheckUnderlyingType(field *matter.Field, de types.DataTypeExtreme, dataExtremePurpose types.DataExtremePurpose) (out types.DataTypeExtreme, redundant bool) {
	out = de
	var nullability types.Nullability
	if field.Quality.Has(matter.QualityNullable) {
		nullability = types.NullabilityNullable
	}
	switch dataExtremePurpose {
	case types.DataExtremePurposeMinimum:

		fieldMinimum := types.Min(ToUnderlyingType(FindBaseType(field.Type)), nullability)
		if cmp, ok := de.Compare(fieldMinimum); ok && cmp == -1 {
			slog.Warn("Field has minimum lower than the range of its data type; overriding", slog.String("name", field.Name), log.Path("source", field), slog.String("specifiedMinimum", de.ZapString(field.Type)), slog.String("fieldMinimum", fieldMinimum.ZapString(field.Type)))
			out = fieldMinimum
		}
		if types.Min(ToUnderlyingType(FindBaseType(field.Type)), types.NullabilityNonNull).ValueEquals(out) {
			redundant = true
		}
	case types.DataExtremePurposeMaximum:
		fieldMaximum := types.Max(ToUnderlyingType(FindBaseType(field.Type)), nullability)
		if cmp, ok := de.Compare(fieldMaximum); ok && cmp == 1 {
			slog.Warn("Field has maximum greater than the range of its data type; overriding", slog.String("name", field.Name), log.Path("source", field), slog.String("specifiedMaximum", de.ZapString(field.Type)), slog.String("fieldMaximum", fieldMaximum.ZapString(field.Type)))
			out = fieldMaximum
		}
		if types.Max(ToUnderlyingType(FindBaseType(field.Type)), types.NullabilityNonNull).ValueEquals(out) {
			redundant = true
		}
	}
	return
}
