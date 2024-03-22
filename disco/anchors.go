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

var properAnchorPattern = regexp.MustCompile(`^ref_[A-Z]+[a-z]+(?:[A-Z]+[a-z]*)*([A-Z]+[a-z]*(?:[A-Z]+[a-z]*)*)*$`)

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
			setAnchorID(info.Element, info.ID, info.Label)
		}
	}

	b.rewriteCrossReferences(doc.CrossReferences(), anchors)
	return nil
}

func normalizeAnchor(info *ascii.Anchor) {
	if properAnchorPattern.Match([]byte(info.ID)) {
		if len(info.Label) == 0 {
			info.Label = normalizeAnchorLabel(info.Name(), info.Element)
		}
	} else {
		name := info.Name()
		if len(name) == 0 {
			name = info.Label
		}
		id, label := normalizeAnchorID(name, info.Element, info.Parent)
		info.ID = id
		info.Label = label
	}
	name := info.Name()
	if info.Label == name {
		label := normalizeAnchorLabel(name, info.Element)
		if len(label) > 0 && label != name {
			info.Label = label
		} else {
			info.Label = ""
		}
	}
}

var anchorInvalidCharacters = strings.NewReplacer(".", "", "(", "", ")", "")

func normalizeAnchorID(name string, element any, parent any) (id string, label string) {
	var parentName string

	label = normalizeAnchorLabel(name, element)

	switch p := parent.(type) {
	case *ascii.Section:
		switch p.SecType {
		case matter.SectionDataTypeStruct, matter.SectionCommand, matter.SectionDataTypeEnum, matter.SectionDataTypeBitmap, matter.SectionEvent:
			parentName = ascii.ReferenceName(p.Base)
			parentName = matter.StripTypeSuffixes(parentName)
			parentName, _ = normalizeAnchorID(parentName, p.Base, p.Parent)
			parentName = strings.TrimPrefix(parentName, "ref_")
		case matter.SectionUnknown:
		case matter.SectionEvents:
		default:
			slog.Debug("unexpected parent section type", slog.String("sectionType", p.SecType.String()))
		}
	}

	var ref strings.Builder
	ref.WriteString("ref_")
	ref.WriteString(parentName)
	ref.WriteString(matter.Case(anchorInvalidCharacters.Replace(label)))
	id = ref.String()
	return
}

func normalizeAnchorLabel(name string, element any) (label string) {
	switch element.(type) {
	case *types.Table:
		label = strings.TrimSpace(name)
	default:
		name = strings.TrimSuffix(name, " Type")
		label = strings.TrimSpace(matter.StripReferenceSuffixes(name))
	}
	return
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
	parents := make([]any, len(infos))
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
			parentName = matter.StripTypeSuffixes(parentName)
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
