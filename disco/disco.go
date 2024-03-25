package disco

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/internal/parse"
	"github.com/hasty/alchemy/matter"
)

type Ball struct {
	doc *ascii.Doc

	options options
}

func NewBall(doc *ascii.Doc) *Ball {
	return &Ball{
		doc:     doc,
		options: defaultOptions,
	}
}

func (b *Ball) disco(cxt context.Context) error {

	dc := newContext(cxt)
	doc := b.doc

	precleanStrings(doc.Elements)

	for _, top := range parse.Skim[*ascii.Section](doc.Elements) {
		err := ascii.AssignSectionTypes(doc, top)
		if err != nil {
			return err
		}
	}

	docType, err := doc.DocType()
	if err != nil {
		return fmt.Errorf("error assigning section types in %s: %w", doc.Path, err)
	}

	topLevelSection := parse.FindFirst[*ascii.Section](doc.Elements)
	if topLevelSection == nil {
		return EmptyDocError
	}

	dp, err := b.parseDoc(doc, docType, topLevelSection)
	if err != nil {
		return fmt.Errorf("error pre-parsing for disco ball in %s: %w", doc.Path, err)
	}

	getExistingDataTypes(dc, dp)

	err = b.getPotentialDataTypes(dc, dp)
	if err != nil {
		return fmt.Errorf("error getting potential data types in %s: %w", doc.Path, err)
	}

	var promotedDataTypes bool
	promotedDataTypes, err = b.promoteDataTypes(dc, topLevelSection)
	if err != nil {
		return fmt.Errorf("error promoting data types in %s: %w", doc.Path, err)
	}
	if promotedDataTypes {
		dp, err = b.parseDoc(doc, docType, topLevelSection)
		if err != nil {
			return fmt.Errorf("error re-parsing for disco ball in %s: %w", doc.Path, err)
		}
	}

	err = b.organizeSubSections(dc, dp)
	if err != nil {
		return fmt.Errorf("error organizing subsections in %s: %w", doc.Path, err)
	}

	err = b.discoBallTopLevelSection(doc, topLevelSection, docType)

	if err != nil {
		return fmt.Errorf("error disco balling top level section in %s: %w", doc.Path, err)
	}
	return b.normalizeAnchors(doc)
}

func (b *Ball) discoBallTopLevelSection(doc *ascii.Doc, top *ascii.Section, docType matter.DocType) error {
	if b.options.reorderSections {
		sectionOrder, ok := matter.TopLevelSectionOrders[docType]
		if !ok {
			slog.Debug("could not determine section order", "docType", docType)

		} else {
			err := reorderSection(top, sectionOrder)
			if err != nil {
				return fmt.Errorf("error reordering sections in %s: %w", doc.Path, err)
			}
		}
		dataTypesSection := ascii.FindSectionByType(top, matter.SectionDataTypes)
		if dataTypesSection != nil {
			err := reorderSection(dataTypesSection, matter.DataTypeSectionOrder)
			if err != nil {
				return fmt.Errorf("error reordering data types section in %s: %w", doc.Path, err)
			}
		}
	}
	b.ensureTableOptions(top.Elements)
	b.postCleanUpStrings(top.Elements)
	return nil
}
