package types

import "math"

var minInt24 int64 = ^maxInt24
var minInt40 int64 = ^maxInt40
var minInt48 int64 = ^maxInt48
var minInt56 int64 = ^maxInt56
var minInt62 int64 = ^maxInt62

var fromRanges = map[BaseDataType]DataTypeExtreme{
	BaseDataTypeMap8:              {Type: DataTypeExtremeTypeUInt64, UInt64: 0},
	BaseDataTypeMap16:             {Type: DataTypeExtremeTypeUInt64, UInt64: 0},
	BaseDataTypeMap32:             {Type: DataTypeExtremeTypeUInt64, UInt64: 0},
	BaseDataTypeMap64:             {Type: DataTypeExtremeTypeUInt64, UInt64: 0},
	BaseDataTypeEnum8:             {Type: DataTypeExtremeTypeUInt64, UInt64: 0},
	BaseDataTypeEnum16:            {Type: DataTypeExtremeTypeUInt64, UInt64: 0},
	BaseDataTypeInt8:              {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt8, Format: NumberFormatInt},
	BaseDataTypeInt16:             {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt16, Format: NumberFormatInt},
	BaseDataTypeInt24:             {Type: DataTypeExtremeTypeInt64, Int64: minInt24, Format: NumberFormatInt},
	BaseDataTypeInt32:             {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt32, Format: NumberFormatInt},
	BaseDataTypeInt40:             {Type: DataTypeExtremeTypeInt64, Int64: minInt40, Format: NumberFormatInt},
	BaseDataTypeInt48:             {Type: DataTypeExtremeTypeInt64, Int64: minInt48, Format: NumberFormatInt},
	BaseDataTypeInt56:             {Type: DataTypeExtremeTypeInt64, Int64: minInt56, Format: NumberFormatInt},
	BaseDataTypeInt64:             {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt64, Format: NumberFormatInt},
	BaseDataTypeUInt8:             {Type: DataTypeExtremeTypeUInt64, UInt64: 0, Format: NumberFormatInt},
	BaseDataTypeUInt16:            {Type: DataTypeExtremeTypeUInt64, UInt64: 0, Format: NumberFormatInt},
	BaseDataTypeUInt24:            {Type: DataTypeExtremeTypeUInt64, UInt64: 0, Format: NumberFormatInt},
	BaseDataTypeUInt32:            {Type: DataTypeExtremeTypeUInt64, UInt64: 0, Format: NumberFormatInt},
	BaseDataTypeUInt40:            {Type: DataTypeExtremeTypeUInt64, UInt64: 0, Format: NumberFormatInt},
	BaseDataTypeUInt48:            {Type: DataTypeExtremeTypeUInt64, UInt64: 0, Format: NumberFormatInt},
	BaseDataTypeUInt56:            {Type: DataTypeExtremeTypeUInt64, UInt64: 0, Format: NumberFormatInt},
	BaseDataTypeUInt64:            {Type: DataTypeExtremeTypeUInt64, UInt64: 0, Format: NumberFormatInt},
	BaseDataTypeEpochMicroseconds: {Type: DataTypeExtremeTypeUInt64, UInt64: 0, Format: NumberFormatHex},
	BaseDataTypeEpochSeconds:      {Type: DataTypeExtremeTypeUInt64, UInt64: 0, Format: NumberFormatHex},
	BaseDataTypePosixMilliseconds: {Type: DataTypeExtremeTypeUInt64, UInt64: 0, Format: NumberFormatHex},

	BaseDataTypePercent:           {Type: DataTypeExtremeTypeUInt64, UInt64: 0},
	BaseDataTypePercentHundredths: {Type: DataTypeExtremeTypeUInt64, UInt64: 0},
	BaseDataTypeTemperature:       {Type: DataTypeExtremeTypeInt64, Int64: -27315, Format: NumberFormatInt},
	BaseDataTypeAmperage:          {Type: DataTypeExtremeTypeInt64, Int64: minInt62, Format: NumberFormatHex},
	BaseDataTypeVoltage:           {Type: DataTypeExtremeTypeInt64, Int64: minInt62, Format: NumberFormatHex},
	BaseDataTypePower:             {Type: DataTypeExtremeTypeInt64, Int64: minInt62, Format: NumberFormatHex},
	BaseDataTypeEnergy:            {Type: DataTypeExtremeTypeInt64, Int64: minInt62, Format: NumberFormatHex},
	BaseDataTypeApparentPower:     {Type: DataTypeExtremeTypeInt64, Int64: minInt62, Format: NumberFormatHex},
	BaseDataTypeApparentEnergy:    {Type: DataTypeExtremeTypeInt64, Int64: minInt62, Format: NumberFormatHex},
	BaseDataTypeReactivePower:     {Type: DataTypeExtremeTypeInt64, Int64: minInt62, Format: NumberFormatHex},
	BaseDataTypeReactiveEnergy:    {Type: DataTypeExtremeTypeInt64, Int64: minInt62, Format: NumberFormatHex},

	BaseDataTypeMoney: {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt64, Format: NumberFormatInt},

	BaseDataTypeTemperatureDifference: {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt16},
	BaseDataTypeSignedTemperature:     {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt8 + 1}, // Spec doesn't allow for -12.8°C, even if not nullable
	BaseDataTypeUnsignedTemperature:   {Type: DataTypeExtremeTypeUInt64, UInt64: 0},
}

func Min(baseType BaseDataType, nullability Nullability) (from DataTypeExtreme) {
	var ok bool
	if nullability == NullabilityNullable {
		from, ok = fromRangesNullable[baseType]
	}
	if !ok {
		from = fromRanges[baseType]
	}
	return
}

func (dt *DataType) Min(nullability Nullability) (from DataTypeExtreme) {
	return Min(dt.BaseType, nullability)
}
