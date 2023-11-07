package db

import (
	"context"
	"log/slog"
	"strings"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
	"github.com/hasty/matterfmt/parse"
)

func (h *Host) indexCommandModels(cxt context.Context, parent *sectionInfo, cluster *matter.Cluster) error {
	for _, c := range cluster.Commands {
		row := newDBRow()
		row.values[matter.TableColumnID] = c.ID
		row.values[matter.TableColumnName] = c.Name
		row.values[matter.TableColumnDirection] = c.Direction
		row.values[matter.TableColumnResponse] = c.Response
		row.values[matter.TableColumnAccess] = ascii.AccessToAsciiString(c.Access)
		row.values[matter.TableColumnConformance] = c.Conformance
		ci := &sectionInfo{id: h.nextId(commandTable), parent: parent, values: row, children: make(map[string][]*sectionInfo)}
		parent.children[commandTable] = append(parent.children[commandTable], ci)
		for _, ef := range c.Fields {
			h.readField(ef, ci, commandFieldTable)
		}
	}
	return nil
}

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
	for _, s := range parse.Skim[*ascii.Section](es.Elements) {
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