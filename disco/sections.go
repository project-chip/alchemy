package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

type sectionDataType struct {
	name     string
	dataType string
	section  *ascii.Section
}

func assignTopLevelSectionTypes(top *ascii.Section) {
	for _, e := range top.Elements {
		switch el := e.(type) {
		case *ascii.Section:
			el.SecType = getSectionType(el)
		}
	}
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
	sec.Elements = newOrder
	return nil
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

func findSectionByType(top *ascii.Section, sectionType matter.Section) *ascii.Section {
	var found *ascii.Section
	find(top.Elements, func(el interface{}) bool {
		if s, ok := el.(*ascii.Section); ok {
			if s.SecType == sectionType {
				found = s
				return true
			}
		}
		return false
	})
	return found
}

func getSectionDataTypeInfo(section *ascii.Section, rows []*types.TableRow, columnMap map[matter.TableColumn]int) (sectionDataMap map[string]*sectionDataType) {
	sectionDataMap = make(map[string]*sectionDataType)
	nameIndex, ok := columnMap[matter.TableColumnName]
	if !ok {
		return
	}
	typeIndex, ok := columnMap[matter.TableColumnType]
	if !ok {
		return
	}
	for _, row := range rows {
		name := strings.TrimSpace(getCellValue(row.Cells[nameIndex]))
		nameKey := strings.ToLower(name)
		if _, ok := sectionDataMap[nameKey]; !ok {
			sectionDataMap[nameKey] = &sectionDataType{name: name, dataType: strings.TrimSpace(getCellValue(row.Cells[typeIndex]))}
		}
	}
	for _, el := range section.Elements {
		if s, ok := el.(*ascii.Section); ok {
			name := s.Name
			fmt.Printf("Attribute: %s\n", name)
			if strings.HasSuffix(name, " Attribute") {
				name = name[0 : len(name)-len(" Attribute")]
				fmt.Printf("\tAttribute: %s\n", name)
			}
			ai, ok := sectionDataMap[strings.ToLower(name)]
			if !ok {
				continue
			}
			ai.section = s
		}
	}
	return
}
