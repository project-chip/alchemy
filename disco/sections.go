package disco

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

var topLevelSectionOrders map[matter.Doc][]matter.Section

func init() {
	topLevelSectionOrders = make(map[matter.Doc][]matter.Section)
	topLevelSectionOrders[matter.DocAppCluster] = []matter.Section{
		matter.SectionPrefix,
		matter.SectionRevisionHistory,
		matter.SectionClassification,
		matter.SectionClusterID,
		matter.SectionFeatures,
		matter.SectionDependencies,
		matter.SectionDataTypes,
		matter.SectionStatusCodes,
		matter.SectionAttributes,
		matter.SectionCommands,
		matter.SectionEvents,
	}
	topLevelSectionOrders[matter.DocDeviceType] = []matter.Section{
		matter.SectionPrefix,
		matter.SectionRevisionHistory,
		matter.SectionClassification,
		matter.SectionConditions,
		matter.SectionClusterRequirements,
		matter.SectionClusterRestrictions,
		matter.SectionElementRequirements,
		matter.SectionEndpointComposition,
	}
}

func reorderTopLevelSection(sec *ascii.Section, docType matter.Doc) {
	sectionOrder, ok := topLevelSectionOrders[docType]
	if !ok {
		fmt.Printf("could not determine section order from doc type %d\n", docType)
		return
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
		panic(fmt.Errorf("non-empty section list after reordering"))
	}
	sec.Elements = newOrder
}

func divyUpSection(sec *ascii.Section, validSectionTypes map[matter.Section]struct{}) map[matter.Section][]interface{} {
	sections := make(map[matter.Section][]interface{})
	lastSectionType := matter.SectionPrefix
	for _, e := range sec.Elements {
		switch el := e.(type) {
		case *ascii.Section:
			el.SecType = getSectionType(el)
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

func getSectionType(section *ascii.Section) matter.Section {
	name := strings.ToLower(strings.TrimSpace(section.Name))
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
		return matter.SectionUnknown
	}
}

func setSectionTitle(sec *ascii.Section, title string) {
	for _, e := range sec.Base.Title {
		switch el := e.(type) {
		case *types.StringElement:
			el.Content = title
		}
	}
}
