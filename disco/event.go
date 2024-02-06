package disco

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (b *Ball) organizeEventsSection(cxt *discoContext, dp *docParse) (err error) {
	for _, events := range dp.events {
		eventsTable := events.table
		if eventsTable.element == nil {
			err = fmt.Errorf("no events table found")
			return
		}
		if eventsTable.columnMap == nil {
			return fmt.Errorf("can't rearrange events table without header row in section %s in %s", events.section.Name, dp.doc.Path)
		}

		if len(eventsTable.columnMap) < 2 {
			return fmt.Errorf("can't rearrange events table with so few matches in section %s in %s", events.section.Name, dp.doc.Path)
		}

		err = b.fixAccessCells(dp.doc, eventsTable.rows, eventsTable.columnMap, mattertypes.EntityTypeEvent)
		if err != nil {
			return fmt.Errorf("error fixing access cells in section %s in %s: %w", events.section.Name, dp.doc.Path, err)
		}

		err = fixConformanceCells(dp.doc, eventsTable.rows, eventsTable.columnMap)
		if err != nil {
			return fmt.Errorf("error fixing conformance cells for event table in section %s in %s: %w", events.section.Name, dp.doc.Path, err)
		}

		err = b.renameTableHeaderCells(eventsTable.rows, eventsTable.headerRow, eventsTable.columnMap, nil)
		if err != nil {
			return fmt.Errorf("error renaming table header cells in section %s in %s: %w", events.section.Name, dp.doc.Path, err)
		}

		b.addMissingColumns(dp.doc, events.section, eventsTable.element, eventsTable.rows, matter.EventsTableColumnOrder[:], nil, eventsTable.headerRow, eventsTable.columnMap)

		b.reorderColumns(dp.doc, events.section, eventsTable.rows, matter.EventsTableColumnOrder[:], eventsTable.columnMap, eventsTable.extraColumns)

		for _, event := range events.children {
			eventTable := event.table
			if eventTable.element == nil {
				continue
			}
			err = fixConstraintCells(dp.doc, eventTable.rows, eventTable.columnMap)
			if err != nil {
				return fmt.Errorf("error fixing constraint cells for event table in section %s in %s: %w", event.section.Name, dp.doc.Path, err)
			}

			err = fixConformanceCells(dp.doc, eventTable.rows, eventTable.columnMap)
			if err != nil {
				return fmt.Errorf("error fixing conformance cells for event table in section %s in %s: %w", event.section.Name, dp.doc.Path, err)
			}

			err = b.renameTableHeaderCells(eventTable.rows, eventTable.headerRow, eventTable.columnMap, nil)
			if err != nil {
				return fmt.Errorf("error renaming table header cells in event table in section %s in %s: %w", event.section.Name, dp.doc.Path, err)
			}

			b.addMissingColumns(dp.doc, event.section, eventTable.element, eventTable.rows, matter.EventTableColumnOrder[:], nil, eventTable.headerRow, eventTable.columnMap)

			b.reorderColumns(dp.doc, event.section, eventTable.rows, matter.EventTableColumnOrder[:], eventTable.columnMap, eventTable.extraColumns)

			b.appendSubsectionTypes(event.section, eventTable.columnMap, eventTable.rows)
		}
	}
	return
}

/*func (b *Ball) organizeEventsSection(cxt *discoContext, doc *ascii.Doc, events *ascii.Section) error {
	t := ascii.FindFirstTable(events)
	if t == nil {
		return fmt.Errorf("no events table found")
	}
	return b.organizeEventsTable(cxt, doc, events, t)
}

func (b *Ball) organizeEventsTable(cxt *discoContext, doc *ascii.Doc, events *ascii.Section, eventsTable *types.Table) error {

	rows := ascii.TableRows(eventsTable)

	headerRowIndex, columnMap, extraColumns, err := ascii.MapTableColumns(b.doc, rows)
	if err != nil {
		return fmt.Errorf("failed mapping table columns for events table in section %s in %s: %w", events.Name, doc.Path, err)
	}

	if columnMap == nil {
		return fmt.Errorf("can't rearrange events table without header row in section %s in %s", events.Name, doc.Path)
	}

	if len(columnMap) < 2 {
		return fmt.Errorf("can't rearrange events table with so few matches in section %s in %s", events.Name, doc.Path)
	}

	err = b.fixAccessCells(doc, rows, columnMap, mattertypes.EntityTypeEvent)
	if err != nil {
		return fmt.Errorf("error fixing access cells in section %s in %s: %w", events.Name, doc.Path, err)
	}

	err = fixConformanceCells(doc, rows, columnMap)
	if err != nil {
		return fmt.Errorf("error fixing conformance cells for event table in section %s in %s: %w", events.Name, doc.Path, err)
	}

	err = b.renameTableHeaderCells(rows, headerRowIndex, columnMap, nil)
	if err != nil {
		return fmt.Errorf("error renaming table header cells in section %s in %s: %w", events.Name, doc.Path, err)
	}

	err = b.organizeEvents(cxt, doc, events, eventsTable, columnMap)
	if err != nil {
		return fmt.Errorf("error organizing events in section %s in %s: %w", events.Name, doc.Path, err)
	}

	b.addMissingColumns(doc, events, eventsTable, rows, matter.EventsTableColumnOrder[:], nil, headerRowIndex, columnMap)

	b.reorderColumns(doc, events, rows, matter.EventsTableColumnOrder[:], columnMap, extraColumns)
	return nil
}

func (b *Ball) organizeEvents(cxt *discoContext, doc *ascii.Doc, events *ascii.Section, eventsTable *types.Table, columnMap ascii.ColumnIndex) error {
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

		hri, cm, ec, err := ascii.MapTableColumns(b.doc, rows)
		if err != nil {
			return fmt.Errorf("error mapping table columns for event table in section %s in %s: %w", ss.Name, doc.Path, err)

		}
		err = fixConstraintCells(doc, rows, cm)
		if err != nil {
			return fmt.Errorf("error fixing constraint cells for event table in section %s in %s: %w", ss.Name, doc.Path, err)
		}

		err = fixConformanceCells(doc, rows, cm)
		if err != nil {
			return fmt.Errorf("error fixing conformance cells for event table in section %s in %s: %w", ss.Name, doc.Path, err)
		}

		err = b.renameTableHeaderCells(rows, hri, cm, nil)
		if err != nil {
			return fmt.Errorf("error renaming table header cells in event table in section %s in %s: %w", ss.Name, doc.Path, err)
		}

		b.addMissingColumns(doc, ss, t, rows, matter.EventTableColumnOrder[:], nil, hri, cm)

		b.reorderColumns(doc, ss, rows, matter.EventTableColumnOrder[:], cm, ec)

		b.appendSubsectionTypes(ss, cm, rows)

	}

	return nil
}
*/
