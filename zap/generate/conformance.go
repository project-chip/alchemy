package generate

import (
	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/dm"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
)

func renderConformance(doc *spec.Doc, identifierStore conformance.IdentifierStore, c conformance.Conformance, parent *etree.Element) error {
	removeConformance(parent)
	if conformance.IsMandatory(c) {
		return nil
	}
	return dm.RenderConformanceElement(doc, identifierStore, c, parent)
}

func removeConformance(parent *etree.Element) {
	var trash []*etree.Element
	for _, child := range parent.Child {
		switch child := child.(type) {
		case *etree.Element:
			switch child.Tag {
			case "mandatoryConform", "optionalConform", "disableConform", "disallowConform", "provisionalConform", "deprecateConform", "otherwiseConform":
				trash = append(trash, child)
			}
		}
	}
	for _, child := range trash {
		parent.RemoveChild(child)
	}
}
