package spec

import (
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (library *Library) toClusterRequirements(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, deviceType *matter.DeviceType) (clusterRequirements []*matter.ClusterRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(reader, d, s)
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
		cr, err = library.toClusterRequirement(reader, deviceType, ti, row)
		if err != nil {
			return
		}
		clusterRequirements = append(clusterRequirements, cr)
	}
	return
}

func (library *Library) toClusterRequirement(reader asciidoc.Reader, deviceType *matter.DeviceType, ti *TableInfo, row *asciidoc.TableRow) (cr *matter.ClusterRequirement, err error) {
	cr = matter.NewClusterRequirement(deviceType, row)
	cr.ClusterID, err = ti.ReadID(reader, row, matter.TableColumnClusterID, matter.TableColumnID)
	if err != nil {
		return
	}
	cr.ClusterName, err = ti.ReadValue(library, row, matter.TableColumnClusterName, matter.TableColumnCluster)
	if err != nil {
		return
	}
	if cr.ClusterName == "" {
		cr.ClusterName, _, err = ti.ReadName(library, row, matter.TableColumnName)
		if err != nil {
			return
		}
	}
	var q string
	q, err = ti.ReadString(reader, row, matter.TableColumnQuality)
	if err != nil {
		return
	}
	var cs string
	cs, err = ti.ReadString(reader, row, matter.TableColumnClientServer)
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
	cr.Conformance = ti.ReadConformance(library, row, matter.TableColumnConformance)
	return
}

func (library *Library) toElementRequirements(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, deviceType *matter.DeviceType) (elementRequirements []*matter.ElementRequirement, clusterRequirements []*matter.ClusterRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(reader, d, s)
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
		er, err = library.toElementRequirement(reader, d, ti, row, deviceType)
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

func (library *Library) toDeviceTypeRequirements(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, deviceType *matter.DeviceType) (deviceTypeRequirements []*matter.DeviceTypeRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(reader, d, s)
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
		deviceId, err = ti.ReadString(reader, row, matter.TableColumnDeviceID, matter.TableColumnID)
		if err != nil {
			return
		}
		if strings.HasSuffix(deviceId, "+") {
			deviceId = deviceId[:len(deviceId)-1]
			dtr.AllowsSuperset = true
		}
		dtr.DeviceTypeID = matter.ParseNumber(deviceId)

		dtr.DeviceTypeName, _, err = ti.ReadName(library, row, matter.TableColumnDeviceName, matter.TableColumnName)
		if err != nil {
			return
		}
		if strings.HasSuffix(dtr.DeviceTypeName, "+") {
			dtr.AllowsSuperset = true
			dtr.DeviceTypeName = dtr.DeviceTypeName[:len(dtr.DeviceTypeName)-1]
		}
		dtr.Constraint = ti.ReadConstraint(library, row, matter.TableColumnConstraint)
		dtr.Conformance = ti.ReadConformance(library, row, matter.TableColumnConformance)
		dtr.Location, err = ti.ReadLocation(reader, row, matter.TableColumnLocation)
		if err != nil {
			return
		}
		deviceTypeRequirements = append(deviceTypeRequirements, dtr)
	}
	return
}

func (library *Library) toComposedDeviceTypeClusterRequirements(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, deviceType *matter.DeviceType) (composedClusterRequirements []*matter.DeviceTypeClusterRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(reader, d, s)
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
		cr, err = library.toClusterRequirement(reader, deviceType, ti, row)
		if err != nil {
			return
		}

		dtcr := matter.NewDeviceTypeClusterRequirement(deviceType, cr, row)
		dtcr.DeviceTypeID, err = ti.ReadID(reader, row, matter.TableColumnDeviceID)
		if err != nil {
			return
		}
		dtcr.DeviceTypeName, _, err = ti.ReadName(library, row, matter.TableColumnDeviceName, matter.TableColumnDevice)
		if err != nil {
			return
		}
		composedClusterRequirements = append(composedClusterRequirements, dtcr)
	}
	return
}

func (library *Library) toConditionRequirements(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, deviceType *matter.DeviceType) (conditionRequirements []*matter.ConditionRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(reader, d, s)
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
		cr.DeviceTypeID, err = ti.ReadID(reader, row, matter.TableColumnDeviceID)
		if err != nil {
			return
		}
		cr.DeviceTypeName, _, err = ti.ReadName(library, row, matter.TableColumnDeviceName, matter.TableColumnDevice)
		if err != nil {
			return
		}
		cr.ConditionName, _, err = ti.ReadName(library, row, matter.TableColumnCondition)
		if err != nil {
			return
		}
		cr.Conformance = ti.ReadConformance(library, row, matter.TableColumnConformance)
		cr.Location, err = ti.ReadLocation(library, row, matter.TableColumnLocation)
		if err != nil {
			return
		}
		conditionRequirements = append(conditionRequirements, cr)
	}
	return
}

func (library *Library) toComposedDeviceTypeElementRequirements(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, deviceType *matter.DeviceType) (composedElementRequirements []*matter.DeviceTypeElementRequirement, composedClusterRequirements []*matter.DeviceTypeClusterRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(reader, d, s)
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
		er, err = library.toElementRequirement(reader, d, ti, row, deviceType)
		if err != nil {
			return
		}
		dter := matter.NewDeviceTypeElementRequirement(deviceType, &er, row)
		dter.DeviceTypeID, err = ti.ReadID(reader, row, matter.TableColumnDeviceID)
		if err != nil {
			return
		}
		dter.DeviceTypeName, _, err = ti.ReadName(library, row, matter.TableColumnDeviceName, matter.TableColumnDevice)
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

func (library *Library) toElementRequirement(reader asciidoc.Reader, d *asciidoc.Document, ti *TableInfo, row *asciidoc.TableRow, deviceType *matter.DeviceType) (cr matter.ElementRequirement, err error) {
	cr = matter.NewElementRequirement(deviceType, row)
	cr.ClusterID, err = ti.ReadID(reader, row, matter.TableColumnClusterID, matter.TableColumnID)
	if err != nil {
		return
	}
	cr.ClusterName, _, err = ti.ReadName(library, row, matter.TableColumnClusterName, matter.TableColumnCluster)
	if err != nil {
		return
	}
	var e string
	e, err = ti.ReadString(reader, row, matter.TableColumnElement)
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

	cr.Name, err = ti.ReadString(reader, row, matter.TableColumnName)
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
		cr.Field, err = ti.ReadString(reader, row, matter.TableColumnField)
		if err != nil {
			return
		}
	}
	cr.Quality, err = ti.ReadQuality(reader, row, cr.Element, matter.TableColumnQuality)
	if err != nil {
		return
	}
	cr.Constraint = ti.ReadConstraint(library, row, matter.TableColumnConstraint)
	var a string
	a, err = ti.ReadString(reader, row, matter.TableColumnAccess)
	if err != nil {
		return
	}
	cr.Access, _ = ParseAccess(a, types.EntityTypeElementRequirement)
	cr.Conformance = ti.ReadConformance(library, row, matter.TableColumnConformance)
	return
}

func (library *Library) toTagRequirements(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, deviceType *matter.DeviceType) (tagRequirements []*matter.TagRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(reader, d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = newGenericParseError(s, "error reading device type tag requirements table: %w", err)
		}
		return
	}
	for row := range ti.ContentRows() {
		var tr *matter.TagRequirement
		tr, err = library.toTagRequirement(reader, d, s, deviceType, ti, row)
		if err != nil {
			return
		}
		tagRequirements = append(tagRequirements, tr)
	}
	return
}

func (library *Library) toTagRequirement(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, deviceType *matter.DeviceType, ti *TableInfo, row *asciidoc.TableRow) (tr *matter.TagRequirement, err error) {
	tr = matter.NewTagRequirement(deviceType, row)

	tr.NamespaceID, err = ti.ReadID(reader, row, matter.TableColumnNamespaceID)
	if err != nil {
		return
	}
	tr.NamespaceName, _, err = ti.ReadName(library, row, matter.TableColumnNamespace)
	if err != nil {
		return
	}
	tr.SemanticTagID, err = ti.ReadID(reader, row, matter.TableColumnTagID)
	if err != nil {
		return
	}
	tr.SemanticTagName, _, err = ti.ReadName(library, row, matter.TableColumnTag)
	if err != nil {
		return
	}

	if ti.ColumnMap.HasAny(matter.TableColumnConstraint) {
		tr.Constraint = ti.ReadConstraint(library, row, matter.TableColumnConstraint)
	}
	if ti.ColumnMap.HasAny(matter.TableColumnConformance) {
		tr.Conformance = ti.ReadConformance(library, row, matter.TableColumnConformance)
	}
	if err != nil {
		return
	}
	return
}

func (library *Library) toDeviceTypeTagRequirements(reader asciidoc.Reader, d *asciidoc.Document, s *asciidoc.Section, deviceType *matter.DeviceType) (deviceTypeTagRequirements []*matter.DeviceTypeTagRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(reader, d, s)
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
		dttr.DeviceTypeID, err = ti.ReadID(reader, row, matter.TableColumnDeviceID)
		if err != nil {
			return
		}
		dttr.DeviceTypeName, _, err = ti.ReadName(library, row, matter.TableColumnDeviceName, matter.TableColumnDevice)
		if err != nil {
			return
		}
		dttr.TagRequirement, err = library.toTagRequirement(reader, d, s, deviceType, ti, row)
		if err != nil {
			return
		}
		deviceTypeTagRequirements = append(deviceTypeTagRequirements, dttr)
	}
	return
}
