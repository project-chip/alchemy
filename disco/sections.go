package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func assignSectionTypes(top *ascii.Section) {
	top.SecType = matter.SectionTop

	ascii.Traverse(top, top.Elements, func(el interface{}, parent ascii.HasElements, index int) bool {
		section, ok := el.(*ascii.Section)
		if !ok {
			//	fmt.Printf("not an ascii.Section: %T\n", parent)
			return false
		}
		ps, ok := parent.(*ascii.Section)
		if !ok {
			//	fmt.Printf("parent not an ascii.Section: %T\n", parent)
			return false
		}

		fmt.Printf("ascii.Section: %s\n", matter.SectionTypeName(ps.SecType))
		section.SecType = getSectionType(ps, section)
		fmt.Printf("%s -> %s!\n", section.Name, matter.SectionTypeName(section.SecType))
		return false
	})
}

func reorderTopLevelSection(sec *ascii.Section, docType matter.DocType) error {
	sectionOrder, ok := matter.TopLevelSectionOrders[docType]
	if !ok {
		slog.Warn("could not determine section order", "docType", docType)
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

func getSectionType(parent *ascii.Section, section *ascii.Section) matter.Section {
	name := strings.ToLower(strings.TrimSpace(section.Name))
	switch parent.SecType {
	case matter.SectionTop:
		switch name {
		case "introduction":
			return matter.SectionIntroduction
		case "revision history":
			return matter.SectionRevisionHistory
		case "classification":
			return matter.SectionClassification
		case "cluster identifiers", "cluster id", "cluster ids":
			return matter.SectionClusterID
		case "features":
			return matter.SectionFeatures
		case "dependencies":
			return matter.SectionDependencies
		case "data types":
			return matter.SectionDataTypes
		case "status codes":
			return matter.SectionStatusCodes
		case "attributes":
			return matter.SectionAttributes
		case "commands":
			return matter.SectionCommands
		case "conditions":
			return matter.SectionConditions
		case "cluster requirements":
			return matter.SectionClusterRequirements
		case "cluster restrictions":
			return matter.SectionClusterRestrictions
		case "element requirements":
			return matter.SectionElementRequirements
		case "endpoint composition":
			return matter.SectionEndpointComposition
		default:
			if strings.HasSuffix(name, " attribute set") {
				return matter.SectionAttributes
			}
			return matter.SectionUnknown
		}
	case matter.SectionAttributes:
		if strings.HasSuffix(name, " attribute") {
			return matter.SectionAttribute
		}
	case matter.SectionDataTypes:
		if strings.HasSuffix(name, "bitmap type") {
			return matter.SectionDataTypeBitmap
		}
		if strings.HasSuffix(name, "enum type") {
			return matter.SectionDataTypeEnum
		}
		if strings.HasSuffix(name, "struct type") {
			return matter.SectionDataTypeStruct
		}
	case matter.SectionCommand:
		if strings.HasSuffix(name, " field") {
			return matter.SectionField
		}
	case matter.SectionCommands:
		if strings.HasSuffix(name, " command") {
			return matter.SectionCommand
		}

	default:

	}
	return matter.SectionUnknown
}

func setSectionTitle(sec *ascii.Section, title string) {
	for _, e := range sec.Base.Title {
		switch el := e.(type) {
		case *types.StringElement:
			el.Content = title
		}
	}
}

func findSectionByType(top *ascii.Section, sectionType matter.Section) *ascii.Section {
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
