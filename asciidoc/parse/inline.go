package parse

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal"
	"github.com/project-chip/alchemy/internal/parse"
)

func Inline(context PreParseContext, path string, reader io.Reader, opts ...Option) (doc *asciidoc.Document, err error) {
	var vals any
	vals, err = ParseReader(path, reader, opts...)
	if err != nil {
		slog.Error("error parsing file", slog.String("path", path), slog.Any("error", err))
		return nil, err
	}
	var set asciidoc.Set
	set, ok := vals.(asciidoc.Set)
	if !ok {
		return nil, fmt.Errorf("unexpected type in UnifiedParse: %T", vals)
	}
	var suppressStack internal.Stack[bool]
	var suppress bool
	var lastTableCell *asciidoc.TableCell
	var addToCell bool

	parse.Filter(&set, func(parent parse.HasElements, el asciidoc.Element) (remove bool, replace asciidoc.Set, shortCircuit bool) {
		switch el := el.(type) {
		case *asciidoc.AttributeEntry:
			if !suppress {
				context.Set(string(el.Name), el.Set)
			}
			remove = true
		case *asciidoc.AttributeReset:
			if !suppress {
				context.Unset(string(el.Name))
			}
			remove = true
		case *asciidoc.IfDef:
			suppressStack.Push(suppress)
			suppress = suppress || !el.Eval(context)
			remove = true
			addToCell = el.Inline
		case *asciidoc.IfNDef:
			suppressStack.Push(suppress)
			suppress = suppress || !el.Eval(context)
			remove = true
			addToCell = el.Inline
		case *asciidoc.IfEval:
			suppressStack.Push(suppress)
			if !suppress {
				var include bool
				include, err = el.Eval(context)
				if err != nil {
					shortCircuit = true
					return
				}
				suppress = !include
			}
			remove = true
			addToCell = el.Inline
		case *asciidoc.IfDefBlock, *asciidoc.IfNDefBlock, *asciidoc.IfEvalBlock:
			err = fmt.Errorf("unexpected type in preparse: %T", el)
			shortCircuit = true
		case *asciidoc.EndIf:
			var ok bool
			suppress, ok = suppressStack.Pop()
			if !ok {
				err = fmt.Errorf("unexpected endif")
				shortCircuit = true
				return
			}
			remove = true
			addToCell = false
		case *asciidoc.InlineIfDef:
			if el.Eval(context) {
				replace = el.Set
			}
		case *asciidoc.InlineIfNDef:
			if el.Eval(context) {
				replace = el.Set
			}
		case *asciidoc.TableCell:
			lastTableCell = el
			remove = suppress
		case *asciidoc.FileInclude:
			remove = suppress
			if !remove {
				replace, err = newIncludeFile(context, parent, el)
				if err != nil {
					shortCircuit = true
					return
				}
				if addToCell {
					if lastTableCell != nil {
						lastTableCell.Append(replace...)
						replace = nil
					}
				}
			}
		default:

			if suppress {
				//				slog.Info("Removing", log.Type("type", el))
			} else {
				//				slog.Info("Keeping", log.Type("type", el))
			}
			remove = suppress
			if !suppress && addToCell {
				if lastTableCell != nil {
					lastTableCell.Append(el)
					remove = true
				}
			}
		}
		return
	})

	if err != nil {
		return
	}
	doc, err = setToDoc(set)
	return
}

func newIncludeFile(context PreParseContext, parent parse.HasElements, include *asciidoc.FileInclude) (elements asciidoc.Set, err error) {
	rawPathWriter := asciidoc.NewWriter(nil)
	err = preparseElements(context, asciidoc.NewReader(include.Set), rawPathWriter)

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
	if !context.ShouldIncludeFile(path) {
		return
	}

	var contents *os.File
	contents, err = os.Open(path.Absolute)
	if err != nil {
		return
	}
	defer contents.Close()

	var options []Option
	table, isTable := parent.(*asciidoc.Table)
	if isTable {
		options = append(options, Entrypoint("IncludedTableElements"))
		options = append(options, GlobalStore("table", table))
	}
	var doc *asciidoc.Document
	doc, err = Inline(context, path.Relative, contents, options...)
	if err != nil {
		return
	}
	elements = doc.Elements()
	return
}
