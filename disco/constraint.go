package disco

import (
	"log/slog"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/constraint"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func fixConstraintCells(doc *ascii.Doc, rows []*types.TableRow, columnMap ascii.ColumnIndex) (err error) {
	if len(rows) < 2 {
		return
	}
	constraintIndex, ok := columnMap[matter.TableColumnConstraint]
	if !ok {
		return
	}
	for _, row := range rows[1:] {
		cell := row.Cells[constraintIndex]
		vc, e := ascii.RenderTableCell(cell)
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
		fixed := c.AsciiDocString(dataType)
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
