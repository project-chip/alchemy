package dm

import (
	"log/slog"
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/types"
)

func dataModelName(dataType *types.DataType) string {
	if dataType.Entity != nil {
		switch e := dataType.Entity.(type) {
		case *matter.Bitmap:
			return e.Name
		case *matter.Enum:
			return e.Name
		case *matter.Struct:
			return e.Name
		case *matter.Command:
			return e.Name
		case *matter.TypeDef:
			return e.Name
		}
	}
	if dataType.IsEnum() || dataType.IsMap() {
		return dataType.Name
	}
	switch dataType.BaseType {
	case types.BaseDataTypeCustom:
		return dataType.Name
	case types.BaseDataTypeBoolean:
		return "bool"
	case types.BaseDataTypeUInt8:
		return "uint8"
	case types.BaseDataTypeUInt16:
		return "uint16"
	case types.BaseDataTypeUInt24:
		return "uint24"
	case types.BaseDataTypeUInt32:
		return "uint32"
	case types.BaseDataTypeUInt40:
		return "uint40"
	case types.BaseDataTypeUInt48:
		return "uint48"
	case types.BaseDataTypeUInt56:
		return "uint56"
	case types.BaseDataTypeUInt64:
		return "uint64"
	case types.BaseDataTypeInt8:
		return "int8"
	case types.BaseDataTypeInt16:
		return "int16"
	case types.BaseDataTypeInt24:
		return "int24"
	case types.BaseDataTypeInt32:
		return "int32"
	case types.BaseDataTypeInt40:
		return "int40"
	case types.BaseDataTypeInt48:
		return "int48"
	case types.BaseDataTypeInt56:
		return "int56"
	case types.BaseDataTypeInt64:
		return "int64"
	case types.BaseDataTypeSingle:
		return "single"
	case types.BaseDataTypeDouble:
		return "double"
	case types.BaseDataTypeString:
		return "string"
	case types.BaseDataTypeOctStr:
		return "octstr"
	case types.BaseDataTypeEpochSeconds:
		return "epoch-s"
	case types.BaseDataTypeEpochMicroseconds:
		return "epoch-us"
	case types.BaseDataTypeElapsedSeconds:
		return "elapsed-s"
	case types.BaseDataTypeSystimeMilliseconds:
		return "systime-ms"
	case types.BaseDataTypePosixMilliseconds:
		return "posix-ms"
	case types.BaseDataTypeSystimeMicroseconds:
		return "systemtime-us"
	case types.BaseDataTypeAmperage:
		return "amperage-mA"
	case types.BaseDataTypeVoltage:
		return "voltage-mV"
	case types.BaseDataTypePower:
		return "power-mW"
	case types.BaseDataTypeEnergy:
		return "energy-mWh"
	case types.BaseDataTypeApparentPower:
		return "power-mVA"
	case types.BaseDataTypeApparentEnergy:
		return "energy-mVAh"
	case types.BaseDataTypeReactivePower:
		return "power-mVAR"
	case types.BaseDataTypeReactiveEnergy:
		return "energy-mVARh"
	case types.BaseDataTypeMoney:
		return "money"
	case types.BaseDataTypeVendorID:
		return "vendor-id"
	case types.BaseDataTypeSubjectID:
		return "subject-id"
	case types.BaseDataTypeNodeID:
		return "node-id"
	case types.BaseDataTypeGroupID:
		return "group-id"
	case types.BaseDataTypeFabricID:
		return "fabric-id"
	case types.BaseDataTypeFabricIndex:
		return "fabric-idx"
	case types.BaseDataTypeActionID:
		return "action-id"
	case types.BaseDataTypeEndpointNumber:
		return "endpoint-no"
	case types.BaseDataTypeSignedTemperature:
		return "int8s"
	case types.BaseDataTypeUnsignedTemperature:
		return "uint8"
	case types.BaseDataTypeTemperatureDifference:
		return "int16s"
	case types.BaseDataTypeSemanticTag:
		return "semtag"
	case types.BaseDataTypeHardwareAddress:
		return "hwadr"
	case types.BaseDataTypeIPAddress:
		return "ipadr"
	case types.BaseDataTypeIPv4Address:
		return "ipv4adr"
	case types.BaseDataTypeIPv6Address:
		return "ipv6adr"
	case types.BaseDataTypeIPv6Prefix:
		return "ipv6pre"
	case types.BaseDataTypePercent:
		return "percent"
	case types.BaseDataTypePercentHundredths:
		return "percent100ths"
	case types.BaseDataTypeTemperature:
		return "temperature"
	case types.BaseDataTypeMessageID:
		return "message-id"
	case types.BaseDataTypeClusterID:
		return "cluster-id"
	case types.BaseDataTypeAttributeID:
		return "attribute-id"
	case types.BaseDataTypeDeviceTypeID:
		return "devtype-id"
	case types.BaseDataTypeStatus:
		return "status"
	case types.BaseDataTypeEndpointID:
		return "endpoint-id"
	case types.BaseDataTypeTag:
		return "tag"
	case types.BaseDataTypeNamespaceID:
		return "namespace"
	default:
		slog.Warn("unknown data model type", "name", dataType.Name)
		return dataType.Name
	}
}

func renderDataTypes(cluster *matter.Cluster, c *etree.Element) (err error) {
	if len(cluster.Enums) == 0 && len(cluster.Bitmaps) == 0 && len(cluster.Structs) == 0 {
		return
	}
	dt := c.CreateElement("dataTypes")
	err = renderTypeDefs(cluster.TypeDefs, dt)
	if err != nil {
		return
	}
	err = renderEnums(cluster.Enums, dt)
	if err != nil {
		return
	}
	err = renderBitmaps(cluster.Bitmaps, dt)
	if err != nil {
		return
	}

	err = renderStructs(cluster.Structs, dt)
	return
}

func renderDataType(f *matter.Field, i *etree.Element) (err error) {
	if f.Type == nil {
		return
	}
	if !f.Type.IsArray() {
		i.CreateAttr("type", dataModelName(f.Type))
		return
	}
	i.CreateAttr("type", "list")
	e := i.CreateElement("entry")
	e.CreateAttr("type", dataModelName(f.Type.EntryType))
	err = renderListConstraint(f.Constraint, f.Type.EntryType, e)
	return
}

func renderListConstraint(c constraint.Constraint, entryType *types.DataType, parent *etree.Element) (err error) {
	switch c := c.(type) {
	case constraint.Set:
		for _, c := range c {
			err = renderListConstraint(c, entryType, parent)
			if err != nil {
				return
			}
		}
	case *constraint.ListConstraint:
		err = renderConstraint(c.EntryConstraint, entryType, parent, nil)
	}
	return
}

func renderFallback(fs matter.FieldSet, f *matter.Field, e *etree.Element) {
	if constraint.IsBlankLimit(f.Fallback) || constraint.IsGenericLimit(f.Fallback) {
		return
	}
	switch fb := f.Fallback.(type) {
	case *constraint.ManufacturerLimit:
		e.CreateAttr("default", "MS")
		return
	case *constraint.ReferenceLimit:
		e.CreateAttr("default", fb.Reference)
		return
	case *constraint.StatusCodeLimit:
		e.CreateAttr("default", fb.StatusCode.String())
		return
	case *constraint.IdentifierLimit:
		switch entity := fb.Entity.(type) {
		case *matter.EnumValue:
			e.CreateAttr("default", entity.Name)
			return
		default:
			if entity != nil {
				slog.Warn("Identifier fallback with unexpected entity", log.Path("field", f), log.Type("type", entity))
			}
		}
	}
	def := f.Fallback.Fallback(matter.NewConstraintContext(f, fs))
	if !def.Defined() {
		return
	}
	e.CreateAttr("default", def.DataModelString(f.Type))
}

func renderTypeDefs(typeDefs []*matter.TypeDef, dt *etree.Element) (err error) {
	tds := make([]*matter.TypeDef, len(typeDefs))
	copy(tds, typeDefs)
	slices.SortStableFunc(tds, func(a, b *matter.TypeDef) int {
		return strings.Compare(a.Name, b.Name)
	})
	for _, e := range tds {
		en := dt.CreateElement("number")
		en.CreateAttr("name", e.Name)
		en.CreateAttr("type", dataModelName(e.Type))
	}
	return
}
