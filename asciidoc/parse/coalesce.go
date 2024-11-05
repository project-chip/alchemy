package parse

import (
	"log/slog"
	"slices"

	"github.com/project-chip/alchemy/asciidoc"
)

type coalesceState struct {
	inlineElements  asciidoc.Set
	text            []*asciidoc.String
	admonition      *asciidoc.Admonition
	blockAttributes *asciidoc.BlockAttributes
}

func coalesce(els asciidoc.Set) (out asciidoc.Set, err error) {
	if len(els) < 2 { // Single element doesn't need coalescing
		return els, nil
	}
	var state coalesceState
	for _, e := range els {
		//fmt.Printf("coalesce: %s %T\n", asciidoc.Describe(e), e)
		var el asciidoc.Element
		switch e := e.(type) {
		case *asciidoc.Admonition:
			state.flushText()
			out = state.flushInline(out)
			state.admonition = e
			continue
		case *asciidoc.BlockAttributes:
			state.flushText()
			out = state.flushInline(out)
			state.appendInlineElement(e)
			state.blockAttributes = e
		case asciidoc.Element:
			el = e
		default:
			//fmt.Printf("unexpected non-element in coalesce: %T\n", e)
			continue
		}

		if el == nil {
			continue
		}

		if hasElements, ok := el.(asciidoc.HasElements); ok {
			var els asciidoc.Set
			els, err = coalesce(hasElements.Elements())
			if err != nil {
				return
			}
			hasElements.SetElements(els)
		}

		if hasChild, ok := el.(asciidoc.HasChild); ok {
			child := hasChild.Child()
			if hasElements, ok := child.(asciidoc.HasElements); ok {
				var els asciidoc.Set
				els, err = coalesce(hasElements.Elements())
				if err != nil {
					return
				}
				hasElements.SetElements(els)
			}
		}

		switch el.Type() {
		case asciidoc.ElementTypeBlock:
			state.flushText()
			out = state.flushInline(out)
			if state.blockAttributes != nil {
				if attributableElement, ok := el.(asciidoc.AttributableElement); ok {
					err = attributableElement.ReadAttributes(attributableElement, state.blockAttributes.AttributeList...)
					if err != nil {
						return
					}
				}
				state.blockAttributes = nil
			}
			out = append(out, el)

		case asciidoc.ElementTypeInline, asciidoc.ElementTypeInlineLiteral:
			switch e := el.(type) {
			case *asciidoc.String:
				state.text = append(state.text, e)
			default:
				state.appendInlineElement(el)
			}
		default:
			//fmt.Printf("unexpected element type in coalesce: %T\n", e)
		}
	}
	state.flushText()
	out = state.flushInline(out)
	//fmt.Printf("coalesced: %d elements\n", len(els))
	return
}

func (cs *coalesceState) appendInlineElement(el asciidoc.Element) {
	cs.flushText()
	cs.inlineElements = append(cs.inlineElements, el)
}

func (cs *coalesceState) flushText() {
	if len(cs.text) == 0 {
		return
	}
	cs.inlineElements = append(cs.inlineElements, asciidoc.JoinStrings(cs.text))
	cs.text = nil
}

func (cs *coalesceState) flushInline(out asciidoc.Set) asciidoc.Set {
	if len(cs.inlineElements) == 0 {
		return out
	}
	if cs.blockAttributes == nil && cs.admonition == nil {
		out = append(out, cs.inlineElements...)
		cs.inlineElements = nil
		return out
	}
	p := copyPosition(out, asciidoc.NewParagraph())
	if cs.blockAttributes != nil {
		err := p.ReadAttributes(p, cs.blockAttributes.AttributeList...)
		if err != nil {
			slog.Warn("error reading attributes while flushing inline elements", slog.Any("error", err))
		}
		bsIndex := slices.IndexFunc(cs.inlineElements, func(e asciidoc.Element) bool {
			return e == cs.blockAttributes
		})
		if bsIndex >= 0 {
			cs.inlineElements.SetElements(slices.Delete(cs.inlineElements, bsIndex, bsIndex+1))
		}
		cs.blockAttributes = nil
	}
	if cs.admonition != nil {
		p.Admonition = cs.admonition.AdmonitionType
		attributes := cs.admonition.Attributes()
		if len(attributes) > 0 {
			err := p.ReadAttributes(p, attributes...)
			if err != nil {
				slog.Warn("error reading attributes while flushing inline elements", slog.Any("error", err))
			}
		}
		cs.admonition = nil
	}

	p.Set = cs.inlineElements
	out = append(out, p)
	cs.inlineElements = nil
	return out
}
