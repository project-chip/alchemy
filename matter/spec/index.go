package spec

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/project-chip/alchemy/internal/log"
)

type referenceIndex struct {
	anchorsParsed  bool
	anchors        map[string][]*Anchor
	anchorsByLabel map[string][]*Anchor

	crossReferencesParsed bool
	crossReferences       map[string][]*CrossReference
}

func newReferenceIndex() referenceIndex {
	return referenceIndex{
		anchors:         make(map[string][]*Anchor),
		anchorsByLabel:  make(map[string][]*Anchor),
		crossReferences: make(map[string][]*CrossReference),
	}
}

func (ri *referenceIndex) changeAnchor(anchor *Anchor, id string) {
	anchors, ok := ri.anchors[anchor.ID]
	if ok {
		i := slices.IndexFunc(anchors, func(a *Anchor) bool { return a == anchor })
		if i >= 0 {
			anchors = append(anchors[:i], anchors[i+1:]...)
			if len(anchors) > 0 {
				ri.anchors[anchor.ID] = anchors
			} else {
				delete(ri.anchors, anchor.ID)
			}
		}
	}
	ri.anchors[id] = append(ri.anchors[id], anchor)
}

func (ri *referenceIndex) changeCrossReference(reference *CrossReference, id string) {
	refs, ok := ri.crossReferences[reference.Reference.ID]
	if ok {
		i := slices.IndexFunc(refs, func(cr *CrossReference) bool { return cr == reference })
		if i >= 0 {
			refs = append(refs[:i], refs[i+1:]...)
			if len(refs) > 0 {
				ri.crossReferences[reference.Reference.ID] = refs
			} else {
				delete(ri.crossReferences, reference.Reference.ID)
			}
		}
	}
	ri.crossReferences[id] = append(ri.crossReferences[id], reference)
}

func (ri *referenceIndex) findAnchor(path fmt.Stringer, id string) *Anchor {

	anchors := ri.findAnchorsByID(id)
	switch len(anchors) {
	case 0:
		return nil
	case 1:
		return anchors[0]
	default:
		args := []any{slog.String("anchorId", id), slog.String("source", path.String())}
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

func (ri *referenceIndex) findAnchorByLabel(path fmt.Stringer, label string) *Anchor {
	anchors := ri.findAnchorsByLabel(label)
	switch len(anchors) {
	case 0:
		return nil
	case 1:
		return anchors[0]
	default:
		args := []any{slog.String("label", label), slog.String("source", path.String())}
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
