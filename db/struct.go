package db

import (
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/matter/types"
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
		ei := &sectionInfo{id: h.nextID(structTable), parent: parent, values: row, children: make(map[string][]*sectionInfo)}
		parent.children[structTable] = append(parent.children[structTable], ei)
		for _, env := range s.Fields {
			h.readField(env, ei, structField, types.EntityTypeStruct)
		}
	}
}
