package db

import (
	"context"

	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
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
