package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

func (b *Ball) organizeBitmapSection(doc *ascii.Doc, section *ascii.Section) error {
	bitsTable := ascii.FindFirstTable(section)
	if bitsTable == nil {
		return fmt.Errorf("no attributes table found")
	}
	name := strings.TrimSpace(section.Name)
	if strings.HasSuffix(strings.ToLower(name), "bitmap") {
		setSectionTitle(section, name+" Type")
	}
	return b.organizeBitmapTable(doc, section, bitsTable)
}

func (b *Ball) organizeBitmapTable(doc *ascii.Doc, section *ascii.Section, bitsTable *types.Table) error {
	rows := ascii.TableRows(bitsTable)

	headerRowIndex, columnMap, extraColumns, err := ascii.MapTableColumns(rows)
	if err != nil {
		return fmt.Errorf("failed mapping table columns for bitmap table in section %s: %w", section.Name, err)
	}

	if columnMap == nil {
		slog.Debug("can't rearrange bitmap table without header row")
		return nil
	}

	if len(columnMap) < 2 {
		slog.Debug("can't rearrange bitmap table with so few matches")
		return nil
	}

	_, ok := columnMap[matter.TableColumnBit]
	if !ok {
		idIndex, ok := columnMap[matter.TableColumnID]
		if ok {
			delete(columnMap, matter.TableColumnID)
			columnMap[matter.TableColumnBit] = idIndex
		}
	}

	err = renameTableHeaderCells(rows, headerRowIndex, columnMap, matter.BitmapTableColumnNames)
	if err != nil {
		return err
	}

	addMissingColumns(doc, section, rows, matter.BitmapTableColumnOrder[:], matter.BitmapTableColumnNames, headerRowIndex, columnMap)

	reorderColumns(doc, section, rows, matter.BitmapTableColumnOrder[:], columnMap, extraColumns)

	nameIndex, ok := columnMap[matter.TableColumnName]
	if ok {

		bitNames := make(map[string]struct{}, len(rows))
		for _, row := range rows {
			bitName, err := ascii.GetTableCellValue(row.Cells[nameIndex])
			if err != nil {
				slog.Warn("could not get cell value for command", "err", err)
				continue
			}
			bitNames[bitName] = struct{}{}
		}
		subSections := parse.FindAll[*ascii.Section](section.Elements)
		for _, ss := range subSections {
			name := strings.TrimSuffix(ss.Name, " Bit")
			if _, ok := bitNames[name]; !ok {
				continue
			}
			if !strings.HasSuffix(strings.ToLower(ss.Name), " bit") {
				setSectionTitle(ss, ss.Name+" Bit")
			}
		}
	}

	return nil
}
