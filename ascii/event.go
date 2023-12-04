package ascii

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

func (s *Section) toEvents(d *Doc) (events []*matter.Event, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap ColumnIndex
	rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
	if err != nil {
		return nil, fmt.Errorf("failed reading events: %w", err)
	}

	eventMap := make(map[string]*matter.Event)
	for i := headerRowIndex + 1; i < len(rows); i++ {
		row := rows[i]
		e := &matter.Event{}
		e.Name, err = readRowValue(row, columnMap, matter.TableColumnName)
		if err != nil {
			return
		}
		e.ID, err = readRowID(row, columnMap, matter.TableColumnID)
		if err != nil {
			return
		}
		e.Description, err = readRowValue(row, columnMap, matter.TableColumnDescription)
		if err != nil {
			return
		}
		e.Priority, err = readRowValue(row, columnMap, matter.TableColumnPriority)
		if err != nil {
			return
		}
		e.Conformance = d.getRowConformance(row, columnMap, matter.TableColumnConformance)
		var a string
		a, err = readRowValue(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		e.Access = ParseAccess(a)
		if e.Access.Invoke == matter.PrivilegeUnknown {
			// Sometimes the invoke access is omitted; we assume it's operate
			e.Access.Invoke = matter.PrivilegeOperate
		}
		events = append(events, e)

		eventMap[e.Name] = e
	}

	for _, s := range parse.Skim[*Section](s.Elements) {
		switch s.SecType {
		case matter.SectionEvent:

			name := strings.TrimSuffix(s.Name, " Event")
			e, ok := eventMap[name]
			if !ok {
				slog.Debug("unknown event", "event", name)
				continue
			}
			var rows []*types.TableRow
			var headerRowIndex int
			var columnMap ColumnIndex
			rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
			if headerRowIndex > 0 {
				firstRow := rows[0]
				if len(firstRow.Cells) > 0 {
					cv, rowErr := GetTableCellValue(rows[0].Cells[0])
					if rowErr == nil {
						cv = strings.ToLower(cv)
						if strings.Contains(cv, "fabric sensitive") || strings.Contains(cv, "fabric-sensitive") {
							e.FabricSensitive = true
						}
					}
				}
			}
			if err != nil {
				if err == NoTableFound {
					err = nil
					continue
				}
				err = fmt.Errorf("failed reading %s event fields: %w", s.Name, err)
				return
			}
			e.Fields, err = d.readFields(headerRowIndex, rows, columnMap)
		}
	}
	return
}
