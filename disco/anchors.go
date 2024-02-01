package disco

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
)

var properAnchorPattern = regexp.MustCompile(`^ref_[A-Z][a-z]+(?:[A-Z][a-z]*)*([A-Z][a-z]*(?:[A-Z][a-z]*)*)*$`)
var acronymPattern = regexp.MustCompile(`[A-Z]{3,}`)

func (b *Ball) normalizeAnchors(doc *ascii.Doc) error {

	anchors, err := doc.Anchors()
	if err != nil {
		return fmt.Errorf("error fetching anchors in %s: %w", doc.Path, err)
	}

	newAnchors := make(map[string][]*ascii.Anchor)
	for _, info := range anchors {
		// Fix all the bad references, and add to list of new anchors, ignoring duplicates for now
		normalizeAnchor(info)
		newAnchors[info.ID] = append(newAnchors[info.ID], info)
	}

	for a, infos := range newAnchors {
		if len(infos) > 1 { // We ended up with some duplicate anchors
			err := disambiguateAnchorSet(infos)
			if err != nil {
				return fmt.Errorf("error disambiguating anchors %s in %s: %w", a, doc.Path, err)
			}
		}
		for _, info := range infos {
			setAnchor(info)
		}
	}

	b.rewriteCrossReferences(doc.CrossReferences(), anchors)
	return nil
}

func normalizeAnchor(info *ascii.Anchor) {
	if properAnchorPattern.Match([]byte(info.ID)) {
		if len(info.Label) == 0 {
			info.Label = strings.TrimSpace(matter.StripReferenceSuffixes(ascii.ReferenceName(info.Element)))
		}
		return
	}
	id, label := normalizeAnchorID(info.Name, info.Element, info.Parent)
	info.ID = id
	info.Label = label
}

var pascalCasePattern = regexp.MustCompile(`^[A-Z][a-z]+([A-Z][a-z]+)+$`)
var anchorInvalidCharacters = strings.NewReplacer(".", "", "(", "", ")", "")

func normalizeAnchorID(name string, element any, parent any) (id string, label string) {
	var parentName string
	switch element.(type) {
	case *types.Table:
		label = strings.TrimSpace(name)
	default:
		label = strings.TrimSpace(matter.StripReferenceSuffixes(name))
	}

	switch p := parent.(type) {
	case *ascii.Section:
		switch p.SecType {
		case matter.SectionDataTypeStruct, matter.SectionCommand, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap, matter.SectionEvent:
			parentName = ascii.ReferenceName(p.Base)
			parentName = ascii.StripTypeSuffixes(parentName)
			parentName, _ = normalizeAnchorID(parentName, p.Base, p.Parent)
			parentName = strings.TrimPrefix(parentName, "ref_")
		}
	}

	var ref strings.Builder

	ref.WriteString(parentName)

	parts := strings.Split(label, " ")
	for _, p := range parts {
		p = anchorInvalidCharacters.Replace(p)
		if pascalCasePattern.MatchString(p) {
			ref.WriteString(p)
		} else {
			ref.WriteString(titleCaser.String(p))
		}
	}
	id = ref.String()
	id = "ref_" + acronymPattern.ReplaceAllStringFunc(id, func(match string) string {
		return string(match[0]) + strings.ToLower(string(match[1:]))
	})
	return
}

func setAnchor(info *ascii.Anchor) {
	setAnchorID(info.Element, info.ID, info.Label)
}

func setAnchorID(element types.WithAttributes, id string, label string) {
	newAttr := make(types.Attributes)
	if len(label) > 0 {
		id += ", " + label
	}
	newAttr[types.AttrID] = id
	element.AddAttributes(newAttr)
}

func disambiguateAnchorSet(infos []*ascii.Anchor) error {
	parents := make([]interface{}, len(infos))
	refIds := make([]string, len(infos))
	for i, info := range infos {
		parents[i] = info.Parent
		refIds[i] = info.ID
	}
	parentSections := make([]*ascii.Section, len(infos))
	for {
		for i := range infos {
			parentSection := findRefSection(parents[i])
			if parentSection == nil {
				return fmt.Errorf("duplicate anchor: %s with invalid parent", refIds[i])
			}
			parentSections[i] = parentSection
			parentName := ascii.ReferenceName(parentSection.Base)
			parentName = ascii.StripTypeSuffixes(parentName)
			refParentId, _ := normalizeAnchorID(parentName, parentSection.Base, parentSection.Parent)
			refIds[i] = refParentId + strings.TrimPrefix(refIds[i], "ref_")
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
		slog.Debug("Switching duplicate anchor", "from", info.ID, "to", refIds[i])
		info.ID = refIds[i]
	}
	return nil
}
