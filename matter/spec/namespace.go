package spec

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func (s *Section) toNamespace(d *Doc, entityMap map[asciidoc.Attributable][]types.Entity) (entities []types.Entity, err error) {
	var namespaceTable *TableInfo
	var valuesTable *TableInfo
	parse.SkimFunc(s.Elements(), func(t *asciidoc.Table) bool {
		ti, err := ReadTable(d, t)
		if err != nil {
			return true
		}
		if ti.ColumnMap.HasAll(matter.TableColumnID, matter.TableColumnNamespace) {
			namespaceTable = ti
			return valuesTable != nil
		}
		if ti.ColumnMap.HasAll(matter.TableColumnID, matter.TableColumnName) {
			valuesTable = ti
			return namespaceTable != nil
		}
		return true
	})

	if namespaceTable == nil || valuesTable == nil {
		return
	}
	ns := matter.NewNamespace(namespaceTable.Element)
	for row := range namespaceTable.Body() {
		ns.ID, err = namespaceTable.ReadID(row, matter.TableColumnID)
		if err != nil {
			return
		}
		var name string
		name, err = namespaceTable.ReadString(row, matter.TableColumnNamespace)
		if err != nil {
			return
		}
		name = text.TrimCaseInsensitiveSuffix(name, " Namespace")
		ns.Name = name
		break
	}
	for row := range valuesTable.Body() {
		st := &matter.SemanticTag{}
		st.ID, err = valuesTable.ReadID(row, matter.TableColumnID)
		if err != nil {
			return
		}
		if !st.ID.Valid() {
			continue
		}
		st.Name, err = valuesTable.ReadString(row, matter.TableColumnName)
		if err != nil {
			return
		}
		st.Description, _ = valuesTable.ReadString(row, matter.TableColumnSummary)
		ns.SemanticTags = append(ns.SemanticTags, st)
	}
	entities = append(entities, ns)
	entityMap[s.Base] = append(entityMap[s.Base], ns)
	return
}
