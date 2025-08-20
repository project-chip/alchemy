package spec

import "github.com/project-chip/alchemy/asciidoc"

type overlayAction uint8

const (
	overlayActionNone   overlayAction = iota
	overlayActionRemove               = (1 << iota)
	overlayActionReplace
	overlayActionOverrideParent
	overlayActionOverrideChildren
	overlayActionAppendElements
)

func (ppa overlayAction) Remove() bool {
	return ppa&overlayActionRemove != 0
}

func (ppa overlayAction) Replace() bool {
	return ppa&overlayActionReplace != 0
}

func (ppa overlayAction) OverrideChildren() bool {
	return ppa&overlayActionOverrideChildren != 0
}

func (ppa overlayAction) OverrideParent() bool {
	return ppa&overlayActionOverrideParent != 0
}

func (ppa overlayAction) Append() bool {
	return ppa&overlayActionAppendElements != 0
}

type elementOverlay struct {
	action   overlayAction
	replace  asciidoc.Elements
	parent   asciidoc.Element
	children asciidoc.Elements
	append   asciidoc.Elements
}

type elementOverlays map[asciidoc.Element]*elementOverlay

func (eo elementOverlays) remove(el asciidoc.Element) {
	if existing, ok := eo[el]; ok {
		existing.action |= overlayActionRemove
	} else {
		eo[el] = &elementOverlay{action: overlayActionRemove}
	}
}

func (eo elementOverlays) replace(el asciidoc.Element, replacement asciidoc.Elements) {
	existing, ok := eo[el]
	if !ok {
		existing = &elementOverlay{action: overlayActionReplace}
		eo[el] = existing
	} else {
		existing.action |= overlayActionReplace
	}
	existing.replace = replacement
}

func (eo elementOverlays) append(parent asciidoc.Element, children ...asciidoc.Element) {
	existing, ok := eo[parent]
	if !ok {
		existing = &elementOverlay{action: overlayActionAppendElements}
		eo[parent] = existing
	} else {
		existing.action |= overlayActionAppendElements
	}
	existing.append = append(existing.append, children...)
}

func (eo elementOverlays) appendChild(parent asciidoc.Element, child asciidoc.Element) {
	existing, ok := eo[parent]
	if !ok {
		existing = &elementOverlay{action: overlayActionOverrideChildren}
		eo[parent] = existing
	} else {
		existing.action |= overlayActionOverrideChildren
	}
	existing.children = append(existing.children, child)
}

func (eo elementOverlays) setParent(child asciidoc.Element, parent asciidoc.Element) {
	existing, ok := eo[child]
	if !ok {
		existing = &elementOverlay{action: overlayActionOverrideParent}
		eo[child] = existing
	} else {
		existing.action |= overlayActionOverrideParent
	}
	existing.parent = parent
}
