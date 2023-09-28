package disco

import (
	"fmt"
	"log/slog"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func Ball(cxt *Context, doc *ascii.Doc) error {
	docType, err := doc.DocType()
	if err != nil {
		return err
	}

	precleanStrings(doc.Elements)
	fixUnrecognizedReferences(doc)

	topLevelSection := ascii.FindFirst[*ascii.Section](doc.Elements)
	if topLevelSection == nil {
		return fmt.Errorf("missing top level section")
	}

	assignSectionTypes(topLevelSection)

	ascii.Search(topLevelSection.Elements, func(s *ascii.Section) bool {
		err := organizeSubSection(doc, docType, topLevelSection, s)
		if err != nil {
			slog.Warn("error organizing subsection", "docType", docType, "sectionType", s.SecType, "error", err)
		}
		return false
	})

	err = normalizeAnchors(doc)
	if err != nil {
		return err
	}

	return discoBallTopLevelSection(doc, topLevelSection, docType)
}

func discoBallTopLevelSection(doc *ascii.Doc, top *ascii.Section, docType matter.DocType) error {
	err := reorderTopLevelSection(top, docType)
	if err != nil {
		return err
	}
	ensureTableOptions(top.Elements)
	postCleanUpStrings(top.Elements)
	return nil
}

func organizeSubSection(doc *ascii.Doc, docType matter.DocType, top *ascii.Section, section *ascii.Section) error {
	var err error
	switch section.SecType {
	case matter.SectionAttributes:
		switch docType {
		case matter.DocTypeAppCluster:
			err = organizeAttributesSection(doc, top, section)
		}
	case matter.SectionCommands:
		err = organizeCommandsSection(doc, section)
	case matter.SectionClassification:
		err = organizeClassificationSection(doc, section)
	case matter.SectionClusterID:
		err = organizeClusterIDSection(doc, section)
	}
	return err
}
