package render

import (
	"fmt"
	"math"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

type AttributeFilter uint32

const (
	AttributeFilterNone AttributeFilter = 0
	AttributeFilterAll  AttributeFilter = math.MaxUint32
)

const (
	AttributeFilterID AttributeFilter = 1 << iota
	AttributeFilterTitle
	AttributeFilterStyle
	AttributeFilterCols
	AttributeFilterText
	AttributeFilterAlt
	AttributeFilterHeight
	AttributeFilterWidth
	AttributeFilterPDFWidth
	AttributeFilterRole
	AttributeFilterAlign
	AttributeFilterFloat
)

type attributeRenderType uint8

const (
	attributeRenderTypeBlock attributeRenderType = iota
	attributeRenderTypeInline
)

func shouldRenderAttributeType(at AttributeFilter, include AttributeFilter, exclude AttributeFilter) bool {
	return ((at & include) == at) && ((at & exclude) != at)
}

func renderAttributes(cxt Target, attributes []asciidoc.Attribute, renderType attributeRenderType) error {
	_, err := renderSelectAttributes(cxt, attributes, AttributeFilterAll, AttributeFilterNone, renderType)
	return err
}

func getAttributeType(name asciidoc.AttributeName) AttributeFilter {
	switch name {
	case asciidoc.AttributeNameTitle:
		return AttributeFilterTitle
	case asciidoc.AttributeNameID:
		return AttributeFilterID
	case asciidoc.AttributeNameColumns:
		return AttributeFilterCols
	case asciidoc.AttributeNameStyle:
		return AttributeFilterStyle
	case asciidoc.AttributeNameReferenceText:
		return AttributeFilterText
	case asciidoc.AttributeNameAlternateText:
		return AttributeFilterAlt
	case asciidoc.AttributeNameHeight:
		return AttributeFilterHeight
	case asciidoc.AttributeNameWidth:
		return AttributeFilterWidth
	case asciidoc.AttributeNamePDFWidth:
		return AttributeFilterPDFWidth
	case asciidoc.AttributeNameAlign:
		return AttributeFilterAlign
	case asciidoc.AttributeNameFloat:
		return AttributeFilterFloat
	}
	return AttributeFilterNone
}

func renderSelectAttributes(cxt Target, attributes []asciidoc.Attribute, include AttributeFilter, exclude AttributeFilter, renderType attributeRenderType) (n int, err error) {
	if len(attributes) == 0 {
		return
	}

	type attributeClass uint8

	const (
		attributeClassNone attributeClass = iota
		attributeClassTitle
		attributeClassAnchor
		attributeClassInline
	)

	var list []attributeClass
	var titleAttributes []asciidoc.Attribute
	var anchors []*asciidoc.AnchorAttribute
	var inlineAttributes []asciidoc.Attribute
	for _, a := range attributes {
		switch a := a.(type) {
		case *asciidoc.TitleAttribute:
			if len(titleAttributes) == 0 {
				list = append(list, attributeClassTitle)
			}
			titleAttributes = append(titleAttributes, a)
		case *asciidoc.AnchorAttribute:
			if len(anchors) == 0 {
				list = append(list, attributeClassAnchor)
			}
			anchors = append(anchors, a)
		case *asciidoc.NamedAttribute:
			if a.Name == asciidoc.AttributeNameTitle {
				if len(titleAttributes) == 0 {
					list = append(list, attributeClassTitle)
				}
				titleAttributes = append(titleAttributes, a)
				break
			}
			if len(inlineAttributes) == 0 {
				list = append(list, attributeClassInline)
			}
			inlineAttributes = append(inlineAttributes, a)
		case *asciidoc.PositionalAttribute, *asciidoc.TableColumnsAttribute, *asciidoc.ShorthandAttribute:
			if len(inlineAttributes) == 0 {
				list = append(list, attributeClassInline)
			}
			inlineAttributes = append(inlineAttributes, a)
		default:
			err = fmt.Errorf("unexpected attribute type: %T", a)
			return
		}
	}
	for _, al := range list {
		switch al {
		case attributeClassTitle:
			if !shouldRenderAttributeType(AttributeFilterTitle, include, exclude) {
				continue
			}
			for _, ta := range titleAttributes {
				err = renderAttributeTitle(cxt, ta.Value().(asciidoc.Elements), include, exclude)
				if err != nil {
					return
				}
			}
		case attributeClassAnchor:
			if !shouldRenderAttributeType(AttributeFilterID, include, exclude) {
				continue
			}
			for _, ta := range anchors {
				err = renderAttributeAnchor(cxt, ta, include, exclude, renderType)
				if err != nil {
					return
				}
			}
		case attributeClassInline:
			filtered := make([]asciidoc.Attribute, 0, len(inlineAttributes))
			for _, ia := range inlineAttributes {
				switch ia := ia.(type) {
				case *asciidoc.NamedAttribute:
					if af := getAttributeType(ia.Name); af != AttributeFilterNone && !shouldRenderAttributeType(af, include, exclude) {
						continue
					}
					filtered = append(filtered, ia)
				case *asciidoc.PositionalAttribute:
					if af := getAttributeType(ia.ImpliedName); af != AttributeFilterNone && !shouldRenderAttributeType(af, include, exclude) {
						continue
					}
					filtered = append(filtered, ia)
				case *asciidoc.TableColumnsAttribute:
					if shouldRenderAttributeType(AttributeFilterCols, include, exclude) {
						filtered = append(filtered, ia)
					}
				case *asciidoc.ShorthandAttribute:
					if shouldRenderAttributeType(AttributeFilterID, include, exclude) {
						filtered = append(filtered, ia)
					}
				default:
					err = fmt.Errorf("unexpected inline attribute type: %T", ia)
				}
			}
			if len(filtered) == 0 {
				continue
			}
			if renderType != attributeRenderTypeInline {
				cxt.EnsureNewLine()
			}
			cxt.FlushWrap()
			cxt.StartBlock()
			cxt.WriteString("[")
			for i, ia := range filtered {
				if i > 0 {
					cxt.WriteRune(',')
				}
				switch ia := ia.(type) {
				case *asciidoc.NamedAttribute:
					var s string
					s, err = quoteAttributeValue(ia.Value())
					if err != nil {
						return
					}
					cxt.WriteString(string(ia.Name))
					cxt.WriteString("=")
					var quoteType string
					switch ia.QuoteType() {
					case asciidoc.AttributeQuoteTypeDouble:
						quoteType = "\""
					case asciidoc.AttributeQuoteTypeSingle:
						quoteType = "'"
					}
					cxt.WriteString(quoteType)
					cxt.WriteString(s)
					cxt.WriteString(quoteType)
				case *asciidoc.PositionalAttribute:
					err = renderNakedAttributeValue(cxt, ia.Value())
				case *asciidoc.TableColumnsAttribute:
					cxt.WriteString("cols=\"")
					cxt.WriteString(ia.AsciiDocString())
					cxt.WriteString("\"")
				case *asciidoc.ShorthandAttribute:
					cxt.WriteString(ia.AsciiDocString())
				default:
					err = fmt.Errorf("unexpected inline attribute type: %T", ia)
				}
				if err != nil {
					return
				}
				n++
			}
			cxt.WriteString("]")
			if renderType != attributeRenderTypeInline {
				cxt.WriteString("\n")
			}
			cxt.EndBlock()
		default:
			err = fmt.Errorf("unexpected attribute list element: %T", al)
			return
		}

	}
	return
}

func renderAttributeAnchor(cxt Target, anchor *asciidoc.AnchorAttribute, include AttributeFilter, exclude AttributeFilter, renderType attributeRenderType) (err error) {
	id := anchor.ID
	if len(id) > 0 && shouldRenderAttributeType(AttributeFilterID, include, exclude) {
		if renderType != attributeRenderTypeInline {
			cxt.EnsureNewLine()
		}
		cxt.FlushWrap()
		cxt.StartBlock()
		cxt.WriteString("[[")
		err = Elements(cxt, "", id...)
		if err != nil {
			return
		}
		if len(anchor.Label) > 0 {
			label := cxt.Subtarget()
			err = Elements(label, "", anchor.Label...)
			if err != nil {
				return
			}
			lbl := label.String()
			if len(lbl) > 0 {
				cxt.WriteString(",")
				cxt.WriteString(lbl)
			}
		}
		cxt.WriteString("]]")
		if renderType != attributeRenderTypeInline {
			cxt.WriteRune('\n')
		}
		cxt.EndBlock()
	}
	return
}

func renderAttributeTitle(cxt Target, title asciidoc.Elements, include AttributeFilter, exclude AttributeFilter) (err error) {
	if len(title) > 0 && shouldRenderAttributeType(AttributeFilterTitle, include, exclude) {
		cxt.EnsureNewLine()
		cxt.FlushWrap()
		cxt.StartBlock()
		cxt.WriteRune('.')
		err = Elements(cxt, "", title...)
		if err != nil {
			return
		}
		cxt.EnsureNewLine()
		cxt.EndBlock()
	}
	return
}

func quoteAttributeValue(val any) (string, error) {
	switch val := val.(type) {
	case string:
		return escapeQuotes(val), nil
	case *asciidoc.String:
		return escapeQuotes(val.Value), nil
	case asciidoc.AttributeReference:
		return "{" + val.Name() + "}", nil
	case asciidoc.Elements:
		var sb strings.Builder
		for _, a := range val {
			s, err := quoteAttributeValue(a)
			if err != nil {
				return "", err
			}
			sb.WriteString(s)
		}
		return sb.String(), nil
	default:
		return "", fmt.Errorf("unexpected attribute value type: %T", val)
	}
}

func renderNakedAttributeValue(cxt Target, val any) (err error) {
	switch val := val.(type) {
	case *asciidoc.String:
		cxt.WriteString(escapeQuotes(val.Value))
	case asciidoc.AttributeReference:
		cxt.WriteRune('{')
		cxt.WriteString(val.Name())
		cxt.WriteRune('}')
	case asciidoc.Elements:
		for _, a := range val {
			err = renderNakedAttributeValue(cxt, a)
			if err != nil {
				return
			}
		}
	default:
		err = fmt.Errorf("unexpected attribute value type: %T", val)
	}
	return
}

func escapeQuotes(s string) string {
	return strings.ReplaceAll(s, "\"", "\\\"")
}

func renderAttributeEntry(cxt Target, ad *asciidoc.AttributeEntry) (err error) {
	cxt.FlushWrap()
	cxt.DisableWrap()
	switch ad.Name {
	case "authors":
		/*if authors, ok := ad.Value().(asciidoc.DocumentAuthors); ok {
			for _, author := range authors {
				if len(author.Email) > 0 {
					cxt.WriteString(author.Email)
					cxt.WriteString(" ")
				}
				if author.DocumentAuthorFullName != nil {
					cxt.WriteString(author.DocumentAuthorFullName.FullName())
				}
				cxt.WriteRune('\n')
			}
		}*/
	default:
		cxt.WriteRune(':')
		cxt.WriteString(string(ad.Name))
		cxt.WriteString(": ")
		err = Elements(cxt, "", ad.Children()...)

		cxt.WriteRune('\n')
	}
	cxt.EnableWrap()
	return
}

func renderAttributeReset(cxt Target, ar *asciidoc.AttributeReset) {
	cxt.FlushWrap()
	cxt.DisableWrap()
	cxt.WriteRune(':')
	cxt.WriteString(string(ar.Name))
	cxt.WriteString("!:\n")
	cxt.EnableWrap()
}
