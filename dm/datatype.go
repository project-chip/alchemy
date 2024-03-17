package dm

import (
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/constraint"
	"github.com/hasty/alchemy/matter/types"
)

func dataModelName(dataType *types.DataType) string {
	if dataType.IsEnum() || dataType.IsMap() {
		return dataType.Name
	}
	switch dataType.BaseType {
	case types.BaseDataTypeCustom:
		return dataType.Name
	case types.BaseDataTypeEpochSeconds:
		return "epoch-s"
	case types.BaseDataTypeEpochMicroseconds:
		return "epoch-us"
	case types.BaseDataTypeSystimeMicroseconds:
		return "systemtime-us"
	case types.BaseDataTypeAmperage:
		return "amperage-ma"
	case types.BaseDataTypeVoltage:
		return "voltage-mv"
	case types.BaseDataTypeEnergy:
		return "energy-mwh"
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
		return "fabric-index"
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
	default:
		return strings.ToLower(dataType.Name)
	}
}

func renderDataTypes(doc *ascii.Doc, cluster *matter.Cluster, c *etree.Element) (err error) {
	if len(cluster.Enums) == 0 && len(cluster.Bitmaps) == 0 && len(cluster.Structs) == 0 {
		return
	}
	dt := c.CreateElement("dataTypes")
	err = renderEnums(doc, cluster, dt)
	if err != nil {
		return
	}
	err = renderBitmaps(doc, cluster, dt)
	if err != nil {
		return
	}

	err = renderStructs(doc, cluster, dt)
	return
}

func renderDataType(f *matter.Field, i *etree.Element) {
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
	if lc, ok := f.Constraint.(*constraint.ListConstraint); ok {
		renderConstraint(lc.EntryConstraint, f.Type.EntryType, e)
	}
}

func renderDefault(fs matter.FieldSet, f *matter.Field, e *etree.Element) {
	if f.Default == "" {
		return
	}
	cons, err := constraint.ParseString(f.Default)
	if err != nil {
		cons = &constraint.GenericConstraint{Value: f.Default}
		return
	}
	ec, ok := cons.(*constraint.ExactConstraint)
	if ok {
		switch limit := ec.Value.(type) {
		case *constraint.ManufacturerLimit:
			e.CreateAttr("default", "MS")
			return
		case *constraint.ReferenceLimit:
			e.CreateAttr("default", limit.Value)
			return
		}
	}
	def := cons.Default(&matter.ConstraintContext{Fields: fs, Field: f})
	if !def.Defined() {
		return
	}
	e.CreateAttr("default", def.DataModelString(f.Type))
}
