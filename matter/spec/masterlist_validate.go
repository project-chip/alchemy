package spec

import (
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func validateMasterLists(spec *Specification) {
	for entityType, ml := range spec.MasterLists {
		switch entityType {
		case types.EntityTypeCluster:
			validateClusterMasterList(spec, ml)
		case types.EntityTypeDeviceType:
			validateDeviceTypeMasterList(spec, ml)
		case types.EntityTypeNamespace:
			validateNamespaceMasterList(spec, ml)
		}
	}
}

func validateClusterMasterList(spec *Specification, ml *matter.MasterList) {
	iu := make(idUniqueness[*matter.MasterListEntry])
	nu := make(nameUniqueness[*matter.MasterListEntry])
	picsu := make(picsUniqueness[*matter.MasterListEntry])
	for _, mle := range ml.Entries {
		iu.check(spec, mle.ID, mle)
		nu.check(spec, mle)
		if mle.PICS != "" {
			picsu.check(spec, mle.PICS, mle)
		}
	}
	for cluster := range spec.Clusters {
		if cluster.ID.Text() == "n/a" {
			continue
		}
		var mle *matter.MasterListEntry
		if cluster.ID.Valid() {
			mle = iu[cluster.ID.Value()]
		} else {
			mle = nu[cluster.Name]
		}
		if mle == nil {
			spec.addError(&MasterListMissingEntryError{Entity: cluster})
		} else {
			switch mle.Name {
			case "Reserved":
				spec.addError(&MasterListReservedIDError{Entity: cluster})
			default:
				if mle.Name != cluster.Name {
					spec.addError(&ReferenceNameMismatchError[*matter.Cluster, *matter.MasterListEntry]{Source: mle, Entity: cluster, Name: mle.Name})
				}
			}
			if mle.PICS != cluster.PICS {
				spec.addError(&MasterListMismatchError{Entity: cluster, Property: "PICS", EntityValue: cluster.PICS, ListValue: mle.PICS})
			}
			if cluster.Hierarchy != "Base" {
				if cluster.ParentCluster != nil && cluster.ParentCluster.PICS != mle.Base {
					spec.addError(&MasterListMismatchError{Entity: cluster, Property: "Base", EntityValue: cluster.ParentCluster.PICS, ListValue: mle.Base})
				}
			} else if mle.Base != "" {
				spec.addError(&MasterListMismatchError{Entity: cluster, Property: "Base", EntityValue: cluster.Hierarchy, ListValue: mle.Base})
			}
			if cluster.Role != mle.Classification {
				spec.addError(&MasterListMismatchError{Entity: cluster, Property: "Classification", EntityValue: cluster.Role.String(), ListValue: mle.Classification.String()})
			}
			if cluster.Scope != mle.Scope {
				spec.addError(&MasterListMismatchError{Entity: cluster, Property: "Scope", EntityValue: cluster.Scope.String(), ListValue: mle.Scope.String()})
			}
		}
	}
}

func validateDeviceTypeMasterList(spec *Specification, ml *matter.MasterList) {
	iu := make(idUniqueness[*matter.MasterListEntry])
	nu := make(nameUniqueness[*matter.MasterListEntry])
	for _, mle := range ml.Entries {
		iu.check(spec, mle.ID, mle)
		nu.check(spec, mle)
	}
	for _, deviceType := range spec.DeviceTypes {
		var mle *matter.MasterListEntry
		if deviceType.ID.Valid() {
			mle = iu[deviceType.ID.Value()]
		} else {
			mle = nu[deviceType.Name]
		}
		if mle == nil {
			spec.addError(&MasterListMissingEntryError{Entity: deviceType})
		} else {
			switch mle.Name {
			case "Reserved":
				spec.addError(&MasterListReservedIDError{Entity: deviceType})
			default:
				if mle.Name != deviceType.Name {
					spec.addError(&ReferenceNameMismatchError[*matter.DeviceType, *matter.MasterListEntry]{Source: mle, Entity: deviceType, Name: mle.Name})
				}
			}
			if deviceType.Scope != mle.Scope {
				spec.addError(&MasterListMismatchError{Entity: deviceType, Property: "Scope", EntityValue: deviceType.Scope.String(), ListValue: mle.Scope.String()})
			}
		}
	}
}

func validateNamespaceMasterList(spec *Specification, ml *matter.MasterList) {
	iu := make(idUniqueness[*matter.MasterListEntry])
	nu := make(nameUniqueness[*matter.MasterListEntry])
	for _, mle := range ml.Entries {
		iu.check(spec, mle.ID, mle)
		nu.check(spec, mle)
	}
	for _, ns := range spec.Namespaces {
		var mle *matter.MasterListEntry
		if ns.ID.Valid() {
			mle = iu[ns.ID.Value()]
		} else {
			mle = nu[ns.Name]
		}
		if mle == nil {
			spec.addError(&MasterListMissingEntryError{Entity: ns})
		} else {
			switch mle.Name {
			case "Reserved":
				spec.addError(&MasterListReservedIDError{Entity: ns})
			default:
				if mle.Name != ns.Name {
					spec.addError(&ReferenceNameMismatchError[*matter.Namespace, *matter.MasterListEntry]{Source: mle, Entity: ns, Name: mle.Name})
				}
			}
		}
	}
}
