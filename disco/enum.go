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

	err = renameTableHeaderCells(rows, headerRowIndex, columnMap, matter.EnumTableColumnNames)
	if err != nil {
		slog.Info("failed renaming", section.Name, err)
		return err
	}

	addMissingColumns(doc, section, rows, matter.EnumTableColumnOrder[:], matter.EnumTableColumnNames, headerRowIndex, columnMap)

	reorderColumns(doc, section, rows, matter.EnumTableColumnOrder[:], columnMap, extraColumns)

	nameIndex, ok := columnMap[matter.TableColumnName]
	if ok {

		valueNames := make(map[string]struct{}, len(rows))
		for _, row := range rows {
			valueName, err := ascii.GetTableCellValue(row.Cells[nameIndex])
			if err != nil {
				slog.Warn("could not get cell value for enum value", "err", err)
				continue
			}
			valueNames[valueName] = struct{}{}
		}
		subSections := parse.FindAll[*ascii.Section](section.Elements)
		for _, ss := range subSections {
			name := strings.TrimSuffix(ss.Name, " Value")
			if _, ok := valueNames[name]; !ok {
				continue
			}
			if !strings.HasSuffix(strings.ToLower(ss.Name), " value") {
				setSectionTitle(ss, ss.Name+" Value")
			}
		}
	}
	return nil
}
