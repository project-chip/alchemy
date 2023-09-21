package disco

import (
	"fmt"
	"strings"

	"github.com/hasty/matterfmt/ascii"
)

type MatterDoc uint8

var topLevelSectionOrders map[MatterDoc][]ascii.MatterSection

const (
	MatterDocUnknown MatterDoc = iota
	MatterDocAppCluster
	MatterDocAppClusterIndex
	MatterDocDeviceType
	MatterDocDeviceTypeIndex
	MatterDocCommonProtocol
	MatterDocDataModel
	MatterDocDeviceAttestation
	MatterDocServiceDeviceManagement
)

func init() {
	topLevelSectionOrders = make(map[MatterDoc][]ascii.MatterSection)
	topLevelSectionOrders[MatterDocAppCluster] = []ascii.MatterSection{
		ascii.MatterSectionPrefix,
		ascii.MatterSectionRevisionHistory,
		ascii.MatterSectionClassification,
		ascii.MatterSectionClusterID,
		ascii.MatterSectionFeatures,
		ascii.MatterSectionDependencies,
		ascii.MatterSectionDataTypes,
		ascii.MatterSectionStatusCodes,
		ascii.MatterSectionAttributes,
		ascii.MatterSectionCommands,
		ascii.MatterSectionEvents,
	}
	topLevelSectionOrders[MatterDocDeviceType] = []ascii.MatterSection{
		ascii.MatterSectionPrefix,
		ascii.MatterSectionRevisionHistory,
		ascii.MatterSectionClassification,
		ascii.MatterSectionConditions,
		ascii.MatterSectionClusterRequirements,
		ascii.MatterSectionClusterRestrictions,
		ascii.MatterSectionElementRequirements,
		ascii.MatterSectionEndpointComposition,
	}
}

func reorderTopLevelSection(sec *ascii.Section, docType MatterDoc) {
	//reorderedList := make([]interface{}, 0, len(sec.Elements))
	sectionOrder, ok := topLevelSectionOrders[docType]
	if !ok {
		fmt.Printf("could not determine section order from doc type %d\n", docType)
		return
	}
	validSectionTypes := make(map[ascii.MatterSection]struct{}, len(sectionOrder)+1)
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

	/*for _, st := range sectionOrder[0:1] {
		fmt.Printf("looking for section %v...\n", st)
		var sectionStart int = -1
		var sectionEnd int = -1
		for i, e := range sec.Elements {
			switch el := e.(type) {
			case *ascii.Section:
				cst := getSectionType(el)
				fmt.Printf("\t%s:\t%s\n", el.Name, ascii.SectionTypeString(cst))
				if sectionStart == -1 {
					if cst == st {
						sectionStart = i
					}
				} else if sectionEnd == -1 {
					if cst != ascii.MatterSectionUnknown {
						sectionEnd = i
					}
				}
			}
			if sectionStart != -1 && sectionEnd != -1 {
				//break
			}
		}
		if sectionStart == -1 {
			//fmt.Printf("failed to find section %v...\n", st)
		} else {
			if sectionEnd != -1 {
				//	fmt.Printf("found section %v between %d and %d...\n", st, sectionStart, sectionEnd)
			} else {
				//fmt.Printf("found section %v between %d and end...\n", st, sectionStart)
			}
		}
	}*/
}

func divyUpSection(sec *ascii.Section, validSectionTypes map[ascii.MatterSection]struct{}) map[ascii.MatterSection][]interface{} {
	sections := make(map[ascii.MatterSection][]interface{})
	lastSectionType := ascii.MatterSectionPrefix
	for _, e := range sec.Elements {
		switch el := e.(type) {
		case *ascii.Section:
			el.SecType = getSectionType(el)
			if el.SecType != ascii.MatterSectionUnknown {
				_, ok := validSectionTypes[el.SecType]
				if ok {
					lastSectionType = el.SecType
				}
			}
		}
		sections[lastSectionType] = append(sections[lastSectionType], e)
	}
	for st, elements := range sections {
		fmt.Printf("\t%s:\n", ascii.SectionTypeString(st))
		for _, e := range elements {
			switch el := e.(type) {
			case *ascii.Section:
				fmt.Printf("\t\t%s\n", el.Name)
			case *ascii.Element:
				fmt.Printf("\t\t{el} %T\n", el.Base)
			default:
				fmt.Printf("\t\t{el} %T\n", el)
			}
		}
	}
	return sections
}

func getSectionType(section *ascii.Section) ascii.MatterSection {
	name := strings.ToLower(strings.TrimSpace(section.Name))
	switch name {
	case "introduction":
		return ascii.MatterSectionIntroduction
	case "revision history":
		return ascii.MatterSectionRevisionHistory
	case "classification":
		return ascii.MatterSectionClassification
	case "cluster identifiers", "cluster id", "cluster ids":
		return ascii.MatterSectionClusterID
	case "features":
		return ascii.MatterSectionFeatures
	case "dependencies":
		return ascii.MatterSectionDependencies
	case "data types":
		return ascii.MatterSectionDataTypes
	case "status codes":
		return ascii.MatterSectionStatusCodes
	case "attributes":
		return ascii.MatterSectionAttributes
	case "commands":
		return ascii.MatterSectionCommands
	case "conditions":
		return ascii.MatterSectionConditions
	case "cluster requirements":
		return ascii.MatterSectionClusterRequirements
	case "cluster restrictions":
		return ascii.MatterSectionClusterRestrictions
	case "element requirements":
		return ascii.MatterSectionElementRequirements
	case "endpoint composition":
		return ascii.MatterSectionEndpointComposition
	default:
		return ascii.MatterSectionUnknown
	}
}
