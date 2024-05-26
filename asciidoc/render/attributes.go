package render

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/hasty/alchemy/asciidoc"
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

func shouldRenderAttributeType(at AttributeFilter, include AttributeFilter, exclude AttributeFilter) bool {
	return ((at & include) == at) && ((at & exclude) != at)
}

func renderAttributes(cxt *Context, el any, attributes []asciidoc.Attribute, inline bool) error {
	return renderSelectAttributes(cxt, el, attributes, AttributeFilterAll, AttributeFilterNone, inline)
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

func renderSelectAttributes(cxt *Context, el any, attributes []asciidoc.Attribute, include AttributeFilter, exclude AttributeFilter, inline bool) (err error) {
	if len(attributes) == 0 {
		return
	}

	type attributeClass uint32

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
		case *asciidoc.PositionalAttribute, *asciidoc.TableColumnsAttribute:
			if len(inlineAttributes) == 0 {
				list = append(list, attributeClassInline)
			}
			inlineAttributes = append(inlineAttributes, a)
		default:
			return fmt.Errorf("unexpected attribute type: %T", a)
		}
	}
	for _, al := range list {
		switch al {
		case attributeClassTitle:
			if !shouldRenderAttributeType(AttributeFilterTitle, include, exclude) {
				continue
			}
			for _, ta := range titleAttributes {
				renderAttributeTitle(cxt, ta.Value().(asciidoc.Set), include, exclude)
			}
		case attributeClassAnchor:
			if !shouldRenderAttributeType(AttributeFilterID, include, exclude) {
				continue
			}
			for _, ta := range anchors {
				renderAttributeAnchor(cxt, ta, include, exclude, inline)
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
				default:
					err = fmt.Errorf("unexpected inline attribute type: %T", ia)
				}
			}
			if len(filtered) == 0 {
				continue
			}
			if !inline {
				cxt.EnsureNewLine()
			}
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
				default:
					err = fmt.Errorf("unexpected inline attribute type: %T", ia)
				}
				if err != nil {
					return
				}
			}
			cxt.WriteString("]")
			if !inline {
				cxt.WriteString("\n")
			}
		default:
			err = fmt.Errorf("unexpected attribute list element: %T", al)
			return
		}

	}
	return
}

func renderAttributeAnchor(cxt *Context, anchor *asciidoc.AnchorAttribute, include AttributeFilter, exclude AttributeFilter, inline bool) {
	id := anchor.ID
	if id != nil && len(id.Value) > 0 && shouldRenderAttributeType(AttributeFilterID, include, exclude) {
		if !inline {
			cxt.EnsureNewLine()
		}
		cxt.WriteString("[[")
		cxt.WriteString(id.Value)
		if len(anchor.Label) > 0 {
			cxt.WriteString(",")
			Elements(cxt, "", anchor.Label...)
		}
		cxt.WriteString("]]")
		if !inline {
			cxt.WriteRune('\n')
		}
	}
}

func renderAttributeTitle(cxt *Context, title asciidoc.Set, include AttributeFilter, exclude AttributeFilter) {
	if len(title) > 0 && shouldRenderAttributeType(AttributeFilterTitle, include, exclude) {
		cxt.EnsureNewLine()
		cxt.WriteRune('.')
		Elements(cxt, "", title...)
		cxt.EnsureNewLine()
	}
}

/*func quoteAttributeValue(cxt *Context, val string) {
	if _, err := strconv.Atoi(strings.TrimSuffix(val, "%")); err == nil {
		cxt.WriteString(val)
	} else {
		cxt.WriteRune('"')
		cxt.WriteString(val)
		cxt.WriteRune('"')
	}
}*/

func renderQuotedAttributeValue(cxt *Context, val any) (err error) {
	var s string
	s, err = quoteAttributeValue(val)
	if err != nil {
		return
	}
	if _, err := strconv.Atoi(strings.TrimSuffix(s, "%")); err == nil {
		cxt.WriteString(s)
		return nil
	}
	cxt.WriteRune('"')
	cxt.WriteString(s)
	cxt.WriteRune('"')
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
	case asciidoc.Set:
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

func renderNakedAttributeValue(cxt *Context, val any) (err error) {
	switch val := val.(type) {
	case *asciidoc.String:
		cxt.WriteString(escapeQuotes(val.Value))
	case asciidoc.AttributeReference:
		cxt.WriteRune('{')
		cxt.WriteString(val.Name())
		cxt.WriteRune('}')
	case asciidoc.Set:
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


func renderAttributeEntry(cxt *Context, ad *asciidoc.AttributeEntry) (err error) {
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
		err = Elements(cxt, "", ad.Elements()...)

		cxt.WriteRune('\n')
	}
	return
}

func renderAttributeReset(cxt *Context, ar *asciidoc.AttributeReset) {
	cxt.WriteRune(':')
	cxt.WriteString(string(ar.Name))
	cxt.WriteString("!:\n")
}

func getAttributeStringValue(cxt *Context, val any) (string, error) {
	switch s := val.(type) {
	case string:
		return s, nil
	case *asciidoc.String:
		return s.Value, nil
	case asciidoc.Set:
		renderContext := NewContext(cxt, cxt.Doc)
		err := Elements(renderContext, "", s...)
		if err != nil {
			return "", err
		}
		return renderContext.String(), nil
	default:
		return "", fmt.Errorf("unknown text attribute value type: %T", val)
	}
}
