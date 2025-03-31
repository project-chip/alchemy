package generate

import (
	"log/slog"

	"github.com/beevik/etree"
	"github.com/project-chip/alchemy/dm"
	"github.com/project-chip/alchemy/internal/xml"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/conformance"
	"github.com/project-chip/alchemy/matter/spec"
	"github.com/project-chip/alchemy/matter/types"
)

func renderConformance(spec *spec.Specification, entity types.Entity, identifierStore conformance.IdentifierStore, c conformance.Conformance, parent *etree.Element, alternatives ...string) error {
	removeConformance(parent)
	if conformance.IsMandatory(c) && !conformance.IsProvisional(c) {
		return nil
	}
	doc, ok := spec.DocRefs[entity]
	if !ok {
		slog.Warn("missing doc ref for entity", matter.LogEntity("entity", entity))
	}
	conformanceElement, err := dm.CreateConformanceElement(doc, c, nil)
	if err != nil {
		return err
	}
	if conformanceElement != nil {
		xml.AppendElement(parent, conformanceElement, alternatives...)
	}
	return nil
}

func removeConformance(parent *etree.Element) {
	var trash []*etree.Element
	for _, child := range parent.Child {
		switch child := child.(type) {
		case *etree.Element:
			switch child.Tag {
			case "mandatoryConform",
				"optionalConform",
				"disableConform",
				"disallowConform",
				"provisionalConform",
				"deprecateConform",
				"otherwiseConform",
				"describedConform",
				"genericConform":
				trash = append(trash, child)
			}
		}
	}
	for _, child := range trash {
		parent.RemoveChild(child)
	}
}
