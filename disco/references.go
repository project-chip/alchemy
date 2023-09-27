package disco

import (
	"fmt"
	"log/slog"
	"regexp"
	"slices"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func getUnrecognizedReference(els []interface{}) string {
	if len(els) < 5 {
		return ""
	}
	sc, ok := els[0].(*types.SpecialCharacter)
	if !ok || sc.Name != "<" {
		return ""
	}
	sc, ok = els[1].(*types.SpecialCharacter)
	if !ok || sc.Name != "<" {
		return ""
	}
	s, ok := els[2].(*types.StringElement)
	if !ok {
		return ""
	}
	sc, ok = els[3].(*types.SpecialCharacter)
	if !ok || sc.Name != ">" {
		return ""
	}
	sc, ok = els[4].(*types.SpecialCharacter)
	if !ok || sc.Name != ">" {
		return ""
	}
	return s.Content
}

// The parser sometimes doesn't recognize inline references, and just returns them as "<", "<", "ref", ">", ">"
// Rather than doing the right thing and fixing the parser, we'll just find these little suckers after the fact
func fixUnrecognizedReferences(doc *ascii.Doc) {
	var elementsWithUnrecognizedReferences []HasElements
	traverse(nil, doc.Base.Elements, func(el interface{}, parent HasElements, index int) bool {
		sc, ok := el.(*types.SpecialCharacter)
		if !ok {
			return false
		}
		if sc.Name != "<" {
			return false
		}
		els := parent.GetElements()[index:]
		if getUnrecognizedReference(els) == "" {
			return false
		}
		elementsWithUnrecognizedReferences = append(elementsWithUnrecognizedReferences, parent)
		return false
	})
	for _, e := range elementsWithUnrecognizedReferences {
		els := e.GetElements()
		for i := 0; i < len(els); i++ {
			sub := els[i:]
			ref := getUnrecognizedReference(sub)
			if ref != "" {
				var id string
				var label string
				parts := strings.Split(ref, ",")
				id = parts[0]
				if len(parts) > 0 {
					label = parts[1]
				}
				icr, _ := types.NewInternalCrossReference(id, label)
				n := slices.Replace(els, i, i+5, interface{}(icr))
				e.SetElements(n)

				break
			}
		}
	}
}

type anchorInfo struct {
	id      string
	label   string
	element types.WithAttributes
	parent  HasElements
	name    string
}

var properAnchorPattern = regexp.MustCompile(`^ref_[A-Z][a-z]+(?:[A-Z][a-z]*)*(_[A-Z][a-z]*(?:[A-Z][a-z]*)*)*$`)
var acronymPattern = regexp.MustCompile(`[A-Z]{3,}`)

func normalizeAnchors(doc *ascii.Doc) error {

	crossReferences := findCrossReferences(doc)

	anchors, err := findAnchors(doc, crossReferences)
	if err != nil {
		return err
	}

	newAnchors := make(map[string][]*anchorInfo)
	for _, info := range anchors {
		// Fix all the bad references, and add to list of new anchors, ignoring duplicates for now
		normalizeAnchor(info)
		newAnchors[info.id] = append(newAnchors[info.id], info)
	}

	for _, infos := range newAnchors {
		if len(infos) > 1 { // We ended up with some duplicate anchors
			err := disambiguateAnchorSet(infos)
			if err != nil {
				return err
			}
		}
		for _, info := range infos {
			setAnchor(info)
		}
	}

	rewriteCrossReferences(crossReferences, anchors)
	return nil
}

func findCrossReferences(doc *ascii.Doc) map[string][]*types.InternalCrossReference {
	crossReferences := make(map[string][]*types.InternalCrossReference)
	traverse(nil, doc.Base.Elements, func(el interface{}, parent HasElements, index int) bool {
		if icr, ok := el.(*types.InternalCrossReference); ok {
			id, ok := icr.ID.(string)
			if !ok {
				return false
			}
			crossReferences[id] = append(crossReferences[id], icr)
		}
		return false
	})
	return crossReferences
}

func findAnchors(doc *ascii.Doc, crossReferences map[string][]*types.InternalCrossReference) (map[string]*anchorInfo, error) {
	anchors := make(map[string]*anchorInfo)
	if traverse(doc, doc.Elements, func(el interface{}, parent HasElements, index int) bool {
		var wa types.WithAttributes
		e, ok := el.(*ascii.Element)
		if ok {
			if wa, ok = e.Base.(types.WithAttributes); !ok {
				return false
			}
		} else if s, ok := el.(*ascii.Section); ok {
			wa = s.Base
		} else {
			return false
		}
		attrs := wa.GetAttributes()
		idAttr, ok := attrs[types.AttrID]
		if !ok {
			return false
		}
		id := strings.TrimSpace(idAttr.(string))
		fmt.Printf("anchor id: \"%s\"\n", id)
		var label string
		if parts := strings.Split(id, ","); len(parts) > 1 {
			id = strings.TrimSpace(parts[0])
			label = strings.TrimSpace(parts[1])
		}
		info := &anchorInfo{
			id:      id,
			label:   label,
			element: wa,
			parent:  parent,
		}
		name := getReferenceName(wa)
		if name != "" {
			info.name = name
		} else if len(label) > 0 {
			info.name = label
		}
		if _, ok := anchors[id]; ok {
			slog.Warn("duplicate anchor; can't fix", "id", id)
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
						slog.Warn("duplicate anchor; can't fix", "id", unescaped)
						return false
					}
					anchors[unescaped] = info
				}
			}
		}
		return false
	}) {
		return nil, fmt.Errorf("error traversing tree")
	}
	return anchors, nil
}

func rewriteCrossReferences(crossReferences map[string][]*types.InternalCrossReference, anchors map[string]*anchorInfo) {
	for id, xrefs := range crossReferences {
		info, ok := anchors[id]
		if !ok {
			slog.Warn("cross reference points to non-existent anchor", "name", id)
			continue
		}

		for _, xref := range xrefs {
			xref.OriginalID = info.id
			xref.ID = info.id
			// If the cross reference has a label that's the same as the one we generated for the anchor, remove it
			if label, ok := xref.Label.(string); ok && label == info.label {
				xref.Label = nil
			}
		}
	}
}

func normalizeAnchor(info *anchorInfo) {
	match := properAnchorPattern.FindStringSubmatch(info.id)
	if len(match) > 0 {
		if len(info.label) == 0 {
			info.label = getReferenceName(info.element)
		}
		return
	}
	id, label := normalizeAnchorID(info.name, info.element)
	info.id = id
	info.label = label
}

func normalizeAnchorID(name string, element interface{}) (id string, label string) {
	switch element.(type) {
	case *types.Table:
		label = strings.TrimSpace(name)
	default:
		label = strings.TrimSpace(stripReferenceSuffixes(name))
	}

	var ref strings.Builder

	parts := strings.Split(label, " ")
	for _, p := range parts {
		ref.WriteString(p)
	}
	id = ref.String()
	id = "ref_" + acronymPattern.ReplaceAllStringFunc(id, func(match string) string {
		return string(match[0]) + strings.ToLower(string(match[1:len(match)-1])) + string(match[len(match)-1:])
	})
	return
}

func setAnchor(info *anchorInfo) {
	newAttr := make(types.Attributes)
	id := info.id
	if len(info.label) > 0 {
		id += ", " + info.label
	}
	newAttr[types.AttrID] = id
	info.element.SetAttributes(newAttr)
}

func disambiguateAnchorSet(infos []*anchorInfo) error {
	parents := make([]interface{}, len(infos))
	refIds := make([]string, len(infos))
	for i, info := range infos {
		parents[i] = info.parent
		refIds[i] = info.id
	}
	parentSections := make([]*ascii.Section, len(infos))
	for {
		for i := range infos {
			parentSection := findRefSection(parents[i])
			if parentSection == nil {
				return fmt.Errorf("duplicate reference: %s with invalid parent", refIds[i])
			}
			parentSections[i] = parentSection
			refParentId, _ := normalizeAnchorID(getReferenceName(parentSection.Base), parentSection.Base)
			refIds[i] = refParentId + "_" + strings.TrimPrefix(refIds[i], "ref_")
		}
		ids := make(map[string]struct{})
		var duplicateIds bool
		for _, refId := range refIds {
			if _, ok := ids[refId]; ok {
				duplicateIds = true
			}
			ids[refId] = struct{}{}
		}
		if duplicateIds {
			for i := range infos {
				parents[i] = parentSections[i].Parent
			}
		} else {
			break
		}
	}
	for i, info := range infos {
		slog.Debug("Switching duplicate anchor", "from", info.id, "to", refIds[i])
		info.id = refIds[i]
	}
	return nil
}

func stripReferenceSuffixes(newId string) string {
	for _, suffix := range matter.DisallowedReferenceSuffixes {
		if strings.HasSuffix(newId, suffix) {
			newId = newId[0 : len(newId)-len(suffix)]
			break
		}
	}
	return newId
}

func findRefSection(parent interface{}) *ascii.Section {
	switch p := parent.(type) {
	case *ascii.Section:
		return p
	case *ascii.Element:
		return findRefSection(p.Parent)
	case *ascii.Doc:
		return nil
	default:
		return nil
	}

}

func getReferenceName(element interface{}) string {
	switch el := element.(type) {
	case *types.Section:
		name := types.Reduce(el.Title)
		if s, ok := name.(string); ok {
			return s
		}
	case *types.Table:
		if el.Attributes != nil {

			if title, ok := el.Attributes["title"]; ok {
				if name, ok := title.(string); ok {
					return name
				}
			}
		}
	default:
		//slog.
		slog.Debug("Unknown type to get reference name", "type", element)
	}
	return ""
}
