package spec

import (
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/suggest"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func toNamespace(spec *Specification, d *Doc, s *asciidoc.Section, pc *parseContext) (err error) {
	var namespaceTable *TableInfo
	var valuesTable *TableInfo
	parse.SkimFunc(d.Reader(), s, s.Children(), func(t *asciidoc.Table) bool {
		ti, err := ReadTable(d, d.Reader(), t)
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
	for row := range namespaceTable.ContentRows() {
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
	for row := range valuesTable.ContentRows() {
		st := matter.NewSemanticTag(ns, row)
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
		if !text.IsAlphanumeric(st.Name) {
			slog.Warn("Semantic tag name is not alphanumeric", slog.String("name", st.Name), log.Path("source", row))
		}
		if text.HasCaseInsensitivePrefix(st.Name, "Reserved") {
			slog.Debug("Skipping reserved semantic tag", slog.String("name", st.Name))
			continue
		}
		st.Description, _ = valuesTable.ReadString(row, matter.TableColumnSummary)
		ns.SemanticTags = append(ns.SemanticTags, st)
	}
	pc.entities = append(pc.entities, ns)
	pc.orderedEntities = append(pc.orderedEntities, ns)
	pc.entitiesByElement[s] = append(pc.entitiesByElement[s], ns)
	return
}

func validateNamespaces(spec *Specification) {
	deviceTypeIds := make(idUniqueness[*matter.Namespace])
	for _, ns := range spec.Namespaces {
		deviceTypeIds.check(spec, ns.ID, ns)
	}
}

type tagFinder struct {
	entityFinderCommon

	namespace *matter.Namespace
}

func newTagFinder(namespace *matter.Namespace, inner entityFinder) *tagFinder {
	return &tagFinder{entityFinderCommon: entityFinderCommon{inner: inner}, namespace: namespace}
}

func (tf *tagFinder) findEntityByIdentifier(identifier string, source log.Source) types.Entity {
	for _, c := range tf.namespace.SemanticTags {
		if c.Name == identifier && c != tf.identity {
			return c
		}
	}
	if tf.inner != nil {
		return tf.inner.findEntityByIdentifier(identifier, source)
	}
	return nil
}

func (tf *tagFinder) suggestIdentifiers(identifier string, suggestions map[types.Entity]int) {
	suggest.PossibleEntities(identifier, suggestions, func(yield func(string, types.Entity) bool) {
		for _, t := range tf.namespace.SemanticTags {

			if t == tf.identity {
				continue
			}
			if !yield(t.Name, t) {
				return
			}

		}
	})
	if tf.inner != nil {
		tf.inner.suggestIdentifiers(identifier, suggestions)
	}
	return
}
