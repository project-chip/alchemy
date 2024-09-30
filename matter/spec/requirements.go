package spec

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
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
		cr := &matter.ClusterRequirement{}
		cr.ClusterID, err = ti.ReadID(row, matter.TableColumnID)
		if err != nil {
			return
		}
		cr.ClusterName, err = ti.ReadString(row, matter.TableColumnCluster)
		if err != nil {
			return
		}
		if cr.ClusterName == "" {
			cr.ClusterName, err = ti.ReadString(row, matter.TableColumnName)
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
		cr.ClusterID, err = ti.ReadID(row, matter.TableColumnDeviceID)
		if err != nil {
			return
		}
		cr.DeviceTypeName, err = ti.ReadString(row, matter.TableColumnDeviceName)
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
	cr.ClusterID, err = ti.ReadID(row, matter.TableColumnClusterID, matter.TableColumnID)
	if err != nil {
		return
	}
	cr.ClusterName, err = ti.ReadString(row, matter.TableColumnCluster)
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
	var q string
	q, err = ti.ReadString(row, matter.TableColumnQuality)
	if err != nil {
		return
	}
	cr.Quality = parseQuality(q, cr.Element, d, row)
	var c string
	c, err = ti.ReadString(row, matter.TableColumnConstraint)
	if err != nil {
		return
	}
	cr.Constraint, err = constraint.ParseString(c)
	if err != nil {
		slog.Warn("failed parsing constraint", log.Element("path", d.Path, row), slog.String("constraint", c))
		cr.Constraint = &constraint.GenericConstraint{Value: c}
	}
	var a string
	a, err = ti.ReadString(row, matter.TableColumnAccess)
	if err != nil {
		return
	}
	cr.Access, _ = ParseAccess(a, types.EntityTypeElementRequirement)
	cr.Conformance = ti.ReadConformance(row, matter.TableColumnConformance)
	return
}

func (s *Section) toConditions(d *Doc) (conditions []*matter.Condition, err error) {
	var ti *TableInfo
	ti, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = fmt.Errorf("error reading conditions table: %w", err)
		}
		return
	}
	featureIndex, ok := ti.ColumnMap[matter.TableColumnFeature]
	if !ok {
		featureIndex, ok = ti.ColumnMap[matter.TableColumnCondition]
		if !ok {
			featureIndex = -1
			for _, col := range ti.ExtraColumns {
				if strings.HasSuffix(col.Name, "Tag") {
					featureIndex = col.Offset
					break
				}
			}
			if featureIndex == -1 {
				err = fmt.Errorf("failed to find tag column in section %s", s.Name)
				return
			}
		}
	}
	for row := range ti.Body() {
		c := matter.NewCondition(s.Base)
		c.Feature, err = ti.ReadStringAtOffset(row, featureIndex)
		if err != nil {
			return
		}
		c.Description, err = ti.ReadString(row, matter.TableColumnDescription)
		if err != nil {
			return
		}
		conditions = append(conditions, c)
	}
	return
}
