package spec

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/errata"
	"github.com/project-chip/alchemy/internal"
	"github.com/project-chip/alchemy/internal/pipeline"
)

func NewPreparseContext(path asciidoc.Path, specRoot string, attributes ...asciidoc.AttributeName) parse.PreParseContext {

	ac := &preparseContext{
		docPath:  path,
		rootPath: specRoot,
	}

	for _, a := range attributes {
		ac.Set(string(a), nil)
	}
	return ac
}

type preparseContext struct {
	docPath    asciidoc.Path
	rootPath   string
	attributes map[string]any
	counters   map[string]*asciidoc.CounterState
}

func (ac *preparseContext) IsSet(name string) bool {
	_, ok := ac.attributes[name]
	return ok
}

func (ac *preparseContext) Get(name string) any {
	return ac.attributes[name]
}

func (ac *preparseContext) Set(name string, value any) {
	if ac.attributes == nil {
		ac.attributes = make(map[string]any)
	}
	ac.attributes[name] = value
}

func (ac *preparseContext) Unset(name string) {
	delete(ac.attributes, name)
}

func (ac *preparseContext) GetCounterState(name string, initialValue string) (*asciidoc.CounterState, error) {
	if ac.counters == nil {
		ac.counters = make(map[string]*asciidoc.CounterState)
	}
	cc, ok := ac.counters[name]
	if ok {
		return cc, nil
	}
	cc = &asciidoc.CounterState{}
	ac.counters[name] = cc
	switch len(initialValue) {
	case 0:
		cc.Value = 1
		cc.CounterType = asciidoc.CounterTypeInteger
	case 1:
		r := initialValue[0]
		if r >= 'a' && r <= 'z' {
			cc.Value = int(r) - int('a')
			cc.CounterType = asciidoc.CounterTypeLowerCase
		} else if r >= 'A' && r <= 'Z' {
			cc.Value = int(r) - int('A')
			cc.CounterType = asciidoc.CounterTypeUpperCase
		} else {
			var err error
			cc.Value, err = strconv.Atoi(initialValue)
			if err != nil {
				return nil, err
			}
			cc.CounterType = asciidoc.CounterTypeInteger
		}
	default:
		var err error
		cc.Value, err = strconv.Atoi(initialValue)
		if err != nil {
			return nil, err
		}
		cc.CounterType = asciidoc.CounterTypeInteger
	}
	return cc, nil

}

func (ac *preparseContext) ResolvePath(root string, path string) (asciidoc.Path, error) {
	linkPath := filepath.Join(root, path)
	return NewSpecPath(linkPath, ac.rootPath)
}

func (ac *preparseContext) ShouldIncludeFile(path asciidoc.Path) bool {
	return errata.GetSpec(path.Relative).UtilityInclude
}

type PreParser struct {
	specRoot string

	attributes []asciidoc.AttributeName
}

func NewPreParser(specRoot string, attributes []asciidoc.AttributeName) (*PreParser, error) {

	return &PreParser{specRoot: specRoot, attributes: attributes}, nil
}

func (r PreParser) Name() string {
	return "Pre-parsing docs"
}

func (r PreParser) Process(cxt context.Context, input *pipeline.Data[*DocGroup], index int32, total int32) (outputs []*pipeline.Data[asciidoc.Reader], extras []*pipeline.Data[*DocGroup], err error) {
	var reader asciidoc.Reader
	reader, err = preparse(input.Content, input.Content.Root, r.specRoot, r.attributes)
	if err != nil {
		return
	}
	input.Content.Reader = reader
	outputs = append(outputs, &pipeline.Data[asciidoc.Reader]{Path: input.Content.Root.Path.Relative, Content: reader})
	return
}

type conditionalBlock struct {
	open     asciidoc.Element
	suppress bool
}

type preparseState struct {
	attributes map[string]any
	counters   map[string]*asciidoc.CounterState

	actions map[asciidoc.Element]*preparseAction
}

type preparseFileState struct {
	state *preparseState

	docPath  asciidoc.Path
	rootPath string
}

func (pps *preparseFileState) IsSet(name string) bool {
	_, ok := pps.state.attributes[name]
	return ok
}

func (pps *preparseFileState) Get(name string) any {
	return pps.state.attributes[name]
}

func (pps *preparseFileState) Set(name string, value any) {
	if pps.state.attributes == nil {
		pps.state.attributes = make(map[string]any)
	}
	pps.state.attributes[name] = value
}

func (pps *preparseFileState) Unset(name string) {
	delete(pps.state.attributes, name)
}

func (ac *preparseFileState) ResolvePath(root string, path string) (asciidoc.Path, error) {
	linkPath := filepath.Join(root, path)
	return NewSpecPath(linkPath, ac.rootPath)
}

func (ac *preparseFileState) GetCounterState(name string, initialValue string) (*asciidoc.CounterState, error) {
	if ac.state.counters == nil {
		ac.state.counters = make(map[string]*asciidoc.CounterState)
	}
	cc, ok := ac.state.counters[name]
	if ok {
		return cc, nil
	}
	cc = &asciidoc.CounterState{}
	ac.state.counters[name] = cc
	switch len(initialValue) {
	case 0:
		cc.Value = 1
		cc.CounterType = asciidoc.CounterTypeInteger
	case 1:
		r := initialValue[0]
		if r >= 'a' && r <= 'z' {
			cc.Value = int(r) - int('a')
			cc.CounterType = asciidoc.CounterTypeLowerCase
		} else if r >= 'A' && r <= 'Z' {
			cc.Value = int(r) - int('A')
			cc.CounterType = asciidoc.CounterTypeUpperCase
		} else {
			var err error
			cc.Value, err = strconv.Atoi(initialValue)
			if err != nil {
				return nil, err
			}
			cc.CounterType = asciidoc.CounterTypeInteger
		}
	default:
		var err error
		cc.Value, err = strconv.Atoi(initialValue)
		if err != nil {
			return nil, err
		}
		cc.CounterType = asciidoc.CounterTypeInteger
	}
	return cc, nil

}

func (ac *preparseFileState) ShouldIncludeFile(path asciidoc.Path) bool {
	return errata.GetSpec(path.Relative).UtilityInclude
}

func preparse(docGroup *DocGroup, doc *Doc, specRoot string, attributes []asciidoc.AttributeName) (reader asciidoc.Reader, err error) {
	pps := &preparseState{
		attributes: make(map[string]any),
		counters:   map[string]*asciidoc.CounterState{},
		actions:    map[asciidoc.Element]*preparseAction{},
	}

	for _, a := range attributes {
		pps.attributes[string(a)] = nil
	}

	ppfs := &preparseFileState{
		state:    pps,
		docPath:  doc.Path,
		rootPath: specRoot,
	}
	err = preparseFile(ppfs, docGroup, doc, doc, doc.Children())
	if err != nil {
		return
	}
	reader = &PreParseReader{elements: pps.actions}
	return
}

func preparseFile(pps *preparseFileState, docGroup *DocGroup, doc *Doc, parent asciidoc.Parent, els asciidoc.Elements) (err error) {
	var suppressStack internal.Stack[*conditionalBlock]
	var suppress bool
	var lastTableCell *asciidoc.TableCell
	var addToCell bool
	parse.Traverse(asciidoc.RawReader, parent, els, func(el asciidoc.Element, parent asciidoc.Parent, index int) (should parse.SearchShould) {
		var remove, replace bool
		var replaceElements asciidoc.Elements
		switch el := el.(type) {
		case *asciidoc.AttributeEntry:
			if !suppress {
				pps.Set(string(el.Name), el.Elements)
			}
			remove = true
		case *asciidoc.UserAttributeReference:
			val := pps.Get(string(el.Name()))
			if val != nil {
				switch attr := val.(type) {
				case string:
					replace = true
					replaceElements = asciidoc.Elements{asciidoc.NewString(attr)}
				case asciidoc.Elements:
					replace = true
					replaceElements = attr
				default:
					err = fmt.Errorf("unexpected user attribute type: %T", val)
					should = parse.SearchShouldStop
					return
				}
			}
		case *asciidoc.AttributeReset:
			if !suppress {
				pps.Unset(string(el.Name))
			}
			remove = true
		case *asciidoc.IfDef:
			suppressStack.Push(&conditionalBlock{suppress: suppress, open: el})
			suppress = suppress || !el.Eval(pps)
			remove = true
			addToCell = el.Inline
		case *asciidoc.IfNDef:
			suppressStack.Push(&conditionalBlock{suppress: suppress, open: el})
			suppress = suppress || !el.Eval(pps)
			remove = true
			addToCell = el.Inline
		case *asciidoc.IfEval:
			suppressStack.Push(&conditionalBlock{suppress: suppress, open: el})
			if !suppress {
				var include bool
				include, err = el.Eval(pps)
				if err != nil {
					should = parse.SearchShouldStop
					return
				}
				suppress = !include
			}
			remove = true
			addToCell = el.Inline
		case *asciidoc.IfDefBlock, *asciidoc.IfNDefBlock, *asciidoc.IfEvalBlock:
			err = fmt.Errorf("unexpected type in preparse: %T", el)
			should = parse.SearchShouldStop
			return
		case *asciidoc.EndIf:
			var ok bool
			var cb *conditionalBlock
			cb, ok = suppressStack.Pop()
			if !ok {
				err = fmt.Errorf("unexpected endif")
				should = parse.SearchShouldStop
				return
			}
			suppress = cb.suppress
			el.Open = cb.open
			remove = true
			addToCell = false
		case *asciidoc.InlineIfDef:
			if !suppress {
				if el.Eval(pps) {
					replace = true
					replaceElements = el.Elements
					err = preparseFile(pps, docGroup, doc, el, el.Elements)
					if err != nil {
						should = parse.SearchShouldStop
						return
					}
				} else {
					remove = true
					addToCell = false
				}
			}
		case *asciidoc.InlineIfNDef:
			if el.Eval(pps) {
				pps.state.actions[el] = &preparseAction{replace: el.Elements}
				err = preparseFile(pps, docGroup, doc, el, el.Elements)
				if err != nil {
					should = parse.SearchShouldStop
					return
				}
			} else {
				remove = true
				addToCell = false
			}
		case *asciidoc.TableCell:
			lastTableCell = el
			if suppress {
				remove = true
			}
		case *asciidoc.FileInclude:
			if suppress {
				remove = true
				break
			}
			var path asciidoc.Path
			var shouldInclude bool
			path, err = includePath(pps, doc, el)
			if err != nil {
				should = parse.SearchShouldStop
				return
			}

			shouldInclude = pps.ShouldIncludeFile(path)

			var includedDoc *Doc
			includedDoc, err = readIncludeFile(docGroup, path, parent, shouldInclude)
			if err != nil {
				should = parse.SearchShouldStop
				return
			}
			if shouldInclude {
				if addToCell {
					tableCellAction, ok := pps.state.actions[lastTableCell]
					if !ok {
						tableCellAction = &preparseAction{}
						pps.state.actions[lastTableCell] = tableCellAction
					}
					tableCellAction.append = append(tableCellAction.append, includedDoc.Children()...)
					remove = true
				} else {
					replace = true
					replaceElements = includedDoc.Children()
				}
			}
			ippfs := &preparseFileState{
				state:    pps.state,
				docPath:  path,
				rootPath: pps.rootPath,
			}
			err = preparseFile(ippfs, docGroup, includedDoc, el, includedDoc.Children())
			if err != nil {
				should = parse.SearchShouldStop
				return
			}
		case *asciidoc.Section:
			if suppress {
				remove = true
			} else {
				iterator := &PreParseReader{elements: pps.state.actions}
				var title strings.Builder
				err = buildSectionTitle(pps, el, iterator, &title, el.Title...)
				if err != nil {
					should = parse.SearchShouldStop
					return
				}
				doc.SetSectionName(el, title.String())
			}
		default:
			if suppress {
				remove = true
			} else if addToCell {
				if lastTableCell != nil {
					tableCellAction, ok := pps.state.actions[lastTableCell]
					if !ok {
						tableCellAction = &preparseAction{}
						pps.state.actions[lastTableCell] = tableCellAction
					}
					tableCellAction.append = append(tableCellAction.append, el)
					remove = true
				}
			}

		}
		if remove {
			pps.state.actions[el] = &preparseAction{remove: true}
		} else if replace {
			pps.state.actions[el] = &preparseAction{replace: replaceElements}
		}
		return parse.SearchShouldContinue
	})
	return
}

type preparseAction struct {
	remove  bool
	replace asciidoc.Elements
	append  asciidoc.Elements
}
type PreParseReader struct {
	elements map[asciidoc.Element]*preparseAction
}

func (ppi *PreParseReader) Iterate(parent asciidoc.Parent, elements asciidoc.Elements) asciidoc.ElementIterator {
	return func(yield func(asciidoc.Element) bool) {
		for _, el := range elements {
			action, ok := ppi.elements[el]
			if !ok {
				if !yield(el) {
					return
				}
			} else {
				if action.remove {
					continue
				}
				if len(action.replace) > 0 {
					for _, el := range action.replace {
						if !yield(el) {
							return
						}
					}
				}
			}
		}
		switch parent := parent.(type) {
		case asciidoc.Element:
			if action, ok := ppi.elements[parent]; ok {
				if len(action.append) > 0 {
					for _, el := range action.append {
						if !yield(el) {
							return
						}
					}
				}
			}
		}
	}
}

func (ppi *PreParseReader) Count(elements asciidoc.Elements) int {
	count := 0
	for _, el := range elements {
		action, ok := ppi.elements[el]
		if ok {
			if action.remove {
				continue
			}
			if len(action.replace) > 0 {
				count += ppi.Count(action.replace)
				continue
			}
		}
		count++
	}
	return count
}

func renderPreParsedDoc(els asciidoc.Elements) (string, error) {
	var sb strings.Builder
	for _, el := range els {
		switch el := el.(type) {
		case *asciidoc.String:
			sb.WriteString(el.Value)
		case *asciidoc.NewLine:
			sb.WriteRune('\n')
		case *asciidoc.EmptyLine:
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
