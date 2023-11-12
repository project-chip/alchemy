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
	var columnMap map[matter.TableColumn]int
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
	if err != nil {
		if err == NoTableFound {
			err = nil
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
		cr.Cluster, err = readRowValue(row, columnMap, matter.TableColumnCluster)
		if err != nil {
			return
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
		cr.Conformance, err = readRowValue(row, columnMap, matter.TableColumnConformance)
		if err != nil {
			return
		}
		clusterRequirements = append(clusterRequirements, cr)
	}
	return
}

func (s *Section) toElementRequirements(d *Doc) (elementRequirements []*matter.ElementRequirement, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap map[matter.TableColumn]int
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
	if err != nil {
		if err == NoTableFound {
			err = nil
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
		cr.Cluster, err = readRowValue(row, columnMap, matter.TableColumnCluster)
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
		default:
			err = fmt.Errorf("unknown element type: %s", e)
		}
		cr.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		var c string
		c, err = readRowValue(row, columnMap, matter.TableColumnConformance)
		if err != nil {
			return
		}
		cr.Constraint = constraint.ParseConstraint(c)
		var a string
		a, err = readRowValue(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		cr.Access = ParseAccess(a)
		cr.Conformance, err = readRowValue(row, columnMap, matter.TableColumnConformance)
		if err != nil {
			return
		}
		elementRequirements = append(elementRequirements, cr)
	}
	return
}
