package disco

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
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

func (b *Ball) Run(cxt context.Context) error {

	dc := newContext(cxt)
	doc := b.doc
	docType, err := doc.DocType()
	if err != nil {
		return err
	}

	precleanStrings(doc.Elements)
	ascii.PatchUnrecognizedReferences(doc)

	for _, top := range parse.Skim[*ascii.Section](doc.Elements) {
		ascii.AssignSectionTypes(docType, top)
	}

	topLevelSection := parse.FindFirst[*ascii.Section](doc.Elements)
	if topLevelSection == nil {
		return fmt.Errorf("missing top level section")
	}

	getExistingDataTypes(dc, topLevelSection)

	for _, s := range parse.FindAll[*ascii.Section](topLevelSection.Elements) {
		err := b.organizeSubSection(dc, doc, docType, topLevelSection, s)
		if err != nil {
			return err
		}
	}

	err = b.promoteDataTypes(dc, topLevelSection)
	if err != nil {
		return err
	}

	err = b.discoBallTopLevelSection(doc, topLevelSection, docType)

	if err != nil {
		return err
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
				return err
			}
		}
		dataTypesSection := ascii.FindSectionByType(top, matter.SectionDataTypes)
		if dataTypesSection != nil {
			err := reorderSection(dataTypesSection, matter.DataTypeSectionOrder)
			if err != nil {
				return err
			}
		}
	}
	b.ensureTableOptions(top.Elements)
	b.postCleanUpStrings(top.Elements)
	return nil
}

func (b *Ball) organizeSubSection(cxt *discoContext, doc *ascii.Doc, docType matter.DocType, top *ascii.Section, section *ascii.Section) error {
	var err error
	switch section.SecType {
	case matter.SectionAttributes:
		switch docType {
		case matter.DocTypeAppCluster:
			err = b.organizeAttributesSection(cxt, doc, top, section)
		}
	case matter.SectionCommands:
		err = b.organizeCommandsSection(cxt, doc, section)
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
		err = b.organizeEventsSection(cxt, doc, section)
	}
	return err
}
