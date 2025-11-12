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

func reorderSection(dc *discoContext, doc *asciidoc.Document, sec *asciidoc.Section, sectionOrder []matter.Section) error {
	validSectionTypes := make(map[matter.Section]struct{}, len(sectionOrder)+1)
	for _, st := range sectionOrder {
		validSectionTypes[st] = struct{}{}
	}
	sections := divyUpSection(dc, doc, sec, validSectionTypes)

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

func divyUpSection(dc *discoContext, doc *asciidoc.Document, sec *asciidoc.Section, validSectionTypes map[matter.Section]struct{}) map[matter.Section]asciidoc.Elements {
	sections := make(map[matter.Section]asciidoc.Elements)
	lastSectionType := matter.SectionPrefix
	for _, e := range sec.Children() {
		switch el := e.(type) {
		case *asciidoc.Section:
			st := dc.library.SectionType(el)
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

func setSectionTitle(dc *discoContext, doc *asciidoc.Document, sec *asciidoc.Section, title string) bool {
	if len(sec.Title) != 1 {
		return false
	}
	switch sec.Title[0].(type) {
	case *asciidoc.String:
		sec.Title[0] = asciidoc.NewString(title)
		dc.library.SetSectionName(sec, title)
		return true
	}
	return false
}

func (b *Baller) appendSubsectionTypes(cxt *discoContext, section *asciidoc.Section, columnMap spec.ColumnIndex, rows []*asciidoc.TableRow) {
	if cxt.errata.IgnoreSection(cxt.library.SectionName(section), errata.DiscoPurposeDataTypeAppendSuffix) {
		return
	}
	var subsectionSuffix string
	var subsectionType matter.Section
	switch cxt.library.SectionType(section) {
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
			name, err := spec.RenderTableCell(cxt.library, row.TableCells()[nameIndex])
			if err != nil {
				slog.Debug("could not get cell value for subsection", "err", err)
				continue
			}
			subSectionNames[name] = struct{}{}
		}
		suffix := " " + subsectionSuffix
		for ss := range parse.FindAll[*asciidoc.Section](cxt.doc, asciidoc.RawReader, section) {
			name := text.TrimCaseInsensitiveSuffix(cxt.library.SectionName(ss), suffix)
			if _, ok := subSectionNames[name]; !ok {
				continue
			}
			if !b.options.AppendSubsectionTypes {
				continue
			}
			if cxt.library.SectionType(ss) == matter.SectionUnknown {
				cxt.library.SetSectionType(ss, subsectionType)
			}
			name = cxt.library.SectionName(ss)
			if !strings.HasSuffix(strings.ToLower(name), strings.ToLower(suffix)) {
				setSectionTitle(cxt, cxt.doc, ss, name+suffix)
			}
		}
	}
}
