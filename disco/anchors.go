package disco

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

type anchorInfo struct {
	id      string
	label   string
	element types.WithAttributes
	parent  ascii.HasElements
	name    string
}

var properAnchorPattern = regexp.MustCompile(`^ref_[A-Z][a-z]+(?:[A-Z][a-z]*)*(_[A-Z][a-z]*(?:[A-Z][a-z]*)*)*$`)
var acronymPattern = regexp.MustCompile(`[A-Z]{3,}`)

func (b *Ball) normalizeAnchors(doc *ascii.Doc) error {

	crossReferences := b.findCrossReferences(doc)

	anchors, err := b.findAnchors(doc, crossReferences)
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

	b.rewriteCrossReferences(crossReferences, anchors)
	return nil
}

func (b *Ball) findAnchors(doc *ascii.Doc, crossReferences map[string][]*types.InternalCrossReference) (map[string]*anchorInfo, error) {
	anchors := make(map[string]*anchorInfo)
	ascii.Traverse(doc, doc.Elements, func(el interface{}, parent ascii.HasElements, index int) bool {
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
		var label string
		if parts := strings.Split(id, ","); len(parts) > 1 {
			id = strings.TrimSpace(parts[0])
			label = strings.TrimSpace(parts[1])
		}
		reftext, ok := attrs.GetAsString("reftext")
		if ok {
			label = reftext
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
	})

	return anchors, nil
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

var pascalCasePattern = regexp.MustCompile(`^[A-Z][a-z]+([A-Z][a-z]+)+$`)
var anchorInvalidCharacters = strings.NewReplacer(".", "", "(", "", ")", "")

func normalizeAnchorID(name string, element interface{}) (id string, label string) {
	switch element.(type) {
	case *types.Table:
		label = strings.TrimSpace(name)
	default:
		label = strings.TrimSpace(matter.StripReferenceSuffixes(name))
	}

	var ref strings.Builder

	parts := strings.Split(label, " ")
	for i, p := range parts {
		p = anchorInvalidCharacters.Replace(p)
		if pascalCasePattern.MatchString(p) {
			if i > 0 {
				ref.WriteString("_")
			}
			ref.WriteString(p)
		} else {
			ref.WriteString(titleCaser.String(p))
		}
	}
	id = ref.String()
	id = "ref_" + acronymPattern.ReplaceAllStringFunc(id, func(match string) string {
		return string(match[0]) + strings.ToLower(string(match[1:len(match)-1])) + string(match[len(match)-1:])
	})
	return
}

func setAnchor(info *anchorInfo) {
	setAnchorID(info.element, info.id, info.label)
}

func setAnchorID(element types.WithAttributes, id string, label string) {
	newAttr := make(types.Attributes)
	if len(label) > 0 {
		id += ", " + label
	}
	newAttr[types.AttrID] = id
	element.AddAttributes(newAttr)
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
		slog.Info("info", "i", infos)
		for i := range infos {
			slog.Info("info", "i", infos[i].name, "e", infos[i].element)

			parentSection := findRefSection(parents[i])
			if parentSection == nil {
				return fmt.Errorf("duplicate anchor: %s with invalid parent", refIds[i])
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
