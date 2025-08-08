package db

import (
	"context"
	"log/slog"

	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func (h *Host) indexCommandModels(cxt context.Context, parent *sectionInfo, cluster *matter.Cluster) error {
	for _, c := range cluster.Commands {
		row := newDBRow()
		row.values[matter.TableColumnID] = c.ID.IntString()
		row.values[matter.TableColumnName] = c.Name
		switch c.Direction {
		case matter.InterfaceClient:
			row.values[matter.TableColumnDirection] = "client"
		case matter.InterfaceServer:
			row.values[matter.TableColumnDirection] = "server"
		default:
			row.values[matter.TableColumnDirection] = "unknown"

		}
		if c.Response != nil {
			row.values[matter.TableColumnResponse] = c.Response.Name
		}
		row.values[matter.TableColumnAccess] = spec.AccessToASCIIDocString(c.Access, types.EntityTypeCommand)
		if c.Conformance != nil {
			row.values[matter.TableColumnConformance] = c.Conformance.ASCIIDocString()
		}
		row.values[matter.TableColumnQuality] = c.Quality
		ci := &sectionInfo{id: h.nextID(commandTable), parent: parent, values: row, children: make(map[string][]*sectionInfo)}
		parent.children[commandTable] = append(parent.children[commandTable], ci)
		for _, ef := range c.Fields {
			h.readField(ef, ci, commandFieldTable, types.EntityTypeCommandField)
		}
	}
	return nil
}

func (h *Host) indexCommands(cxt context.Context, doc *spec.Doc, ci *sectionInfo, es *spec.Section) error {
	if ci.children == nil {
		ci.children = make(map[string][]*sectionInfo)
	}
	err := h.readTableSection(cxt, doc, ci, es, commandTable)
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
	for s := range parse.Skim[*spec.Section](es.Children()) {
		switch s.SecType {
		case matter.SectionCommand:
			name := text.TrimCaseInsensitiveSuffix(s.Name, " Command")
			p, ok := em[name]
			if !ok {
				slog.Debug("no matching command", "name", s.Name)
				continue
			}
			if p.children == nil {
				p.children = make(map[string][]*sectionInfo)
			}
			err = h.readTableSection(cxt, doc, p, s, commandFieldTable)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
