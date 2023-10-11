package parse

import (
	"strings"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func AssignSectionTypes(top *ascii.Section) {
	top.SecType = matter.SectionTop

	ascii.Traverse(top, top.Elements, func(el interface{}, parent ascii.HasElements, index int) bool {
		section, ok := el.(*ascii.Section)
		if !ok {
			return false
		}
		ps, ok := parent.(*ascii.Section)
		if !ok {
			return false
		}

		section.SecType = getSectionType(ps, section)
		return false
	})
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
