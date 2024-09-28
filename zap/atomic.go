package zap

import (
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

var atomicRequestTypeEnum = &matter.Enum{
	Name: "AtomicRequestTypeEnum",
	Type: types.NewDataType(types.BaseDataTypeEnum8, false),
	Values: matter.EnumValueSet{
		&matter.EnumValue{
			Value:       matter.NewNumber(0),
			Name:        "BeginWrite",
			Summary:     "Begin an atomic write",
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
		&matter.EnumValue{
			Value:       matter.NewNumber(1),
			Name:        "CommitWrite",
			Summary:     "Commit an atomic write",
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
		&matter.EnumValue{
			Value:       matter.NewNumber(2),
			Name:        "RollbackWrite",
			Summary:     "Rollback an atomic write, discarding any pending changes",
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
	},
}

var atomicAttributeStatusStruct = &matter.Struct{
	Name:        "AtomicAttributeStatusStruct",
	Description: "This struct indicates the status of an attribute during an atomic write.",
	Fields: matter.FieldSet{
		&matter.Field{
			ID:   matter.NewNumber(0),
			Name: "AttributeID",
			Type: types.NewDataType(types.BaseDataTypeAttributeID, false),
		},
		&matter.Field{
			ID:   matter.NewNumber(1),
			Name: "StatusCode",
			Type: types.NewDataType(types.BaseDataTypeStatus, false),
		},
	},
}

var atomicResponse = &matter.Command{
	ID:          matter.AtomicResponseCommandID,
	Name:        "AtomicResponse",
	Description: "Returns the status of an atomic write",
	Direction:   matter.InterfaceClient,
	Conformance: conformance.Set{&conformance.Optional{}},
	Access:      matter.Access{Invoke: matter.PrivilegeManage},
	Fields: matter.FieldSet{
		&matter.Field{
			ID:          matter.NewNumber(0),
			Name:        "StatusCode",
			Type:        types.NewDataType(types.BaseDataTypeStatus, false),
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
		&matter.Field{
			ID:          matter.NewNumber(0),
			Name:        "AttributeStatus",
			Type:        types.NewCustomDataType("AtomicAttributeStatusStruct", true),
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
		&matter.Field{
			ID:          matter.NewNumber(0),
			Name:        "Timeout",
			Type:        types.NewDataType(types.BaseDataTypeUInt16, false),
			Conformance: conformance.Set{&conformance.Optional{}},
		},
	},
}

var atomicRequest = &matter.Command{

	ID:          matter.AtomicRequestCommandID,
	Name:        "AtomicRequest",
	Description: "Begins, Commits or Cancels an atomic write",
	Direction:   matter.InterfaceServer,
	Response: &types.DataType{
		Name:     "AtomicResponse",
		BaseType: types.BaseDataTypeCustom,
		Entity:   atomicResponse,
	},
	Conformance: conformance.Set{&conformance.Optional{}},
	Access:      matter.Access{Invoke: matter.PrivilegeManage},
	Fields: matter.FieldSet{
		&matter.Field{
			ID:          matter.NewNumber(0),
			Name:        "RequestType",
			Type:        types.NewCustomDataType("AtomicRequestTypeEnum", false),
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
		&matter.Field{
			ID:          matter.NewNumber(0),
			Name:        "AttributeRequests",
			Type:        types.NewDataType(types.BaseDataTypeAttributeID, true),
			Conformance: conformance.Set{&conformance.Mandatory{}},
		},
		&matter.Field{
			ID:          matter.NewNumber(0),
			Name:        "Timeout",
			Type:        types.NewDataType(types.BaseDataTypeUInt16, false),
			Conformance: conformance.Set{&conformance.Optional{}},
		},
	},
}

func hasAtomicAttributes(cluster *matter.Cluster) bool {
	for _, f := range cluster.Attributes {
		if f.Quality.Has(matter.QualityAtomicWrite) {
			return true
		}
	}
	return false
}
