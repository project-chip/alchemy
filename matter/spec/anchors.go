package spec

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
)

type Anchor struct {
	Document      *asciidoc.Document
	Source        matter.Source
	ID            asciidoc.Elements
	LabelElements asciidoc.Elements
	Element       asciidoc.Element
	Parent        asciidoc.Parent
}

func NewAnchor(doc *asciidoc.Document, id asciidoc.Elements, element asciidoc.Element, parent asciidoc.Parent, label ...asciidoc.Element) *Anchor {
	return &Anchor{
		Document:      doc,
		Source:        NewSource(doc, element),
		ID:            id,
		Element:       element,
		Parent:        parent,
		LabelElements: label,
	}
}

func (library *Library) Identifier(reader asciidoc.Reader, a *Anchor) string {
	return library.anchorId(reader, a.Parent, a.Element, a.ID)
}

func (a *Anchor) Name(reader asciidoc.Reader) string {
	name := ReferenceName(reader, a.Element)
	if len(name) > 0 {
		return name
	}
	return ""
}

func (library *Library) SyncToDoc(a *Anchor, id asciidoc.Elements) {
	if !id.Equals(a.ID) {
		library.changeAnchor(library, a, a.Parent, id)
		a.ID = id
	}
	switch e := a.Element.(type) {
	case *asciidoc.Anchor:
		e.ID = a.ID
		e.Elements = make(asciidoc.Elements, len(a.LabelElements))
		copy(e.Elements, a.LabelElements)
	case asciidoc.Attributable:
		var idAttribute asciidoc.Attribute
		var refTextAttribute *asciidoc.NamedAttribute
		for _, attr := range e.Attributes() {
			switch attr := attr.(type) {
			case *asciidoc.ShorthandAttribute:
				if attr.ID != nil {
					idAttribute = attr
				}
			case *asciidoc.AnchorAttribute:
				attr.ID = a.ID
				if len(a.LabelElements) > 0 {
					attr.Label = make(asciidoc.Elements, len(a.LabelElements))
					copy(attr.Label, a.LabelElements)
				} else {
					attr.Label = nil
				}
				return
			case *asciidoc.NamedAttribute:
				switch attr.Name {
				case asciidoc.AttributeNameID:
					idAttribute = attr
				case asciidoc.AttributeNameReferenceText:
					refTextAttribute = attr
				}
			}
		}
		if idAttribute != nil {
			switch idAttribute := idAttribute.(type) {
			case *asciidoc.ShorthandAttribute:
				idAttribute.ID.Elements = a.ID
				return
			case *asciidoc.NamedAttribute:
				idAttribute.Val = a.ID
				if len(a.LabelElements) > 0 {
					if refTextAttribute != nil {
						refTextAttribute.Val = make(asciidoc.Elements, len(a.LabelElements))
						copy(refTextAttribute.Val, a.LabelElements)
					} else {
						e.AppendAttribute(asciidoc.NewNamedAttribute(string(asciidoc.AttributeNameReferenceText), a.LabelElements, asciidoc.AttributeQuoteTypeDouble))
					}
				}
				return
			}
		}
		e.AppendAttribute(asciidoc.NewAnchorAttribute(a.ID, a.LabelElements))
	}

}

func (library *Library) Anchors(reader asciidoc.Reader) (map[string][]*Anchor, error) {
	if library.referenceIndex.anchorsParsed {
		return library.referenceIndex.anchors, nil
	}
	library.findAnchors(reader)
	return library.anchors, nil
}

func (library *Library) findAnchors(reader asciidoc.Reader) (anchors map[string][]*Anchor) {
	crossReferences := library.crossReferences
	parse.Search(library.Root, reader, library.Root, reader.Children(library.Root), func(doc *asciidoc.Document, el any, parent asciidoc.ParentElement, index int) parse.SearchShould {
		var anchor *Anchor
		var label string
		switch el := el.(type) {
		case *asciidoc.Anchor:
			anchor = NewAnchor(doc, el.ID, el, parent, el.Elements...)
		case *asciidoc.Section:
			anchor = library.makeAnchor(doc, parent, el, crossReferences)
			if anchor != nil {
				label = library.SectionName(el)
				library.anchorsByLabel[label] = append(library.anchorsByLabel[label], anchor)
			}
		case asciidoc.Element:
			anchor = library.makeAnchor(doc, parent, el, crossReferences)
		default:
			slog.Warn("unexpected anchor element", "type", fmt.Sprintf("%T", el))
			return parse.SearchShouldSkip
		}
		if anchor != nil {
			anchorID := library.anchorId(reader, anchor.Parent, anchor.Element, anchor.ID)
			library.anchors[anchorID] = append(library.anchors[anchorID], anchor)
			if len(anchor.LabelElements) > 0 {
				anchorLabel := strings.TrimSpace(asciidoc.AttributeAsciiDocString(anchor.LabelElements))
				if len(anchorLabel) > 0 && anchorLabel != label {
					library.anchorsByLabel[anchorLabel] = append(library.anchorsByLabel[anchorLabel], anchor)
				}
			}

		}
		return parse.SearchShouldContinue
	})
	library.anchorsParsed = true
	return
}

func (library *Library) makeAnchor(doc *asciidoc.Document, parent asciidoc.Parent, element asciidoc.Element, crossReferences map[string][]*CrossReference) *Anchor {
	// If there's a cross-reference for it, then we'll need to make an anchor
	id, labelSet := library.getAnchorElements(doc, element, crossReferences)
	if len(id) == 0 {
		return nil
	}
	//anchorId := library.anchorId(library, parent, element, id)
	//slog.Info("Creating anchor for section with cross reference", slog.String("id", anchorId), slog.String("path", doc.Path.Relative), log.Address("elementAddress", element))
	a := NewAnchor(doc, id, element, parent, labelSet...)
	return a
}

func (library *Library) getAnchorElements(doc *asciidoc.Document, element asciidoc.Element, crossReferences map[string][]*CrossReference) (id asciidoc.Elements, labelSet asciidoc.Elements) {
	var idAttr asciidoc.Attribute
	var refTextAttr *asciidoc.NamedAttribute
	if wa, ok := element.(asciidoc.Attributable); ok {
		for _, a := range wa.Attributes() {
			switch a := a.(type) {
			case *asciidoc.AnchorAttribute:
				idAttr = a
			case *asciidoc.ShorthandAttribute:
				if a.ID != nil {
					idAttr = a
				}
			case *asciidoc.NamedAttribute:
				switch a.Name {
				case asciidoc.AttributeNameID:
					idAttr = a
				case asciidoc.AttributeNameReferenceText:
					refTextAttr = a
				}
			}
		}
	}

	if idAttr == nil {
		if s, ok := element.(*asciidoc.Section); ok && crossReferences != nil {
			sectionName := library.SectionName(s)
			id = asciidoc.NewStringElements(sectionName)
			if _, ok := crossReferences[sectionName]; ok {
				return
			}
		}
		return nil, nil
	}
	switch idAttr := idAttr.(type) {
	case *asciidoc.ShorthandAttribute:
		id = idAttr.ID.Elements
	case *asciidoc.AnchorAttribute:
		id = idAttr.ID
		labelSet = idAttr.Label
	case *asciidoc.NamedAttribute:
		id = idAttr.Val
	}
	if refTextAttr != nil {
		labelSet = refTextAttr.Val
	}
	return
}

func (library *Library) FindAnchor(id string, source log.Source) *Anchor {
	a := library.findAnchor(source, id)
	if a != nil {
		return a
	}
	a = library.findAnchorByLabel(source, id)
	if a != nil {
		return a
	}

	return nil
}

func (library *Library) FindAnchorByID(id asciidoc.Elements, element asciidoc.ParentElement, source log.Source) *Anchor {
	anchorID := library.anchorId(library, element, element, id)
	return library.FindAnchor(anchorID, source)
}

func (library *Library) FindAnchors(id string) []*Anchor {
	a := library.findAnchorsByID(id)
	if a != nil {
		return a
	}
	a = library.findAnchorsByLabel(id)
	if a != nil {
		return a
	}

	return nil
}
