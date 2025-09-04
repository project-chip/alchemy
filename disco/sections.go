package disco

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal/text"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func (b *Baller) organizeSubSections(dc *discoContext) (err error) {
	organizers := []func(dc *discoContext) error{
		b.organizeAttributesSection,
		b.organizeClassificationSection,
		b.organizeClusterIDSection,
		b.organizeFeaturesSection,
		b.organizeBitmapSections,
		b.organizeEnumSections,
		b.organizeStructSections,
		b.organizeCommandsSection,
		b.organizeEventsSection,
		b.organizeDeviceIDSection,
		b.organizeClusterRequirementsSection,
		b.organizeElementRequirementsSection,
		b.organizeComposedDeviceClusterRequirementsSection,
		b.organizeComposedDeviceElementRequirementsSection,
		b.organizeConditionRequirementsSection,
	}
	for _, organizer := range organizers {
		err = organizer(dc)
		if err != nil {
			return
		}
	}
	return
}

func reorderSection(doc *spec.Doc, sec *asciidoc.Section, sectionOrder []matter.Section) error {
	validSectionTypes := make(map[matter.Section]struct{}, len(sectionOrder)+1)
	for _, st := range sectionOrder {
		validSectionTypes[st] = struct{}{}
	}
	sections := divyUpSection(doc, sec, validSectionTypes)

	newOrder := make(asciidoc.Elements, 0, len(sec.Children()))
	for _, st := range sectionOrder {
		if els, ok := sections[st]; ok {

			newOrder = append(newOrder, els...)
			delete(sections, st)
		}
	}
	if len(sections) > 0 {
		return fmt.Errorf("non-empty section list after reordering")
	}

	sec.SetChildren(newOrder)
	return nil
}

func divyUpSection(doc *spec.Doc, sec *asciidoc.Section, validSectionTypes map[matter.Section]struct{}) map[matter.Section]asciidoc.Elements {
	sections := make(map[matter.Section]asciidoc.Elements)
	lastSectionType := matter.SectionPrefix
	for _, e := range sec.Children() {
		switch el := e.(type) {
		case *asciidoc.Section:
			st := doc.SectionType(el)
			if st != matter.SectionUnknown {
				_, ok := validSectionTypes[st]
				if ok {
					lastSectionType = st
				}
			}
		}
		sections[lastSectionType] = append(sections[lastSectionType], e)
	}
	return sections
}

func setSectionTitle(doc *spec.Doc, sec *asciidoc.Section, title string) {
	for i, e := range sec.Title {
		switch e.(type) {
		case *asciidoc.String:
			sec.Title[i] = asciidoc.NewString(title)
			doc.SetSectionName(sec, title)
		}
	}
}

func (b *Baller) appendSubsectionTypes(cxt *discoContext, section *asciidoc.Section, columnMap spec.ColumnIndex, rows []*asciidoc.TableRow) {
	if cxt.errata.IgnoreSection(cxt.doc.SectionName(section), errata.DiscoPurposeDataTypeAppendSuffix) {
		return
	}
	var subsectionSuffix string
	var subsectionType matter.Section
	switch cxt.doc.SectionType(section) {
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
		suffix := " " + subsectionSuffix
		for ss := range parse.FindAll[*asciidoc.Section](asciidoc.RawReader, section) {
			name := text.TrimCaseInsensitiveSuffix(cxt.doc.SectionName(ss), suffix)
			if _, ok := subSectionNames[name]; !ok {
				continue
			}
			if !b.options.AppendSubsectionTypes {
				continue
			}
			if cxt.doc.SectionType(ss) == matter.SectionUnknown {
				cxt.doc.SetSectionType(ss, subsectionType)
			}
			name = cxt.doc.SectionName(ss)
			if !strings.HasSuffix(strings.ToLower(name), strings.ToLower(suffix)) {
				setSectionTitle(cxt.doc, ss, name+suffix)
			}
		}
	}
}
