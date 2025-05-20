package dm

import (
	"slices"
	"strings"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/types"
)

func renderStructs(structs []*matter.Struct, dt *etree.Element) (err error) {
	ss := make([]*matter.Struct, len(structs))
	copy(ss, structs)
	slices.SortStableFunc(ss, func(a, b *matter.Struct) int {
		return strings.Compare(a.Name, b.Name)
	})
	for _, s := range ss {
		en := dt.CreateElement("struct")
		en.CreateAttr("name", s.Name)
		err = renderFields(s.Fields, en, s)
		if err != nil {
			return
		}
		if s.FabricScoping == matter.FabricScopingScoped {
			en.CreateElement("access").CreateAttr("fabricScoped", "true")
		}
	}
	return
}

func renderFields(fs matter.FieldSet, parent *etree.Element, parentEntity types.Entity) (err error) {
	for _, f := range fs {
		err = renderField(fs, f, parent, parentEntity)
		if err != nil {
			return
		}
	}
	return
}

func renderField(fs matter.FieldSet, f *matter.Field, parent *etree.Element, parentEntity types.Entity) (err error) {
	if !f.ID.Valid() {
		return
	}
	i := parent.CreateElement("field")
	i.CreateAttr("id", f.ID.IntString())
	i.CreateAttr("name", f.Name)
	err = renderDataType(f, i)
	if err != nil {
		return
	}
	err = renderAnonymousType(i, f)
	if err != nil {
		return
	}
	renderAttributeAccess(i, f.Access)
	renderQuality(i, f.Quality)
	err = renderConformanceElement(f.Conformance, i, parentEntity)
	if err != nil {
		return
	}
	err = renderConstraint(f.Constraint, f.Type, i, parentEntity)
	if err != nil {
		return
	}
	renderFallback(fs, f, i)

	return
}
