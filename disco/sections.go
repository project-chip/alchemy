package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

type sectionOrganizer func(dc *discoContext, dp *docParse) error

func (b *Ball) organizeSubSections(dc *discoContext, dp *docParse) (err error) {
	organizers := []sectionOrganizer{
		b.organizeAttributesSection,
		b.organizeClassificationSection,
		b.organizeClusterIDSection,
		b.organizeBitmapSections,
		b.organizeEnumSections,
		b.organizeStructSections,
		b.organizeCommandsSection,
		b.organizeEventsSection,
	}
	for _, organizer := range organizers {
		err = organizer(dc, dp)
		if err != nil {
			return
		}
	}
	/*for _, section := range parse.FindAll[*ascii.Section](topLevelSection.Elements) {
		var err error
		switch section.SecType {
		case matter.SectionAttributes:
			switch docType {
			case matter.DocTypeCluster:
				err = b.organizeAttributesSection(dc, doc, topLevelSection, section)
			}
		case matter.SectionCommands:
			err = b.organizeCommandsSection(dc, doc, section)
		case matter.SectionClassification:
			err = b.organizeClassificationSection(doc, section)
		case matter.SectionClusterID:
			err = b.organizeClusterIDSection(doc, section)
		case matter.SectionDataTypeBitmap:
			err = b.organizeBitmapSection(doc, section)
		case matter.SectionDataTypeEnum:
			err = b.organizeEnumSection(doc, section)
		case matter.SectionDataTypeStruct:
			err = b.organizeStructSection(doc, section)
		case matter.SectionEvents:
			err = b.organizeEventsSection(dc, doc, section)
		}
		if err != nil {
			return fmt.Errorf("error organizing subsections of section %s in %s: %w", section.Name, doc.Path, err)
		}
	}*/
	return
}

/*
func (b *Ball) organizeSubSectionsOld(dc *discoContext, doc *ascii.Doc, docType matter.DocType, topLevelSection *ascii.Section) error {
	for _, section := range parse.FindAll[*ascii.Section](topLevelSection.Elements) {
		var err error
		switch section.SecType {
		case matter.SectionAttributes:
			switch docType {
			case matter.DocTypeCluster:
				//err = b.organizeAttributesSection(dc, doc, topLevelSection, section)
			}
		case matter.SectionCommands:
			err = b.organizeCommandsSection(dc, doc, section)
		case matter.SectionClassification:
			err = b.organizeClassificationSection(doc, section)
		case matter.SectionClusterID:
			err = b.organizeClusterIDSection(doc, section)
		case matter.SectionDataTypeBitmap:
			err = b.organizeBitmapSection(doc, section)
		case matter.SectionDataTypeEnum:
			err = b.organizeEnumSection(doc, section)
		case matter.SectionDataTypeStruct:
			err = b.organizeStructSection(doc, section)
		case matter.SectionEvents:
			err = b.organizeEventsSection(dc, doc, section)
		}
		if err != nil {
			return fmt.Errorf("error organizing subsections of section %s in %s: %w", section.Name, doc.Path, err)
		}
	}
	return nil
}*/

func reorderSection(sec *ascii.Section, sectionOrder []matter.Section) error {
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

	return sec.SetElements(newOrder)
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
			sec.Name = title
		}
	}
}

func (b *Ball) appendSubsectionTypes(section *ascii.Section, columnMap ascii.ColumnIndex, rows []*types.TableRow) {
	var subsectionSuffix string
	var subsectionType matter.Section
	switch section.SecType {
	case matter.SectionDataTypeBitmap:
		subsectionSuffix = "Bit"
		subsectionType = matter.SectionBit
	case matter.SectionDataTypeEnum:
		subsectionSuffix = "Value"
		subsectionType = matter.SectionValue
	case matter.SectionCommand, matter.SectionEvent, matter.SectionDataTypeStruct:
		subsectionSuffix = "Field"
		subsectionType = matter.SectionField
	}
	nameIndex, ok := columnMap[matter.TableColumnName]
	if ok {

		subSectionNames := make(map[string]struct{}, len(rows))
		for _, row := range rows {
			name, err := ascii.GetTableCellValue(row.Cells[nameIndex])
			if err != nil {
				slog.Debug("could not get cell value for subsection", "err", err)
				continue
			}
			subSectionNames[name] = struct{}{}
		}
		subSections := parse.FindAll[*ascii.Section](section.Elements)
		suffix := " " + subsectionSuffix
		for _, ss := range subSections {
			name := strings.TrimSuffix(ss.Name, suffix)
			if _, ok := subSectionNames[name]; !ok {
				continue
			}
			if ss.SecType == matter.SectionUnknown {
				ss.SecType = subsectionType
			}
			if !b.options.appendSubsectionTypes {
				continue
			}
			if !strings.HasSuffix(strings.ToLower(ss.Name), strings.ToLower(suffix)) {
				setSectionTitle(ss, ss.Name+suffix)
			}
		}
	}
}
