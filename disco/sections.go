package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

type sectionOrganizer func(dc *discoContext, dp *docParse) error

func (b *Ball) organizeSubSections(dc *discoContext, dp *docParse) (err error) {
	organizers := []sectionOrganizer{
		b.organizeAttributesSection,
		b.organizeClassificationSection,
		b.organizeClusterIDSection,
		b.organizeFeaturesSection,
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

func reorderSection(sec *spec.Section, sectionOrder []matter.Section) error {
	validSectionTypes := make(map[matter.Section]struct{}, len(sectionOrder)+1)
	for _, st := range sectionOrder {
		validSectionTypes[st] = struct{}{}
	}
	sections := divyUpSection(sec, validSectionTypes)

	newOrder := make(asciidoc.Set, 0, len(sec.Elements()))
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

func divyUpSection(sec *spec.Section, validSectionTypes map[matter.Section]struct{}) map[matter.Section]asciidoc.Set {
	sections := make(map[matter.Section]asciidoc.Set)
	lastSectionType := matter.SectionPrefix
	for _, e := range sec.Elements() {
		switch el := e.(type) {
		case *spec.Section:
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

func setSectionTitle(sec *spec.Section, title string) {
	for i, e := range sec.Base.Title {
		switch e.(type) {
		case *asciidoc.String:
			sec.Base.Title[i] = asciidoc.NewString(title)
			sec.Name = title
		}
	}
}

func (b *Ball) appendSubsectionTypes(section *spec.Section, columnMap spec.ColumnIndex, rows []*asciidoc.TableRow) {
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
			name, err := spec.RenderTableCell(row.TableCells()[nameIndex])
			if err != nil {
				slog.Debug("could not get cell value for subsection", "err", err)
				continue
			}
			subSectionNames[name] = struct{}{}
		}
		subSections := parse.FindAll[*spec.Section](section.Elements())
		suffix := " " + subsectionSuffix
		for _, ss := range subSections {
			name := text.TrimCaseInsensitiveSuffix(ss.Name, suffix)
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
