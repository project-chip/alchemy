package disco

import (
	"log/slog"

	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func Ball(doc *ascii.Doc) error {
	return discoBallDoc(doc)
}

func discoBallDoc(doc *ascii.Doc) error {
	docType, err := doc.DocType()
	if err != nil {
		return err
	}
	precleanStrings(doc.Elements)
	fixUnrecognizedReferences(doc)
	err = normalizeReferences(doc)
	if err != nil {
		return err
	}
	find(doc.Elements, func(s *ascii.Section) bool {
		discoBallTopLevelSection(doc, s, docType)
		return true
	})
	return nil
}

func discoBallTopLevelSection(doc *ascii.Doc, top *ascii.Section, docType matter.DocType) {
	assignTopLevelSectionTypes(top)
	find(top.Elements, func(s *ascii.Section) bool {
		err := organizeSubSection(doc, docType, top, s)
		if err != nil {
			slog.Warn("error organizing subsection", "docType", docType, "sectionType", s.SecType, "error", err)
		}
		return false
	})
	reorderTopLevelSection(top, docType)
	ensureTableOptions(top.Elements)
	postCleanUpStrings(top.Elements)
}

func organizeSubSection(doc *ascii.Doc, docType matter.DocType, top *ascii.Section, section *ascii.Section) error {
	var err error
	switch section.SecType {
	case matter.SectionAttributes:
		switch docType {
		case matter.DocTypeAppCluster:
			err = organizeAttributesSection(doc, top, section)
		}
	case matter.SectionClassification:
		err = organizeClassificationSection(doc, section)
	case matter.SectionClusterID:
		err = organizeClusterIDSection(doc, section)
	}
	return err
}
