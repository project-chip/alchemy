package db

import (
	"context"
	"log/slog"
	"strings"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

func (h *Host) indexEventModels(cxt context.Context, parent *sectionInfo, cluster *matter.Cluster) error {
	for _, e := range cluster.Events {
		row := newDBRow()
		row.values[matter.TableColumnID] = e.ID
		row.values[matter.TableColumnName] = e.Name
		row.values[matter.TableColumnPriority] = e.Priority
		row.values[matter.TableColumnAccess] = ascii.AccessToAsciiString(e.Access)
		row.values[matter.TableColumnConformance] = e.Conformance
		ei := &sectionInfo{id: h.nextId(eventTable), parent: parent, values: row, children: make(map[string][]*sectionInfo)}
		parent.children[eventTable] = append(parent.children[eventTable], ei)
		for _, ef := range e.Fields {
			h.readField(ef, ei, eventFieldTable)
		}
	}
	return nil
}

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
