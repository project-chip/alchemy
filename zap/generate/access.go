package generate

import (
	"encoding/xml"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/matter"
)

func (cr *configuratorRenderer) setAccessAttributes(element *etree.Element, op string, p matter.Privilege) {
	element.CreateAttr("op", op)
	role := element.SelectAttr("role")
	var name string
	if role != nil {
		name = "role"
	} else if cr.configurator.Errata.WritePrivilegeAsRole {
		name = "role"
		element.RemoveAttr("privilege")
	} else {
		name = "privilege"
		element.RemoveAttr("role")
	}
	px, _ := p.MarshalXMLAttr(xml.Name{Local: name})
	element.CreateAttr(name, px.Value)
}
