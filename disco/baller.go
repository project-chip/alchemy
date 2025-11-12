package disco

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

type Baller struct {
	spec    *spec.Specification
	options DiscoOptions
}

func NewBaller(specification *spec.Specification, discoOptions DiscoOptions) *Baller {
	b := &Baller{
		spec:    specification,
		options: discoOptions,
	}
	return b
}

func (r Baller) Name() string {
	return "Disco balling"
}

func (r Baller) Process(cxt context.Context, input *pipeline.Data[*asciidoc.Document], index int32, total int32) (outputs []*pipeline.Data[*asciidoc.Document], extras []*pipeline.Data[*asciidoc.Document], err error) {
	err = r.disco(cxt, input.Content)
	if err != nil {
		if err == ErrEmptyDoc {
			err = nil
			return
		}
		slog.Warn("Error disco balling document", "path", input.Path, "error", err)
		err = nil
		return
	}
	outputs = append(outputs, pipeline.NewData(input.Path, input.Content))
	return
}

func (b *Baller) disco(cxt context.Context, doc *asciidoc.Document) error {

	var ok bool
	var library *spec.Library
	library, ok = b.spec.LibraryForDocument(doc)
	if !ok {
		return fmt.Errorf("unable to find library for doc %s", doc.Path.Relative)
	}

	dc := newContext(cxt, library, doc)

	precleanStrings(doc)

	docType, err := dc.library.DocType(doc)
	if err != nil {
		return fmt.Errorf("error assigning section types in %s: %w", doc.Path, err)
	}

	topLevelSection := parse.FindFirst[*asciidoc.Section](doc, asciidoc.RawReader, doc)
	if topLevelSection == nil {
		return ErrEmptyDoc
	}

	dc.parsed, err = b.parseDoc(dc.library, doc, docType, topLevelSection)
	if err != nil {
		return fmt.Errorf("error pre-parsing for disco ball in %s: %w", doc.Path, err)
	}

	getExistingDataTypes(dc)

	err = b.getPotentialDataTypes(dc)
	if err != nil {
		return fmt.Errorf("error getting potential data types in %s: %w", doc.Path, err)
	}

	var promotedDataTypes bool
	promotedDataTypes, err = b.promoteDataTypes(dc, topLevelSection)
	if err != nil {
		return fmt.Errorf("error promoting data types in %s: %w", doc.Path, err)
	}
	if promotedDataTypes {
		dc.parsed, err = b.parseDoc(dc.library, doc, docType, topLevelSection)
		if err != nil {
			return fmt.Errorf("error re-parsing for disco ball in %s: %w", doc.Path, err)
		}
	}

	err = b.organizeSubSections(dc)
	if err != nil {
		return fmt.Errorf("error organizing subsections in %s: %w", doc.Path, err)
	}

	err = b.discoBallTopLevelSection(dc, topLevelSection, docType)

	if err != nil {
		return fmt.Errorf("error disco balling top level section in %s: %w", doc.Path, err)
	}

	if b.options.DisambiguateConformanceChoice {
		err = disambiguateConformance(dc)
		if err != nil {
			return fmt.Errorf("error disambiguating conformance in %s: %w", doc.Path, err)
		}
	}

	if b.options.ReorderSections {
		//	slog.Info("Renaming sections!")
		parse.Search(doc, asciidoc.RawReader, topLevelSection, topLevelSection.Children(), func(doc *asciidoc.Document, section *asciidoc.Section, parent asciidoc.ParentElement, index int) parse.SearchShould {
			sectionType := dc.library.SectionType(section)
			if sectionType != matter.SectionUnknown {
				canonicalSectionName := matter.CanonicalSectionTypeName(sectionType)
				//	slog.Info("Renaming section!", "from", doc.SectionName(section), "to", canonicalSectionName)
				if canonicalSectionName != "" && canonicalSectionName != dc.library.SectionName(section) {
					slog.Info("Renaming section!", "new name", canonicalSectionName, log.Path("source", section))
					if setSectionTitle(dc, doc, section, canonicalSectionName) {
						slog.Info("Renamed section!", "new name", canonicalSectionName)

					}
				}
			}
			return parse.SearchShouldContinue
		})

	}
	return nil
}

func (b *Baller) discoBallTopLevelSection(dc *discoContext, top *asciidoc.Section, docType matter.DocType) error {
	if b.options.ReorderSections {
		sectionOrder, ok := matter.TopLevelSectionOrders[docType]
		if !ok {
			slog.Debug("could not determine section order", slog.String("path", dc.doc.Path.Relative), slog.String("docType", docType.String()))

		} else {
			err := reorderSection(dc, dc.doc, top, sectionOrder)
			if err != nil {
				return fmt.Errorf("error reordering sections in %s: %w", dc.doc.Path, err)
			}
		}
		dataTypesSection := dc.library.FindSectionByType(asciidoc.RawReader, dc.doc, top, matter.SectionDataTypes)
		if dataTypesSection != nil {
			err := reorderSection(dc, dc.doc, dataTypesSection, matter.DataTypeSectionOrder)
			if err != nil {
				return fmt.Errorf("error reordering data types section in %s: %w", dc.doc.Path, err)
			}
		}
		deviceRequirementsSection := dc.library.FindSectionByType(dc.library, dc.doc, top, matter.SectionDeviceTypeRequirements)
		if deviceRequirementsSection != nil {
			err := reorderSection(dc, dc.doc, deviceRequirementsSection, matter.DeviceRequirementsSectionOrder)
			if err != nil {
				return fmt.Errorf("error reordering device requirements section in %s: %w", dc.doc.Path.Relative, err)
			}
		}
	}
	b.ensureTableOptions(dc.doc, top)
	b.postCleanUpStrings(dc.doc, top)
	return nil
}
