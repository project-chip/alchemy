package db

import (
	"context"

	"github.com/project-chip/alchemy/matter"
)

func (h *Host) indexNamepsace(cxt context.Context, parent *sectionInfo, namespace *matter.Namespace) error {
	if !namespace.ID.Valid() {
		return nil
	}
	namespaceRow := newDBRow()
	namespaceRow.values[matter.TableColumnID] = namespace.ID.IntString()
	namespaceRow.values[matter.TableColumnName] = namespace.Name

	ci := &sectionInfo{id: h.nextID(namespaceTable), parent: parent, values: namespaceRow, children: make(map[string][]*sectionInfo)}

	for _, t := range namespace.SemanticTags {
		tagRow := newDBRow()
		tagRow.values[matter.TableColumnID] = t.ID.IntString()
		tagRow.values[matter.TableColumnName] = t.Name
		fci := &sectionInfo{id: h.nextID(tagTable), parent: ci, values: tagRow}
		ci.children[tagTable] = append(ci.children[tagTable], fci)

	}

	parent.children[namespaceTable] = append(parent.children[namespaceTable], ci)
	return nil
}
