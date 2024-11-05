package parse

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
)

type PreParseContext interface {
	IsSet(name string) bool
	Get(name string) any
	Set(name string, value any)
	Unset(name string)
	GetCounterState(name string, initialValue string) (*CounterState, error)
	ResolvePath(path string) (asciidoc.Path, error)
	ShouldIncludeFile(path asciidoc.Path) bool
}

type CounterType uint8

const (
	CounterTypeInteger CounterType = iota
	CounterTypeUpperCase
	CounterTypeLowerCase
)

type CounterState struct {
	CounterType CounterType
	Value       int
}

func PreParseFile(context PreParseContext, path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		slog.Error("error reading file for preparse", slog.String("path", path), slog.Any("error", err))
		return "", err
	}
	return PreParseReader(context, path, file)
}

func PreParseReader(context PreParseContext, path string, reader io.Reader) (string, error) {
	vals, err := preParseReaderToSet(context, path, reader)
	if err != nil {
		return "", err
	}
	return renderPreParsedDoc(vals)
}

func preParseReaderToSet(context PreParseContext, path string, reader io.Reader) (asciidoc.Set, error) {
	vals, err := ParseReader(path, reader, Entrypoint("PreParse"))
	if err != nil {
		slog.Error("error preparsing file", slog.String("path", path), slog.Any("error", err))
		return nil, err
	}
	set, ok := vals.(asciidoc.Set)
	if ok {
		result := asciidoc.NewWriter(nil)
		err = preparseElements(context, asciidoc.NewReader(set), result)
		if err != nil {
			return nil, err
		}
		return result.Set(), nil
	}
	return nil, fmt.Errorf("unexpected type in PreParseReader: %T", vals)
}

func preparseElements(context PreParseContext, r *asciidoc.Reader, w *asciidoc.Writer) (err error) {
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
				err = preparseElements(context, asciidoc.NewReader(el.Set), w)
			}
		case *asciidoc.IfNDefBlock:
			if el.Eval(context) {
				err = preparseElements(context, asciidoc.NewReader(el.Set), w)
			}
		case *asciidoc.IfEvalBlock:
			var include bool
			include, err = el.Eval(context)
			if err == nil && include {
				err = preparseElements(context, asciidoc.NewReader(el.Set), w)
			}
		case *asciidoc.Counter:
			err = renderCounter(context, el, w)
		case *asciidoc.FileInclude:
			rawPathWriter := asciidoc.NewWriter(nil)
			err = preparseElements(context, asciidoc.NewReader(el.Set), rawPathWriter)

			if err != nil {
				return
			}
			var rawPath string
			rawPath, err = renderPreParsedDoc(rawPathWriter.Set())
			if err != nil {
				return
			}
			var path asciidoc.Path
			path, err = context.ResolvePath(rawPath)
			if err != nil {
				return
			}
			if context.ShouldIncludeFile(path) {
				err = includeFile(context, path.Absolute, w)
			} else {
				w.Write(el)
			}
		default:
			w.Write(el)
		}
		if err != nil {
			return
		}
	}
}

func includeFile(context PreParseContext, path string, w *asciidoc.Writer) error {
	file, err := os.Open(path)
	if err != nil {
		slog.Error("error reading file for preparse", slog.String("path", path), slog.Any("error", err))
		return err
	}
	set, err := preParseReaderToSet(context, path, file)
	if err != nil {
		return err
	}
	return preparseElements(context, asciidoc.NewReader(set), w)
}

func renderReference(context PreParseContext, name asciidoc.AttributeName, w *asciidoc.Writer) error {
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
		return preparseElements(context, asciidoc.NewReader(a), w)
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
		case *asciidoc.FileInclude:
			sb.WriteString(el.Raw())
			sb.WriteRune('\n')
		default:
			return "", fmt.Errorf("unexpected type rendering preparsed doc: %T", el)
		}
	}
	return sb.String(), nil
}

func renderCounter(context PreParseContext, c *asciidoc.Counter, w *asciidoc.Writer) error {

	cc, err := context.GetCounterState(c.Name, c.InitialValue)
	if err != nil {
		return err
	}

	if c.Display {
		switch cc.CounterType {
		case CounterTypeInteger:
			w.Write(asciidoc.NewString(strconv.Itoa(cc.Value)))
		case CounterTypeLowerCase:
			r := rune(int('a') + cc.Value)
			w.Write(asciidoc.NewString(string(r)))
		case CounterTypeUpperCase:
			r := rune(int('A') + cc.Value)
			w.Write(asciidoc.NewString(string(r)))
		}
	}
	switch cc.CounterType {
	case CounterTypeInteger:
		cc.Value += 1
	case CounterTypeUpperCase:
		if cc.Value < 'Z' {
			cc.Value += 1
		}
	case CounterTypeLowerCase:
		if cc.Value < 'z' {
			cc.Value += 1
		}
	}
	return nil
}
