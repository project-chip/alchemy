package types

import (
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

type DataTypeRank uint8

const (
	DataTypeRankScalar DataTypeRank = iota
	DataTypeRankList
)

func (r DataTypeRank) IsList() bool {
	return r == DataTypeRankList
}

type DataType struct {
	BaseType BaseDataType     `json:"baseType"`
	Name     string           `json:"name"`
	Entity   Entity           `json:"-"`
	Source   asciidoc.Element `json:"-"`

	EntryType *DataType `json:"entryType,omitempty"`
}

func NewDataType(baseType BaseDataType, rank DataTypeRank) *DataType {
	if rank.IsList() {
		return &DataType{Name: "list", BaseType: BaseDataTypeList, EntryType: NewDataType(baseType, DataTypeRankScalar)}
	}
	return &DataType{Name: BaseDataTypeName(baseType), BaseType: baseType}
}

func NewCustomDataType(dataType string, rank DataTypeRank) *DataType {
	if rank.IsList() {
		return &DataType{Name: "list", BaseType: BaseDataTypeList, EntryType: NewCustomDataType(dataType, DataTypeRankScalar)}
	}
	return &DataType{Name: dataType, BaseType: BaseDataTypeCustom}
}

func NewNamedDataType(name string, baseType BaseDataType, rank DataTypeRank) *DataType {
	if rank.IsList() {
		return &DataType{Name: "list", BaseType: BaseDataTypeList, EntryType: NewNamedDataType(name, baseType, DataTypeRankScalar)}
	}
	return &DataType{Name: name, BaseType: baseType}
}

func ParseDataType(typeName string, rank DataTypeRank) *DataType {
	if len(typeName) == 0 {
		return nil
	}
	if rank.IsList() {
		return &DataType{Name: "list", BaseType: BaseDataTypeList, EntryType: ParseDataType(typeName, DataTypeRankScalar)}
	}
	typeName = strings.TrimSpace(typeName)
	//typeName = strings.TrimPrefix(typeName, "ref_")
	//typeName = strings.TrimPrefix(typeName, "DataType")
	dt := &DataType{Name: typeName}

	var nameOverride string
	dt.BaseType, nameOverride = ParseDataTypeName(typeName)
	if nameOverride != "" {
		dt.Name = nameOverride
	}
	return dt
}

func ParseDataTypeName(typeName string) (baseType BaseDataType, name string) {
	switch strings.ToLower(typeName) {
	case "bool", "boolean":
		baseType = BaseDataTypeBoolean
	case "uint8":
		baseType = BaseDataTypeUInt8
	case "uint16":
		baseType = BaseDataTypeUInt16
	case "uint24":
		baseType = BaseDataTypeUInt24
	case "uint32":
		baseType = BaseDataTypeUInt32
	case "uint40":
		baseType = BaseDataTypeUInt40
	case "uint48":
		baseType = BaseDataTypeUInt48
	case "uint56":
		baseType = BaseDataTypeUInt56
	case "uint64":
		baseType = BaseDataTypeUInt64

	case "int8":
		baseType = BaseDataTypeInt8
	case "int16":
		baseType = BaseDataTypeInt16
	case "int24":
		baseType = BaseDataTypeInt24
	case "int32":
		baseType = BaseDataTypeInt32
	case "int40":
		baseType = BaseDataTypeInt40
	case "int48":
		baseType = BaseDataTypeInt48
	case "int56":
		baseType = BaseDataTypeInt56
	case "int64":
		baseType = BaseDataTypeInt64

	case "single":
		baseType = BaseDataTypeSingle
	case "double":
		baseType = BaseDataTypeDouble

	case "enum8":
		baseType = BaseDataTypeEnum8
	case "enum16":
		baseType = BaseDataTypeEnum16

	case "map8":
		baseType = BaseDataTypeMap8
	case "map16":
		baseType = BaseDataTypeMap16
	case "map32":
		baseType = BaseDataTypeMap32
	case "map64":
		baseType = BaseDataTypeMap64
	case "string", "character string":
		baseType = BaseDataTypeString
	case "octstr", "octet string":
		baseType = BaseDataTypeOctStr
	case "percent":
		baseType = BaseDataTypePercent
	case "percent100ths":
		baseType = BaseDataTypePercentHundredths
	case "temperature":
		baseType = BaseDataTypeTemperature
	case "amperage-ma":
		baseType = BaseDataTypeAmperage
	case "voltage-mv":
		baseType = BaseDataTypeVoltage
	case "power-mw":
		baseType = BaseDataTypePower
	case "energy-mwh":
		baseType = BaseDataTypeEnergy
	case "power-mva", "energy-mva":
		baseType = BaseDataTypeApparentPower
	case "energy-mvah":
		baseType = BaseDataTypeApparentEnergy
	case "power-mvar", "energy-mvar":
		baseType = BaseDataTypeReactivePower
	case "energy-mvarh":
		baseType = BaseDataTypeReactiveEnergy
	case "money":
		baseType = BaseDataTypeMoney
	case "elapsed-s":
		baseType = BaseDataTypeElapsedSeconds
	case "epoch-s", "utc": // utc is deprecated
		baseType = BaseDataTypeEpochSeconds
	case "epoch-us", "epochus", "epoch time in microseconds":
		baseType = BaseDataTypeEpochMicroseconds
	case "systime_ms", "systime-ms", "systemtimems", "system time in milliseconds":
		baseType = BaseDataTypeSystimeMilliseconds
	case "systime_us", "systime-us", "systemtimeus", "system time in microseconds":
		baseType = BaseDataTypeSystimeMicroseconds
	case "posix-ms", "posixms", "posix time in milliseconds":
		baseType = BaseDataTypePosixMilliseconds
	case "date":
		baseType = BaseDataTypeDate
	case "action-id":
		baseType = BaseDataTypeActionID
	case "attrib-id", "attribute-id":
		baseType = BaseDataTypeAttributeID
	case "cluster-id":
		baseType = BaseDataTypeClusterID
	case "command-id":
		baseType = BaseDataTypeCommandID
	case "data-ver":
		baseType = BaseDataTypeDataVersion
	case "devtype-id":
		baseType = BaseDataTypeDeviceTypeID
	case "entry-idx", "entryidx":
		baseType = BaseDataTypeEntryIndex
	case "event-id", "eventid":
		baseType = BaseDataTypeEventID
	case "event-no", "eventnumber":
		baseType = BaseDataTypeEventNumber
	case "fabric-id", "fabricid":
		baseType = BaseDataTypeFabricID
	case "fabric-idx", "fabricidx":
		baseType = BaseDataTypeFabricIndex
	case "field-id", "fieldid", "field id":
		baseType = BaseDataTypeFieldID
	case "group-id", "groupid", "group id":
		baseType = BaseDataTypeGroupID
	case "node-id", "nodeid", "node id":
		baseType = BaseDataTypeNodeID
	case "subject-id", "subjectid", "subject id":
		baseType = BaseDataTypeSubjectID
	case "transaction-id":
		baseType = BaseDataTypeTransactionID
	case "vendor-id", "vendorid", "vendor id":
		baseType = BaseDataTypeVendorID
	case "endpoint-id":
		baseType = BaseDataTypeEndpointID
	case "endpoint-no", "endpointnumber", "endpoint number":
		baseType = BaseDataTypeEndpointNumber
	case "eui64":
		baseType = BaseDataTypeIeeeAddress
	case "temperaturedifference", "tempdiff":
		baseType = BaseDataTypeTemperatureDifference
	case "unsignedtemperature":
		baseType = BaseDataTypeUnsignedTemperature
	case "signedtemperature":
		baseType = BaseDataTypeSignedTemperature
	case "hwadr", "hardware address":
		baseType = BaseDataTypeHardwareAddress
	case "ipadr":
		baseType = BaseDataTypeIPAddress
	case "ipv4adr", "ipv4address", "ipv4 address":
		baseType = BaseDataTypeIPv4Address
	case "ipv6adr", "ipv6address", "ipv6 address":
		baseType = BaseDataTypeIPv6Address
	case "ipv6pre", "ipv6 prefix":
		baseType = BaseDataTypeIPv6Prefix
	case "semtag":
		baseType = BaseDataTypeCustom
		name = "SemanticTagStruct"
	case "status", "statuscode", "status code":
		baseType = BaseDataTypeStatus
	case "priority":
		baseType = BaseDataTypePriority
	case "messageid":
		baseType = BaseDataTypeMessageID
	case "tag":
		baseType = BaseDataTypeTag
	case "namespace":
		baseType = BaseDataTypeNamespaceID
	case "locationdesc":
		baseType = BaseDataTypeCustom
		name = "LocationDescriptorStruct"
	case "currency", "currencystruct":
		baseType = BaseDataTypeCustom
		name = "CurrencyStruct"
	case "price":
		baseType = BaseDataTypeCustom
		name = "PriceStruct"
	case "threelevelautoenum":
		baseType = BaseDataTypeCustom
		name = "ThreeLevelAutoEnum"
	case "list":
		baseType = BaseDataTypeList
	default:
		baseType = BaseDataTypeCustom
	}
	return
}

func (dt *DataType) Clone() *DataType {
	ndt := &DataType{Name: dt.Name, BaseType: dt.BaseType, Source: dt.Source}
	if dt.EntryType != nil {
		ndt.EntryType = dt.EntryType.Clone()
	}
	return ndt
}

func (dt *DataType) HasLength() bool {
	return dt != nil && (dt.BaseType == BaseDataTypeString || dt.BaseType == BaseDataTypeOctStr || dt.BaseType == BaseDataTypeMessageID)
}

func (dt *DataType) IsArray() bool {
	return dt != nil && dt.BaseType == BaseDataTypeList
}

func (dt *DataType) IsCustom() bool {
	return dt != nil && dt.BaseType == BaseDataTypeCustom
}

func (dt *DataType) IsMap() bool {
	if dt == nil {
		return false
	}
	switch dt.BaseType {
	case BaseDataTypeMap8, BaseDataTypeMap16, BaseDataTypeMap32, BaseDataTypeMap64:
		return true
	}
	return false
}

func (dt *DataType) IsEnum() bool {
	if dt == nil {
		return false
	}
	switch dt.BaseType {
	case BaseDataTypeEnum8, BaseDataTypeEnum16:
		return true
	}
	return false
}

func (dt *DataType) Equals(other *DataType) bool {
	if dt == nil {
		return other == nil
	} else if other == nil {
		return false
	}
	if dt.BaseType != other.BaseType {
		return false
	}
	if dt.Name != other.Name {
		return false
	}
	if dt.EntryType == nil {
		return other.EntryType == nil
	}
	if other.EntryType == nil {
		return false
	}
	return dt.EntryType.Equals(other.EntryType)
}
