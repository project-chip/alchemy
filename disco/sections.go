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
	return
}

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
			name, err := ascii.RenderTableCell(row.Cells[nameIndex])
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
