package render

import (
	"fmt"
	"sort"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/matterfmt/output"
)

func renderAttributes(cxt *output.Context, el interface{}, attributes types.Attributes) {
	if len(attributes) == 0 {
		return
	}

	var id string
	var title string
	var style string
	var keys []string
	for key, val := range attributes {
		switch key {
		case "id":
			id = val.(string)
		case "style":
			style = val.(string)
		case "title":
			switch v := val.(type) {
			case string:
				title = v
			case []interface{}:
				renderContext := output.NewContext(cxt, cxt.Doc)
				RenderElements(renderContext, "", v)
				title = renderContext.String()
				for _, p := range v {
					fmt.Printf("title element: %T\n", p)
				}
			default:
				panic(fmt.Sprintf("unknown title type: %T", v))
			}
		default:
			keys = append(keys, key)
		}
	}
	if len(style) > 0 {
		switch style {
		case "NOTE", "IMPORTANT", "TIP", "CAUTION", "WARNING":
			switch el.(type) {
			case *types.Paragraph:
				cxt.WriteString(fmt.Sprintf("%s: ", style))
			default:
				cxt.WriteString(fmt.Sprintf("[%s]\n", style))
			}
		case "none":
			cxt.WriteString("[none]\n")
		case "lowerroman":
			cxt.WriteString("[lowerroman]\n")
		case "arabic":
			cxt.WriteString("[arabic]\n")
		default:
			fmt.Printf("Unknown style: %s\n", style)
		}
	}
	if len(title) > 0 {
		cxt.WriteRune('.')
		cxt.WriteString(title)
		cxt.WriteNewline()
	}
	if len(id) > 0 && id[0] != '_' {
		cxt.WriteString("[[")
		cxt.WriteString(id)
		cxt.WriteString("]]")
		cxt.WriteRune('\n')
	}
	if len(keys) > 0 {
		sort.Strings(keys)
		switch el.(type) {
		case *types.ImageBlock:
		default:
			cxt.WriteNewline()
		}
		cxt.WriteString("[")
		count := 0
		for _, key := range keys {
			switch key {
			case "cols":
				continue
			}
			val := attributes[key]
			if count > 0 {
				cxt.WriteRune(',')
			}
			cxt.WriteString(key)
			cxt.WriteRune('=')
			switch v := val.(type) {
			case string:
				cxt.WriteRune('"')
				cxt.WriteString(v)
				cxt.WriteRune('"')
			case types.Options:
				cxt.WriteRune('"')
				for _, o := range v {
					switch opt := o.(type) {
					case string:
						cxt.WriteString(opt)
					default:
						fmt.Printf("unknown attribute option type: %T\n", o)
					}
				}
				cxt.WriteRune('"')
			default:
				fmt.Printf("unknown attribute type: %T\n", val)
			}
			count++
		}
		cxt.WriteRune(']')
		cxt.WriteRune('\n')
	}
}
