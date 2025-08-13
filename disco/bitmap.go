package disco

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Baller) organizeBitmapSections(cxt *discoContext) (err error) {
	for _, bms := range cxt.parsed.bitmaps {
		err = b.organizeBitmapSection(cxt, bms)
		if err != nil {
			return
		}
	}
	return
}

func (b *Baller) organizeBitmapSection(cxt *discoContext, bms *subSection) (err error) {
	b.canonicalizeDataTypeSectionName(cxt, bms.section, "Bitmap")
	bitsTable := bms.table
	if bitsTable == nil || bitsTable.Element == nil {
		slog.Warn("Could not organize bitmap section, as no table of bitmap values was found", log.Path("source", bms.section))
		return
	}
	if bitsTable.ColumnMap == nil {
		slog.Debug("can't rearrange bitmap table without header row")
		return nil
	}

	if len(bitsTable.ColumnMap) < 2 {
		slog.Debug("can't rearrange bitmap table with so few matches")
		return nil
	}

	err = b.renameTableHeaderCells(cxt, bms.section, bitsTable, nil)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in section %s in %s: %w", cxt.doc.SectionName(bms.section), cxt.doc.Path, err)
	}

	err = b.addMissingColumns(cxt, bms.section, bitsTable, matter.Tables[matter.TableTypeBitmap], types.EntityTypeBitmapValue)
	if err != nil {
		return fmt.Errorf("error adding missing table columns in bitmap section %s in %s: %w", cxt.doc.SectionName(bms.section), cxt.doc.Path, err)
	}

	err = b.reorderColumns(cxt, bms.section, bitsTable, matter.TableTypeBitmap)
	if err != nil {
		return err
	}

	b.appendSubsectionTypes(cxt, bms.section, bitsTable.ColumnMap, bitsTable.Rows)
	b.removeMandatoryFallbacks(bitsTable)

	b.fixBitmapRange(cxt, bms)
	return
}

var bitRangePattern = regexp.MustCompile(`^(?P<From>[0-9]+)(?<Separator>\.{2,}|\s*\-\s*)(?P<To>[0-9]+)$`)

func (b *Baller) fixBitmapRange(cxt *discoContext, bms *subSection) {
	if cxt.errata.IgnoreSection(cxt.doc.SectionName(bms.section), errata.DiscoPurposeDataTypeBitmapFixRange) {
		return
	}
	bitIndex, ok := bms.table.ColumnIndex(matter.TableColumnBit, matter.TableColumnValue)
	if !ok {
		return
	}
	for i, row := range bms.table.Rows {
		if i == bms.table.HeaderRowIndex {
			continue
		}
		cell := row.Cell(bitIndex)
		bit, err := spec.RenderTableCell(cell)
		if err != nil {
			continue
		}
		matches := bitRangePattern.FindStringSubmatch(bit)
		if len(matches) < 4 {
			continue
		}
		start, err := parse.HexOrDec(matches[1])
		if err != nil {
			continue
		}
		end, err := parse.HexOrDec(matches[3])
		if err != nil {
			continue
		}
		if start < end {
			setCellString(cell, fmt.Sprintf("%d..%d", end, start))
			continue
		}
		if strings.TrimSpace(matches[2]) == "-" {
			setCellString(cell, fmt.Sprintf("%d..%d", start, end))
			continue
		}
	}
}
