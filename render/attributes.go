package render

import (
	"fmt"
	"log/slog"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
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
)

func shouldRenderAttributeType(at AttributeFilter, include AttributeFilter, exclude AttributeFilter) bool {
	return ((at & include) == at) && ((at & exclude) != at)
}

func renderAttributes(cxt *output.Context, el interface{}, attributes types.Attributes, inline bool) error {
	return renderSelectAttributes(cxt, el, attributes, AttributeFilterAll, AttributeFilterCols, inline)
}

func renderSelectAttributes(cxt *output.Context, el interface{}, attributes types.Attributes, include AttributeFilter, exclude AttributeFilter, inline bool) (err error) {
	if len(attributes) == 0 {
		return
	}

	var id string
	var title string
	var style string
	var keys []string
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
				renderContext := output.NewContext(cxt, cxt.Doc)
				RenderElements(renderContext, "", v)
				title = renderContext.String()
			default:
				err = fmt.Errorf("unknown title type: %T", v)
				return
			}
		default:
			keys = append(keys, key)
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
		case types.UpperRoman, types.LowerRoman, types.Arabic, types.UpperAlpha, types.LowerAlpha:
			cxt.WriteRune('[')
			cxt.WriteString(style)
			cxt.WriteString("]\n")
			return
		case "a2s", "actdiag", "plantuml", "qrcode", "blockdiag", "d2", "lilypond":
			renderDiagramAttributes(cxt, style, id, keys, attributes)
			return
		case "literal_paragraph":
		case "source":
			cxt.WriteRune('[')
			cxt.WriteString(style)
			lang, ok := attributes[types.AttrLanguage]
			if ok {
				cxt.WriteString(", ")
				cxt.WriteString(lang.(string))
			}
			cxt.WriteString("]\n")
			return
		default:
			err = fmt.Errorf("unknown style: %s", style)
			return
		}
	}
	if len(title) > 0 && shouldRenderAttributeType(AttributeFilterTitle, include, exclude) {
		cxt.WriteNewline()
		cxt.WriteRune('.')
		cxt.WriteString(title)
		cxt.WriteNewline()
	}
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
	if len(keys) > 0 {
		sort.Strings(keys)
		if !inline {
			cxt.WriteNewline()
		}

		count := 0
		for _, key := range keys {
			var attributeType AttributeFilter
			var skipKey = false
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
			case "pdfwidth":
				attributeType = AttributeFilterPDFWidth
			}
			if !shouldRenderAttributeType(AttributeFilterAlt, include, exclude) {
				continue
			}
			val := attributes[key]
			var keyVal string

			switch attributeType {
			case AttributeFilterText:
				if s, ok := val.(string); ok {
					keyVal = s
					skipKey = true
				}
			case AttributeFilterAlt:
				if s, ok := val.(string); ok {
					keyVal = s
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
					for i, e := range v {
						switch tc := e.(type) {
						case *types.TableColumn:
							var val strings.Builder
							if tc.Multiplier > 1 {
								val.WriteString(strconv.Itoa(tc.Multiplier))
								val.WriteRune('*')
							}
							if tc.HAlign != types.HAlignDefault {
								val.WriteString(string(tc.HAlign))
							}
							if tc.VAlign != types.VAlignDefault {
								val.WriteString(string(tc.VAlign))
							}
							if tc.Autowidth {
								val.WriteRune('~')
							} else if tc.Weight > 1 {
								val.WriteString(strconv.Itoa(tc.Weight))
							}
							if len(tc.Style) > 0 {
								val.WriteString(string(tc.Style))
							}
							columns = append(columns, val.String())
							if i == len(v)-1 && val.Len() == 0 {
								// The parser looks for tokens ending with commas, but these values
								// are actually joined with commas; if the last value is blank,
								// the parser will report one fewer column def, so we add it back
								columns = append(columns, "")
							}
						default:
							err = fmt.Errorf("unknown attribute: %T", e)
							return
						}
					}
					keyVal = strings.Join(columns, ",")
				default:
					err = fmt.Errorf("unknown attribute type: %T", val)
					return
				}
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
					if _, err := strconv.Atoi(strings.TrimSuffix(keyVal, "%")); err == nil {
						cxt.WriteString(keyVal)
					} else {
						cxt.WriteRune('"')
						cxt.WriteString(keyVal)
						cxt.WriteRune('"')
					}
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

func renderDiagramAttributes(cxt *output.Context, style string, id string, keys []string, attributes types.Attributes) {
	cxt.WriteString("[")
	cxt.WriteString(style)
	if len(id) > 0 {
		cxt.WriteString(", id=\"")
		cxt.WriteString(id)
		cxt.WriteRune('"')
	}
	for _, k := range keys {
		v, ok := attributes[k]
		if !ok {
			continue
		}
		cxt.WriteString(", ")
		cxt.WriteString(k)
		s, ok := v.(string)
		if ok && len(s) > 0 {
			cxt.WriteString(`="`)
			cxt.WriteString(s)
			cxt.WriteRune('"')
		}
	}
	cxt.WriteRune(']')
	cxt.WriteRune('\n')
}

func renderAttributeDeclaration(cxt *output.Context, ad *types.AttributeDeclaration) (err error) {
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
		case nil:
		default:
			err = fmt.Errorf("unknown attribute declaration value type: %T", ad.Value)
		}
		cxt.WriteRune('\n')
	}
	return
}
