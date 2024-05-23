package db

import (
	"context"
	"log/slog"
	"strings"

	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/spec"
	"github.com/hasty/alchemy/matter/types"
)

func (h *Host) indexEventModels(cxt context.Context, parent *sectionInfo, cluster *matter.Cluster) error {
	for _, e := range cluster.Events {
		row := newDBRow()
		row.values[matter.TableColumnID] = e.ID.HexString()
		row.values[matter.TableColumnName] = e.Name
		row.values[matter.TableColumnPriority] = e.Priority
		row.values[matter.TableColumnAccess] = spec.AccessToASCIIDocString(e.Access, types.EntityTypeEvent)
		if e.Conformance != nil {
			row.values[matter.TableColumnConformance] = e.Conformance.ASCIIDocString()
		}
		ei := &sectionInfo{id: h.nextID(eventTable), parent: parent, values: row, children: make(map[string][]*sectionInfo)}
		parent.children[eventTable] = append(parent.children[eventTable], ei)
		for _, ef := range e.Fields {
			h.readField(ef, ei, eventFieldTable, types.EntityTypeEvent)
		}
	}
	return nil
}

func (h *Host) indexEvents(cxt context.Context, doc *spec.Doc, ci *sectionInfo, es *spec.Section) error {
	if ci.children == nil {
		ci.children = make(map[string][]*sectionInfo)
	}
	err := h.readTableSection(cxt, doc, ci, es, eventTable)
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
	for _, s := range parse.Skim[*spec.Section](es.Elements()) {
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
			err = h.readTableSection(cxt, doc, p, s, eventFieldTable)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
