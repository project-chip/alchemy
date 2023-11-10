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
	attributesTable := ascii.FindFirstTable(section)
	if attributesTable == nil {
		return fmt.Errorf("no attributes table found")
	}
	fmt.Printf("enum! %s\n", section.Name)
	name := strings.TrimSpace(section.Name)
	if strings.HasSuffix(strings.ToLower(name), "enum") {
		setSectionTitle(section, name+" Type")
	}
	b.organizeEnumTable(doc, section, attributesTable)
	return nil
}

func (b *Ball) organizeEnumTable(doc *ascii.Doc, section *ascii.Section, attributesTable *types.Table) error {
	rows := ascii.TableRows(attributesTable)

	headerRowIndex, columnMap, extraColumns, err := ascii.MapTableColumns(rows)
	if err != nil {
		return err
	}

	if columnMap == nil {
		slog.Debug("can't rearrange enum table without header row")
		return nil
	}

	if len(columnMap) < 2 {
		slog.Debug("can't rearrange enum table with so few matches")
		return nil
	}

	err = renameTableHeaderCells(rows, headerRowIndex, columnMap, matter.EnumTableColumnNames)
	if err != nil {
		return err
	}

	addMissingColumns(doc, section, rows, matter.EnumTableColumnOrder[:], matter.EnumTableColumnNames, headerRowIndex, columnMap)

	reorderColumns(doc, section, rows, matter.EnumTableColumnOrder[:], columnMap, extraColumns)

	return nil
}
