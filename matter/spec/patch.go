package spec

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/types"
)

// PatchSpecForSdk is a grab bag of oddities in the spec that need to be corrected for use in the SDK
func PatchSpecForSdk(spec *Specification) error {
	patchDescriptorCluster(spec)
	patchScenesCluster(spec)
	patchLabelCluster(spec)
	patchLevelControlCluster(spec)

	for m := range spec.DocRefs {
		switch v := m.(type) {
		case *matter.ClusterGroup:
			for _, cl := range v.Clusters {
				if hasAtomicAttributes(cl) {
					addAtomicOperations(spec, cl)
				}
			}
		case *matter.Cluster:
			if hasAtomicAttributes(v) {
				addAtomicOperations(spec, v)
			}
		}
	}

	// We have to rebuild these indicies after we make the above changes
	buildClusterReferences(spec)
	associateDeviceTypeRequirementWithClusters(spec)
	return nil
}

func patchScenesCluster(spec *Specification) {
	scenesCluster, ok := spec.ClustersByName["Scenes Management"]
	if !ok {
		slog.Warn("Could not find Scenes cluster")
		return
	}
	var addSceneCommand *matter.Command
	for _, c := range scenesCluster.Commands {
		if strings.EqualFold(c.Name, "AddScene") {
			addSceneCommand = c
			break
		}
	}
	if addSceneCommand == nil {
		slog.Warn("Could not find AddScene command on Scenes cluster")
		return
	}
	for _, f := range addSceneCommand.Fields {
		if strings.EqualFold(f.Name, "ExtensionFieldSetStructs") {
			f.Name = "ExtensionFieldSets"
			break
		}
	}
}

func patchDescriptorCluster(spec *Specification) {
	/* This is a hacky workaround for a spec problem: SemanticTagStruct is defined twice, in two different ways.
	The first is a global struct that's used by the Descriptor cluster
	The second is a cluster-level struct on Mode Select
	Inserting one as a global object, and the other as a struct on Mode Select breaks zap
	*/
	desc, ok := spec.ClustersByName["Descriptor"]
	if !ok {
		slog.Warn("Could not find Descriptor cluster")
		return
	}
	for o := range spec.GlobalObjects {
		s, ok := o.(*matter.Struct)
		if !ok {
			continue
		}

		if s.Name == "SemanticTagStruct" {
			desc.AddStructs(s)
			delete(spec.GlobalObjects, s)
			break
		}
	}
}

func patchLabelCluster(spec *Specification) {
	/*
		Another hacky workaround: the spec defines LabelStruct under a base cluster called Label Cluster, but the
		ZAP XML has this struct under Fixed Label
	*/
	fixedLabelCluster, ok := spec.ClustersByName["Fixed Label"]
	if !ok {
		slog.Warn("Could not find Fixed Label cluster")
		return
	}
	labelCluster, ok := spec.ClustersByName["Label"]
	if !ok {
		slog.Warn("Could not find Label cluster")
		return
	}
	for _, s := range labelCluster.Structs {
		if s.Name == "LabelStruct" {
			fixedLabelCluster.MoveStruct(s)
			break
		}
	}
}

func patchLevelControlCluster(spec *Specification) {
	levelControlCluster, ok := spec.ClustersByID[0x0008]
	if !ok {
		slog.Warn("Unable to patch Level Control cluster; not found")
		return
	}
	// Level Control cluster has a series of OnOff commands that are spec'd to have the same parameters as their non-OnOff versions,
	// but those parameters aren't explicitly listed in the spec, so we copy them over
	for _, c := range levelControlCluster.Commands {
		if !text.HasCaseInsensitiveSuffix(c.Name, "WithOnOff") {
			continue
		}
		if len(c.Fields) > 0 {
			// Assume that if we have fields already, this has been fixed in the spec
			continue
		}
		baseName := text.TrimCaseInsensitiveSuffix(c.Name, "WithOnOff")
		var matchingCommand *matter.Command
		for _, mc := range levelControlCluster.Commands {
			if strings.EqualFold(mc.Name, baseName) {
				matchingCommand = mc
				break
			}
		}
		if matchingCommand == nil {
			slog.Warn("Unable to find matching command for Level Control cluster", slog.String("commandName", c.Name))
			continue
		}
		for _, f := range matchingCommand.Fields {
			c.Fields = append(c.Fields, f.Clone())
		}
	}
}

func hasAtomicAttributes(cluster *matter.Cluster) bool {
	for _, f := range cluster.Attributes {
		if f.Quality.Has(matter.QualityAtomicWrite) {
			return true
		}
	}
	return false
}

func addAtomicOperations(spec *Specification, cluster *matter.Cluster) {

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
