package dm

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/constraint"
	"github.com/hasty/alchemy/matter"
)

func dataModelName(s string) string {
	switch s {
	case "octstr": // Everybody's gotta be just a little bit different
		return "octets"
	default:
		return s
	}
}

func renderDataTypes(cluster *matter.Cluster, c *etree.Element) (err error) {
	if len(cluster.Enums) == 0 && len(cluster.Bitmaps) == 0 && len(cluster.Structs) == 0 {
		return
	}
	dt := c.CreateElement("dataTypes")
	err = renderEnums(cluster, dt)
	if err != nil {
		return
	}
	err = renderBitmaps(cluster, dt)
	if err != nil {
		return
	}

	err = renderStructs(cluster, dt)
	return
}

func renderDataType(f *matter.Field, i *etree.Element) {
	if f.Type != nil {
		if !f.Type.IsArray() {
			i.CreateAttr("type", dataModelName(f.Type.Name))
		} else {
			i.CreateAttr("type", "list")
			e := i.CreateElement("entry")
			e.CreateAttr("type", dataModelName(f.Type.EntryType.Name))
			if lc, ok := f.Constraint.(*constraint.ListConstraint); ok {
				renderConstraint(lc.EntryConstraint, f.Type.EntryType, e)
			}
		}
	}
}

func renderDefault(fs matter.FieldSet, f *matter.Field, e *etree.Element) {
	if f.Default == "" {
		return
	}
	cons := constraint.ParseConstraint(f.Default)
	ec, ok := cons.(*constraint.ExactConstraint)
	if ok {
		_, ok = ec.Value.(*constraint.ManufacturerLimit)
		if ok {
			e.CreateAttr("default", "MS")
		}
	}
	def := cons.Default(&matter.ConstraintContext{Fields: fs, Field: f})
	if !def.Defined() {
		return
	}
	e.CreateAttr("default", def.DataModelString(f.Type))
}
