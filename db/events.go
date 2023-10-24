package db

import (
	"context"
	"log/slog"
	"strings"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func (h *Host) indexEvents(cxt context.Context, ci *sectionInfo, es *ascii.Section) error {
	if ci.children == nil {
		ci.children = make(map[string][]*sectionInfo)
	}
	err := h.readTableSection(cxt, ci, es, eventTable)
	if err != nil {
		return err
	}
	events := ci.children[eventTable]
	if len(events) == 0 {
		return nil
	}
	em := make(map[string]*sectionInfo)
	for _, si := range ci.children[eventTable] {
		name, ok := si.values.values[matter.TableColumnName]
		if ok {
			if ns, ok := name.(string); ok {
				em[ns] = si
			}
		}
	}
	for _, s := range parse.Skim[*ascii.Section](es.Elements) {
		switch s.SecType {
		case matter.SectionEvent:
			name := strings.TrimSuffix(s.Name, " Event")
			p, ok := em[name]
			if !ok {
				slog.Error("no matching event", "name", s.Name)
				continue
			}
			if p.children == nil {
				p.children = make(map[string][]*sectionInfo)
			}
			err = h.readTableSection(cxt, p, s, eventFieldTable)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
