package disco

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/hasty/adoc/elements"
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
			setAnchorID(info.Element, info.ID, info.LabelElements)
		}
	}

	b.rewriteCrossReferences(doc.CrossReferences(), anchors)
	return nil
}

func normalizeAnchor(info *ascii.Anchor) {
	if properAnchorPattern.Match([]byte(info.ID)) {
		if len(info.LabelElements) == 0 {
			info.LabelElements = normalizeAnchorLabel(info.Name(), info.Element)
		}
	} else {
		name := info.Name()
		if len(name) == 0 {
			name = elements.AttributeAsciiDocString(info.LabelElements)
		}
		id, label := normalizeAnchorID(name, info.Element, info.Parent)
		info.ID = id
		info.LabelElements = label
	}
	name := info.Name()
	if strings.TrimSpace(elements.AttributeAsciiDocString(info.LabelElements)) == name {
		label := normalizeAnchorLabel(name, info.Element)
		if len(label) > 0 && strings.TrimSpace(elements.AttributeAsciiDocString(label)) != name {
			info.LabelElements = label
		} else {
			info.LabelElements = nil
		}
	}
}

var anchorInvalidCharacters = strings.NewReplacer(".", "", "(", "", ")", "")

func normalizeAnchorID(name string, element any, parent any) (id string, label elements.Set) {
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
	labelString := elements.AttributeAsciiDocString(label)
	ref.WriteString(matter.Case(anchorInvalidCharacters.Replace(labelString)))
	id = ref.String()
	return
}

func normalizeAnchorLabel(name string, element any) (label elements.Set) {
	switch element.(type) {
	case *elements.Table:
		label = elements.Set{elements.NewString(strings.TrimSpace(name))}
	default:
		name = strings.TrimSuffix(name, " Type")
		label = elements.Set{elements.NewString(strings.TrimSpace(matter.StripReferenceSuffixes(name)))}
	}
	return
}

func setAnchorID(element elements.Attributable, id string, label elements.Set) {
	var idAttribute *elements.NamedAttribute
	var refTextAttribute *elements.NamedAttribute
	for _, a := range element.Attributes() {
		switch a := a.(type) {
		case *elements.AnchorAttribute:
			a.ID = elements.NewString(id)
			if len(label) > 0 {
				a.Label = label
			} else {
				a.Label = nil
			}
			return
		case *elements.NamedAttribute:
			switch a.Name {
			case elements.AttributeNameID:
				idAttribute = a
			case elements.AttributeNameReferenceText:
				refTextAttribute = a
			}
		}
	}

	if idAttribute != nil {
		idAttribute.Val = elements.Set{elements.NewString(id)}
		if len(label) > 0 {
			if refTextAttribute != nil {
				refTextAttribute.Val = label
			} else {
				element.AppendAttribute(elements.NewNamedAttribute(string(elements.AttributeNameReferenceText), label, elements.AttributeQuoteTypeDouble))
			}
		}
		return
	}
	element.AppendAttribute(elements.NewAnchorAttribute(elements.NewString(id), label))
}

func disambiguateAnchorSet(infos []*ascii.Anchor) error {
	parents := make([]any, len(infos))
	refIDs := make([]string, len(infos))
	for i, info := range infos {
		parents[i] = info.Parent
		refIDs[i] = info.ID
	}
	parentSections := make([]*ascii.Section, len(infos))
	for {
		for i := range infos {
			parentSection := findRefSection(parents[i])
			if parentSection == nil {
				return fmt.Errorf("duplicate anchor: %s with invalid parent", refIDs[i])
			}
			parentSections[i] = parentSection
			parentName := ascii.ReferenceName(parentSection.Base)
			parentName = matter.StripTypeSuffixes(parentName)
			refParentID, _ := normalizeAnchorID(parentName, parentSection.Base, parentSection.Parent)
			refIDs[i] = refParentID + strings.TrimPrefix(refIDs[i], "ref_")
		}
		ids := make(map[string]struct{})
		var duplicateIds bool
		for _, refID := range refIDs {
			if _, ok := ids[refID]; ok {
				duplicateIds = true
			}
			ids[refID] = struct{}{}
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
		slog.Debug("Switching duplicate anchor", "from", info.ID, "to", refIDs[i])
		info.ID = refIDs[i]
	}
	return nil
}
