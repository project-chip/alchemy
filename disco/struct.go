package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
)

func (b *Ball) organizeStructSection(doc *ascii.Doc, section *ascii.Section) error {
	fieldsTable := ascii.FindFirstTable(section)
	if fieldsTable == nil {
		return fmt.Errorf("no attributes table found")
	}
	name := strings.TrimSpace(section.Name)
	if strings.HasSuffix(strings.ToLower(name), "struct") {
		setSectionTitle(section, name+" Type")
	}
	return b.organizeStructTable(doc, section, fieldsTable)
}

func (b *Ball) organizeStructTable(doc *ascii.Doc, section *ascii.Section, fieldsTable *types.Table) error {
	rows := ascii.TableRows(fieldsTable)

	headerRowIndex, columnMap, extraColumns, err := ascii.MapTableColumns(rows)
	if err != nil {
		return err
	}

	if columnMap == nil {
		slog.Debug("can't rearrange struct table without header row")
		return nil
	}

	if len(columnMap) < 2 {
		slog.Debug("can't rearrange struct table with so few matches")
		return nil
	}

	err = renameTableHeaderCells(rows, headerRowIndex, columnMap, matter.StructTableColumnNames)
	if err != nil {
		return err
	}

	addMissingColumns(doc, section, rows, matter.StructTableColumnOrder[:], matter.StructTableColumnNames, headerRowIndex, columnMap)

	reorderColumns(doc, section, rows, matter.StructTableColumnOrder[:], columnMap, extraColumns)

	return nil
}
