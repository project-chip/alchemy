package mle

import (
	"fmt"
	"path/filepath"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter/spec"
)

func Process(root string, s *spec.Specification) (violations map[string][]spec.Violation, err error) {
	var masterClusterListPath, masterDeviceTypeListPath asciidoc.Path
	masterClusterListPath, err = asciidoc.NewPath(filepath.Join(root, "master_lists/MasterClusterList.adoc"), root)
	if err != nil {
		return
	}
	masterDeviceTypeListPath, err = asciidoc.NewPath(filepath.Join(root, "master_lists/MasterDeviceTypeList.adoc"), root)
	if err != nil {
		return
	}

	masterClusterMap, reservedClusterIDs, vc, err := parseMasterClusterList(masterClusterListPath)
	if err != nil {
		return
	}
	masterDeviceTypeMap, reservedDeviceTypeIDs, vd, err := parseMasterDeviceTypeList(masterDeviceTypeListPath)
	if err != nil {
		return
	}

	violations = spec.MergeViolations(vc, vd)

	for c := range s.Clusters {
		if !c.ID.Valid() {
			continue
		}

		if source, isReserved := reservedClusterIDs[c.ID.Value()]; isReserved {
			v := spec.Violation{Entity: c, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster ID is reserved. ID='%s' Violating cluster name='%s' ", c.ID.HexString(), c.Name)}
			v.Path, v.Line = source.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}
		ci, ok := masterClusterMap[c.ID.Value()]
		if ok {
			if ci.Name != c.Name {
				v := spec.Violation{Entity: c, Type: spec.ViolationMasterList, Text: fmt.Sprintf("This Cluster ID is present on Master List with another name. ID='%s' cluster name='%s' Name on master list='%s'", c.ID.HexString(), c.Name, ci.Name)}
				v.Path, v.Line = c.Origin()
				violations[v.Path] = append(violations[v.Path], v)

			}
			if ci.PICScode != c.PICS {
				v := spec.Violation{Entity: c, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Cluster PICS mismatch. Cluster='%s' PICS in data model='%s' PICS on master list='%s'", c.Name, c.PICS, ci.PICScode)}
				v.Path, v.Line = c.Origin()
				violations[v.Path] = append(violations[v.Path], v)
			}
		} else {
			v := spec.Violation{Entity: c, Type: spec.ViolationMasterList, Text: fmt.Sprintf("This Cluster ID is missing from the Master List of clusters. ID='%s' cluster name='%s'", c.ID.HexString(), c.Name)}
			v.Path, v.Line = c.Origin()
			violations[v.Path] = append(violations[v.Path], v)
		}
	}

	for _, dt := range s.DeviceTypes {
		if !dt.ID.Valid() {
			continue
		}
		if source, isReserved := reservedDeviceTypeIDs[dt.ID.Value()]; isReserved {
			v := spec.Violation{Entity: dt, Type: spec.ViolationMasterList, Text: fmt.Sprintf("Device Type ID is reserved. ID='%s' Violating device type name='%s' ", dt.ID.HexString(), dt.Name)}
			v.Path, v.Line = source.Origin()
			violations[v.Path] = append(violations[v.Path], v)
			continue
		}
		di, ok := masterDeviceTypeMap[dt.ID.Value()]
		if ok {
			if di.Name != dt.Name {
				v := spec.Violation{Entity: dt, Type: spec.ViolationMasterList, Text: fmt.Sprintf("This Device Type ID is present on Master List with another name. ID='%s' device type name='%s' Name on master list='%s'", dt.ID.HexString(), dt.Name, di.Name)}
				v.Path, v.Line = dt.Origin()
				violations[v.Path] = append(violations[v.Path], v)

			}
		} else {
			v := spec.Violation{Entity: dt, Type: spec.ViolationMasterList, Text: fmt.Sprintf("This device type is missing from the Master List of device types. ID='%s' device type name='%s'", dt.ID.HexString(), dt.Name)}
			v.Path, v.Line = dt.Origin()
			violations[v.Path] = append(violations[v.Path], v)
		}
	}
	return
}
