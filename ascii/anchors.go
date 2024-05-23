package ascii

import (
	"log/slog"

	"github.com/hasty/adoc/asciidoc"
	"github.com/hasty/alchemy/internal/parse"
)

type Anchor struct {
	ID            string
	LabelElements asciidoc.Set
	Element       asciidoc.Attributable
	Parent        parse.HasElements
}

func (a *Anchor) Name() string {
	name := ReferenceName(a.Element)
	if len(name) > 0 {
		return name
	}
	return ""
}

func (doc *Doc) Anchors() (map[string]*Anchor, error) {
	if doc.anchors != nil {
		return doc.anchors, nil
	}
	anchors := make(map[string]*Anchor)
	crossReferences := doc.CrossReferences()
	parse.Traverse(doc, doc.Elements(), func(el any, parent parse.HasElements, index int) parse.SearchShould {
		var wa asciidoc.Attributable
		e, ok := el.(*Element)
		if ok {
			if wa, ok = e.Base.(asciidoc.Attributable); !ok {
				return parse.SearchShouldContinue
			}
		} else if s, ok := el.(*Section); ok {
			wa = s.Base
		} else {
			return parse.SearchShouldSkip
		}
		var idAttr asciidoc.Attribute
		var refTextAttr *asciidoc.NamedAttribute
		for _, a := range wa.Attributes() {
			switch a := a.(type) {
			case *asciidoc.AnchorAttribute:
				idAttr = a

			case *asciidoc.NamedAttribute:
				switch a.Name {
				case asciidoc.AttributeNameID:
					idAttr = a
				case asciidoc.AttributeNameReferenceText:
					refTextAttr = a
				}
			}
		}
		if idAttr == nil {
			if s, ok := wa.(*asciidoc.Section); ok {
				id := s.Name()

				if _, ok := crossReferences[id]; ok { // If there's a cross-reference for it, then we'll need to make an anchor

					if _, ok := anchors[id]; ok {
						slog.Debug("duplicate anchor; can't fix", "id", id)
						return parse.SearchShouldContinue
					}

					info := &Anchor{
						ID:      id,
						Element: wa,
						Parent:  parent,
					}

					anchors[id] = info
				}
			}
			return parse.SearchShouldContinue
		}
		var id string
		var labelSet asciidoc.Set
		switch idAttr := idAttr.(type) {
		case *asciidoc.AnchorAttribute:
			id = idAttr.ID.Value
			labelSet = idAttr.Label
		case *asciidoc.NamedAttribute:
			id = idAttr.AsciiDocString()
		}
		if refTextAttr != nil {
			labelSet = refTextAttr.Val
		}
		/*id := strings.TrimSpace(idAttr.AsciiDocString())
		if parts := strings.Split(id, ","); len(parts) > 1 {
			id = strings.TrimSpace(parts[0])
			label = strings.TrimSpace(parts[1])
		}
		refText := wa.GetAttributeByName(asciidoc.AttributeNameReferenceText)
		if refText != nil {
			label = refText.AsciiDocString()
		}*/
		info := &Anchor{
			ID:            id,
			LabelElements: labelSet,
			Element:       wa,
			Parent:        parent,
		}
		if _, ok := anchors[id]; ok {
			slog.Debug("duplicate anchor; can't fix", "id", id)
			return parse.SearchShouldContinue
		}

		anchors[id] = info
		/*if !strings.HasPrefix(id, "_") {

		} else { // Anchors prefaced with "_" may have been created by the parser
			if _, ok := crossReferences[id]; ok { // If there's a cross-reference for it, then we'll render it
				anchors[id] = info
			} else { // If there isn't a cross reference to the id, there might be one to its original version
				unescaped := strings.TrimSpace(strings.ReplaceAll(id, "_", " "))
				if _, ok = crossReferences[unescaped]; ok {
					if _, ok := anchors[unescaped]; ok {
						slog.Debug("duplicate anchor; can't fix", "id", unescaped)
						return false
					}
					anchors[unescaped] = info
				}
			}
		}*/
		return parse.SearchShouldContinue
	})
	doc.anchors = anchors
	return anchors, nil
}
