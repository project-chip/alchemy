package disco

import (
	"github.com/hasty/matterfmt/ascii"
)

func Ball(doc *ascii.Doc) {
	docType, _ := getDocType(doc)
	reorderDoc(doc, docType)
}

func reorderDoc(doc *ascii.Doc, docType MatterDoc) {
	for _, e := range doc.Elements {
		switch el := e.(type) {
		case *ascii.Section:

			reorderTopLevelSection(el, docType)
			return
		}
	}
}
