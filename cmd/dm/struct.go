package dm

import (
	"github.com/beevik/etree"
	"github.com/hasty/alchemy/matter"
)

func renderStructs(cluster *matter.Cluster, dt *etree.Element) (err error) {
	for _, s := range cluster.Structs {
		en := dt.CreateElement("struct")
		en.CreateAttr("name", s.Name)
		for _, f := range s.Fields {
			if !f.ID.Valid() {
				continue
			}
			i := en.CreateElement("field")
			i.CreateAttr("id", f.ID.IntString())
			i.CreateAttr("name", f.Name)
			if f.Type != nil {
				i.CreateAttr("type", f.Type.Name)
			}
			if f.Access.Read != matter.PrivilegeUnknown {
				i.CreateAttr("read", "true")
			}
			if f.Access.Write != matter.PrivilegeUnknown {
				i.CreateAttr("write", "true")
			}
			err = renderConformanceString(f.Conformance, i)
			if err != nil {
				return
			}
			err = renderConstraint(f.Constraint, f.Type, i)
			if err != nil {
				return
			}
		}
	}
	return
}
