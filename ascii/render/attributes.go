package render

import (
	"fmt"
	"log/slog"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
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

func renderAttributes(cxt *Context, el any, attributes types.Attributes, inline bool) error {
	_, err := renderSelectAttributes(cxt, el, attributes, AttributeFilterAll, AttributeFilterNone, inline)
	return err
}

func renderSelectAttributes(cxt *Context, el any, attributes types.Attributes, include AttributeFilter, exclude AttributeFilter, inline bool) (n int, err error) {
	if len(attributes) == 0 {
		return
	}

	var id string
	var title string
	var style string
	var keys []string
	var roles types.Roles
	for key, val := range attributes {
		switch key {
		case types.AttrID:
			id = val.(string)
		case types.AttrStyle:
			style = val.(string)
		case types.AttrTitle:
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
		case types.AttrPositional1:
			if s, ok := val.(string); ok {
				style = s
			}
		case types.AttrRoles:
			switch v := val.(type) {
			case types.Roles:
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
		case types.Tip, types.Note, types.Important, types.Warning, types.Caution:
			switch el.(type) {
			case *types.Paragraph:
				cxt.WriteString(fmt.Sprintf("%s: ", style))
			default:
				n++
				cxt.WriteString(fmt.Sprintf("[%s]\n", style))
			}
		case "none":
			cxt.WriteString("[none]\n")
			n++
			renderAttributeTitle(cxt, title, include, exclude)
			renderAttributeAnchor(cxt, id, include, exclude, inline)
		case types.UpperRoman, types.LowerRoman, types.Arabic, types.UpperAlpha, types.LowerAlpha:
			if !inline {
				cxt.WriteNewline()
			}
			renderAttributeTitle(cxt, title, include, exclude)
			renderAttributeAnchor(cxt, id, include, exclude, inline)
			cxt.WriteRune('[')
			cxt.WriteString(style)
			n++
			if start, ok := attributes[types.AttrStart]; ok {
				if start, ok := start.(string); ok {
					n++
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
			n++
			lang, ok := attributes[types.AttrLanguage]
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
			n++
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
				n++
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
				n++
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
	n = 0
	if len(keys) > 0 {
		sort.Strings(keys)
		if !inline {
			cxt.WriteNewline()
		}

		for _, key := range keys {
			var keyVal string
			var skipKey bool
			keyVal, skipKey, err = getKeyValue(cxt, key, attributes[key], include, exclude)
			if err != nil {
				return
			}

			if len(keyVal) != 0 {
				if n == 0 {
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

				n++
			}

		}
		if n > 0 {
			cxt.WriteRune(']')
			if !inline {
				cxt.WriteRune('\n')
			}
		}
	}
	return
}

func renderAttributeAnchor(cxt *Context, id string, include AttributeFilter, exclude AttributeFilter, inline bool) {
	if len(id) > 0 && id[0] != '_' && shouldRenderAttributeType(AttributeFilterID, include, exclude) {
		if !inline {
			cxt.WriteNewline()
		}
		cxt.WriteString("[[")
		cxt.WriteString(id)
		cxt.WriteString("]]")
		if !inline {
			cxt.WriteRune('\n')
		}
	}
}

func renderAttributeTitle(cxt *Context, title string, include AttributeFilter, exclude AttributeFilter) {
	if len(title) > 0 && shouldRenderAttributeType(AttributeFilterTitle, include, exclude) {
		cxt.WriteNewline()
		cxt.WriteRune('.')
		cxt.WriteString(title)
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

func getKeyValue(cxt *Context, key string, val any, include AttributeFilter, exclude AttributeFilter) (keyVal string, skipKey bool, err error) {
	var attributeType AttributeFilter
	switch key {
	case types.AttrCols:
		attributeType = AttributeFilterCols
	case types.AttrInlineLinkText:
		attributeType = AttributeFilterText
	case types.AttrImageAlt:
		attributeType = AttributeFilterAlt
		skipKey = true
	case types.AttrHeight:
		attributeType = AttributeFilterHeight
	case types.AttrWidth:
		attributeType = AttributeFilterWidth
	case types.AttrRoles:
		attributeType = AttributeFilterRole
		skipKey = true
	case types.AttrFloat:
		attributeType = AttributeFilterFloat
	case types.AttrImageAlign:
		attributeType = AttributeFilterAlign
	case "pdfwidth":
		attributeType = AttributeFilterPDFWidth
	case types.AttrPositional1, types.AttrPositional2, types.AttrPositional3:
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

		case types.Options:
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
				case *types.TableColumn:
					var val strings.Builder
					if tc.Multiplier > 1 || tc.MultiplierSpecified {
						val.WriteString(strconv.Itoa(tc.Multiplier))
						val.WriteRune('*')
					}
					if tc.HAlign != types.HAlignDefault || tc.HAlignSpecified {
						val.WriteString(string(tc.HAlign))
					}
					if tc.VAlign != types.VAlignDefault || tc.VAlignSpecified {
						val.WriteRune('.')
						val.WriteString(string(tc.VAlign))
					}
					if tc.Autowidth {
						val.WriteRune('~')
					} else if tc.Weight > 1 || tc.WeightSpecified {
						val.WriteString(strconv.Itoa(tc.Weight))
					} else if tc.PercentageSpecified {
						val.WriteString(strconv.Itoa(tc.Percentage))
						val.WriteRune('%')
					}
					if len(tc.Style) > 0 {
						val.WriteString(string(tc.Style))
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

func renderDiagramAttributes(cxt *Context, style string, id string, title string, keys []string, inline bool, attributes types.Attributes, include AttributeFilter, exclude AttributeFilter) (n int, err error) {

	renderAttributeTitle(cxt, title, include, exclude)
	renderAttributeAnchor(cxt, id, include, exclude, inline)
	cxt.WriteString("[")
	cxt.WriteString(style)
	slices.Sort(keys)
	n = 1
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
		n++
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
}

func renderAttributeDeclaration(cxt *Context, ad *types.AttributeDeclaration) (err error) {
	switch ad.Name {
	case "authors":
		if authors, ok := ad.Value.(types.DocumentAuthors); ok {
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
		}
	default:
		cxt.WriteNewline()
		cxt.WriteString(ad.RawText())
	}
	return
}

func renderAttributeReset(cxt *Context, ar *types.AttributeReset) {
	cxt.WriteRune(':')
	cxt.WriteString(ar.Name)
	cxt.WriteString("!:\n")
}

func getAttributeStringValue(cxt *Context, val any) (string, error) {
	switch s := val.(type) {
	case string:
		return s, nil
	case *types.StringElement:
		return s.Content, nil
	case []any:
		renderContext := NewContext(cxt, cxt.Doc)
		err := Elements(renderContext, "", s)
		if err != nil {
			return "", err
		}
		return renderContext.String(), nil
	default:
		return "", fmt.Errorf("unknown text attribute value type: %T", val)
	}
}
