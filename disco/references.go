package disco

import (
	"fmt"
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
	var elementsWithUnrecognizedReferences []types.WithElements
	traverse(nil, doc.Base.Elements, func(el interface{}, parent types.WithElements, index int) bool {
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
	element interface{}
	name    string
}

var properReferencePattern = regexp.MustCompile(`^ref_[A-Z][a-z]+(?:[A-Z][a-z]*)*(_[A-Z][a-z]*(?:[A-Z][a-z]*)*)*$`)
var acronymPattern = regexp.MustCompile(`[A-Z]{3,}`)

func normalizeReferences(doc *ascii.Doc) {

	referenceCount := 0

	crossReferences := make(map[string][]*types.InternalCrossReference)
	traverse(nil, doc.Base.Elements, func(el interface{}, parent types.WithElements, index int) bool {
		if icr, ok := el.(*types.InternalCrossReference); ok {
			//fmt.Printf("ICR! %s %s %s\n", icr.ID, icr.OriginalID, icr.Label)
			id, ok := icr.ID.(string)
			if !ok {
				return false
			}
			list := crossReferences[id]
			list = append(list, icr)
			crossReferences[id] = list
			referenceCount++
		}
		return false
	})

	references := make(map[string]*refInfo)
	referenceNames := make(map[string]*refInfo)
	traverse(nil, doc.Base.Elements, func(el interface{}, parent types.WithElements, index int) bool {
		wa, ok := el.(types.WithAttributes)
		if !ok {
			return false
		}
		attrs := wa.GetAttributes()
		idAttr, ok := attrs["id"]
		if !ok {
			return false
		}
		id := strings.TrimSpace(idAttr.(string))
		fmt.Printf("element: %T id: %s\n", el, id)
		var label string
		if parts := strings.Split(id, ","); len(parts) > 1 {
			id = strings.TrimSpace(parts[0])
			label = strings.TrimSpace(parts[1])
		}
		info := &refInfo{
			id:      id,
			label:   label,
			element: el,
		}
		name := getReferenceName(el)
		if name != "" {
			fmt.Printf("Setting id: %s to %s\n", id, name)
			info.name = name
		} else if len(label) > 0 {
			info.name = label
		}
		if _, ok := references[id]; ok {
			fmt.Printf("duplicate reference! %s\n", id)
			return false
		}
		if info.name != "" {
			referenceNames[info.name] = info

		}
		if strings.HasPrefix(id, "_") {
			_, ok = crossReferences[id]
			if ok {
				references[id] = info
			} else {
				unescaped := strings.ReplaceAll(id, "_", " ")[1:]
				_, ok = crossReferences[unescaped]
				if !ok {
					fmt.Printf("ID not in use: %s -> %s\n", id, unescaped)
					return false
				}
				references[unescaped] = info
			}
		} else {
			references[id] = info
		}
		return false
	})

	for _, ref := range references {
		fixRef(ref)
	}

	newIDs := make(map[string]struct{})
	for _, ref := range references {
		_, ok := newIDs[ref.id]
		if ok {
			fmt.Printf("Duplicate id! %s\n", ref.id)
			return
		}
		newIDs[ref.id] = struct{}{}
	}

	fmt.Printf("reference count: %d names: %d\n", referenceCount, len(referenceNames))
	for id, xrefs := range crossReferences {
		info, ok := references[id]
		if !ok {
			fmt.Printf("missing cross reference target id: %s\n", id)
			info, ok = referenceNames[id]
			if !ok {
				fmt.Printf("missing cross reference target name: %s\n", id)
				continue
			}

		}
		for _, xref := range xrefs {
			xref.ID = info.id
		}
	}

}

func fixRef(info *refInfo) {
	match := properReferencePattern.FindStringSubmatch(info.id)
	if len(match) > 0 {
		fmt.Printf("Good reference: %s\n", info.id)
		return
	}
	var ref strings.Builder
	ref.WriteString("ref_")
	parts := strings.Split(info.name, " ")
	for _, p := range parts {
		ref.WriteString(p)
	}
	newId := ref.String()
	newId = acronymPattern.ReplaceAllStringFunc(newId, func(match string) string {
		return string(match[0]) + strings.ToLower(string(match[1:len(match)-1])) + string(match[len(match)-1:])
	})
	for _, suffix := range matter.DisallowedReferenceSuffixes {
		if strings.HasSuffix(newId, suffix) {
			newId = newId[0 : len(newId)-len(suffix)]
			break
		}
	}
	fmt.Printf("Bad reference: %s (%s) to %s\n", info.id, info.name, newId)
	info.id = newId
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
