package disco

import (
	"fmt"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Ball) organizeEventsSection(cxt *discoContext, dp *docParse) (err error) {
	for _, events := range dp.events {
		eventsTable := events.table
		if eventsTable == nil || eventsTable.Element == nil {
			err = fmt.Errorf("no events table found")
			return
		}
		if eventsTable.ColumnMap == nil {
			return fmt.Errorf("can't rearrange events table without header row in section %s in %s", events.section.Name, dp.doc.Path)
		}

		if len(eventsTable.ColumnMap) < 2 {
			return fmt.Errorf("can't rearrange events table with so few matches in section %s in %s", events.section.Name, dp.doc.Path)
		}

		err = b.fixAccessCells(dp, events, types.EntityTypeEvent)
		if err != nil {
			return fmt.Errorf("error fixing access cells in section %s in %s: %w", events.section.Name, dp.doc.Path, err)
		}

		err = b.fixConformanceCells(dp, events, eventsTable.Rows, eventsTable.ColumnMap)
		if err != nil {
			return fmt.Errorf("error fixing conformance cells for event table in section %s in %s: %w", events.section.Name, dp.doc.Path, err)
		}

		err = b.renameTableHeaderCells(dp.doc, events.section, eventsTable, nil)
		if err != nil {
			return fmt.Errorf("error renaming table header cells in section %s in %s: %w", events.section.Name, dp.doc.Path, err)
		}

		b.addMissingColumns(events.section, eventsTable, matter.Tables[matter.TableTypeEvents], types.EntityTypeEvent)

		err = b.reorderColumns(dp.doc, events.section, eventsTable, matter.TableTypeEvents)
		if err != nil {
			return err
		}

		err = b.linkIndexTables(cxt, events)
		if err != nil {
			return err
		}

		for _, event := range events.children {
			eventTable := event.table
			if eventTable == nil || eventTable.Element == nil {
				continue
			}
			err = b.fixConstraintCells(event.section, eventTable)
			if err != nil {
				return fmt.Errorf("error fixing constraint cells for event table in section %s in %s: %w", event.section.Name, dp.doc.Path, err)
			}

			err = b.fixConformanceCells(dp, event, eventTable.Rows, eventTable.ColumnMap)
			if err != nil {
				return fmt.Errorf("error fixing conformance cells for event table in section %s in %s: %w", event.section.Name, dp.doc.Path, err)
			}

			err = b.renameTableHeaderCells(dp.doc, event.section, eventTable, nil)
			if err != nil {
				return fmt.Errorf("error renaming table header cells in event table in section %s in %s: %w", event.section.Name, dp.doc.Path, err)
			}

			err = b.addMissingColumns(event.section, eventTable, matter.Tables[matter.TableTypeEvent], types.EntityTypeStructField)
			if err != nil {
				return fmt.Errorf("error adding missing columns to event table in section %s in %s: %w", event.section.Name, dp.doc.Path, err)
			}

			err = b.reorderColumns(dp.doc, event.section, eventTable, matter.TableTypeEvent)
			if err != nil {
				return fmt.Errorf("error reordering columns in event table in section %s in %s: %w", event.section.Name, dp.doc.Path, err)
			}

			b.appendSubsectionTypes(event.section, eventTable.ColumnMap, eventTable.Rows)

			err = b.linkIndexTables(cxt, event)
			if err != nil {
				return fmt.Errorf("error linking event index tables in section %s in %s: %w", event.section.Name, dp.doc.Path, err)
			}
		}
	}
	return
}
