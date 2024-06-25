package parse

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/project-chip/alchemy/asciidoc"
)

func PreParseFile(context *AttributeContext, path string) (string, error) {
	fmt.Printf("path: %s\n", path)
	//v, err := os.ReadFile(path)
	//	fmt.Printf("file: %s\n", string(v))
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error reading: %v\n", err)
		return "", err
	}
	return PreParseReader(context, path, file)
}

func PreParseReader(context *AttributeContext, path string, reader io.Reader) (string, error) {
	start := time.Now()
	vals, err := ParseReader(path, reader, Entrypoint("PreParse"))
	if err != nil {
		fmt.Printf("error parsing: %v\n", err)
		return "", err
	}
	elapsed := time.Since(start)

	switch vals := vals.(type) {
	case asciidoc.Set:
		//		fmt.Printf("coalescing asciidoc...\n")

		result := asciidoc.NewWriter(nil)
		err = preparse(context, asciidoc.NewReader(vals), result)
		if err != nil {
			return "", err
		}
		if debugParser {
			fmt.Printf("\n\n\n\n\n\n")
			dump(0, result.Set()...)
			fmt.Printf("elapsed: %s\n", elapsed.String())
		}
		return renderPreParsedDoc(result.Set())
	default:
		return "", fmt.Errorf("unexpected type in PreParseReader: %T", vals)
	}
}

func preparse(context *AttributeContext, r *asciidoc.Reader, w *asciidoc.Writer) (err error) {

	for {
		el := r.Read()
		if el == nil {
			return
		}
		switch el := el.(type) {
		case *asciidoc.AttributeEntry:
			context.Set(string(el.Name), el.Set)
		case *asciidoc.AttributeReset:
			context.Unset(string(el.Name))
		case *asciidoc.UserAttributeReference:
			err = renderReference(context, asciidoc.AttributeName(el.Name()), w)
		case *asciidoc.IfDefBlock:
			if el.Eval(context) {
				err = preparse(context, asciidoc.NewReader(el.Set), w)
			}
		case *asciidoc.IfNDefBlock:
			if el.Eval(context) {
				err = preparse(context, asciidoc.NewReader(el.Set), w)
			}
		case *asciidoc.IfEvalBlock:
			var include bool
			include, err = el.Eval(context)
			if err == nil && include {
				err = preparse(context, asciidoc.NewReader(el.Set), w)
			}
		case *asciidoc.Counter:
			err = renderCounter(context, el, w)
		default:
			w.Write(el)
		}
		if err != nil {
			return
		}
	}
}

func renderReference(context *AttributeContext, name asciidoc.AttributeName, w *asciidoc.Writer) error {
	a := context.Get(string(name))
	if a == nil {
		w.Write(asciidoc.NewString(fmt.Sprintf("{%s}", name)))
		return nil
	}
	switch a := a.(type) {
	case string:
		w.Write(asciidoc.NewString(a))
	case *asciidoc.String:
		w.Write(a)
	case asciidoc.Set:
		return preparse(context, asciidoc.NewReader(a), w)
	default:
		return fmt.Errorf("unknown type rendering reference: %T", a)
	}
	return nil
}

func renderPreParsedDoc(els asciidoc.Set) (string, error) {
	var sb strings.Builder
	for _, el := range els {
		switch el := el.(type) {
		case *asciidoc.String:
			sb.WriteString(el.Value)
		case *asciidoc.NewLine:
			sb.WriteRune('\n')
		case asciidoc.EmptyLine:
			sb.WriteRune('\n')
		case *asciidoc.CharacterReplacementReference:
			sb.WriteString(el.ReplacementValue())
		default:
			return "", fmt.Errorf("unexpected type rendering preparsed doc: %T", el)
		}
	}
	return sb.String(), nil
}

func renderCounter(context *AttributeContext, c *asciidoc.Counter, w *asciidoc.Writer) error {

	cc, ok := context.counters[c.Name]
	if !ok {
		if context.counters == nil {
			context.counters = make(map[string]*counter)
		}
		cc = &counter{}
		context.counters[c.Name] = cc
		switch len(c.InitialValue) {
		case 0:
			cc.value = 1
			cc.counterType = counterTypeInteger
		case 1:
			r := c.InitialValue[0]
			if r >= 'a' && r <= 'z' {
				cc.value = int(r) - int('a')
				cc.counterType = counterTypeLowerCase
			} else if r >= 'A' && r <= 'Z' {
				cc.value = int(r) - int('A')
				cc.counterType = counterTypeUpperCase
			} else {
				var err error
				cc.value, err = strconv.Atoi(c.InitialValue)
				if err != nil {
					return err
				}
				cc.counterType = counterTypeInteger
			}
		default:
			var err error
			cc.value, err = strconv.Atoi(c.InitialValue)
			if err != nil {
				return err
			}
			cc.counterType = counterTypeInteger
		}
	}
	if c.Display {
		switch cc.counterType {
		case counterTypeInteger:
			w.Write(asciidoc.NewString(strconv.Itoa(cc.value)))
		case counterTypeLowerCase:
			r := rune(int('a') + cc.value)
			w.Write(asciidoc.NewString(string(r)))
		case counterTypeUpperCase:
			r := rune(int('A') + cc.value)
			w.Write(asciidoc.NewString(string(r)))
		}
	}
	switch cc.counterType {
	case counterTypeInteger:
		cc.value += 1
	case counterTypeUpperCase:
		if cc.value < 'Z' {
			cc.value += 1
		}
	case counterTypeLowerCase:
		if cc.value < 'z' {
			cc.value += 1
		}
	}
	return nil
}
