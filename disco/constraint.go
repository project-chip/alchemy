package disco

import (
	"log/slog"

	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/constraint"
	"github.com/hasty/alchemy/matter/spec"
	"github.com/hasty/alchemy/matter/types"
)

func fixConstraintCells(doc *spec.Doc, ti *tableInfo) (err error) {
	if len(ti.rows) < 2 {
		return
	}
	constraintIndex, ok := ti.getColumnIndex(matter.TableColumnConstraint)
	if !ok {
		return
	}
	for _, row := range ti.rows[1:] {
		cell := row.Cell(constraintIndex)
		vc, e := spec.RenderTableCell(cell)
		if e != nil {
			continue
		}

		dataType, e := doc.ReadRowDataType(row, ti.columnMap, matter.TableColumnType)
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

func simplifyConstraints(cons constraint.Constraint, dataType *types.DataType) constraint.Constraint {
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
