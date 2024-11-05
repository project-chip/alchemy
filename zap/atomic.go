package zap

import (
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func hasAtomicAttributes(cluster *matter.Cluster) bool {
	for _, f := range cluster.Attributes {
		if f.Quality.Has(matter.QualityAtomicWrite) {
			return true
		}
	}
	return false
}

func addAtomicOperations(spec *spec.Specification, cluster *matter.Cluster) {

	var atomicAttributeStatusStruct *matter.Struct
	var atomicRequestTypeEnum *matter.Enum
	for o := range spec.GlobalObjects {
		switch o := o.(type) {
		case *matter.Struct:
			if o.Name == "AtomicAttributeStatusStruct" {
				atomicAttributeStatusStruct = o
			}
		case *matter.Enum:
			if o.Name == "AtomicRequestTypeEnum" {
				atomicRequestTypeEnum = o
			}
		}
	}

	var attributeStatusDataType = types.NewCustomDataType("AtomicAttributeStatusStruct", true)
	attributeStatusDataType.Entity = atomicAttributeStatusStruct

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
				Type:        attributeStatusDataType,
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

	var requestTypeEnumDataType = types.NewCustomDataType("AtomicRequestTypeEnum", false)
	requestTypeEnumDataType.Entity = atomicRequestTypeEnum

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
				Type:        requestTypeEnumDataType,
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
	spec.ClusterRefs.Add(cluster, atomicRequest)
	spec.ClusterRefs.Add(cluster, atomicResponse)
	cluster.Commands = append(cluster.Commands, atomicRequest)
	cluster.Commands = append(cluster.Commands, atomicResponse)

}
