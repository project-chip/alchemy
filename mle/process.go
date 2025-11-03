package mle

import (
	"fmt"
	"path/filepath"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func Process(root string, s *spec.Specification) (violations map[string][]spec.Violation, err error) {

	masterClusterMap, reservedClusterIDs, vc, err := parseMasterClusterList(filepath.Join(root, "master_lists", "MasterClusterList.adoc"))
	if err != nil {
		return
	}
	masterDeviceTypeMap, reservedDeviceTypeIDs, vd, err := parseMasterDeviceTypeList(filepath.Join(root, "master_lists", "MasterDeviceTypeList.adoc"))
	if err != nil {
		return
	}

	violations = spec.MergeViolations(vc, vd)

	clusters := make([]*matter.Cluster, 0)
	for c := range s.Clusters {
		clusters = append(clusters, c)
	}
	deviceTypes := s.DeviceTypes

	for _, c := range clusters {
		if ci, ok := masterClusterMap[c.Name]; ok {
			if ci.ClusterID != c.ID.HexString() {
				v := spec.Violation{Entity: c, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster ID mismatch. Cluster='%s' ID in data model='%s' ID on master list='%s'", c.Name, c.ID.HexString(), ci.ClusterID)}
				v.Path, v.Line = c.Origin()
				violations[v.Path] = append(violations[v.Path], v)
			}
			if ci.PICScode != c.PICS {
				v := spec.Violation{Entity: c, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster PICS mismatch. Cluster='%s' PICS in data model='%s' PICS on master list='%s'", c.Name, c.PICS, ci.PICScode)}
				v.Path, v.Line = c.Origin()
				violations[v.Path] = append(violations[v.Path], v)
			}
		} else {
			if taken, name := clusterIdTaken(c.ID.HexString(), masterClusterMap); taken {
				v := spec.Violation{Entity: c, Type: spec.ViolationMasterList, Text: fmt.Sprintf("This Cluster ID is present on Master List with another name. ID='%s' cluster name='%s' Name on master list='%s'", c.ID.HexString(), c.Name, name)}
				v.Path, v.Line = c.Origin()
				violations[v.Path] = append(violations[v.Path], v)
			} else if reserved := clusterIdReserved(c.ID.HexString(), reservedClusterIDs); reserved {
				v := spec.Violation{Entity: c, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster ID is reserved. ID='%s' Violating cluster name='%s' ", c.ID.HexString(), c.Name)}
				v.Path, v.Line = c.Origin()
				violations[v.Path] = append(violations[v.Path], v)
			}
		}
	}

	for _, dt := range deviceTypes {
		if dti, ok := masterDeviceTypeMap[dt.Name]; ok {
			if dti.DeviceTypeID != dt.ID.HexString() {
				v := spec.Violation{Entity: dt, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Device Type ID mismatch. Device Type Name='%s' ID in data model='%s' ID on master list='%s'", dt.Name, dt.ID.HexString(), dti.DeviceTypeID)}
				v.Path, v.Line = dt.Origin()
				violations[v.Path] = append(violations[v.Path], v)
			}
		} else {
			if taken, name := deviceTypeIdTaken(dt.ID.HexString(), masterDeviceTypeMap); taken {
				v := spec.Violation{Entity: dt, Type: spec.ViolationMasterList, Text: fmt.Sprintf("This Device Type ID is present on Master List with another name. ID='%s' device type name='%s' Name on master list='%s'", dt.ID.HexString(), dt.Name, name)}
				v.Path, v.Line = dt.Origin()
				violations[v.Path] = append(violations[v.Path], v)
			} else if reserved := deviceTypeIdReserved(dt.ID.HexString(), reservedDeviceTypeIDs); reserved {
				v := spec.Violation{Entity: dt, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Device Type ID is reserved. ID='%s' Violating device type name='%s' ", dt.ID.HexString(), dt.Name)}
				v.Path, v.Line = dt.Origin()
				violations[v.Path] = append(violations[v.Path], v)
			}
		}
	}

	return
}
