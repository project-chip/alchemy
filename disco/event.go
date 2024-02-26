package disco

import (
	"fmt"

	"github.com/hasty/alchemy/matter"
	mattertypes "github.com/hasty/alchemy/matter/types"
)

func (b *Ball) organizeEventsSection(cxt *discoContext, dp *docParse) (err error) {
	for _, events := range dp.events {
		eventsTable := &events.table
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

		err = b.fixAccessCells(dp.doc, eventsTable, mattertypes.EntityTypeEvent)
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

		err = b.linkIndexTables(cxt, events)
		if err != nil {
			return err
		}

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

			err = b.linkIndexTables(cxt, event)
			if err != nil {
				return err
			}
		}
	}
	return
}
