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
	fmt.Printf("found unrecognized reference: %s\n", s.Content)

	return s.Content
}

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

type refInfo struct {
	id      string
	label   string
	element types.WithAttributes
	parent  HasElements
	name    string
}

var properReferencePattern = regexp.MustCompile(`^ref_[A-Z][a-z]+(?:[A-Z][a-z]*)*(_[A-Z][a-z]*(?:[A-Z][a-z]*)*)*$`)
var acronymPattern = regexp.MustCompile(`[A-Z]{3,}`)

func normalizeReferences(doc *ascii.Doc) error {

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

	anchors := make(map[string]*refInfo)
	newAnchors := make(map[string][]*refInfo)

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
		idAttr, ok := attrs["id"]
		if !ok {
			return false
		}
		id := strings.TrimSpace(idAttr.(string))
		var label string
		if parts := strings.Split(id, ","); len(parts) > 1 {
			id = strings.TrimSpace(parts[0])
			label = strings.TrimSpace(parts[1])
		}
		info := &refInfo{
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
		return fmt.Errorf("error traversing tree")
	}

	for _, info := range anchors {
		// Fix all the bad references, and add to list of new anchors, ignoring duplicates for now
		fixRef(info)
		newAnchors[info.id] = append(newAnchors[info.id], info)
	}

	for _, infos := range newAnchors {
		if len(infos) > 1 {
			err := disambiguateRefs(infos)
			if err != nil {
				return err
			}
		}
		for _, info := range infos {
			setRef(info)
		}
	}

	for id, xrefs := range crossReferences {
		info, ok := anchors[id]
		if !ok {
			slog.Warn("missing cross reference target", "name", id)
			continue
		}

		for _, xref := range xrefs {
			xref.OriginalID = info.id
			xref.ID = info.id
		}
	}
	return nil
}

func fixRef(info *refInfo) {
	match := properReferencePattern.FindStringSubmatch(info.id)
	if len(match) > 0 {
		return
	}
	id, label := getRefId(info.name)
	info.id = id
	info.label = label
}

func getRefId(name string) (id string, label string) {
	var ref strings.Builder
	parts := strings.Split(name, " ")
	for _, p := range parts {
		ref.WriteString(p)
	}
	label = ref.String()
	label = stripReferenceSuffixes(label)
	id = "ref_" + acronymPattern.ReplaceAllStringFunc(label, func(match string) string {
		return string(match[0]) + strings.ToLower(string(match[1:len(match)-1])) + string(match[len(match)-1:])
	})
	return
}

func setRef(info *refInfo) {
	newAttr := make(types.Attributes)
	newAttr[types.AttrID] = info.id + ", " + info.label
	info.element.SetAttributes(newAttr)
}

func disambiguateRefs(infos []*refInfo) error {
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
			refParentId, _ := getRefId(getReferenceName(parentSection.Base))
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
		fmt.Printf("switching duplicate id %s to %s\n", info.id, refIds[i])
		info.id = refIds[i]
	}
	return nil
	/*
		var refParent interface{} = ref.parent
				var dupeParent interface{} = dupe.parent
				var refId = ref.id
				var dupeId = dupe.id
				for {
					refParentSection := findRefSection(refParent)
					dupeParentSection := findRefSection(dupeParent)
					if refParentSection == nil {
						return fmt.Errorf("duplicate reference: %s with invalid parent", ref.id)
					}
					if dupeParentSection == nil {
						return fmt.Errorf("duplicate reference: %s with invalid parent", ref.id)
					}
					refParentId, _ := getRefId(getReferenceName(refParentSection.Base))
					refId = refParentId + "_" + strings.TrimPrefix(refId, "ref_")
					dupeParentId, _ := getRefId(getReferenceName(dupeParentSection.Base))
					dupeId = dupeParentId + "_" + strings.TrimPrefix(dupeId, "ref_")
					if refId == dupeId {
						refParent = refParentSection.Parent
						dupeParent = dupeParentSection.Parent
						continue
						//return fmt.Errorf("duplicate reference: %s with identical ids %s", ref.id, refId)
					}
					fmt.Printf("splitting id %s into %s and %s\n", ref.id, refId, dupeId)
					ref.id = refId
					dupe.id = dupeId
					setRef(ref)
					setRef(dupe)
					delete(existingReferences, id)
					existingReferences[refId] = ref
					existingReferences[dupeId] = dupe
					newReferences[dupe.id] = dupe
					break
				}*/
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
		fmt.Printf("Unknown type to get reference name: %T\n", element)
	}
	return ""
}
