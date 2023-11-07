package disco

import (
	"context"
	"fmt"

	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
	"github.com/hasty/alchemy/parse"
)

type Ball struct {
	doc *ascii.Doc

	ShouldLinkAttributes bool
}

func NewBall(doc *ascii.Doc) *Ball {
	return &Ball{
		doc: doc,
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

	topLevelSection := parse.FindFirst[*ascii.Section](doc.Elements)
	if topLevelSection == nil {
		return fmt.Errorf("missing top level section")
	}

	ascii.AssignSectionTypes(docType, topLevelSection)

	getExistingDataTypes(dc, topLevelSection)

	for _, s := range parse.FindAll[*ascii.Section](topLevelSection.Elements) {
		err := b.organizeSubSection(dc, doc, docType, topLevelSection, s)
		if err != nil {
			return err
		}
	}

	err = promoteDataTypes(dc, topLevelSection)
	if err != nil {
		return err
	}

	err = b.normalizeAnchors(doc)
	if err != nil {
		return err
	}

	return b.discoBallTopLevelSection(doc, topLevelSection, docType)
}

func (b *Ball) discoBallTopLevelSection(doc *ascii.Doc, top *ascii.Section, docType matter.DocType) error {
	err := reorderTopLevelSection(top, docType)
	if err != nil {
		return err
	}
	ensureTableOptions(top.Elements)
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
	}
	return err
}
