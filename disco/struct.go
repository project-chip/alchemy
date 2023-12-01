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
		slog.Debug("no struct table found")
		return nil
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
		return fmt.Errorf("failed mapping field table for %s struct: %w", section.Name, err)
	}

	if columnMap == nil {
		slog.Debug("can't rearrange struct table without header row")
		return nil
	}

	if len(columnMap) < 2 {
		slog.Debug("can't rearrange struct table with so few matches")
		return nil
	}

	err = b.renameTableHeaderCells(rows, headerRowIndex, columnMap, nil)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in struct table in section %s in %s: %w", section.Name, doc.Path, err)
	}

	b.addMissingColumns(doc, section, fieldsTable, rows, matter.StructTableColumnOrder[:], nil, headerRowIndex, columnMap)

	b.reorderColumns(doc, section, rows, matter.StructTableColumnOrder[:], columnMap, extraColumns)

	b.appendSubsectionTypes(section, columnMap, rows)

	return nil
}
