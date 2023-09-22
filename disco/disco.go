package disco

import (
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/matter"
)

func Ball(doc *ascii.Doc) {
	docType, _ := getDocType(doc)
	discoBallDoc(doc, docType)
}

func discoBallDoc(doc *ascii.Doc, docType matter.Doc) {
	precleanStrings(doc.Elements)
	for _, e := range doc.Elements {
		switch el := e.(type) {
		case *ascii.Section:
			discoBallTopLevelSection(doc, el, docType)
			return
		}
	}
}

func discoBallTopLevelSection(doc *ascii.Doc, top *ascii.Section, docType matter.Doc) {

	for _, e := range top.Elements {
		switch el := e.(type) {
		case *ascii.Section:
			el.SecType = getSectionType(el)
			organizeSubSection(doc, docType, el)
		}
	}
	reorderTopLevelSection(top, docType)
	ensureTableOptions(top.Elements)
	postCleanUpStrings(top.Elements)
}

func organizeSubSection(doc *ascii.Doc, docType matter.Doc, section *ascii.Section) {
	switch section.SecType {
	case matter.SectionAttributes:
		switch docType {
		case matter.DocAppCluster:
			organizeAttributesSection(doc, section)

		}
	case matter.SectionClassification:
		organizeClassificationSection(doc, section)
	case matter.SectionClusterID:
		organizeClusterIDSection(doc, section)
	}
}
