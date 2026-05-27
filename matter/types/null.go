package types

import "math"

type Nullability uint8

const (
	NullabilityNonNull Nullability = iota
	NullabilityNullable
)

type Nullable interface {
	NullValue() uint64
}

func (dt *DataType) NullValue() uint64 {
	if dt == nil {
		return 0
	}
	switch dt.BaseType {
	case BaseDataTypeInt8, BaseDataTypeSignedTemperature, BaseDataTypeFabricIndex:
		return 0x80
	case BaseDataTypeInt16, BaseDataTypeTemperature, BaseDataTypeTemperatureDifference:
		return 0x8000
	case BaseDataTypeInt24:
		return 0x800000
	case BaseDataTypeInt32:
		return 0x80000000
	case BaseDataTypeInt40:
		return 0x8000000000
	case BaseDataTypeInt48:
		return 0x800000000000
	case BaseDataTypeInt56:
		return 0x80000000000000
	case BaseDataTypeInt64, BaseDataTypeAmperage, BaseDataTypeVoltage, BaseDataTypePower, BaseDataTypeEnergy, BaseDataTypeApparentEnergy, BaseDataTypeApparentPower, BaseDataTypeReactiveEnergy, BaseDataTypeReactivePower, BaseDataTypeMoney:
		return 0x8000000000000000
	case BaseDataTypeUInt8, BaseDataTypeBoolean, BaseDataTypeMap8, BaseDataTypeEnum8, BaseDataTypePercent, BaseDataTypePriority, BaseDataTypeStatus, BaseDataTypeUnsignedTemperature, BaseDataTypeActionID, BaseDataTypeNamespaceID, BaseDataTypeTag:
		return math.MaxUint8
	case BaseDataTypeUInt16, BaseDataTypeMap16, BaseDataTypeEnum16, BaseDataTypePercentHundredths, BaseDataTypeGroupID, BaseDataTypeEndpointID, BaseDataTypeEndpointNumber, BaseDataTypeVendorID, BaseDataTypeEntryIndex:
		return math.MaxUint16
	case BaseDataTypeUInt24:
		return maxUint24
	case BaseDataTypeUInt32, BaseDataTypeMap32, BaseDataTypeSingle, BaseDataTypeEpochSeconds, BaseDataTypeElapsedSeconds, BaseDataTypeDeviceTypeID, BaseDataTypeClusterID, BaseDataTypeAttributeID, BaseDataTypeFieldID, BaseDataTypeEventID, BaseDataTypeCommandID, BaseDataTypeTransactionID, BaseDataTypeDataVersion:
		return math.MaxUint32
	case BaseDataTypeUInt40:
		return maxUint40
	case BaseDataTypeUInt48:
		return maxUint48
	case BaseDataTypeUInt56:
		return maxUint56
	case BaseDataTypeUInt64, BaseDataTypeMap64, BaseDataTypeDouble, BaseDataTypeEpochMicroseconds, BaseDataTypePosixMilliseconds, BaseDataTypeSystimeMicroseconds, BaseDataTypeSystimeMilliseconds, BaseDataTypeFabricID, BaseDataTypeNodeID, BaseDataTypeIeeeAddress, BaseDataTypeEventNumber, BaseDataTypeSubjectID:
		return math.MaxUint64
	case BaseDataTypeCustom:
		if dt.Entity != nil {
			if m, ok := dt.Entity.(Nullable); ok {
				return m.NullValue()
			}
		}
	}
	return 0
}

var fromRangesNullable = map[BaseDataType]DataTypeExtreme{
	BaseDataTypeInt8:  {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt8 + 1, Format: NumberFormatInt},
	BaseDataTypeInt16: {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt16 + 1, Format: NumberFormatInt},
	BaseDataTypeInt24: {Type: DataTypeExtremeTypeInt64, Int64: minInt24 + 1, Format: NumberFormatInt},
	BaseDataTypeInt32: {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt32 + 1, Format: NumberFormatInt},
	BaseDataTypeInt40: {Type: DataTypeExtremeTypeInt64, Int64: minInt40 + 1, Format: NumberFormatInt},
	BaseDataTypeInt48: {Type: DataTypeExtremeTypeInt64, Int64: minInt48 + 1, Format: NumberFormatInt},
	BaseDataTypeInt56: {Type: DataTypeExtremeTypeInt64, Int64: minInt56 + 1, Format: NumberFormatInt},
	BaseDataTypeInt64: {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt64 + 1, Format: NumberFormatInt},

	BaseDataTypeAmperage:       {Type: DataTypeExtremeTypeInt64, Int64: minInt62 + 1, Format: NumberFormatHex},
	BaseDataTypeVoltage:        {Type: DataTypeExtremeTypeInt64, Int64: minInt62 + 1, Format: NumberFormatHex},
	BaseDataTypePower:          {Type: DataTypeExtremeTypeInt64, Int64: minInt62 + 1, Format: NumberFormatHex},
	BaseDataTypeEnergy:         {Type: DataTypeExtremeTypeInt64, Int64: minInt62 + 1, Format: NumberFormatHex},
	BaseDataTypeApparentPower:  {Type: DataTypeExtremeTypeInt64, Int64: minInt62 + 1, Format: NumberFormatHex},
	BaseDataTypeApparentEnergy: {Type: DataTypeExtremeTypeInt64, Int64: minInt62 + 1, Format: NumberFormatHex},
	BaseDataTypeReactivePower:  {Type: DataTypeExtremeTypeInt64, Int64: minInt62 + 1, Format: NumberFormatHex},
	BaseDataTypeReactiveEnergy: {Type: DataTypeExtremeTypeInt64, Int64: minInt62 + 1, Format: NumberFormatHex},
	BaseDataTypeMoney:          {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt64 + 1, Format: NumberFormatInt},

	BaseDataTypeTemperatureDifference: {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt16 + 1, Format: NumberFormatInt},
	BaseDataTypeSignedTemperature:     {Type: DataTypeExtremeTypeInt64, Int64: math.MinInt8 + 1, Format: NumberFormatInt},
	BaseDataTypeUnsignedTemperature:   {Type: DataTypeExtremeTypeUInt64, UInt64: 0, Format: NumberFormatInt},
}

var toRangesNullable = map[BaseDataType]DataTypeExtreme{
	BaseDataTypeMap8:              {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint8 & ^(1 << 7)},
	BaseDataTypeMap16:             {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint16 & ^(1 << 15)},
	BaseDataTypeMap32:             {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint32 & ^(1 << 31)},
	BaseDataTypeMap64:             {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint64 & ^(1 << 64)},
	BaseDataTypeEnum8:             {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint8 - 1},
	BaseDataTypeEnum16:            {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint16 - 1},
	BaseDataTypeUInt8:             {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint8 - 1, Format: NumberFormatInt},
	BaseDataTypeUInt16:            {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint16 - 1, Format: NumberFormatInt},
	BaseDataTypeUInt24:            {Type: DataTypeExtremeTypeUInt64, UInt64: maxUint24 - 1, Format: NumberFormatInt},
	BaseDataTypeUInt32:            {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint32 - 1, Format: NumberFormatInt},
	BaseDataTypeUInt40:            {Type: DataTypeExtremeTypeUInt64, UInt64: maxUint40 - 1, Format: NumberFormatInt},
	BaseDataTypeUInt48:            {Type: DataTypeExtremeTypeUInt64, UInt64: maxUint48 - 1, Format: NumberFormatInt},
	BaseDataTypeUInt56:            {Type: DataTypeExtremeTypeUInt64, UInt64: maxUint56 - 1, Format: NumberFormatInt},
	BaseDataTypeUInt64:            {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint64 - 1, Format: NumberFormatInt},
	BaseDataTypeEpochMicroseconds: {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint64 - 1, Format: NumberFormatHex},
	BaseDataTypeEpochSeconds:      {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint32 - 1, Format: NumberFormatHex},
	BaseDataTypePosixMilliseconds: {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint64 - 1, Format: NumberFormatHex},
	BaseDataTypeAmperage:          {Type: DataTypeExtremeTypeInt64, Int64: maxInt62, Format: NumberFormatHex},
	BaseDataTypeVoltage:           {Type: DataTypeExtremeTypeInt64, Int64: maxInt62, Format: NumberFormatHex},
	BaseDataTypePower:             {Type: DataTypeExtremeTypeInt64, Int64: maxInt62, Format: NumberFormatHex},
	BaseDataTypeEnergy:            {Type: DataTypeExtremeTypeInt64, Int64: maxInt62, Format: NumberFormatHex},
	BaseDataTypeApparentPower:     {Type: DataTypeExtremeTypeInt64, Int64: minInt62, Format: NumberFormatHex},
	BaseDataTypeApparentEnergy:    {Type: DataTypeExtremeTypeInt64, Int64: minInt62, Format: NumberFormatHex},
	BaseDataTypeReactivePower:     {Type: DataTypeExtremeTypeInt64, Int64: minInt62, Format: NumberFormatHex},
	BaseDataTypeReactiveEnergy:    {Type: DataTypeExtremeTypeInt64, Int64: minInt62, Format: NumberFormatHex},

	BaseDataTypeUnsignedTemperature: {Type: DataTypeExtremeTypeUInt64, UInt64: math.MaxUint8 - 1, Format: NumberFormatInt},
}
