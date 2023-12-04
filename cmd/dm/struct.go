package dm

import (
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
)

func renderStructs(cluster *matter.Cluster, dt *etree.Element) (err error) {
	structs := make([]*matter.Struct, len(cluster.Structs))
	copy(structs, cluster.Structs)
	slices.SortFunc(structs, func(a, b *matter.Struct) int {
		return strings.Compare(a.Name, b.Name)
	})
	for _, s := range structs {
		en := dt.CreateElement("struct")
		en.CreateAttr("name", s.Name)
		err = renderFields(cluster, s.Fields, en)
		if err != nil {
			return
		}
	}
	return
}

func renderFields(cluster *matter.Cluster, fs matter.FieldSet, parent *etree.Element) (err error) {
	for _, f := range fs {
		err = renderField(cluster, fs, f, parent)
	}
	return
}

func renderField(cluster *matter.Cluster, fs matter.FieldSet, f *matter.Field, parent *etree.Element) (err error) {
	if !f.ID.Valid() {
		return
	}
	i := parent.CreateElement("field")
	i.CreateAttr("id", f.ID.IntString())
	i.CreateAttr("name", f.Name)
	renderDataType(f, i)
	if f.Access.Read != matter.PrivilegeUnknown {
		i.CreateAttr("read", "true")
	}
	if f.Access.Write != matter.PrivilegeUnknown {
		i.CreateAttr("write", "true")
	}
	if f.Quality.Has(matter.QualityNullable) {
		i.CreateElement("quality").CreateAttr("nullable", "true")
	}
	err = renderConformanceString(cluster, f.Conformance, i)
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
