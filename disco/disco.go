package disco

import (
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
	normalizeReferences(doc)
	find(doc.Elements, func(s *ascii.Section) bool {
		discoBallTopLevelSection(doc, s, docType)
		return true
	})
	return nil
}

func discoBallTopLevelSection(doc *ascii.Doc, top *ascii.Section, docType matter.DocType) {
	assignTopLevelSectionTypes(top)
	find(top.Elements, func(s *ascii.Section) bool {
		organizeSubSection(doc, docType, top, s)
		return false
	})
	reorderTopLevelSection(top, docType)
	ensureTableOptions(top.Elements)
	postCleanUpStrings(top.Elements)
}

func organizeSubSection(doc *ascii.Doc, docType matter.DocType, top *ascii.Section, section *ascii.Section) {
	switch section.SecType {
	case matter.SectionAttributes:
		switch docType {
		case matter.DocTypeAppCluster:
			organizeAttributesSection(doc, top, section)
		}
	case matter.SectionClassification:
		organizeClassificationSection(doc, section)
	case matter.SectionClusterID:
		organizeClusterIDSection(doc, section)
	}
}
