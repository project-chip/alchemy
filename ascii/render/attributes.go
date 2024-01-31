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
)

func shouldRenderAttributeType(at AttributeFilter, include AttributeFilter, exclude AttributeFilter) bool {
	return ((at & include) == at) && ((at & exclude) != at)
}

func renderAttributes(cxt *Context, el interface{}, attributes types.Attributes, inline bool) error {
	return renderSelectAttributes(cxt, el, attributes, AttributeFilterAll, AttributeFilterNone, inline)
}

func renderSelectAttributes(cxt *Context, el interface{}, attributes types.Attributes, include AttributeFilter, exclude AttributeFilter, inline bool) (err error) {
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
			case []interface{}:
				renderContext := NewContext(cxt, cxt.Doc)
				err = RenderElements(renderContext, "", v)
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
			case []interface{}:
				roles = v
			case interface{}:
				roles = []interface{}{v}
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
				cxt.WriteString(fmt.Sprintf("[%s]\n", style))
			}
		case "none":
			cxt.WriteString("[none]\n")
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

func getKeyValue(cxt *Context, key string, val interface{}, include AttributeFilter, exclude AttributeFilter) (keyVal string, skipKey bool, err error) {
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
		case []interface{}:

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

func renderDiagramAttributes(cxt *Context, style string, id string, title string, keys []string, inline bool, attributes types.Attributes, include AttributeFilter, exclude AttributeFilter) (err error) {

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
		cxt.WriteRune(':')
		cxt.WriteString(ad.Name)
		cxt.WriteString(":")
		switch val := ad.Value.(type) {
		case string:
			cxt.WriteRune(' ')
			cxt.WriteString(val)
		case *types.Paragraph:
			var previous interface{}
			err = renderParagraph(cxt, val, &previous)
		case []interface{}:
			err = RenderElements(cxt, "", val)
		case nil:
		default:
			err = fmt.Errorf("unknown attribute declaration value type: %T", ad.Value)
		}
		cxt.WriteRune('\n')
	}
	return
}

func renderAttributeReset(cxt *Context, ar *types.AttributeReset) {
	cxt.WriteRune(':')
	cxt.WriteString(ar.Name)
	cxt.WriteString("!:\n")
}

func getAttributeStringValue(cxt *Context, val interface{}) (string, error) {
	switch s := val.(type) {
	case string:
		return s, nil
	case *types.StringElement:
		return s.Content, nil
	case []interface{}:
		renderContext := NewContext(cxt, cxt.Doc)
		err := RenderElements(renderContext, "", s)
		if err != nil {
			return "", err
		}
		return renderContext.String(), nil
	default:
		return "", fmt.Errorf("unknown text attribute value type: %T", val)
	}
}
