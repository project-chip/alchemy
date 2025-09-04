package spec

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func toClusterRequirements(d *Doc, s *asciidoc.Section, deviceType *matter.DeviceType) (clusterRequirements []*matter.ClusterRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = newGenericParseError(s, "error reading cluster requirements table: %w", err)
		}
		return
	}
	for row := range ti.ContentRows() {
		var cr *matter.ClusterRequirement
		cr, err = toClusterRequirement(deviceType, ti, row)
		if err != nil {
			return
		}
		clusterRequirements = append(clusterRequirements, cr)
	}
	return
}

func toClusterRequirement(deviceType *matter.DeviceType, ti *TableInfo, row *asciidoc.TableRow) (cr *matter.ClusterRequirement, err error) {
	cr = matter.NewClusterRequirement(deviceType, row)
	cr.ClusterID, err = ti.ReadID(row, matter.TableColumnClusterID, matter.TableColumnID)
	if err != nil {
		return
	}
	cr.ClusterName, err = ti.ReadValue(row, matter.TableColumnClusterName, matter.TableColumnCluster)
	if err != nil {
		return
	}
	if cr.ClusterName == "" {
		cr.ClusterName, _, err = ti.ReadName(row, matter.TableColumnName)
		if err != nil {
			return
		}
	}
	var q string
	q, err = ti.ReadString(row, matter.TableColumnQuality)
	if err != nil {
		return
	}
	var cs string
	cs, err = ti.ReadString(row, matter.TableColumnClientServer)
	if err != nil {
		return
	}
	switch strings.ToLower(cs) {
	case "server":
		cr.Interface = matter.InterfaceServer
	case "client":
		cr.Interface = matter.InterfaceClient
	default:
		slog.Error("Unknown client/server value", "value", cs, log.Path("source", row))
		err = newGenericParseError(row, "unknown client/server value: \"%s\"", cs)
		return
	}
	cr.Quality = matter.ParseQuality(q)
	cr.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
	return
}

func toElementRequirements(d *Doc, s *asciidoc.Section, deviceType *matter.DeviceType) (elementRequirements []*matter.ElementRequirement, clusterRequirements []*matter.ClusterRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = newGenericParseError(s, "error reading element requirements table: %w", err)
		}
		return
	}
	for row := range ti.ContentRows() {
		var er matter.ElementRequirement
		er, err = toElementRequirement(d, ti, row, deviceType)
		if err != nil {
			return
		}
		if er.Element != types.EntityTypeUnknown {
			elementRequirements = append(elementRequirements, &er)
		} else {
			// The element is blank; we previously expressed some kinds of cluster requirements this way

			cr := matter.NewClusterRequirement(deviceType, row)
			cr.Interface = matter.InterfaceServer
			cr.ClusterID = er.ClusterID
			cr.ClusterName = er.ClusterName
			cr.Quality = er.Quality
			if len(er.Conformance) > 0 {
				cr.Conformance = er.Conformance.CloneSet()
			}
			clusterRequirements = append(clusterRequirements, cr)
		}
	}
	return
}

func toDeviceTypeRequirements(d *Doc, s *asciidoc.Section, deviceType *matter.DeviceType) (deviceTypeRequirements []*matter.DeviceTypeRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = newGenericParseError(s, "error reading element requirements table: %w", err)
		}
		return
	}
	for row := range ti.ContentRows() {
		dtr := matter.NewDeviceTypeRequirement(deviceType, row)
		var deviceId string
		deviceId, err = ti.ReadString(row, matter.TableColumnDeviceID, matter.TableColumnID)
		if err != nil {
			return
		}
		if strings.HasSuffix(deviceId, "+") {
			deviceId = deviceId[:len(deviceId)-1]
			dtr.AllowsSuperset = true
		}
		dtr.DeviceTypeID = matter.ParseNumber(deviceId)

		dtr.DeviceTypeName, _, err = ti.ReadName(row, matter.TableColumnDeviceName, matter.TableColumnName)
		if err != nil {
			return
		}
		if strings.HasSuffix(dtr.DeviceTypeName, "+") {
			dtr.AllowsSuperset = true
			dtr.DeviceTypeName = dtr.DeviceTypeName[:len(dtr.DeviceTypeName)-1]
		}
		dtr.Constraint = ti.ReadConstraint(row, matter.TableColumnConstraint)
		dtr.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
		dtr.Location, err = ti.ReadLocation(row, matter.TableColumnLocation)
		if err != nil {
			return
		}
		deviceTypeRequirements = append(deviceTypeRequirements, dtr)
	}
	return
}

func toComposedDeviceTypeClusterRequirements(d *Doc, s *asciidoc.Section, deviceType *matter.DeviceType) (composedClusterRequirements []*matter.DeviceTypeClusterRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = newGenericParseError(s, "error reading composed device type cluster requirements table: %w", err)
		}
		return
	}
	for row := range ti.ContentRows() {
		var cr *matter.ClusterRequirement
		cr, err = toClusterRequirement(deviceType, ti, row)
		if err != nil {
			return
		}

		dtcr := matter.NewDeviceTypeClusterRequirement(deviceType, cr, row)
		dtcr.DeviceTypeID, err = ti.ReadID(row, matter.TableColumnDeviceID)
		if err != nil {
			return
		}
		dtcr.DeviceTypeName, _, err = ti.ReadName(row, matter.TableColumnDeviceName, matter.TableColumnDevice)
		if err != nil {
			return
		}
		composedClusterRequirements = append(composedClusterRequirements, dtcr)
	}
	return
}

func toConditionRequirements(d *Doc, s *asciidoc.Section, deviceType *matter.DeviceType) (conditionRequirements []*matter.ConditionRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = newGenericParseError(s, "error reading composed device type condition requirements table: %w", err)
		}
		return
	}
	for row := range ti.ContentRows() {
		cr := matter.NewConditionRequirement(deviceType, row)
		cr.DeviceTypeID, err = ti.ReadID(row, matter.TableColumnDeviceID)
		if err != nil {
			return
		}
		cr.DeviceTypeName, _, err = ti.ReadName(row, matter.TableColumnDeviceName, matter.TableColumnDevice)
		if err != nil {
			return
		}
		cr.ConditionName, _, err = ti.ReadName(row, matter.TableColumnCondition)
		if err != nil {
			return
		}
		cr.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
		conditionRequirements = append(conditionRequirements, cr)
	}
	return
}

func toComposedDeviceTypeElementRequirements(d *Doc, s *asciidoc.Section, deviceType *matter.DeviceType) (composedElementRequirements []*matter.DeviceTypeElementRequirement, composedClusterRequirements []*matter.DeviceTypeClusterRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = newGenericParseError(s, "error reading composed device element requirements table: %w", err)
		}
		return
	}
	for row := range ti.ContentRows() {
		var er matter.ElementRequirement
		er, err = toElementRequirement(d, ti, row, deviceType)
		if err != nil {
			return
		}
		dter := matter.NewDeviceTypeElementRequirement(deviceType, &er, row)
		dter.DeviceTypeID, err = ti.ReadID(row, matter.TableColumnDeviceID)
		if err != nil {
			return
		}
		dter.DeviceTypeName, _, err = ti.ReadName(row, matter.TableColumnDeviceName, matter.TableColumnDevice)
		if err != nil {
			return
		}

		if dter.ElementRequirement.Element != types.EntityTypeUnknown {
			composedElementRequirements = append(composedElementRequirements, dter)
		} else {
			// The element is blank; we previously expressed some kinds of composed cluster requirements this way
			cr := matter.NewClusterRequirement(deviceType, row)
			// These always only apply to the server
			cr.Interface = matter.InterfaceServer
			cr.ClusterID = dter.ElementRequirement.ClusterID
			cr.ClusterName = dter.ElementRequirement.ClusterName
			cr.Quality = dter.ElementRequirement.Quality
			dtcr := matter.NewDeviceTypeClusterRequirement(deviceType, cr, row)
			dtcr.DeviceTypeID = dter.DeviceTypeID
			dtcr.DeviceTypeName = dter.DeviceTypeName
			if len(dter.ElementRequirement.Conformance) > 0 {
				dtcr.ClusterRequirement.Conformance = dter.ElementRequirement.Conformance.CloneSet()
			}
			composedClusterRequirements = append(composedClusterRequirements, dtcr)
		}
	}
	return
}

func toElementRequirement(d *Doc, ti *TableInfo, row *asciidoc.TableRow, deviceType *matter.DeviceType) (cr matter.ElementRequirement, err error) {
	cr = matter.NewElementRequirement(deviceType, row)
	cr.ClusterID, err = ti.ReadID(row, matter.TableColumnClusterID, matter.TableColumnID)
	if err != nil {
		return
	}
	cr.ClusterName, _, err = ti.ReadName(row, matter.TableColumnClusterName, matter.TableColumnCluster)
	if err != nil {
		return
	}
	var e string
	e, err = ti.ReadString(row, matter.TableColumnElement)
	if err != nil {
		return
	}
	switch strings.ToLower(e) {
	case "feature":
		cr.Element = types.EntityTypeFeature
	case "attribute":
		cr.Element = types.EntityTypeAttribute
	case "command":
		cr.Element = types.EntityTypeCommand
	case "command field":
		cr.Element = types.EntityTypeCommandField
	case "event":
		cr.Element = types.EntityTypeEvent
	case "":
		slog.Warn("Blank element in element requirements; treating as a cluster requirement", slog.String("clusterName", cr.ClusterName), log.Path("source", row))
	default:
		err = newGenericParseError(row, "unknown element type: \"%s\"", e)
		return
	}

	cr.Name, err = ti.ReadString(row, matter.TableColumnName)
	if err != nil {
		return
	}
	if cr.Element == types.EntityTypeCommandField {
		parts := strings.FieldsFunc(cr.Name, func(r rune) bool { return r == '.' })
		if len(parts) == 2 {
			cr.Name = parts[0]
			cr.Field = parts[1]
		}
	}
	if cr.Field == "" {
		cr.Field, err = ti.ReadString(row, matter.TableColumnField)
		if err != nil {
			return
		}
	}
	cr.Quality, err = ti.ReadQuality(row, cr.Element, matter.TableColumnQuality)
	if err != nil {
		return
	}
	cr.Constraint = ti.ReadConstraint(row, matter.TableColumnConstraint)
	var a string
	a, err = ti.ReadString(row, matter.TableColumnAccess)
	if err != nil {
		return
	}
	cr.Access, _ = ParseAccess(a, types.EntityTypeElementRequirement)
	cr.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
	return
}

func toDeviceTypeTagRequirements(d *Doc, s *asciidoc.Section, deviceType *matter.DeviceType) (deviceTypeTagRequirements []*matter.DeviceTypeTagRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = newGenericParseError(s, "error reading device type tag requirements table: %w", err)
		}
		return
	}
	for row := range ti.ContentRows() {
		dttr := matter.NewDeviceTypeTagRequirement(deviceType, row)
		dttr.DeviceTypeID, err = ti.ReadID(row, matter.TableColumnDeviceID)
		if err != nil {
			return
		}
		dttr.DeviceTypeName, _, err = ti.ReadName(row, matter.TableColumnDeviceName, matter.TableColumnDevice)
		if err != nil {
			return
		}
		dttr.NamespaceID, err = ti.ReadID(row, matter.TableColumnNamespaceID)
		if err != nil {
			return
		}
		dttr.NamespaceName, _, err = ti.ReadName(row, matter.TableColumnNamespace)
		if err != nil {
			return
		}
		dttr.SemanticTagID, err = ti.ReadID(row, matter.TableColumnTagID)
		if err != nil {
			return
		}
		dttr.SemanticTagName, _, err = ti.ReadName(row, matter.TableColumnTag)
		if err != nil {
			return
		}

		if ti.ColumnMap.HasAny(matter.TableColumnConstraint) {
			dttr.Constraint = ti.ReadConstraint(row, matter.TableColumnConstraint)
		}
		if ti.ColumnMap.HasAny(matter.TableColumnConformance) {
			dttr.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
		}
		dttr.Location, err = ti.ReadLocation(row, matter.TableColumnLocation)
		if err != nil {
			return
		}
		deviceTypeTagRequirements = append(deviceTypeTagRequirements, dttr)
	}
	return
}
