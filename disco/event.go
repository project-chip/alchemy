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

func (b *Ball) organizeEventsSection(cxt *discoContext, doc *ascii.Doc, events *ascii.Section) error {
	t := ascii.FindFirstTable(events)
	if t == nil {
		return fmt.Errorf("no events table found")
	}
	return b.organizeEventsTable(cxt, doc, events, t)
}

func (b *Ball) organizeEventsTable(cxt *discoContext, doc *ascii.Doc, events *ascii.Section, eventsTable *types.Table) error {

	rows := ascii.TableRows(eventsTable)

	headerRowIndex, columnMap, extraColumns, err := ascii.MapTableColumns(rows)
	if err != nil {
		return fmt.Errorf("failed mapping table columns for events table in section %s: %w", events.Name, err)
	}

	if columnMap == nil {
		return fmt.Errorf("can't rearrange events table without header row")
	}

	if len(columnMap) < 2 {
		return fmt.Errorf("can't rearrange events table with so few matches")
	}

	err = b.fixAccessCells(doc, rows, columnMap)
	if err != nil {
		return err
	}

	err = b.renameTableHeaderCells(rows, headerRowIndex, columnMap, nil)
	if err != nil {
		return err
	}

	err = b.organizeEvents(cxt, doc, events, eventsTable, columnMap)
	if err != nil {
		return err
	}

	b.addMissingColumns(doc, events, rows, matter.EventsTableColumnOrder[:], nil, headerRowIndex, columnMap)

	b.reorderColumns(doc, events, rows, matter.EventsTableColumnOrder[:], columnMap, extraColumns)
	return nil
}

func (b *Ball) organizeEvents(cxt *discoContext, doc *ascii.Doc, events *ascii.Section, eventsTable *types.Table, columnMap map[matter.TableColumn]int) error {
	nameIndex, ok := columnMap[matter.TableColumnName]
	if !ok {
		return nil
	}
	eventNames := make(map[string]struct{}, len(eventsTable.Rows))
	for _, row := range eventsTable.Rows {
		eventName, err := ascii.GetTableCellValue(row.Cells[nameIndex])
		if err != nil {
			slog.Debug("could not get cell value for event", "err", err)
			continue
		}
		eventNames[eventName] = struct{}{}
	}
	subSections := parse.Skim[*ascii.Section](events.Elements)
	for _, ss := range subSections {
		name := strings.TrimSuffix(ss.Name, " Event")
		if _, ok := eventNames[name]; !ok {
			continue
		}
		t := ascii.FindFirstTable(ss)
		if t == nil {
			continue
		}
		rows := ascii.TableRows(t)

		hri, cm, ec, err := ascii.MapTableColumns(rows)
		if err != nil {
			return fmt.Errorf("failed mapping table columns for event table in section %s: %w", ss.Name, err)

		}
		err = fixConstraintCells(rows, columnMap)
		if err != nil {
			return err
		}
		err = getPotentialDataTypes(cxt, ss, rows, columnMap)
		if err != nil {
			return err
		}

		err = b.renameTableHeaderCells(rows, hri, cm, nil)
		if err != nil {
			return err
		}

		b.addMissingColumns(doc, ss, rows, matter.EventTableColumnOrder[:], nil, hri, cm)

		b.reorderColumns(doc, ss, rows, matter.EventTableColumnOrder[:], cm, ec)

		b.appendSubsectionTypes(ss, cm, rows, "Field")

	}

	return nil
}
