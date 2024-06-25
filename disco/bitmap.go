package disco

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Ball) organizeBitmapSections(cxt *discoContext, dp *docParse) (err error) {
	for _, bms := range dp.bitmaps {
		err = b.organizeBitmapSection(cxt, dp, bms)
		if err != nil {
			return
		}
	}
	return
}

func (b *Ball) organizeBitmapSection(cxt *discoContext, dp *docParse, bms *subSection) (err error) {
	name := strings.TrimSpace(bms.section.Name)
	if strings.HasSuffix(strings.ToLower(name), "bitmap") {
		setSectionTitle(bms.section, name+" Type")
	}
	bitsTable := bms.table
	if bitsTable.element == nil {
		return
	}
	if bitsTable.columnMap == nil {
		slog.Debug("can't rearrange bitmap table without header row")
		return nil
	}

	if len(bitsTable.columnMap) < 2 {
		slog.Debug("can't rearrange bitmap table with so few matches")
		return nil
	}

	err = b.renameTableHeaderCells(b.doc, &bitsTable, nil)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in section %s in %s: %w", bms.section.Name, dp.doc.Path, err)
	}

	err = b.addMissingColumns(&bitsTable, matter.Tables[matter.TableTypeBitmap], nil, types.EntityTypeBitmapValue)
	if err != nil {
		return fmt.Errorf("error adding missing table columns in bitmap section %s in %s: %w", bms.section.Name, dp.doc.Path, err)
	}

	err = b.reorderColumns(dp.doc, bms.section, &bitsTable, matter.TableTypeBitmap)
	if err != nil {
		return err
	}

	b.appendSubsectionTypes(bms.section, bitsTable.columnMap, bitsTable.rows)

	b.fixBitmapRange(bms)
	return
}

var bitRangePattern = regexp.MustCompile(`^(?P<From>[0-9]+)(?<Separator>\.{2,}|\s*\-\s*)(?P<To>[0-9]+)$`)

func (b *Ball) fixBitmapRange(bms *subSection) {
	bitIndex, ok := bms.table.getColumnIndex(matter.TableColumnBit, matter.TableColumnValue)
	if !ok {
		return
	}
	for i, row := range bms.table.rows {
		if i == bms.table.headerRow {
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
