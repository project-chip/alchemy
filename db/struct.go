package db

import (
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (h *Host) indexStructs(cluster *matter.Cluster, parent *sectionInfo) {
	for _, s := range cluster.Structs {
		row := newDBRow()
		row.values[matter.TableColumnName] = s.Name
		row.values[matter.TableColumnDescription] = s.Description
		if s.FabricScoping == matter.FabricScopingScoped {
			row.values[matter.TableColumnScope] = "fabric"
		} else {
			row.values[matter.TableColumnScope] = ""
		}
		ei := h.newSectionInfo(structTable, parent, row, s)
		parent.children[structTable] = append(parent.children[structTable], ei)
		for _, env := range s.Fields {
			h.readField(env, ei, structField, types.EntityTypeStruct)
		}
	}
}
