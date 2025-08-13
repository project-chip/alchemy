package disco

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

type Baller struct {
	options DiscoOptions
}

func NewBaller(discoOptions DiscoOptions) *Baller {
	b := &Baller{
		options: discoOptions,
	}
	return b
}

func (r Baller) Name() string {
	return "Disco balling"
}

func (r Baller) Process(cxt context.Context, input *pipeline.Data[*spec.Doc], index int32, total int32) (outputs []*pipeline.Data[*spec.Doc], extras []*pipeline.Data[*spec.Doc], err error) {
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

func (b *Baller) disco(cxt context.Context, doc *spec.Doc) error {

	dc := newContext(cxt, doc)

	precleanStrings(doc)

	for top := range parse.Skim[*asciidoc.Section](doc.Reader(), doc, doc.Children()) {
		err := spec.AssignSectionTypes(doc, top)
		if err != nil {
			return err
		}
	}

	docType, err := doc.DocType()
	if err != nil {
		return fmt.Errorf("error assigning section types in %s: %w", doc.Path, err)
	}

	topLevelSection := parse.FindFirst[*asciidoc.Section](doc.Reader(), doc)
	if topLevelSection == nil {
		return ErrEmptyDoc
	}

	dc.parsed, err = b.parseDoc(doc, docType, topLevelSection)
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
		dc.parsed, err = b.parseDoc(doc, docType, topLevelSection)
		if err != nil {
			return fmt.Errorf("error re-parsing for disco ball in %s: %w", doc.Path, err)
		}
	}

	err = b.organizeSubSections(dc)
	if err != nil {
		return fmt.Errorf("error organizing subsections in %s: %w", doc.Path, err)
	}

	err = b.discoBallTopLevelSection(doc, topLevelSection, docType)

	if err != nil {
		return fmt.Errorf("error disco balling top level section in %s: %w", doc.Path, err)
	}

	if b.options.DisambiguateConformanceChoice {
		err = disambiguateConformance(dc)
		if err != nil {
			return fmt.Errorf("error disambiguating conformance in %s: %w", doc.Path, err)
		}
	}
	return nil
}

func (b *Baller) discoBallTopLevelSection(doc *spec.Doc, top *asciidoc.Section, docType matter.DocType) error {
	if b.options.ReorderSections {
		sectionOrder, ok := matter.TopLevelSectionOrders[docType]
		if !ok {
			slog.Debug("could not determine section order", slog.String("path", doc.Path.Relative), slog.String("docType", docType.String()))

		} else {
			err := reorderSection(doc, top, sectionOrder)
			if err != nil {
				return fmt.Errorf("error reordering sections in %s: %w", doc.Path, err)
			}
		}
		dataTypesSection := spec.FindSectionByType(doc, top, matter.SectionDataTypes)
		if dataTypesSection != nil {
			err := reorderSection(doc, dataTypesSection, matter.DataTypeSectionOrder)
			if err != nil {
				return fmt.Errorf("error reordering data types section in %s: %w", doc.Path, err)
			}
		}
	}
	b.ensureTableOptions(doc, top)
	b.postCleanUpStrings(doc, top)
	return nil
}
