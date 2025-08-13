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
	Document      *Doc
	Source        matter.Source
	ID            asciidoc.Elements
	LabelElements asciidoc.Elements
	Element       asciidoc.Element
	Parent        asciidoc.Parent
}

func NewAnchor(doc *Doc, id asciidoc.Elements, element asciidoc.Element, parent asciidoc.Parent, label ...asciidoc.Element) *Anchor {
	return &Anchor{
		Document:      doc,
		Source:        NewSource(doc, element),
		ID:            id,
		Element:       element,
		Parent:        parent,
		LabelElements: label,
	}
}

func (a *Anchor) Identifier() string {
	return a.Document.anchorId(a.Document.Reader(), a.Parent, a.Element, a.ID)
}

func (a *Anchor) Name() string {
	name := ReferenceName(a.Document, a.Element)
	if len(name) > 0 {
		return name
	}
	return ""
}

func (a *Anchor) SyncToDoc(id asciidoc.Elements) {
	if !id.Equals(a.ID) {
		a.Document.changeAnchor(a, a.Parent, id)
		if a.Document.group != nil {
			a.Document.group.changeAnchor(a, a.Parent, id)
		}
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
	parse.Search(doc.Reader(), doc, doc.Children(), func(el any, parent asciidoc.Parent, index int) parse.SearchShould {
		var anchor *Anchor
		var label string
		switch el := el.(type) {
		case *asciidoc.Anchor:
			anchor = NewAnchor(doc, el.ID, el, parent, el.Elements...)
		case *asciidoc.Section:
			anchor = doc.makeAnchor(parent, el, crossReferences)
			if anchor != nil {
				label = doc.SectionName(el)
				doc.anchorsByLabel[label] = append(doc.anchorsByLabel[label], anchor)
			}
		case asciidoc.Element:
			anchor = doc.makeAnchor(parent, el, crossReferences)
		default:
			slog.Warn("unexpected anchor element", "type", fmt.Sprintf("%T", el))
			return parse.SearchShouldSkip
		}
		if anchor != nil {
			anchorID := doc.anchorId(doc.Reader(), anchor.Parent, anchor.Element, anchor.ID)
			doc.anchors[anchorID] = append(doc.anchors[anchorID], anchor)
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

func (doc *Doc) makeAnchor(parent asciidoc.Parent, element asciidoc.Element, crossReferences map[string][]*CrossReference) *Anchor {
	// If there's a cross-reference for it, then we'll need to make an anchor
	id, labelSet := getAnchorElements(doc, element, crossReferences)
	if len(id) == 0 {
		return nil
	}
	slog.Debug("Creating anchor for section with cross reference", slog.Any("id", id), slog.String("path", doc.Path.Relative))
	a := NewAnchor(doc, id, element, parent, labelSet...)
	return a
}

func getAnchorElements(doc *Doc, element asciidoc.Element, crossReferences map[string][]*CrossReference) (id asciidoc.Elements, labelSet asciidoc.Elements) {
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
			sectionName := doc.SectionName(s)
			id = asciidoc.NewStringElements(sectionName)
			if _, ok := crossReferences[sectionName]; ok {
				return
			}
		}
		return
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

func (d *Doc) FindAnchorByID(id asciidoc.Elements, element asciidoc.ParentElement, source log.Source) *Anchor {
	anchorID := d.anchorId(d.Reader(), element, element, id)
	return d.FindAnchor(anchorID, source)
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
