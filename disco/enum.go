package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
)

func (b *Ball) organizeEnumSection(doc *ascii.Doc, section *ascii.Section) error {
	enumTable := ascii.FindFirstTable(section)
	if enumTable == nil {
		return fmt.Errorf("no enum table found")
	}
	name := strings.TrimSpace(section.Name)
	if strings.HasSuffix(strings.ToLower(name), "enum") {
		setSectionTitle(section, name+" Type")
	}
	return b.organizeEnumTable(doc, section, enumTable)
}

func (b *Ball) organizeEnumTable(doc *ascii.Doc, section *ascii.Section, attributesTable *types.Table) error {
	rows := ascii.TableRows(attributesTable)

	headerRowIndex, columnMap, extraColumns, err := ascii.MapTableColumns(rows)
	if err != nil {
		return fmt.Errorf("failed mapping table columns for enum table in section %s: %w", section.Name, err)
	}

	if columnMap == nil {
		slog.Debug("can't rearrange enum table without header row")
		return nil
	}

	if len(columnMap) < 2 {
		slog.Debug("can't rearrange enum table with so few matches")
		return nil
	}

	err = b.renameTableHeaderCells(rows, headerRowIndex, columnMap, nil)
	if err != nil {
		slog.Info("failed renaming", section.Name, err)
		return err
	}

	b.addMissingColumns(doc, section, rows, matter.EnumTableColumnOrder[:], nil, headerRowIndex, columnMap)

	b.reorderColumns(doc, section, rows, matter.EnumTableColumnOrder[:], columnMap, extraColumns)

	b.appendSubsectionTypes(section, columnMap, rows, "Value")

	return nil
}
