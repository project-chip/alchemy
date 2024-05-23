package disco

import (
	"log/slog"

	"github.com/hasty/alchemy/asciidoc"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/constraint"
	"github.com/hasty/alchemy/matter/spec"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func fixConstraintCells(doc *spec.Doc, rows []*asciidoc.TableRow, columnMap spec.ColumnIndex) (err error) {
	if len(rows) < 2 {
		return
	}
	constraintIndex, ok := columnMap[matter.TableColumnConstraint]
	if !ok {
		return
	}
	for _, row := range rows[1:] {
		cell := row.Cell(constraintIndex)
		vc, e := spec.RenderTableCell(cell)
		if e != nil {
			continue
		}

		dataType, e := doc.ReadRowDataType(row, columnMap, matter.TableColumnType)
		if e != nil {
			slog.Debug("error reading data type for constraint", slog.String("path", doc.Path), slog.Any("error", e))
			continue
		}
		if dataType == nil {
			continue
		}
		c, e := constraint.ParseString(vc)
		if e != nil {
			continue
		}
		c = simplifyConstraints(c, dataType)
		fixed := c.ASCIIDocString(dataType)
		if fixed != vc {
			err = setCellString(cell, fixed)
			if err != nil {
				return
			}
		}

	}
	return
}

func simplifyConstraints(cons constraint.Constraint, dataType *mattertypes.DataType) constraint.Constraint {
	switch c := cons.(type) {
	case *constraint.RangeConstraint:
		switch from := c.Minimum.(type) {
		case *constraint.IntLimit:
			if from.Value == 0 && dataType.BaseType.IsUnsigned() {
				return &constraint.MaxConstraint{Maximum: c.Maximum}
			}
		}
	}
	return cons
}
