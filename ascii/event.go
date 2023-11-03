package ascii

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func (s *Section) toEvents(d *Doc) (events []*matter.Event, err error) {
	var rows []*types.TableRow
	var headerRowIndex int
	var columnMap map[matter.TableColumn]int
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
		e.ID, err = readRowValue(row, columnMap, matter.TableColumnID)
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
		var a string
		a, err = readRowValue(row, columnMap, matter.TableColumnAccess)
		if err != nil {
			return
		}
		e.Access = ParseAccess(a)
		events = append(events, e)

		eventMap[e.Name] = e
	}

	for _, s := range parse.Skim[*Section](s.Elements) {
		switch s.SecType {
		case matter.SectionEvent:

			name := strings.TrimSuffix(s.Name, " Event")
			e, ok := eventMap[name]
			if !ok {
				slog.Info("unknown event", "event", name)
				continue
			}
			var rows []*types.TableRow
			var headerRowIndex int
			var columnMap map[matter.TableColumn]int
			rows, headerRowIndex, columnMap, _, err = parseFirstTable(s)
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
