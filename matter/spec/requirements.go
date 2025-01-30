package spec

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toClusterRequirements(d *Doc) (clusterRequirements []*matter.ClusterRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = fmt.Errorf("error reading cluster requirements table: %w", err)
		}
		return
	}
	for row := range ti.Body() {
		cr := matter.NewClusterRequirement(row)
		cr.ClusterID, err = ti.ReadID(row, matter.TableColumnID)
		if err != nil {
			return
		}
		cr.ClusterName, err = ti.ReadValue(row, matter.TableColumnCluster)
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
			err = fmt.Errorf("unknown client/server value: %s", cs)
			return
		}
		cr.Quality = matter.ParseQuality(q)
		cr.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
		clusterRequirements = append(clusterRequirements, cr)
	}
	return
}

func (s *Section) toElementRequirements(d *Doc) (elementRequirements []*matter.ElementRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = fmt.Errorf("error reading element requirements table: %w", err)
		}
		return
	}
	for row := range ti.Body() {
		var cr matter.ElementRequirement
		cr, err = s.toElementRequirement(d, ti, row)
		if err != nil {
			return
		}
		elementRequirements = append(elementRequirements, &cr)
	}
	return
}

func (s *Section) toDeviceTypeRequirements(d *Doc) (deviceTypeRequirements []*matter.DeviceTypeRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = fmt.Errorf("error reading element requirements table: %w", err)
		}
		return
	}
	for row := range ti.Body() {
		cr := matter.NewDeviceTypeRequirement(row)
		cr.DeviceTypeID, err = ti.ReadID(row, matter.TableColumnDeviceID)
		if err != nil {
			return
		}
		cr.DeviceTypeName, _, err = ti.ReadName(row, matter.TableColumnName)
		if err != nil {
			return
		}
		cr.Constraint = ti.ReadConstraint(row, matter.TableColumnConstraint)
		cr.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
		deviceTypeRequirements = append(deviceTypeRequirements, cr)
	}
	return
}

func (s *Section) toComposedDeviceTypeRequirements(d *Doc) (composedRequirements []*matter.ComposedDeviceTypeRequirement, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = fmt.Errorf("error reading element requirements table: %w", err)
		}
		return
	}
	for row := range ti.Body() {
		var cr matter.ComposedDeviceTypeRequirement
		cr.DeviceTypeID, err = ti.ReadID(row, matter.TableColumnDeviceID)
		if err != nil {
			return
		}
		cr.DeviceTypeName, _, err = ti.ReadName(row, matter.TableColumnDeviceName)
		if err != nil {
			return
		}
		cr.ElementRequirement, err = s.toElementRequirement(d, ti, row)
		if err != nil {
			return
		}
		composedRequirements = append(composedRequirements, &cr)
	}
	return
}

func (*Section) toElementRequirement(d *Doc, ti *TableInfo, row *asciidoc.TableRow) (cr matter.ElementRequirement, err error) {
	cr = matter.NewElementRequirement(row)
	cr.ClusterID, err = ti.ReadID(row, matter.TableColumnClusterID, matter.TableColumnID)
	if err != nil {
		return
	}
	cr.ClusterName, _, err = ti.ReadName(row, matter.TableColumnCluster)
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
	default:
		if e != "" {
			err = fmt.Errorf("unknown element type: \"%s\"", e)
		}
	}
	if err != nil {
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
