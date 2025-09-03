package overlay

import (
	"fmt"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

type Reader struct {
	overlays elementOverlays
}

func (r *Reader) Iterate(parent asciidoc.Parent, elements asciidoc.Elements) asciidoc.ElementIterator {
	return func(yield func(asciidoc.Element) bool) {
		if !r.iterate(elements, yield) {
			return
		}
		var parentOverlay *elementOverlay
		switch parent := parent.(type) {
		case asciidoc.Element:
			parentOverlay = r.overlays[parent]
		}
		if parentOverlay != nil && len(parentOverlay.append) > 0 {
			for _, el := range parentOverlay.append {
				if !yield(el) {
					return
				}
			}
		}
	}
}

func (r *Reader) iterate(elements asciidoc.Elements, yield func(asciidoc.Element) bool) bool {
	for _, el := range elements {
		elementOverlay, ok := r.overlays[el]
		if !ok {
			if !yield(el) {
				return false
			}
		} else {
			if elementOverlay.action.Remove() {
				continue
			}
			if elementOverlay.action.Replace() {
				if !r.iterate(elementOverlay.replace, yield) {
					return false
				}
				continue
			}
			if !yield(el) {
				return false
			}
		}
	}
	return true
}

func (r *Reader) StringValue(parent asciidoc.Parent, elements asciidoc.Elements) (string, error) {
	return renderSimpleElements(elements)
}

func (r *Reader) Parent(child asciidoc.ChildElement) asciidoc.Element {
	if overlay, ok := r.overlays[child]; ok && overlay.action.OverrideParent() {
		return overlay.parent
	}
	return child.Parent()
}

func (r *Reader) Children(parent asciidoc.ParentElement) asciidoc.Elements {
	if overlay, ok := r.overlays[parent]; ok && overlay.action.OverrideChildren() {
		return overlay.children
	}
	return parent.Children()
}

func renderSimpleElements(els asciidoc.Elements) (string, error) {
	var sb strings.Builder
	for _, el := range els {
		switch el := el.(type) {
		case *asciidoc.String:
			sb.WriteString(el.Value)
		case *asciidoc.NewLine:
			sb.WriteRune('\n')
		case *asciidoc.EmptyLine:
			sb.WriteRune('\n')
		case *asciidoc.CharacterReplacementReference:
			sb.WriteString(el.ReplacementValue())
		case *asciidoc.FileInclude:
			sb.WriteString(el.Raw())
			sb.WriteRune('\n')
		default:
			return "", fmt.Errorf("unexpected type rendering simple elements: %T", el)
		}
	}
	return sb.String(), nil
}
