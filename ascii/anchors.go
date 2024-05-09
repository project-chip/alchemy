package ascii

import (
	"log/slog"
	"strings"

	"github.com/hasty/adoc/elements"
	"github.com/hasty/alchemy/internal/parse"
)

type Anchor struct {
	ID      string
	Label   string
	Element elements.Attributable
	Parent  parse.HasElements
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
	parse.Traverse(doc, doc.Elements, func(el any, parent parse.HasElements, index int) bool {
		var wa elements.Attributable
		e, ok := el.(*Element)
		if ok {
			if wa, ok = e.Base.(elements.Attributable); !ok {
				return false
			}
		} else if s, ok := el.(*Section); ok {
			wa = s.Base
		} else {
			return false
		}
		idAttr := wa.GetAttributeByName(elements.AttributeNameID)
		if idAttr == nil {
			return false
		}
		id := strings.TrimSpace(idAttr.AsciiDocString())
		var label string
		if parts := strings.Split(id, ","); len(parts) > 1 {
			id = strings.TrimSpace(parts[0])
			label = strings.TrimSpace(parts[1])
		}
		refText := wa.GetAttributeByName(elements.AttributeNameReferenceText)
		if refText != nil {
			label = refText.AsciiDocString()
		}
		info := &Anchor{
			ID:      id,
			Label:   label,
			Element: wa,
			Parent:  parent,
		}
		if _, ok := anchors[id]; ok {
			slog.Debug("duplicate anchor; can't fix", "id", id)
			return false
		}

		if !strings.HasPrefix(id, "_") {
			anchors[id] = info
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
		}
		return false
	})
	doc.anchors = anchors
	return anchors, nil
}
