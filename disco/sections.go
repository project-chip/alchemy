package disco

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func reorderTopLevelSection(sec *ascii.Section, docType matter.DocType) error {
	sectionOrder, ok := matter.TopLevelSectionOrders[docType]
	if !ok {
		//slog.Warn("could not determine section order", "docType", docType)
		return nil
	}
	validSectionTypes := make(map[matter.Section]struct{}, len(sectionOrder)+1)
	for _, st := range sectionOrder {
		validSectionTypes[st] = struct{}{}
	}
	sections := divyUpSection(sec, validSectionTypes)
	newOrder := make([]interface{}, 0, len(sec.Elements))
	for _, st := range sectionOrder {
		if els, ok := sections[st]; ok {
			newOrder = append(newOrder, els...)
			delete(sections, st)
		}
	}
	if len(sections) > 0 {
		return fmt.Errorf("non-empty section list after reordering")
	}
	sec.SetElements(newOrder)
	return nil
}

func divyUpSection(sec *ascii.Section, validSectionTypes map[matter.Section]struct{}) map[matter.Section][]interface{} {
	sections := make(map[matter.Section][]interface{})
	lastSectionType := matter.SectionPrefix
	for _, e := range sec.Elements {
		switch el := e.(type) {
		case *ascii.Section:
			if el.SecType != matter.SectionUnknown {
				_, ok := validSectionTypes[el.SecType]
				if ok {
					lastSectionType = el.SecType
				}
			}
		}
		sections[lastSectionType] = append(sections[lastSectionType], e)
	}
	return sections
}

func setSectionTitle(sec *ascii.Section, title string) {
	for _, e := range sec.Base.Title {
		switch el := e.(type) {
		case *types.StringElement:
			el.Content = title
		}
	}
}

func FindSectionByType(top *ascii.Section, sectionType matter.Section) *ascii.Section {
	var found *ascii.Section
	ascii.Search(top.Elements, func(s *ascii.Section) bool {
		if s.SecType == sectionType {
			found = s
			return true
		}
		return false
	})
	return found
}
