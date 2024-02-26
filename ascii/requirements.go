package ascii

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/constraint"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (s *Section) toClusterRequirements(d *Doc) (clusterRequirements []*matter.ClusterRequirement, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		if err == NoTableFound {
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
		cr.ClusterName, err = readRowAsciiDocString(row, columnMap, matter.TableColumnCluster)
		if err != nil {
			return
		}
		if cr.ClusterName == "" {
			cr.ClusterName, err = readRowAsciiDocString(row, columnMap, matter.TableColumnName)
			if err != nil {
				return
			}
		}
		var q string
		q, err = readRowAsciiDocString(row, columnMap, matter.TableColumnQuality)
		if err != nil {
			return
		}
		var cs string
		cs, err = readRowAsciiDocString(row, columnMap, matter.TableColumnClientServer)
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
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		if err == NoTableFound {
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
		cr.ClusterName, err = readRowAsciiDocString(row, columnMap, matter.TableColumnCluster)
		if err != nil {
			return
		}
		var e string
		e, err = readRowAsciiDocString(row, columnMap, matter.TableColumnElement)
		if err != nil {
			return
		}
		switch strings.ToLower(e) {
		case "feature":
			cr.Element = mattertypes.EntityTypeFeature
		case "attribute":
			cr.Element = mattertypes.EntityTypeAttribute
		case "command":
			cr.Element = mattertypes.EntityTypeCommand
		case "command field":
			cr.Element = mattertypes.EntityTypeCommandField
		case "event":
			cr.Element = mattertypes.EntityTypeEvent
		default:
			if e != "" {
				err = fmt.Errorf("unknown element type: \"%s\"", e)
			}
		}
		if err != nil {
			return
		}
		cr.Name, err = readRowAsciiDocString(row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		var q string
		q, err = readRowAsciiDocString(row, columnMap, matter.TableColumnQuality)
		if err != nil {
			return
		}
		cr.Quality = matter.ParseQuality(q)
		var c string
		c, err = readRowAsciiDocString(row, columnMap, matter.TableColumnConstraint)
		if err != nil {
			return
		}
		cr.Constraint = constraint.ParseString(c)
		var a string
		a, err = readRowAsciiDocString(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		cr.Access = ParseAccess(a, mattertypes.EntityTypeElementRequirement)
		cr.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		elementRequirements = append(elementRequirements, cr)
	}
	return
}

func (s *Section) toConditions(d *Doc) (conditions []*matter.Condition, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	var extraColumns []ExtraColumn
	rows, headerRowIndex, columnMap, extraColumns, err = parseFirstTable(d, s)
	if err != nil {
		if err == NoTableFound {
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
		c := &matter.Condition{}
		c.Feature, err = readRowCellAsciiDocString(row, featureIndex)
		if err != nil {
			return
		}
		c.Description, err = readRowAsciiDocString(row, columnMap, matter.TableColumnDescription)
		if err != nil {
			return
		}
		conditions = append(conditions, c)
	}
	return
}
