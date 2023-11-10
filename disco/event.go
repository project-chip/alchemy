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
		return err
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

	err = renameTableHeaderCells(rows, headerRowIndex, columnMap, matter.EventsTableColumnNames)
	if err != nil {
		return err
	}

	err = organizeEvents(cxt, doc, events, eventsTable, columnMap)
	if err != nil {
		return err
	}

	addMissingColumns(doc, events, rows, matter.EventsTableColumnOrder[:], matter.EventTableColumnNames, headerRowIndex, columnMap)

	reorderColumns(doc, events, rows, matter.EventsTableColumnOrder[:], columnMap, extraColumns)
	return nil
}

func organizeEvents(cxt *discoContext, doc *ascii.Doc, events *ascii.Section, eventsTable *types.Table, columnMap map[matter.TableColumn]int) error {
	nameIndex, ok := columnMap[matter.TableColumnName]
	if !ok {
		return nil
	}
	eventNames := make(map[string]struct{}, len(eventsTable.Rows))
	for _, row := range eventsTable.Rows {
		eventName, err := ascii.GetTableCellValue(row.Cells[nameIndex])
		if err != nil {
			slog.Warn("could not get cell value for event", "err", err)
			continue
		}
		eventNames[eventName] = struct{}{}
	}
	subSections := parse.Skim[*ascii.Section](events.Elements)
	for _, ss := range subSections {
		name := strings.TrimSuffix(ss.Name, " Event")
		fmt.Printf("looking for event %s\n", name)
		if _, ok := eventNames[name]; !ok {
			fmt.Printf("didn't find event %s\n", name)
			continue
		}
		t := ascii.FindFirstTable(ss)
		if t == nil {
			fmt.Printf("didn't find table %s\n", name)
			continue
		}
		rows := ascii.TableRows(t)

		hri, cm, ec, err := ascii.MapTableColumns(rows)
		if err != nil {
			return err
		}
		err = fixConstraintCells(rows, columnMap)
		if err != nil {
			return err
		}
		err = getPotentialDataTypes(cxt, ss, rows, columnMap)
		if err != nil {
			return err
		}

		addMissingColumns(doc, ss, rows, matter.EventTableColumnOrder[:], matter.EventTableColumnNames, hri, cm)

		reorderColumns(doc, ss, rows, matter.EventTableColumnOrder[:], cm, ec)

	}

	return nil
}
