package disco

import (
	"log/slog"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/constraint"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Ball) fixConstraintCells(section *spec.Section, ti *tableInfo) (err error) {
	if len(ti.rows) < 2 {
		return
	}
	if b.errata.IgnoreSection(section.Name, errata.DiscoPurposeTableConstraint) {
		return
	}
	constraintIndex, ok := ti.getColumnIndex(matter.TableColumnConstraint)
	if !ok {
		return
	}
	qualityIndex, hasQuality := ti.getColumnIndex(matter.TableColumnQuality)
	for _, row := range ti.rows[1:] {
		cell := row.Cell(constraintIndex)
		vc, e := spec.RenderTableCell(cell)
		if e != nil {
			continue
		}

		dataType, e := b.doc.ReadRowDataType(row, ti.columnMap, matter.TableColumnType)
		if e != nil {
			slog.Debug("error reading data type for constraint", slog.String("path", b.doc.Path.String()), slog.Any("error", e))
			continue
		}
		if dataType == nil {
			continue
		}
		c, e := constraint.ParseString(vc)
		if e != nil {
			continue
		}
		var quality matter.Quality
		if hasQuality {
			cell := row.Cell(qualityIndex)
			vc, e := spec.RenderTableCell(cell)
			if e != nil {
				continue
			}
			quality = matter.ParseQuality(vc)

		}
		c = simplifyConstraints(c, dataType, quality)
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

type constraintContext struct {
	dataType *types.DataType
}

func (cc *constraintContext) DataType() *types.DataType {

	return cc.dataType
}

func (cc *constraintContext) ReferenceConstraint(ref string) constraint.Constraint {

	return nil
}

func (cc *constraintContext) Default(name string) (def types.DataTypeExtreme) {

	return
}

func simplifyConstraints(cons constraint.Constraint, dataType *types.DataType, quality matter.Quality) constraint.Constraint {
	switch c := cons.(type) {
	case *constraint.RangeConstraint:
		if quality != matter.QualityNone { // We know whether or not this type is nullable, so we can check for the full range
			cc := &constraintContext{dataType: dataType}
			nullable := quality.Has(matter.QualityNullable)
			dataTypeMin := dataType.Min(nullable)
			dataTypeMax := dataType.Max(nullable)
			rangeMin := c.Minimum.Min(cc)
			rangeMax := c.Maximum.Max(cc)
			if dataTypeMin.ValueEquals(rangeMin) && dataTypeMax.ValueEquals(rangeMax) {
				// This is a range constraint that happens to provide the full range of the data type
				return constraint.NewAllConstraint("all")
			}
		}
		switch from := c.Minimum.(type) {
		case *constraint.IntLimit:
			if from.Value == 0 && dataType.BaseType.IsUnsigned() {
				return &constraint.MaxConstraint{Maximum: c.Maximum}
			}
		}

	}
	return cons
}
