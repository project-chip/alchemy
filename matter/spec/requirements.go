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
	var rows []*asciidoc.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = fmt.Errorf("error reading cluster requirements table: %w", err)
		}
		return
	}
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		cr := &matter.ClusterRequirement{}
		cr.ID, err = readRowID(row, columnMap, matter.TableColumnID)
		if err != nil {
			return
		}
		cr.ClusterName, err = readRowASCIIDocString(row, columnMap, matter.TableColumnCluster)
		if err != nil {
			return
		}
		if cr.ClusterName == "" {
			cr.ClusterName, err = readRowASCIIDocString(row, columnMap, matter.TableColumnName)
			if err != nil {
				return
			}
		}
		var q string
		q, err = readRowASCIIDocString(row, columnMap, matter.TableColumnQuality)
		if err != nil {
			return
		}
		var cs string
		cs, err = readRowASCIIDocString(row, columnMap, matter.TableColumnClientServer)
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
		cr.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		clusterRequirements = append(clusterRequirements, cr)
	}
	return
}

func (s *Section) toElementRequirements(d *Doc) (elementRequirements []*matter.ElementRequirement, err error) {
	var rows []*asciidoc.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = fmt.Errorf("error reading element requirements table: %w", err)
		}
		return
	}
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		cr := &matter.ElementRequirement{}
		cr.ID, err = readRowID(row, columnMap, matter.TableColumnID)
		if err != nil {
			return
		}
		cr.ClusterName, err = readRowASCIIDocString(row, columnMap, matter.TableColumnCluster)
		if err != nil {
			return
		}
		var e string
		e, err = readRowASCIIDocString(row, columnMap, matter.TableColumnElement)
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
		cr.Name, err = readRowASCIIDocString(row, columnMap, matter.TableColumnName)
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
		q, err = readRowASCIIDocString(row, columnMap, matter.TableColumnQuality)
		if err != nil {
			return
		}
		cr.Quality = parseQuality(q, cr.Element, d, row)
		var c string
		c, err = readRowASCIIDocString(row, columnMap, matter.TableColumnConstraint)
		if err != nil {
			return
		}
		cr.Constraint, err = constraint.ParseString(c)
		if err != nil {
			slog.Warn("failed parsing constraint", log.Element("path", d.Path, row), slog.String("constraint", c))
			cr.Constraint = &constraint.GenericConstraint{Value: c}
		}
		var a string
		a, err = readRowASCIIDocString(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		cr.Access, _ = ParseAccess(a, types.EntityTypeElementRequirement)
		cr.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		elementRequirements = append(elementRequirements, cr)
	}
	return
}

func (s *Section) toConditions(d *Doc) (conditions []*matter.Condition, err error) {
	var rows []*asciidoc.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	var extraColumns []ExtraColumn
	rows, headerRowIndex, columnMap, extraColumns, err = parseFirstTable(d, s)
	if err != nil {
		if err == ErrNoTableFound {
			err = nil
		} else {
			err = fmt.Errorf("error reading conditions table: %w", err)
		}
		return
	}
	featureIndex, ok := columnMap[matter.TableColumnFeature]
	if !ok {
		featureIndex, ok = columnMap[matter.TableColumnCondition]
		if !ok {
			featureIndex = -1
			for _, col := range extraColumns {
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
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		c := matter.NewCondition(s.Base)
		c.Feature, err = readRowCellASCIIDocString(row, featureIndex)
		if err != nil {
			return
		}
		c.Description, err = readRowASCIIDocString(row, columnMap, matter.TableColumnDescription)
		if err != nil {
			return
		}
		conditions = append(conditions, c)
	}
	return
}
