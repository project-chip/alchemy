package db

import (
	"context"

	"github.com/project-chip/alchemy/matter"
)

func (h *Host) indexDeviceTypeModel(cxt context.Context, parent *sectionInfo, deviceType *matter.DeviceType) error {
	deviceTypeRow := newDBRow()
	deviceTypeRow.values[matter.TableColumnID] = deviceType.ID.IntString()
	deviceTypeRow.values[matter.TableColumnName] = deviceType.Name
	deviceTypeRow.values[matter.TableColumnSupersetOf] = deviceType.SupersetOf
	deviceTypeRow.values[matter.TableColumnClass] = deviceType.Class
	deviceTypeRow.values[matter.TableColumnScope] = deviceType.Scope

	dti := h.newSectionInfo(deviceTypeTable, parent, deviceTypeRow, deviceType)

	for _, r := range deviceType.Revisions {
		revisionRow := newDBRow()
		revisionRow.values[matter.TableColumnID] = r.Number
		revisionRow.values[matter.TableColumnDescription] = r.Description
		fci := h.newSectionInfo(deviceTypeRevisionTable, dti, revisionRow, r)
		dti.children[deviceTypeRevisionTable] = append(dti.children[deviceTypeRevisionTable], fci)

	}
	for _, c := range deviceType.Conditions {
		revisionRow := newDBRow()
		revisionRow.values[matter.TableColumnFeature] = c.Feature
		revisionRow.values[matter.TableColumnDescription] = c.Description
		fci := h.newSectionInfo(deviceTypeConditionTable, dti, revisionRow, c)
		dti.children[deviceTypeConditionTable] = append(dti.children[deviceTypeConditionTable], fci)

	}

	for _, c := range deviceType.ClusterRequirements {
		row := newDBRow()
		row.values[matter.TableColumnID] = c.ClusterID.IntString()
		row.values[matter.TableColumnName] = c.ClusterName
		row.values[matter.TableColumnQuality] = c.Quality.String()
		if c.Conformance != nil {
			row.values[matter.TableColumnConformance] = c.Conformance.ASCIIDocString()
		}
		switch c.Interface {
		case matter.InterfaceClient:
			row.values[matter.TableColumnDirection] = "client"
		case matter.InterfaceServer:
			row.values[matter.TableColumnDirection] = "server"
		default:
			row.values[matter.TableColumnDirection] = "unknown"

		}
		fci := h.newSectionInfo(deviceTypeClusterRequirementTable, dti, row, c)
		dti.children[deviceTypeClusterRequirementTable] = append(dti.children[deviceTypeClusterRequirementTable], fci)

	}
	parent.children[deviceTypeTable] = append(parent.children[deviceTypeTable], dti)

	for _, dr := range deviceType.DeviceTypeRequirements {
		row := newDBRow()
		row.values[matter.TableColumnID] = dr.DeviceTypeID.IntString()
		row.values[matter.TableColumnName] = dr.DeviceTypeName
		if dr.Conformance != nil {
			row.values[matter.TableColumnConformance] = dr.Conformance.ASCIIDocString()
		}
		if dr.Constraint != nil {
			row.values[matter.TableColumnConstraint] = dr.Constraint.ASCIIDocString(nil)

		}
		fci := h.newSectionInfo(deviceTypeCompositionTable, dti, row, dr)
		dti.children[deviceTypeCompositionTable] = append(dti.children[deviceTypeCompositionTable], fci)

	}
	return nil
}
