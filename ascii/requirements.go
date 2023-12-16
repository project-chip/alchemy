package ascii

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/constraint"
	"github.com/hasty/alchemy/matter"
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
		cr.ClusterName, err = readRowValue(row, columnMap, matter.TableColumnCluster)
		if err != nil {
			return
		}
		if cr.ClusterName == "" {
			cr.ClusterName, err = readRowValue(row, columnMap, matter.TableColumnName)
			if err != nil {
				return
			}
		}
		var q string
		q, err = readRowValue(row, columnMap, matter.TableColumnQuality)
		if err != nil {
			return
		}
		var cs string
		cs, err = readRowValue(row, columnMap, matter.TableColumnClientServer)
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
		cr.ClusterName, err = readRowValue(row, columnMap, matter.TableColumnCluster)
		if err != nil {
			return
		}
		var e string
		e, err = readRowValue(row, columnMap, matter.TableColumnElement)
		if err != nil {
			return
		}
		switch strings.ToLower(e) {
		case "feature":
			cr.Element = matter.EntityFeature
		case "attribute":
			cr.Element = matter.EntityAttribute
		case "command":
			cr.Element = matter.EntityCommand
		case "command field":
			cr.Element = matter.EntityCommandField
		case "event":
			cr.Element = matter.EntityEvent
		default:
			if e != "" {
				err = fmt.Errorf("unknown element type: \"%s\"", e)
			}
		}
		if err != nil {
			return
		}
		cr.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		var q string
		q, err = readRowValue(row, columnMap, matter.TableColumnQuality)
		if err != nil {
			return
		}
		cr.Quality = matter.ParseQuality(q)
		var c string
		c, err = readRowValue(row, columnMap, matter.TableColumnConstraint)
		if err != nil {
			return
		}
		cr.Constraint = constraint.ParseConstraint(c)
		var a string
		a, err = readRowValue(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		cr.Access = ParseAccess(a, false)
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
		featureIndex, ok := columnMap[matter.TableColumnCondition]
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
		c.Feature, err = readRowCell(row, featureIndex)
		if err != nil {
			return
		}
		c.Description, err = readRowValue(row, columnMap, matter.TableColumnDescription)
		if err != nil {
			return
		}
		conditions = append(conditions, c)
	}
	return
}
