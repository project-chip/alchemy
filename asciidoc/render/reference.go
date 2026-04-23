package render

import (
	"github.com/project-chip/alchemy/asciidoc"
)

func renderInternalCrossReference(cxt Target, cf *asciidoc.CrossReference) (err error) {
	id := cf.ID

	cxt.StartBlock()
	switch cf.Format {
	case asciidoc.CrossReferenceFormatNatural:
		cxt.WriteString("<<")
		err = Elements(cxt, "", id...)
		if err != nil {
			return
		}
		if !cf.Elements.IsWhitespace() {
			cxt.WriteString(",")
			err = Elements(cxt, "", cf.Children()...)
			if err != nil {
				return
			}
		}
		cxt.WriteString(">>")
	case asciidoc.CrossReferenceFormatMacro:
		cxt.WriteString("xref:")
		err = Elements(cxt, "", id...)
		if err != nil {
			return
		}

		cxt.WriteRune('[')
		err = Elements(cxt, "", cf.Children()...)
		if err != nil {
			return
		}
		cxt.WriteRune(']')
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
		err = renderAttributes(cxt, attributes, attributeRenderTypeInline)
		if err != nil {
			return
		}
	}
	cxt.EndBlock()
	return
}
