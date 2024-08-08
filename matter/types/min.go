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

	BaseDataTypePercent:               {Type: DataTypeExtremeTypeUInt64, UInt64: 0},
	BaseDataTypePercentHundredths:     {Type: DataTypeExtremeTypeUInt64, UInt64: 10000},
	BaseDataTypeTemperature:           {Type: DataTypeExtremeTypeInt64, Int64: -27315, Format: NumberFormatInt},
	BaseDataTypeAmperage:              {Type: DataTypeExtremeTypeInt64, Int64: minInt62},
	BaseDataTypeVoltage:               {Type: DataTypeExtremeTypeInt64, Int64: minInt62},
	BaseDataTypePower:                 {Type: DataTypeExtremeTypeInt64, Int64: minInt62},
	BaseDataTypeEnergy:                {Type: DataTypeExtremeTypeInt64, Int64: minInt62},
	BaseDataTypeTemperatureDifference: {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt16},
	BaseDataTypeSignedTemperature:     {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt8 + 1}, // Spec doesn't allow for -12.8°C, even if not nullable
	BaseDataTypeUnsignedTemperature:   {Type: DataTypeExtremeTypeUInt64, UInt64: 0},
}

func (dt *DataType) Min(nullable bool) (from DataTypeExtreme) {
	var ok bool
	if nullable {
		from, ok = fromRangesNullable[dt.BaseType]
	}
	if !ok {
		from = fromRanges[dt.BaseType]
	}
	return
}
