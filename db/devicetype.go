package db

import (
	"context"

	"github.com/hasty/alchemy/matter"
)

func (h *Host) indexDeviceTypeModel(cxt context.Context, parent *sectionInfo, deviceType *matter.DeviceType) error {
	deviceTypeRow := newDBRow()
	deviceTypeRow.values[matter.TableColumnID] = deviceType.ID.IntString()
	deviceTypeRow.values[matter.TableColumnName] = deviceType.Name
	deviceTypeRow.values[matter.TableColumnSuperset] = deviceType.Superset
	deviceTypeRow.values[matter.TableColumnClass] = deviceType.Class
	deviceTypeRow.values[matter.TableColumnScope] = deviceType.Scope

	dti := &sectionInfo{id: h.nextId(deviceTypeTable), parent: parent, values: deviceTypeRow, children: make(map[string][]*sectionInfo)}

	for _, r := range deviceType.Revisions {
		revisionRow := newDBRow()
		revisionRow.values[matter.TableColumnID] = r.Number
		revisionRow.values[matter.TableColumnDescription] = r.Description
		fci := &sectionInfo{id: h.nextId(deviceTypeRevisionTable), parent: dti, values: revisionRow}
		dti.children[deviceTypeRevisionTable] = append(dti.children[deviceTypeRevisionTable], fci)

	}
	for _, c := range deviceType.Conditions {
		revisionRow := newDBRow()
		revisionRow.values[matter.TableColumnFeature] = c.Feature
		revisionRow.values[matter.TableColumnDescription] = c.Description
		fci := &sectionInfo{id: h.nextId(deviceTypeConditionTable), parent: dti, values: revisionRow}
		dti.children[deviceTypeConditionTable] = append(dti.children[deviceTypeConditionTable], fci)

	}

	for _, c := range deviceType.ClusterRequirements {
		row := newDBRow()
		row.values[matter.TableColumnID] = c.ID.IntString()
		row.values[matter.TableColumnName] = c.ClusterName
		row.values[matter.TableColumnQuality] = c.Quality.String()
		if c.Conformance != nil {
			row.values[matter.TableColumnConformance] = c.Conformance.AsciiDocString()
		}
		switch c.Interface {
		case matter.InterfaceClient:
			row.values[matter.TableColumnDirection] = "client"
		case matter.InterfaceServer:
			row.values[matter.TableColumnDirection] = "server"
		default:
			row.values[matter.TableColumnDirection] = "unknown"

		}
		fci := &sectionInfo{id: h.nextId(deviceTypeClusterRequirementTable), parent: dti, values: row}
		dti.children[deviceTypeClusterRequirementTable] = append(dti.children[deviceTypeClusterRequirementTable], fci)

	}
	parent.children[deviceTypeTable] = append(parent.children[deviceTypeTable], dti)
	return nil
}
