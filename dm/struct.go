package dm

import (
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func renderStructs(doc *spec.Doc, cluster *matter.Cluster, dt *etree.Element) (err error) {
	structs := make([]*matter.Struct, len(cluster.Structs))
	copy(structs, cluster.Structs)
	slices.SortStableFunc(structs, func(a, b *matter.Struct) int {
		return strings.Compare(a.Name, b.Name)
	})
	for _, s := range structs {
		en := dt.CreateElement("struct")
		en.CreateAttr("name", s.Name)
		err = renderFields(doc, cluster, s.Fields, en)
		if err != nil {
			return
		}
		if s.FabricScoping == matter.FabricScopingScoped {
			en.CreateElement("access").CreateAttr("fabricScoped", "true")
		}
	}
	return
}

func renderFields(doc *spec.Doc, cluster *matter.Cluster, fs matter.FieldSet, parent *etree.Element) (err error) {
	for _, f := range fs {
		err = renderField(doc, cluster, fs, f, parent)
	}
	return
}

func renderField(doc *spec.Doc, cluster *matter.Cluster, fs matter.FieldSet, f *matter.Field, parent *etree.Element) (err error) {
	if !f.ID.Valid() {
		return
	}
	i := parent.CreateElement("field")
	i.CreateAttr("id", f.ID.IntString())
	i.CreateAttr("name", f.Name)
	renderDataType(f, i)
	err = renderAnonymousType(doc, cluster, i, f)
	if err != nil {
		return
	}
	renderAttributeAccess(i, f.Access)
	renderQuality(i, f.Quality)
	err = renderConformanceElement(doc, fs, f.Conformance, i)
	if err != nil {
		return
	}
	err = renderConstraint(f.Constraint, f.Type, i)
	if err != nil {
		return
	}
	renderFallback(fs, f, i)

	return
}
