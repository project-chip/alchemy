package db

import (
	"context"

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
		ci := h.newSectionInfo(commandTable, parent, row, c)
		parent.children[commandTable] = append(parent.children[commandTable], ci)
		for _, ef := range c.Fields {
			h.readField(ef, ci, commandFieldTable, types.EntityTypeCommandField)
		}
	}
	return nil
}
