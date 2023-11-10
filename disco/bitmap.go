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
		return fmt.Errorf("no attributes table found")
	}
	name := strings.TrimSpace(section.Name)
	if strings.HasSuffix(strings.ToLower(name), "bitmap") {
		setSectionTitle(section, name+" Type")
	}
	return nil
}

func (b *Ball) organizeBitmapTable(doc *ascii.Doc, section *ascii.Section, bitsTable *types.Table) error {
	rows := ascii.TableRows(bitsTable)

	headerRowIndex, columnMap, extraColumns, err := ascii.MapTableColumns(rows)
	if err != nil {
		return err
	}

	if columnMap == nil {
		slog.Debug("can't rearrange bitmap table without header row")
		return nil
	}

	if len(columnMap) < 2 {
		slog.Debug("can't rearrange bitmap table with so few matches")
		return nil
	}

	err = renameTableHeaderCells(rows, headerRowIndex, columnMap, matter.BitmapTableColumnNames)
	if err != nil {
		return err
	}

	addMissingColumns(doc, section, rows, matter.BitmapTableColumnOrder[:], matter.BitmapTableColumnNames, headerRowIndex, columnMap)

	reorderColumns(doc, section, rows, matter.BitmapTableColumnOrder[:], columnMap, extraColumns)

	return nil
}
