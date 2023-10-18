package db

import (
	"context"
	"log/slog"
	"strings"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func (h *Host) indexCommands(cxt context.Context, ci *sectionInfo, es *ascii.Section) error {
	if ci.children == nil {
		ci.children = make(map[string][]*sectionInfo)
	}
	err := h.readTableSection(cxt, ci, es, commandTable)
	if err != nil {
		return err
	}
	commands := ci.children[commandTable]
	if len(commands) == 0 {
		return nil
	}
	em := make(map[string]*sectionInfo)
	for _, si := range commands {
		name, ok := si.values.values[matter.TableColumnName]
		if ok {
			if ns, ok := name.(string); ok {
				em[ns] = si
			}
		}
	}
	for _, s := range ascii.Skim[*ascii.Section](es.Elements) {
		switch s.SecType {
		case matter.SectionCommand:
			name := strings.TrimSuffix(s.Name, " Command")
			p, ok := em[name]
			if !ok {
				slog.Error("no matching command", "name", s.Name)
				continue
			}
			if p.children == nil {
				p.children = make(map[string][]*sectionInfo)
			}
			err = h.readTableSection(cxt, p, s, commandFieldTable)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
