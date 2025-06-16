package render

import (
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

func renderInternalCrossReference(cxt Target, cf *asciidoc.CrossReference) (err error) {
	id := cf.ID

	if strings.HasPrefix(id, "_") {
		return
	}

	cxt.StartBlock()
	switch cf.Format {
	case asciidoc.CrossReferenceFormatNatural:
		cxt.WriteString("<<")
		cxt.WriteString(id)
		if !cf.Set.IsWhitespace() {
			cxt.WriteString(",")
			err = Elements(cxt, "", cf.Elements()...)
			if err != nil {
				return
			}
		}
		cxt.WriteString(">>")
	case asciidoc.CrossReferenceFormatMacro:
		cxt.WriteString("xref:")
		cxt.WriteString(id)
		attributes := cf.Attributes()
		if len(attributes) == 0 {
			cxt.WriteString("[]\n")
		} else {
			err = renderAttributes(cxt, attributes, true)
			if err != nil {
				return
			}
		}
	}

	cxt.EndBlock()
	return
}

func renderDocumentCrossReference(cxt Target, dcf *asciidoc.DocumentCrossReference) (err error) {

	cxt.StartBlock()
	cxt.WriteString("xref:")
	if !dcf.ReferencePath.IsWhitespace() {
		err = Elements(cxt, "", dcf.ReferencePath...)
		if err != nil {
			return
		}
	}
	attributes := dcf.Attributes()
	if len(attributes) == 0 {
		cxt.WriteString("[]\n")
	} else {
		err = renderAttributes(cxt, attributes, true)
		if err != nil {
			return
		}
	}
	cxt.EndBlock()
	return
}
