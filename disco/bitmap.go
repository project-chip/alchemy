package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
)

func (b *Ball) organizeBitmapSection(doc *ascii.Doc, section *ascii.Section) error {
	bitsTable := ascii.FindFirstTable(section)
	if bitsTable == nil {
		return fmt.Errorf("no bitmap table found")
	}
	name := strings.TrimSpace(section.Name)
	if strings.HasSuffix(strings.ToLower(name), "bitmap") {
		setSectionTitle(section, name+" Type")
	}
	return b.organizeBitmapTable(doc, section, bitsTable)
}

func (b *Ball) organizeBitmapTable(doc *ascii.Doc, section *ascii.Section, bitsTable *types.Table) error {
	rows := ascii.TableRows(bitsTable)

	headerRowIndex, columnMap, extraColumns, err := ascii.MapTableColumns(doc, rows)
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

	err = b.renameTableHeaderCells(rows, headerRowIndex, columnMap, nil)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in section %s in %s: %w", section.Name, doc.Path, err)
	}

	b.addMissingColumns(doc, section, bitsTable, rows, matter.BitmapTableColumnOrder[:], nil, headerRowIndex, columnMap)

	b.reorderColumns(doc, section, rows, matter.BitmapTableColumnOrder[:], columnMap, extraColumns)

	b.appendSubsectionTypes(section, columnMap, rows)

	return nil
}
