package spec

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toEvents(d *Doc, entityMap map[asciidoc.Attributable][]types.Entity) (events matter.EventSet, err error) {
	var rows []*asciidoc.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
	if err != nil {
		return nil, fmt.Errorf("failed reading events: %w", err)
	}

	eventMap := make(map[string]*matter.Event)
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		e := &matter.Event{}
		e.Name, err = ReadRowValue(d, row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		e.Name = matter.StripTypeSuffixes(e.Name)
		e.ID, err = readRowID(row, columnMap, matter.TableColumnID)
		if err != nil {
			return
		}
		e.Priority, err = readRowASCIIDocString(row, columnMap, matter.TableColumnPriority)
		if err != nil {
			return
		}
		e.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		var a string
		a, err = readRowASCIIDocString(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		e.Access, _ = ParseAccess(a, types.EntityTypeEvent)
		if e.Access.Read == matter.PrivilegeUnknown {
			// Sometimes the invoke access is omitted; we assume it's view
			e.Access.Read = matter.PrivilegeView
		}
		events = append(events, e)

		eventMap[e.Name] = e
	}

	for _, s := range parse.Skim[*Section](s.Elements()) {
		switch s.SecType {
		case matter.SectionEvent:

			name := strings.TrimSuffix(s.Name, " Event")
			e, ok := eventMap[name]
			if !ok {
				slog.Debug("unknown event", "event", name)
				continue
			}
			e.Description = getDescription(d, s.Set)
			var rows []*asciidoc.TableRow
			var headerRowIndex int
			var columnMap ColumnIndex
			rows, headerRowIndex, columnMap, _, err = parseFirstTable(d, s)
			if headerRowIndex > 0 {
				firstRow := rows[0]
				tableCells := firstRow.TableCells()
				if len(tableCells) > 0 {
					cv, rowErr := RenderTableCell(tableCells[0])
					if rowErr == nil {
						cv = strings.ToLower(cv)
						if strings.Contains(cv, "fabric sensitive") || strings.Contains(cv, "fabric-sensitive") {
							e.Access.FabricSensitivity = matter.FabricSensitivitySensitive
						}
					}
				}
			}
			if err != nil {
				if err == ErrNoTableFound {
					err = nil
					continue
				}
				err = fmt.Errorf("failed reading %s event fields: %w", s.Name, err)
				return
			}
			e.Fields, err = d.readFields(headerRowIndex, rows, columnMap, types.EntityTypeEvent)
			if err != nil {
				return
			}
			entityMap[s.Base] = append(entityMap[s.Base], e)
			fieldMap := make(map[string]*matter.Field, len(e.Fields))
			for _, f := range e.Fields {
				fieldMap[f.Name] = f
			}
			err = s.mapFields(fieldMap, entityMap)
			if err != nil {
				return
			}
		}
	}
	return
}
