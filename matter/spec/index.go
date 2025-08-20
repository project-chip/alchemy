package spec

import (
	"log/slog"
	"slices"
	"strings"
	"sync"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
)

type referenceIndex struct {
	sync.RWMutex

	anchorsParsed  bool
	anchors        map[string][]*Anchor
	identifiers    map[asciidoc.Element]string
	anchorsByLabel map[string][]*Anchor

	crossReferencesParsed bool
	crossReferences       map[string][]*CrossReference
}

func newReferenceIndex() referenceIndex {
	return referenceIndex{
		anchors:         make(map[string][]*Anchor),
		identifiers:     make(map[asciidoc.Element]string),
		anchorsByLabel:  make(map[string][]*Anchor),
		crossReferences: make(map[string][]*CrossReference),
	}
}

func (ri *referenceIndex) anchorId(reader asciidoc.Reader, parent asciidoc.Parent, element asciidoc.Element, id asciidoc.Elements) string {
	if existing, ok := ri.identifiers[element]; ok {
		return existing
	}

	var s strings.Builder
	for el := range reader.Iterate(parent, id) {
		switch el := el.(type) {
		case *asciidoc.String:
			s.WriteString(el.Value)
		default:
			slog.Warn("unexpected type in anchor id", log.Type("type", el))
			for el := range reader.Iterate(parent, id) {
				slog.Warn("unexpected type in anchor id2", log.Type("type", el))

			}
		}
	}
	ri.identifiers[element] = s.String()
	return s.String()
}

func (ri *referenceIndex) changeAnchor(reader asciidoc.Reader, anchor *Anchor, parent asciidoc.Parent, id asciidoc.Elements) {
	anchorID := ri.anchorId(reader, parent, anchor.Element, anchor.ID)

	anchors, ok := ri.anchors[anchorID]
	if ok {
		i := slices.IndexFunc(anchors, func(a *Anchor) bool { return a == anchor })
		if i >= 0 {
			anchors = append(anchors[:i], anchors[i+1:]...)
			if len(anchors) > 0 {
				ri.anchors[anchorID] = anchors
			} else {
				delete(ri.anchors, anchorID)
			}
		}
	}
	anchorID = ri.anchorId(reader, anchor.Parent, anchor.Element, id)
	ri.anchors[anchorID] = append(ri.anchors[anchorID], anchor)
}

func (ri *referenceIndex) changeCrossReference(reader asciidoc.Reader, reference *CrossReference, id asciidoc.Elements) {
	referenceID := ri.anchorId(reader, reference.Reference, reference.Reference, reference.Reference.ID)
	refs, ok := ri.crossReferences[referenceID]
	if ok {
		i := slices.IndexFunc(refs, func(cr *CrossReference) bool { return cr == reference })
		if i >= 0 {
			refs = append(refs[:i], refs[i+1:]...)
			if len(refs) > 0 {
				ri.crossReferences[referenceID] = refs
			} else {
				delete(ri.crossReferences, referenceID)
			}
		}
	}
	referenceID = ri.anchorId(reader, reference.Reference, reference.Reference, id)
	ri.crossReferences[referenceID] = append(ri.crossReferences[referenceID], reference)
}

func (ri *referenceIndex) findAnchor(source log.Source, id string) *Anchor {

	anchors := ri.findAnchorsByID(id)
	switch len(anchors) {
	case 0:
		return nil
	case 1:
		return anchors[0]
	default:
		args := []any{slog.String("anchorId", id), log.Path("source", source)}
		for _, an := range anchors {
			args = append(args, log.Path("target", an.Source))
		}
		slog.Warn("ambiguous anchor id reference", args...)
		return nil
	}
}

func (ri *referenceIndex) findAnchorsByID(id string) []*Anchor {
	return ri.anchors[id]
}

func (ri *referenceIndex) findAnchorByLabel(source log.Source, label string) *Anchor {
	anchors := ri.findAnchorsByLabel(label)
	switch len(anchors) {
	case 0:
		return nil
	case 1:
		return anchors[0]
	default:
		args := []any{slog.String("label", label), log.Path("source", source)}
		for _, an := range anchors {
			args = append(args, log.Path("target", an.Source))
		}
		slog.Warn("ambiguous anchor label reference", args...)
		return nil
	}

}

func (ri *referenceIndex) findAnchorsByLabel(label string) (anchors []*Anchor) {
	anchors = ri.anchorsByLabel[label]
	return
}
