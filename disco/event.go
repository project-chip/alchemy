package disco

import (
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (b *Baller) organizeEventsSection(cxt *discoContext) (err error) {
	for _, events := range cxt.parsed.events {
		eventsTable := events.table
		if eventsTable == nil || eventsTable.Element == nil {
			slog.Warn("Could not organize Events section, as no table of events was found", log.Path("source", events.section))
			return
		}
		if eventsTable.ColumnMap == nil {
			return fmt.Errorf("can't rearrange events table without header row in section %s in %s", cxt.doc.SectionName(events.section), cxt.doc.Path)
		}

		if len(eventsTable.ColumnMap) < 2 {
			return fmt.Errorf("can't rearrange events table with so few matches in section %s in %s", cxt.doc.SectionName(events.section), cxt.doc.Path)
		}

		err = b.renameTableHeaderCells(cxt, events.section, eventsTable, matter.Tables[matter.TableTypeEvents].ColumnRenames)
		if err != nil {
			return fmt.Errorf("error renaming table header cells in section %s in %s: %w", cxt.doc.SectionName(events.section), cxt.doc.Path, err)
		}

		err = b.fixAccessCells(cxt, events, types.EntityTypeEvent)
		if err != nil {
			return fmt.Errorf("error fixing access cells in section %s in %s: %w", cxt.doc.SectionName(events.section), cxt.doc.Path, err)
		}

		err = b.fixQualityCells(cxt, events)
		if err != nil {
			return err
		}

		err = b.fixConformanceCells(cxt, events, eventsTable.Rows, eventsTable.ColumnMap)
		if err != nil {
			return fmt.Errorf("error fixing conformance cells for event table in section %s in %s: %w", cxt.doc.SectionName(events.section), cxt.doc.Path, err)
		}

		err = b.addMissingColumns(cxt, events.section, eventsTable, matter.Tables[matter.TableTypeEvents], types.EntityTypeEvent)
		if err != nil {
			return fmt.Errorf("error adding missing columns in section %s in %s: %w", cxt.doc.SectionName(events.section), cxt.doc.Path, err)
		}

		err = b.reorderColumns(cxt, events.section, eventsTable, matter.TableTypeEvents)
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
			err = b.fixConstraintCells(cxt, event.section, eventTable)
			if err != nil {
				return fmt.Errorf("error fixing constraint cells for event table in section %s in %s: %w", cxt.doc.SectionName(event.section), cxt.doc.Path, err)
			}

			err = b.fixConformanceCells(cxt, event, eventTable.Rows, eventTable.ColumnMap)
			if err != nil {
				return fmt.Errorf("error fixing conformance cells for event table in section %s in %s: %w", cxt.doc.SectionName(event.section), cxt.doc.Path, err)
			}

			err = b.renameTableHeaderCells(cxt, event.section, eventTable, matter.Tables[matter.TableTypeEventFields].ColumnRenames)
			if err != nil {
				return fmt.Errorf("error renaming table header cells in event table in section %s in %s: %w", cxt.doc.SectionName(event.section), cxt.doc.Path, err)
			}

			err = b.addMissingColumns(cxt, event.section, eventTable, matter.Tables[matter.TableTypeEventFields], types.EntityTypeStructField)
			if err != nil {
				return fmt.Errorf("error adding missing columns to event table in section %s in %s: %w", cxt.doc.SectionName(event.section), cxt.doc.Path, err)
			}

			err = b.reorderColumns(cxt, event.section, eventTable, matter.TableTypeEventFields)
			if err != nil {
				return fmt.Errorf("error reordering columns in event table in section %s in %s: %w", cxt.doc.SectionName(event.section), cxt.doc.Path, err)
			}

			b.appendSubsectionTypes(cxt, event.section, eventTable.ColumnMap, eventTable.Rows)
			b.removeMandatoryFallbacks(eventTable)

			err = b.linkIndexTables(cxt, event)
			if err != nil {
				return fmt.Errorf("error linking event index tables in section %s in %s: %w", cxt.doc.SectionName(event.section), cxt.doc.Path, err)
			}
		}
	}
	return
}
