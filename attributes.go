package main

import (
	"fmt"
	"sort"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (d *doc) renderAttributes(el interface{}, attributes types.Attributes, out *output) {
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
				var s output
				d.renderElements("", v, &s)
				title = s.String()
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
				out.WriteString(fmt.Sprintf("%s: ", style))
			default:
				out.WriteString(fmt.Sprintf("[%s]\n", style))
			}
		case "none":
			out.WriteString("[none]\n")
		case "lowerroman":
			out.WriteString("[lowerroman]\n")
		case "arabic":
			out.WriteString("[arabic]\n")
		default:
			fmt.Printf("Unknown style: %s\n", style)
		}
	}
	if len(title) > 0 {
		out.WriteRune('.')
		out.WriteString(title)
		out.WriteNewline()
	}
	if len(id) > 0 && id[0] != '_' {
		out.WriteString("[[")
		out.WriteString(id)
		out.WriteString("]]")
		out.WriteRune('\n')
	}
	if len(keys) > 0 {
		sort.Strings(keys)
		switch el.(type) {
		case *types.ImageBlock:
		default:
			out.WriteNewline()
		}
		out.WriteString("[")
		count := 0
		for _, key := range keys {
			switch key {
			case "cols":
				continue
			}
			val := attributes[key]
			if count > 0 {
				out.WriteRune(',')
			}
			out.WriteString(key)
			out.WriteRune('=')
			switch v := val.(type) {
			case string:
				out.WriteRune('"')
				out.WriteString(v)
				out.WriteRune('"')
			case types.Options:
				out.WriteRune('"')
				for _, o := range v {
					switch opt := o.(type) {
					case string:
						out.WriteString(opt)
					default:
						fmt.Printf("unknown attribute option type: %T\n", o)
					}
				}
				out.WriteRune('"')
			default:
				fmt.Printf("unknown attribute type: %T\n", val)
			}
			count++
		}
		out.WriteRune(']')
		out.WriteRune('\n')
	}
}
