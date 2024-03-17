package dm

import (
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
)

func renderStructs(doc *ascii.Doc, cluster *matter.Cluster, dt *etree.Element) (err error) {
	structs := make([]*matter.Struct, len(cluster.Structs))
	copy(structs, cluster.Structs)
	slices.SortFunc(structs, func(a, b *matter.Struct) int {
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

func renderFields(doc *ascii.Doc, cluster *matter.Cluster, fs matter.FieldSet, parent *etree.Element) (err error) {
	for _, f := range fs {
		err = renderField(doc, cluster, fs, f, parent)
	}
	return
}

func renderField(doc *ascii.Doc, cluster *matter.Cluster, fs matter.FieldSet, f *matter.Field, parent *etree.Element) (err error) {
	if !f.ID.Valid() {
		return
	}
	i := parent.CreateElement("field")
	i.CreateAttr("id", f.ID.IntString())
	i.CreateAttr("name", f.Name)
	renderDataType(f, i)
	renderAttributeAccess(i, f.Access)
	renderQuality(i, f.Quality, matter.QualityNullable)
	err = renderConformanceString(doc, fs, f.Conformance, i)
	if err != nil {
		return
	}
	err = renderConstraint(f.Constraint, f.Type, i)
	if err != nil {
		return
	}
	renderDefault(fs, f, i)

	return
}
