package types

import "math"

var maxUint24 uint64 = math.MaxUint16 | (math.MaxUint8 << 16)
var maxUint40 uint64 = math.MaxUint32 | (math.MaxUint8 << 32)
var maxUint48 uint64 = math.MaxUint32 | (math.MaxUint16 << 32)
var maxUint56 uint64 = math.MaxUint32 | (maxUint24 << 32)

var maxInt24 int64 = math.MaxUint16 | (math.MaxInt8 << 16)
var maxInt40 int64 = math.MaxUint32 | (math.MaxInt8 << 32)
var maxInt48 int64 = math.MaxUint32 | (math.MaxInt16 << 32)
var maxInt56 int64 = math.MaxUint32 | (maxInt24 << 32)
var maxInt62 int64 = math.MaxUint64 & ^(3 << 62)

var toRanges = map[BaseDataType]DataTypeExtreme{
	BaseDataTypeMap8:              {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint8},
	BaseDataTypeMap16:             {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint16},
	BaseDataTypeMap32:             {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint32},
	BaseDataTypeMap64:             {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint64},
	BaseDataTypeEnum8:             {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint8},
	BaseDataTypeEnum16:            {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint16},
	BaseDataTypeUInt8:             {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint8, Format: NumberFormatInt},
	BaseDataTypeUInt16:            {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint16, Format: NumberFormatInt},
	BaseDataTypeUInt24:            {Type: DataTypeExtremeTypeUInt64, UInt64: maxUint24, Format: NumberFormatInt},
	BaseDataTypeUInt32:            {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint32, Format: NumberFormatInt},
	BaseDataTypeUInt40:            {Type: DataTypeExtremeTypeUInt64, UInt64: maxUint40, Format: NumberFormatInt},
	BaseDataTypeUInt48:            {Type: DataTypeExtremeTypeUInt64, UInt64: maxUint48, Format: NumberFormatInt},
	BaseDataTypeUInt56:            {Type: DataTypeExtremeTypeUInt64, UInt64: maxUint56, Format: NumberFormatInt},
	BaseDataTypeUInt64:            {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint64, Format: NumberFormatInt},
	BaseDataTypeInt8:              {Type: DataTypeExtremeTypeInt64, Int64: math.MaxInt8, Format: NumberFormatInt},
	BaseDataTypeInt16:             {Type: DataTypeExtremeTypeInt64, Int64: math.MaxInt16, Format: NumberFormatInt},
	BaseDataTypeInt24:             {Type: DataTypeExtremeTypeInt64, Int64: maxInt24, Format: NumberFormatInt},
	BaseDataTypeInt32:             {Type: DataTypeExtremeTypeInt64, Int64: math.MaxInt32, Format: NumberFormatInt},
	BaseDataTypeInt40:             {Type: DataTypeExtremeTypeInt64, Int64: maxInt40, Format: NumberFormatInt},
	BaseDataTypeInt48:             {Type: DataTypeExtremeTypeInt64, Int64: maxInt48, Format: NumberFormatInt},
	BaseDataTypeInt56:             {Type: DataTypeExtremeTypeInt64, Int64: maxInt56, Format: NumberFormatInt},
	BaseDataTypeInt64:             {Type: DataTypeExtremeTypeInt64, Int64: math.MaxInt64, Format: NumberFormatInt},
	BaseDataTypePercent:           {Type: DataTypeExtremeTypeUInt64, UInt64: 100},
	BaseDataTypePercentHundredths: {Type: DataTypeExtremeTypeUInt64, UInt64: 10000},
	BaseDataTypeEpochMicroseconds: {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint64, Format: NumberFormatHex},
	BaseDataTypeEpochSeconds:      {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint32, Format: NumberFormatHex},
	BaseDataTypePosixMilliseconds: {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint64, Format: NumberFormatHex},

	BaseDataTypeAmperage:       {Type: DataTypeExtremeTypeInt64, Int64: maxInt62, Format: NumberFormatHex},
	BaseDataTypeVoltage:        {Type: DataTypeExtremeTypeInt64, Int64: maxInt62, Format: NumberFormatHex},
	BaseDataTypePower:          {Type: DataTypeExtremeTypeInt64, Int64: maxInt62, Format: NumberFormatHex},
	BaseDataTypeEnergy:         {Type: DataTypeExtremeTypeInt64, Int64: maxInt62, Format: NumberFormatHex},
	BaseDataTypeApparentPower:  {Type: DataTypeExtremeTypeInt64, Int64: maxInt62, Format: NumberFormatHex},
	BaseDataTypeApparentEnergy: {Type: DataTypeExtremeTypeInt64, Int64: maxInt62, Format: NumberFormatHex},
	BaseDataTypeReactivePower:  {Type: DataTypeExtremeTypeInt64, Int64: maxInt62, Format: NumberFormatHex},
	BaseDataTypeReactiveEnergy: {Type: DataTypeExtremeTypeInt64, Int64: maxInt62, Format: NumberFormatHex},

	BaseDataTypeMoney: {Type: DataTypeExtremeTypeInt64, Int64: math.MaxInt64, Format: NumberFormatInt},

	BaseDataTypeTemperature:           {Type: DataTypeExtremeTypeInt64, Int64: math.MaxInt16, Format: NumberFormatInt},
	BaseDataTypeTemperatureDifference: {Type: DataTypeExtremeTypeInt64, Int64: math.MaxInt16, Format: NumberFormatInt},
	BaseDataTypeSignedTemperature:     {Type: DataTypeExtremeTypeInt64, Int64: math.MaxInt8, Format: NumberFormatInt},
	BaseDataTypeUnsignedTemperature:   {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint8, Format: NumberFormatInt},

	BaseDataTypeList: {Type: DataTypeExtremeTypeInt64, Int64: math.MaxUint16 - 1, Format: NumberFormatInt},
}

func Max(baseType BaseDataType, nullable Nullability) (to DataTypeExtreme) {
	var ok bool
	if nullable == NullabilityNullable {
		to, ok = toRangesNullable[baseType]
	}
	if !ok {
		to = toRanges[baseType]
	}
	return
}

func (dt *DataType) Max(nullable Nullability) (to DataTypeExtreme) {
	return Max(dt.BaseType, nullable)
}
