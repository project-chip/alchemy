package render

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/hasty/adoc/elements"
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

func renderAttributes(cxt *Context, el any, attributes []elements.Attribute, inline bool) error {
	return renderSelectAttributes(cxt, el, attributes, AttributeFilterAll, AttributeFilterNone, inline)
}

func renderSelectAttributes(cxt *Context, el any, attributes []elements.Attribute, include AttributeFilter, exclude AttributeFilter, inline bool) (err error) {
	if len(attributes) == 0 {
		return
	}

	var titleAttributes []*elements.TitleAttribute
	var anchors []*elements.Anchor
	var inlineAttributes []elements.Attribute
	for _, a := range attributes {
		switch a := a.(type) {
		case *elements.TitleAttribute:
			titleAttributes = append(titleAttributes, a)
		case *elements.Anchor:
			anchors = append(anchors, a)
		case *elements.NamedAttribute, *elements.PositionalAttribute, *elements.TableColumnsAttribute:
			inlineAttributes = append(inlineAttributes, a)
		default:
			return fmt.Errorf("unexpected attribute type: %T", a)
		}
	}
	if len(inlineAttributes) > 0 {
		cxt.WriteString("[")
		for i, ia := range inlineAttributes {
			if i > 0 {
				cxt.WriteRune(',')
			}
			switch ia := ia.(type) {
			case *elements.NamedAttribute:
				cxt.WriteString(string(ia.Name))
				cxt.WriteString("=\"")
				err = renderQuotedAttributeValue(cxt, ia.Value())
				cxt.WriteString("\"")
			case *elements.PositionalAttribute:
				err = renderNakedAttributeValue(cxt, ia.Value())
			case *elements.TableColumnsAttribute:
				cxt.WriteString("cols=\"")
				renderQuotedAttributeValue(cxt, ia.AsciiDocString())
				cxt.WriteString("\"")
			default:
				err = fmt.Errorf("unexpected inline attribute type: %T", ia)
			}
			if err != nil {
				return
			}
		}
		cxt.WriteString("]\n")
	}
	for _, ta := range titleAttributes {
		renderAttributeTitle(cxt, ta.Value().([]elements.Element), include, exclude)
	}
	for _, aa := range anchors {
		renderAttributeAnchor(cxt, aa, include, exclude, inline)
	}

	/*var id string
	var title string
	var style string
	var keys []string
	var roles elements.Roles
	for _, val := range attributes {
		switch val := val.(type) {
		case *elements.NamedAttribute:
			switch val.Name {
			case elements.AttributeNameID:
			case elements.AttributeNameStyle:
			case elements.AttributeNameTitle:
			case elements.AttributeNameRoles:

			}
		case *elements.PositionalAttribute:
		}
		switch key {
		case elements.AttrID:
			id = val.(string)
		case elements.AttrStyle:
			style = val.(string)
		case elements.AttrTitle:
			switch v := val.(type) {
			case string:
				title = v
			case []any:
				renderContext := NewContext(cxt, cxt.Doc)
				err = Elements(renderContext, "", v)
				title = renderContext.String()
			default:
				err = fmt.Errorf("unknown title type: %T", v)
				return
			}
		case elements.AttrPositional1:
			if s, ok := val.(string); ok {
				style = s
			}
		case elements.AttrRoles:
			switch v := val.(type) {
			case elements.Roles:
				roles = v
			case []any:
				roles = v
			case any:
				roles = []any{v}
			default:
				err = fmt.Errorf("unknown roles type: %T", v)
				return
			}

		default:
			keys = append(keys, key)
		}
		if err != nil {
			return
		}
	}

	if len(style) > 0 && shouldRenderAttributeType(AttributeFilterStyle, include, exclude) {
		switch style {
		case elements.Tip, elements.Note, elements.Important, elements.Warning, elements.Caution:
			switch el.(type) {
			case *elements.Paragraph:
				cxt.WriteString(fmt.Sprintf("%s: ", style))
			default:
				cxt.WriteString(fmt.Sprintf("[%s]\n", style))
			}
		case "none":
			cxt.WriteString("[none]\n")
			renderAttributeTitle(cxt, title, include, exclude)
			renderAttributeAnchor(cxt, id, include, exclude, inline)
		case elements.UpperRoman, elements.LowerRoman, elements.Arabic, elements.UpperAlpha, elements.LowerAlpha:
			if !inline {
				cxt.WriteNewline()
			}
			renderAttributeTitle(cxt, title, include, exclude)
			renderAttributeAnchor(cxt, id, include, exclude, inline)
			cxt.WriteRune('[')
			cxt.WriteString(style)
			if start, ok := attributes[elements.AttrStart]; ok {
				if start, ok := start.(string); ok {
					cxt.WriteString(", start=")
					cxt.WriteString(start)
				}
			}
			cxt.WriteString("]\n")
			renderAttributeTitle(cxt, title, include, exclude)
			renderAttributeAnchor(cxt, id, include, exclude, inline)
			return
		case "a2s", "actdiag", "plantuml", "qrcode", "blockdiag", "d2", "lilypond", "ditaa", "graphviz", "asciimath":
			return renderDiagramAttributes(cxt, style, id, title, keys, inline, attributes, include, exclude)
		case "literal_paragraph":
		case "source":
			renderAttributeTitle(cxt, title, include, exclude)
			renderAttributeAnchor(cxt, id, include, exclude, inline)
			if !inline {
				cxt.WriteNewline()
			}
			cxt.WriteRune('[')
			cxt.WriteString(style)
			lang, ok := attributes[elements.AttrLanguage]
			if ok {
				cxt.WriteString(", ")
				cxt.WriteString(lang.(string))
			}
			cxt.WriteString("]\n")
			renderAttributeTitle(cxt, title, include, exclude)
			renderAttributeAnchor(cxt, id, include, exclude, inline)
			return
		default:
			if !inline {
				cxt.WriteNewline()
			}
			cxt.WriteRune('[')
			cxt.WriteString(style)
			for _, key := range keys {
				var keyVal string
				var skipKey bool
				keyVal, skipKey, err = getKeyValue(cxt, key, attributes[key], include, exclude)
				if err != nil {
					return
				}
				if keyVal == "" {
					continue
				}
				cxt.WriteRune(',')
				if skipKey {
					cxt.WriteString(keyVal)
				} else {
					cxt.WriteString(key)
					cxt.WriteRune('=')
					quoteAttributeValue(cxt, keyVal)
				}
			}
			cxt.WriteString("]\n")
			renderAttributeTitle(cxt, title, include, exclude)
			renderAttributeAnchor(cxt, id, include, exclude, inline)
			return
		}
	}
	if len(roles) > 0 {
		if !inline {
			cxt.WriteNewline()
		}
		cxt.WriteString("[")
		for _, r := range roles {
			switch rs := r.(type) {
			case string:
				cxt.WriteRune('.')
				cxt.WriteString(rs)
			default:
				slog.Debug("unknown role type", "role", r)
			}
		}
		cxt.WriteString("]")
		if !inline {
			cxt.WriteNewline()
		}
		renderAttributeTitle(cxt, title, include, exclude)
		renderAttributeAnchor(cxt, id, include, exclude, inline)
		return
	}
	renderAttributeTitle(cxt, title, include, exclude)
	renderAttributeAnchor(cxt, id, include, exclude, inline)
	if len(keys) > 0 {
		sort.Strings(keys)
		if !inline {
			cxt.WriteNewline()
		}

		count := 0
		for _, key := range keys {
			var keyVal string
			var skipKey bool
			keyVal, skipKey, err = getKeyValue(cxt, key, attributes[key], include, exclude)
			if err != nil {
				return
			}

			if len(keyVal) != 0 {
				if count == 0 {
					cxt.WriteString("[")
				} else {
					cxt.WriteRune(',')
				}
				if skipKey {
					cxt.WriteString(keyVal)
				} else {
					cxt.WriteString(key)
					cxt.WriteRune('=')
					quoteAttributeValue(cxt, keyVal)
				}

				count++
			}

		}
		if count > 0 {
			cxt.WriteRune(']')
			if !inline {
				cxt.WriteRune('\n')
			}
		}
	}*/
	return
}

func renderAttributeAnchor(cxt *Context, anchor *elements.Anchor, include AttributeFilter, exclude AttributeFilter, inline bool) {
	id := anchor.String()
	if len(id) > 0 && id[0] != '_' && shouldRenderAttributeType(AttributeFilterID, include, exclude) {
		if !inline {
			cxt.WriteNewline()
		}
		cxt.WriteString("[[")
		cxt.WriteString(id)
		if len(anchor.Label) > 0 {
			cxt.WriteString(", ")
			Elements(cxt, "", anchor.Label...)
		}
		cxt.WriteString("]]")
		if !inline {
			cxt.WriteRune('\n')
		}
	}
}

func renderAttributeTitle(cxt *Context, title []elements.Element, include AttributeFilter, exclude AttributeFilter) {
	if len(title) > 0 && shouldRenderAttributeType(AttributeFilterTitle, include, exclude) {
		cxt.WriteNewline()
		cxt.WriteRune('.')
		Elements(cxt, "", title...)
		cxt.WriteNewline()
	}
}

func quoteAttributeValue(cxt *Context, val string) {
	if _, err := strconv.Atoi(strings.TrimSuffix(val, "%")); err == nil {
		cxt.WriteString(val)
	} else {
		cxt.WriteRune('"')
		cxt.WriteString(val)
		cxt.WriteRune('"')
	}
}

func renderQuotedAttributeValue(cxt *Context, val any) (err error) {
	switch val := val.(type) {
	case string:
		cxt.WriteString(escapeQuotes(val))
	case *elements.String:
		cxt.WriteString(escapeQuotes(val.Value))
	case elements.AttributeReference:
		cxt.WriteRune('{')
		cxt.WriteString(val.Name())
		cxt.WriteRune('}')
	case []elements.Element:
		for _, a := range val {
			err = renderQuotedAttributeValue(cxt, a)
			if err != nil {
				return
			}
		}
	default:
		err = fmt.Errorf("unexpected attribute value type: %T", val)
	}
	return
}

func renderNakedAttributeValue(cxt *Context, val any) (err error) {
	switch val := val.(type) {
	case *elements.String:
		cxt.WriteString(escapeQuotes(val.Value))
	case elements.AttributeReference:
		cxt.WriteRune('{')
		cxt.WriteString(val.Name())
		cxt.WriteRune('}')
	case []elements.Element:
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

/*
func getKeyValue(cxt *Context, key string, val any, include AttributeFilter, exclude AttributeFilter) (keyVal string, skipKey bool, err error) {
	var attributeType AttributeFilter
	switch key {
	case elements.AttrCols:
		attributeType = AttributeFilterCols
	case elements.AttrInlineLinkText:
		attributeType = AttributeFilterText
	case elements.AttrImageAlt:
		attributeType = AttributeFilterAlt
		skipKey = true
	case elements.AttrHeight:
		attributeType = AttributeFilterHeight
	case elements.AttrWidth:
		attributeType = AttributeFilterWidth
	case elements.AttrRoles:
		attributeType = AttributeFilterRole
		skipKey = true
	case elements.AttrFloat:
		attributeType = AttributeFilterFloat
	case elements.AttrImageAlign:
		attributeType = AttributeFilterAlign
	case "pdfwidth":
		attributeType = AttributeFilterPDFWidth
	case elements.AttrPositional1, elements.AttrPositional2, elements.AttrPositional3:
		skipKey = true
	}
	if attributeType != AttributeFilterNone && !shouldRenderAttributeType(attributeType, include, exclude) {
		return
	}

	switch attributeType {
	case AttributeFilterText:
		keyVal, err = getAttributeStringValue(cxt, val)
		if err != nil {
			return
		}
		skipKey = true
	case AttributeFilterAlt:
		keyVal, err = getAttributeStringValue(cxt, val)
		if err != nil {
			return
		}
	default:
		switch v := val.(type) {
		case string:
			keyVal = v

		case elements.Options:
			for _, o := range v {
				switch opt := o.(type) {
				case string:
					keyVal = opt
				default:
					slog.Debug("unknown attribute option", "type", o)
				}
			}
		case []any:

			var columns []string
			for _, e := range v {
				switch tc := e.(type) {
				case *elements.TableColumn:
					var val strings.Builder
					if tc.Multiplier.IsSet() {
						val.WriteString(strconv.Itoa(tc.Multiplier.Get()))
						val.WriteRune('*')
					}
					if tc.HorizontalAlign.IsSet() {
						switch tc.HorizontalAlign.Get() {
						case elements.TableCellHorizontalAlignLeft:
							val.WriteRune('<')
						case elements.TableCellHorizontalAlignCenter:
							val.WriteRune('^')
						case elements.TableCellHorizontalAlignRight:
							val.WriteRune('>')
						}
					}
					if tc.VerticalAlign.IsSet() {
						val.WriteRune('.')
						switch tc.VerticalAlign.Get() {
						case elements.TableCellVerticalAlignTop:
							val.WriteRune('<')
						case elements.TableCellVerticalAlignMiddle:
							val.WriteRune('^')
						case elements.TableCellVerticalAlignBottom:
							val.WriteRune('>')
						}
					}
					if tc.Width.IsSet() {
						width := tc.Width.Get()
						if width == -1 {
							val.WriteRune('~')
						} else {
							val.WriteString(strconv.Itoa(int(width)))
						}
					}
					if tc.Style.IsSet() {
						switch tc.Style.Get() {
						case elements.TableCellStyleAsciiDoc:
							val.WriteRune('a')
						case elements.TableCellStyleDefault:
							val.WriteRune('d')
						case elements.TableCellStyleEmphasis:
							val.WriteRune('e')
						case elements.TableCellStyleHeader:
							val.WriteRune('h')
						case elements.TableCellStyleLiteral:
							val.WriteRune('l')
						case elements.TableCellStyleMonospace:
							val.WriteRune('m')
						case elements.TableCellStyleStrong:
							val.WriteRune('s')
						}
					}
					columns = append(columns, val.String())

				default:
					err = fmt.Errorf("unknown attribute: %T", e)
					return
				}
			}
			keyVal = strings.Join(columns, ",")
		case nil:
			keyVal = ""
		default:
			err = fmt.Errorf("unknown attribute type: %T", val)
			return
		}
	}
	return
}
*/
/*
func renderDiagramAttributes(cxt *Context, style string, id string, title string, keys []string, inline bool, attributes elements.AttributeList, include AttributeFilter, exclude AttributeFilter) (err error) {

	renderAttributeTitle(cxt, title, include, exclude)
	renderAttributeAnchor(cxt, id, include, exclude, inline)
	cxt.WriteString("[")
	cxt.WriteString(style)
	slices.Sort(keys)
	for _, key := range keys {
		var keyVal string
		var skipKey bool
		keyVal, skipKey, err = getKeyValue(cxt, key, attributes[key], include, exclude)
		if err != nil {
			return
		}
		if keyVal == "" {
			continue
		}
		cxt.WriteRune(',')
		if skipKey {
			cxt.WriteString(keyVal)
		} else {
			cxt.WriteString(key)
			cxt.WriteRune('=')
			quoteAttributeValue(cxt, keyVal)
		}
	}
	cxt.WriteRune(']')
	if !inline {
		cxt.WriteRune('\n')
	}
	return
}*/

func renderAttributeEntry(cxt *Context, ad *elements.AttributeEntry) (err error) {
	switch ad.Name {
	case "authors":
		/*if authors, ok := ad.Value().(elements.DocumentAuthors); ok {
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

func renderAttributeReset(cxt *Context, ar elements.AttributeReset) {
	cxt.WriteRune(':')
	cxt.WriteString(string(ar.Name))
	cxt.WriteString("!:\n")
}

func getAttributeStringValue(cxt *Context, val any) (string, error) {
	switch s := val.(type) {
	case string:
		return s, nil
	case *elements.String:
		return s.Value, nil
	case []elements.Element:
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
