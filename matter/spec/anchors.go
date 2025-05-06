package spec

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/matter"
)

type Anchor struct {
	Document      *Doc
	Source        matter.Source
	ID            string
	LabelElements asciidoc.Set
	Element       asciidoc.Element
	Parent        parse.HasElements
}

func NewAnchor(doc *Doc, id string, element asciidoc.Element, parent parse.HasElements, label ...asciidoc.Element) *Anchor {
	return &Anchor{
		Document:      doc,
		Source:        NewSource(doc, element),
		ID:            id,
		Element:       element,
		Parent:        parent,
		LabelElements: label,
	}
}

func (a *Anchor) Name() string {
	name := ReferenceName(a.Element)
	if len(name) > 0 {
		return name
	}
	return ""
}

func (a *Anchor) SyncToDoc(id string) {
	if id != a.ID {
		a.Document.changeAnchor(a, id)
		if a.Document.group != nil {
			a.Document.group.changeAnchor(a, id)
		}
		a.ID = id
	}
	switch e := a.Element.(type) {
	case *asciidoc.Anchor:
		e.ID = a.ID
		e.Set = make(asciidoc.Set, len(a.LabelElements))
		copy(e.Set, a.LabelElements)
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
				attr.ID = asciidoc.NewString(a.ID)
				if len(a.LabelElements) > 0 {
					attr.Label = make(asciidoc.Set, len(a.LabelElements))
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
				idAttribute.ID.Set = asciidoc.Set{asciidoc.NewString(a.ID)}
				return
			case *asciidoc.NamedAttribute:
				idAttribute.Val = asciidoc.Set{asciidoc.NewString(a.ID)}
				if len(a.LabelElements) > 0 {
					if refTextAttribute != nil {
						refTextAttribute.Val = make(asciidoc.Set, len(a.LabelElements))
						copy(refTextAttribute.Val, a.LabelElements)
					} else {
						e.AppendAttribute(asciidoc.NewNamedAttribute(string(asciidoc.AttributeNameReferenceText), a.LabelElements, asciidoc.AttributeQuoteTypeDouble))
					}
				}
				return
			}
		}
		e.AppendAttribute(asciidoc.NewAnchorAttribute(asciidoc.NewString(a.ID), a.LabelElements))
	}

}

func (doc *Doc) Anchors() (map[string][]*Anchor, error) {
	if doc.referenceIndex.anchorsParsed {
		return doc.referenceIndex.anchors, nil
	}
	doc.findAnchors()
	return doc.anchors, nil
}

func (doc *Doc) findAnchors() {
	var crossReferences map[string][]*CrossReference
	if doc.group != nil {
		crossReferences = doc.group.crossReferences
	} else {
		crossReferences = doc.CrossReferences()
	}
	parse.Traverse(doc, doc.Elements(), func(el any, parent parse.HasElements, index int) parse.SearchShould {
		var anchor *Anchor
		var label string
		switch el := el.(type) {
		case *asciidoc.Anchor:
			anchor = NewAnchor(doc, el.ID, el, parent, el.Set...)
		case *Section:
			anchor = doc.makeAnchor(parent, el.Base, crossReferences)
			if anchor != nil {
				label = el.Name
				doc.anchorsByLabel[label] = append(doc.anchorsByLabel[label], anchor)
			}
		case asciidoc.Element:
			anchor = doc.makeAnchor(parent, el, crossReferences)
		default:
			slog.Warn("unexpected anchor element", "type", fmt.Sprintf("%T", el))
			return parse.SearchShouldSkip
		}
		if anchor != nil {
			doc.anchors[anchor.ID] = append(doc.anchors[anchor.ID], anchor)
			if len(anchor.LabelElements) > 0 {
				anchorLabel := strings.TrimSpace(asciidoc.AttributeAsciiDocString(anchor.LabelElements))
				if len(anchorLabel) > 0 && anchorLabel != label {
					doc.anchorsByLabel[anchorLabel] = append(doc.anchorsByLabel[anchorLabel], anchor)
				}
			}

		}
		return parse.SearchShouldContinue
	})
	doc.anchorsParsed = true
}

func (doc *Doc) makeAnchor(parent parse.HasElements, element asciidoc.Element, crossReferences map[string][]*CrossReference) *Anchor {
	// If there's a cross-reference for it, then we'll need to make an anchor
	id, labelSet := getAnchorElements(element, crossReferences)
	if id == "" {
		return nil
	}
	slog.Debug("Creating anchor for section with cross reference", slog.String("id", id), slog.String("path", doc.Path.Relative))
	a := NewAnchor(doc, id, element, parent, labelSet...)
	return a
}

func getAnchorElements(element asciidoc.Element, crossReferences map[string][]*CrossReference) (id string, labelSet asciidoc.Set) {
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
			id = s.Name()
			if _, ok := crossReferences[id]; ok {
				return
			}
		}
		return "", nil
	}
	switch idAttr := idAttr.(type) {
	case *asciidoc.ShorthandAttribute:
		id = asciidoc.AttributeAsciiDocString(idAttr.ID.Set)
	case *asciidoc.AnchorAttribute:
		id = idAttr.ID.Value
		labelSet = idAttr.Label
	case *asciidoc.NamedAttribute:
		id = idAttr.AsciiDocString()
	}
	if refTextAttr != nil {
		labelSet = refTextAttr.Val
	}
	return
}

func (d *Doc) FindAnchor(id string, source log.Source) *Anchor {
	a := d.findAnchor(source, id)
	if a != nil {
		return a
	}
	a = d.findAnchorByLabel(source, id)
	if a != nil {
		return a
	}
	if d.group != nil {
		a = d.group.findAnchor(source, id)
		if a != nil {
			return a
		}
		a = d.group.findAnchorByLabel(source, id)
		if a != nil {
			return a
		}
	}

	return nil
}

func (d *Doc) FindAnchors(id string) []*Anchor {
	a := d.findAnchorsByID(id)
	if a != nil {
		return a
	}
	a = d.findAnchorsByLabel(id)
	if a != nil {
		return a
	}
	if d.group != nil {
		a = d.group.findAnchorsByID(id)
		if a != nil {
			return a
		}
		a = d.group.findAnchorsByLabel(id)
		if a != nil {
			return a
		}
	}

	return nil
}
