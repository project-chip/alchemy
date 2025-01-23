package types

func (dt *DataType) Size() int {

	switch dt.BaseType {
	case BaseDataTypeBoolean, BaseDataTypeMap8, BaseDataTypeUInt8, BaseDataTypeInt8, BaseDataTypeEnum8, BaseDataTypePercent, BaseDataTypePriority, BaseDataTypeStatus:
		return 1
	case BaseDataTypeMap16, BaseDataTypeUInt16, BaseDataTypeInt16, BaseDataTypeEnum16, BaseDataTypePercentHundredths, BaseDataTypeGroupID, BaseDataTypeEndpointID, BaseDataTypeEndpointNumber, BaseDataTypeVendorID:
		return 2
	case BaseDataTypeUInt24, BaseDataTypeInt24:
		return 3
	case BaseDataTypeMap32, BaseDataTypeUInt32, BaseDataTypeInt32, BaseDataTypeSingle, BaseDataTypeEpochSeconds, BaseDataTypeElapsedSeconds:
		return 4
	case BaseDataTypeUInt40, BaseDataTypeInt40:
		return 5
	case BaseDataTypeUInt48, BaseDataTypeInt48:
		return 6
	case BaseDataTypeUInt56, BaseDataTypeInt56:
		return 7
	case BaseDataTypeMap64, BaseDataTypeUInt64, BaseDataTypeInt64, BaseDataTypeDouble:
		return 8
	case BaseDataTypeTimeOfDay:
		return 4
	case BaseDataTypeDate:
		return 8
	case BaseDataTypeEpochMicroseconds, BaseDataTypePosixMilliseconds, BaseDataTypeSystimeMicroseconds, BaseDataTypeSystimeMilliseconds:
		return 8
	case BaseDataTypeTemperature, BaseDataTypeTemperatureDifference:
		return 2
	case BaseDataTypeSignedTemperature, BaseDataTypeUnsignedTemperature:
		return 1
	case BaseDataTypeAmperage, BaseDataTypeVoltage, BaseDataTypePower, BaseDataTypeEnergy, BaseDataTypeApparentEnergy, BaseDataTypeApparentPower, BaseDataTypeReactiveEnergy, BaseDataTypeReactivePower, BaseDataTypeMoney:
		return 8
	case BaseDataTypeDeviceTypeID:
		return 4
	case BaseDataTypeFabricID, BaseDataTypeNodeID, BaseDataTypeIeeeAddress:
		return 8
	case BaseDataTypeFabricIndex:
		return 1
	case BaseDataTypeClusterID, BaseDataTypeAttributeID, BaseDataTypeFieldID, BaseDataTypeEventID, BaseDataTypeCommandID, BaseDataTypeTransactionID:
		return 4
	case BaseDataTypeActionID:
		return 1
	case BaseDataTypeEntryIndex:
		return 2
	case BaseDataTypeDataVersion:
		return 4
	case BaseDataTypeEventNumber, BaseDataTypeSubjectID:
		return 8
	case BaseDataTypeIPv4Address:
		return 4
	case BaseDataTypeIPv6Address, BaseDataTypeMessageID:
		return 16
	case BaseDataTypeSemanticTag:
		return 4
	case BaseDataTypeNamespaceID:
		return 1
	case BaseDataTypeTag:
		return 1
	case BaseDataTypeCustom:
		if entity, ok := dt.Entity.(HasBaseDataType); ok {
			switch entity.BaseDataType() {
			case BaseDataTypeMap8:
				return 1
			case BaseDataTypeMap16:
				return 2
			case BaseDataTypeMap32:
				return 4
			case BaseDataTypeMap64:
				return 8
			case BaseDataTypeEnum8:
				return 1
			case BaseDataTypeEnum16:
				return 2
			}
		}
	}
	return 0
}
