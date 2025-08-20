package spec

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/pipeline"
	"github.com/project-chip/alchemy/internal/text"
)

/*
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

	func (r PreParser) Process(cxt context.Context, input *pipeline.Data[*Library], index int32, total int32) (outputs []*pipeline.Data[asciidoc.Reader], extras []*pipeline.Data[*Library], err error) {
		var reader asciidoc.Reader
		reader, err = preparse(input.Content, input.Content.Root, r.specRoot, r.attributes)
		if err != nil {
			return
		}
		outputs = append(outputs, &pipeline.Data[asciidoc.Reader]{Path: input.Content.Root.Path.Relative, Content: reader})
		return
	}
*/
type conditionalBlock struct {
	open     asciidoc.Element
	suppress bool
}

type preparseState struct {
	attributes map[string]any
	counters   map[string]*asciidoc.CounterState

	overlays elementOverlays

	lastSection      *asciidoc.Section
	lastSectionLevel int
	//lastSectionAction *elementOverlay
}

type preparseFileState struct {
	state *preparseState

	docPath  asciidoc.Path
	rootPath string

	sectionLevelOffset int
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
	switch path.Relative {
	case "templates/DiscoBallCluster.adoc", "templates/DiscoBallDeviceType.adoc":
		return false
	default:
		return true
	}
}

type PreParseReader struct {
	overlays elementOverlays
}

func (ppi *PreParseReader) Iterate(parent asciidoc.Parent, elements asciidoc.Elements) asciidoc.ElementIterator {
	return func(yield func(asciidoc.Element) bool) {
		if !ppi.iterate(elements, yield) {
			return
		}
		var parentOverlay *elementOverlay
		switch parent := parent.(type) {
		case asciidoc.Element:
			parentOverlay = ppi.overlays[parent]
		}
		if parentOverlay != nil && len(parentOverlay.append) > 0 {
			for _, el := range parentOverlay.append {
				if !yield(el) {
					return
				}
			}
		}
	}
}

func (ppi *PreParseReader) iterate(elements asciidoc.Elements, yield func(asciidoc.Element) bool) bool {
	for _, el := range elements {
		elementOverlay, ok := ppi.overlays[el]
		if !ok {
			if !yield(el) {
				return false
			}
		} else {
			if elementOverlay.action.Remove() {
				continue
			}
			if elementOverlay.action.Replace() {
				if !ppi.iterate(elementOverlay.replace, yield) {
					return false
				}
				continue
			}
			if !yield(el) {
				return false
			}
		}
	}
	return true
}

func (ppi *PreParseReader) StringValue(parent asciidoc.Parent, elements asciidoc.Elements) (string, error) {
	var sb strings.Builder
	for el := range ppi.Iterate(parent, elements) {
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
		case *asciidoc.Superscript:
			// This is usually an asterisk, and should be ignored
		default:
			switch el := el.(type) {
			case log.Source:
				return "", newGenericParseError(el, "unexpected type rendering string value: %T", el)
			default:
				return "", fmt.Errorf("unexpected type rendering string value: %T", el)
			}
		}
	}
	return sb.String(), nil
}

func (ppi *PreParseReader) Parent(child asciidoc.ChildElement) asciidoc.Element {
	if overlay, ok := ppi.overlays[child]; ok && overlay.action.OverrideParent() {
		return overlay.parent
	}
	return child.Parent()
}

func (ppi *PreParseReader) Children(parent asciidoc.ParentElement) asciidoc.Elements {
	if overlay, ok := ppi.overlays[parent]; ok && overlay.action.OverrideChildren() {
		return overlay.children
	}
	return parent.Children()
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

func (library *Library) Preparse(specRoot string, attributes []asciidoc.AttributeName) (err error) {
	pps := &preparseState{
		attributes: make(map[string]any),
		counters:   map[string]*asciidoc.CounterState{},
		overlays:   map[asciidoc.Element]*elementOverlay{},
	}

	for _, a := range attributes {
		pps.attributes[string(a)] = nil
	}

	ppfs := &preparseFileState{
		state:    pps,
		docPath:  library.Root.Path,
		rootPath: specRoot,
	}
	err = library.preparseFile(ppfs, library.Root, library.Root, library.Root.Children())
	if err != nil {
		return
	}
	library.buildDoc(ppfs, library.Root)
	library.Reader = &PreParseReader{overlays: pps.overlays}
	return
}

/*
var current asciidoc.ElementList
	current = d
	var lastSection *asciidoc.Section
	for _, el := range els {
		switch el := el.(type) {
		case *asciidoc.Section:
			if lastSection != nil {
				if el.Level > lastSection.Level {
					lastSection.AddChildSection(el)
					current.Append(el)
				} else if el.Level <= lastSection.Level {
					parent := lastSection.ParentSection()
					var found bool
					for parent != nil {
						if el.Level > parent.Level {
							parent.Append(el)
							parent.AddChildSection(el)
							found = true
							break
						}
						parent = parent.ParentSection()
					}
					if !found { // No parent smaller
						d.Append(el)
					}
				}
			} else {
				current.Append(el)
			}

			lastSection = el
			current = el
		default:
			current.Append(el)
		}
	}
*/

func (library *Library) preparseFile(pps *preparseFileState, d *asciidoc.Document, parent asciidoc.ParentElement, els asciidoc.Elements) (err error) {
	var suppressStack internal.Stack[*conditionalBlock]
	var suppress bool
	var lastTableCell *asciidoc.TableCell
	var addToCell bool
	library.Docs = append(library.Docs, d)

	parse.Traverse(d, asciidoc.RawReader, parent, els, func(doc *asciidoc.Document, el asciidoc.Element, parent asciidoc.ParentElement, index int) (should parse.SearchShould) {
		var remove, replace bool
		var replaceElements asciidoc.Elements
		switch el := el.(type) {
		case *asciidoc.Section:
			iterator := &PreParseReader{overlays: pps.state.overlays}
			var title strings.Builder
			err = buildSectionTitle(pps, el, iterator, &title, el.Title...)
			if err != nil {
				should = parse.SearchShouldStop
				return
			}
			library.SetSectionName(el, title.String())
			sectionLevel := el.Level + pps.sectionLevelOffset
			library.SetSectionLevel(el, sectionLevel)
			if suppress {
				remove = true
			}
		/*case *asciidoc.Section:
		if suppress {
			remove = true
		} else {
			iterator := &PreParseReader{overlays: pps.state.overlays}
			var title strings.Builder
			err = buildSectionTitle(pps, el, iterator, &title, el.Title...)
			if err != nil {
				should = parse.SearchShouldStop
				return
			}
			library.SetSectionName(el, title.String())
			sectionLevel := el.Level + pps.sectionLevelOffset
			library.SetSectionLevel(el, sectionLevel)
			//slog.Info("section", "name", title.String(), "level", el.Level, "offset", pps.sectionLevelOffset, "level", sectionLevel)
			if pps.state.lastSection != nil {
				//	slog.Info("lastSection", "name", library.SectionName(pps.state.lastSection), "lastLevel", pps.state.lastSectionLevel)
				if sectionLevel > pps.state.lastSectionLevel {
					//slog.Info("adding section as child of last section", "name", title.String(), "parent", library.SectionName(pps.state.lastSection), log.Path("source", el))
					pps.state.overlays.appendChild(pps.state.lastSection, el)
					pps.state.overlays.setParent(el, pps.state.lastSection)
				} else if sectionLevel <= pps.state.lastSectionLevel {
					var parentSection *asciidoc.Section
					parent := iterator.Parent(pps.state.lastSection)
					var found bool
					for parent != nil {
						var ok bool
						parentSection, ok = parent.(*asciidoc.Section)
						if !ok {
							break
						}
						if sectionLevel > library.SectionLevel(parentSection) {
							found = true
							break
						}
						parent = iterator.Parent(parentSection)
					}
					if found {
						//	slog.Info("adding section as child of parent", "name", title.String(), "parent", library.SectionName(parentSection), log.Path("source", el))
						pps.state.overlays.appendChild(parentSection, el)
						pps.state.overlays.setParent(el, parentSection)
					} else { // No parent smaller
						//slog.Info("adding section as child of doc", "name", title.String(), log.Path("source", el))
						pps.state.overlays.appendChild(library.Root, el)
						pps.state.lastSection = nil
						//pps.state.overlays.setParent(el, library.Root)
					}
				}
			} else {
				//slog.Info("adding root section as child of doc", "name", title.String(), log.Path("source", el))
				pps.state.overlays.appendChild(library.Root, el)
				//pps.state.overlays.setParent(el, library.Root)
			}
			pps.state.lastSection = el
			pps.state.lastSectionLevel = sectionLevel
			return parse.SearchShouldContinue
		}*/
		case *asciidoc.AttributeEntry:
			if !suppress {
				pps.Set(string(el.Name), el.Elements)
				switch el.Name {
				case "leveloffset":
					var leveloffset int
					var relative bool
					leveloffset, relative, err = parseLevelOffset(el, el.Elements)
					if err != nil {
						return parse.SearchShouldStop
					}
					if relative {
						pps.sectionLevelOffset += leveloffset
					} else {
						pps.sectionLevelOffset = leveloffset
					}
					//	slog.Info("SECTION OFFSET", "now", pps.sectionLevelOffset)
				}
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
			//slog.Info("pushing ifdef", log.Path("source", el), "index", index)
			suppressStack.Push(&conditionalBlock{suppress: suppress, open: el})
			suppress = suppress || !el.Eval(pps)
			remove = true
			addToCell = el.Inline
		case *asciidoc.IfNDef:
			//slog.Info("pushing ifmdef", log.Path("source", el), "index", index)
			suppressStack.Push(&conditionalBlock{suppress: suppress, open: el})
			suppress = suppress || !el.Eval(pps)
			remove = true
			addToCell = el.Inline
		case *asciidoc.IfEval:
			//slog.Info("pushing ifeval", log.Path("source", el), "index", index)
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
			//slog.Info("popping endif", log.Path("source", el))
			cb, ok = suppressStack.Pop()
			if !ok {
				err = newGenericParseError(el, "unexpected endif")
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
					err = library.preparseFile(pps, doc, el, el.Elements)
					if err != nil {
						should = parse.SearchShouldStop
						return
					}
				} else {
					addToCell = false
				}
			}
			remove = true
		case *asciidoc.InlineIfNDef:
			if el.Eval(pps) {
				pps.state.overlays[el] = &elementOverlay{replace: el.Elements}
				err = library.preparseFile(pps, doc, el, el.Elements)
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
			//var shouldInclude bool
			path, err = includePath(pps, el.Document(), el)
			if err != nil {
				should = parse.SearchShouldStop
				return
			}

			sectionLevelOffset := pps.sectionLevelOffset

			for _, attr := range el.Attributes() {
				switch attr := attr.(type) {
				case *asciidoc.NamedAttribute:
					switch attr.Name {
					case "leveloffset":
						var leveloffset int
						var relative bool
						leveloffset, relative, err = parseLevelOffset(attr, attr.Val)
						if err != nil {
							return parse.SearchShouldStop
						}
						if relative {
							sectionLevelOffset += leveloffset
						} else {
							sectionLevelOffset = leveloffset
						}
						//slog.Info("SECTION OFFSET", "include", sectionLevelOffset)
					}
				}
			}

			if !pps.ShouldIncludeFile(path) {
				remove = true
				break
			}

			var includedDoc *asciidoc.Document
			//slog.Info("including file", "lib", library.Root.Path.Relative, "from", el.Document().Path.Dir(), "path", path.Relative, log.Path("source", el))

			includedDoc, err = library.cache.include(path, parent)
			if err != nil {
				should = parse.SearchShouldStop
				return
			}
			if addToCell {
				tableCellOverlay, ok := pps.state.overlays[lastTableCell]
				if !ok {
					tableCellOverlay = &elementOverlay{}
					pps.state.overlays[lastTableCell] = tableCellOverlay
				}
				tableCellOverlay.action |= overlayActionAppendElements
				tableCellOverlay.append = append(tableCellOverlay.append, includedDoc.Children()...)
				remove = true
			} else {
				replace = true
				replaceElements = includedDoc.Children()
			}

			library.parents[includedDoc] = append(library.parents[includedDoc], doc)
			library.children[doc] = append(library.children[doc], includedDoc)

			ippfs := &preparseFileState{
				state:              pps.state,
				docPath:            path,
				rootPath:           pps.rootPath,
				sectionLevelOffset: sectionLevelOffset,
			}
			err = library.preparseFile(ippfs, includedDoc, el, includedDoc.Children())
			if err != nil {
				should = parse.SearchShouldStop
				return
			}

		default:
			if suppress {
				remove = true
			} else if addToCell {
				if lastTableCell != nil {
					pps.state.overlays.append(lastTableCell, el)
					remove = true
				}
			}

		}
		if remove {
			pps.state.overlays.remove(el)
		} else if replace {
			pps.state.overlays.replace(el, replaceElements)
		} /* else {
			if pps.state.lastSection != nil {
				pps.state.overlays.appendChild(pps.state.lastSection, el)
			} else {
				pps.state.overlays.appendChild(library.Root, el)
			}
		}*/
		return parse.SearchShouldContinue
	})
	return
}

func (l *Library) buildDoc(pps *preparseFileState, d *asciidoc.Document) {
	iterator := &PreParseReader{overlays: pps.state.overlays}
	for el := range parse.Skim[asciidoc.Element](iterator, d, iterator.Children(d)) {
		switch el := el.(type) {
		case *asciidoc.Section:
			sectionLevel := l.SectionLevel(el)
			if pps.state.lastSection != nil {
				//		slog.Info("lastSection", "name", l.SectionName(pps.state.lastSection), "lastLevel", pps.state.lastSectionLevel)
				if sectionLevel > pps.state.lastSectionLevel {
					//	slog.Info("adding section as child of last section", "name", l.SectionName(el), "parent", l.SectionName(pps.state.lastSection), log.Path("source", el))
					pps.state.overlays.appendChild(pps.state.lastSection, el)
					pps.state.overlays.setParent(el, pps.state.lastSection)
				} else if sectionLevel <= pps.state.lastSectionLevel {
					var parentSection *asciidoc.Section
					parent := iterator.Parent(pps.state.lastSection)
					var found bool
					for parent != nil {
						var ok bool
						parentSection, ok = parent.(*asciidoc.Section)
						if !ok {
							break
						}
						if sectionLevel > l.SectionLevel(parentSection) {
							found = true
							break
						}
						parent = iterator.Parent(parentSection)
					}
					if found {
						//	slog.Info("adding section as child of parent", "name", l.SectionName(el), "parent", l.SectionName(parentSection), log.Path("source", el))
						pps.state.overlays.appendChild(parentSection, el)
						pps.state.overlays.setParent(el, parentSection)
					} else { // No parent smaller
						//slog.Info("adding section as child of doc", "name", l.SectionName(el), log.Path("source", el))
						pps.state.overlays.appendChild(l.Root, el)
						pps.state.lastSection = nil
						//pps.state.overlays.setParent(el, library.Root)
					}
				}
			} else {
				//slog.Info("adding root section as child of doc", "name", title.String(), log.Path("source", el))
				pps.state.overlays.appendChild(l.Root, el)
				//pps.state.overlays.setParent(el, library.Root)
			}
			pps.state.lastSection = el
			pps.state.lastSectionLevel = sectionLevel
		default:
			if pps.state.lastSection != nil {
				pps.state.overlays.appendChild(pps.state.lastSection, el)
			} else {
				pps.state.overlays.appendChild(l.Root, el)
			}
		}
	}
	/*parse.Search(d, iterator, d, iterator.Children(d), func(doc *asciidoc.Document, section *asciidoc.Section, parent asciidoc.ParentElement, offset int) parse.SearchShould {
		sectionLevel := l.SectionLevel(section)
		if pps.state.lastSection != nil {
			//	slog.Info("lastSection", "name", library.SectionName(pps.state.lastSection), "lastLevel", pps.state.lastSectionLevel)
			if sectionLevel > pps.state.lastSectionLevel {
				//slog.Info("adding section as child of last section", "name", title.String(), "parent", library.SectionName(pps.state.lastSection), log.Path("source", el))
				pps.state.overlays.appendChild(pps.state.lastSection, section)
				pps.state.overlays.setParent(el, pps.state.lastSection)
			} else if sectionLevel <= pps.state.lastSectionLevel {
				var parentSection *asciidoc.Section
				parent := iterator.Parent(pps.state.lastSection)
				var found bool
				for parent != nil {
					var ok bool
					parentSection, ok = parent.(*asciidoc.Section)
					if !ok {
						break
					}
					if sectionLevel > library.SectionLevel(parentSection) {
						found = true
						break
					}
					parent = iterator.Parent(parentSection)
				}
				if found {
					//	slog.Info("adding section as child of parent", "name", title.String(), "parent", library.SectionName(parentSection), log.Path("source", el))
					pps.state.overlays.appendChild(parentSection, el)
					pps.state.overlays.setParent(el, parentSection)
				} else { // No parent smaller
					//slog.Info("adding section as child of doc", "name", title.String(), log.Path("source", el))
					pps.state.overlays.appendChild(library.Root, el)
					pps.state.lastSection = nil
					//pps.state.overlays.setParent(el, library.Root)
				}
			}
		} else {
			//slog.Info("adding root section as child of doc", "name", title.String(), log.Path("source", el))
			pps.state.overlays.appendChild(library.Root, el)
			//pps.state.overlays.setParent(el, library.Root)
		}
		pps.state.lastSection = el
		pps.state.lastSectionLevel = sectionLevel
		return parse.SearchShouldContinue
	})*/
}

func parseLevelOffset(el asciidoc.Parent, elements asciidoc.Elements) (leveloffset int, relative bool, err error) {
	var val string
	val, err = asciidoc.RawReader.StringValue(el, elements)
	if err != nil {
		return
	}
	leveloffset, relative, err = text.ParseRelativeNumber(val)
	return
}

type LibraryParser struct {
	specRoot string

	attributes []asciidoc.AttributeName
}

func NewLibraryParser(specRoot string, attributes []asciidoc.AttributeName) (*LibraryParser, error) {

	return &LibraryParser{specRoot: specRoot, attributes: attributes}, nil
}

func (r LibraryParser) Name() string {
	return "Pre-parsing library docs"
}

func (r LibraryParser) Process(cxt context.Context, input *pipeline.Data[*Library], index int32, total int32) (outputs []*pipeline.Data[*Document], extras []*pipeline.Data[*Library], err error) {
	err = input.Content.Preparse(r.specRoot, r.attributes)
	return
}

/*
	func preparse(docGroup *Library, doc *asciidoc.Document, specRoot string, attributes []asciidoc.AttributeName) (reader asciidoc.Reader, err error) {
		pps := &preparseState{
			attributes:    make(map[string]any),
			counters:      map[string]*asciidoc.CounterState{},
			actions:       map[asciidoc.Element]*preparseAction{},
			includedFiles: map[string]*asciidoc.Document{},
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

	func preparseFile(pps *preparseFileState, docGroup *Library, doc *asciidoc.Document, parent asciidoc.ParentElement, els asciidoc.Elements) (err error) {
		var suppressStack internal.Stack[*conditionalBlock]
		var suppress bool
		var lastTableCell *asciidoc.TableCell
		var addToCell bool
		parse.Traverse(asciidoc.RawReader, parent, els, func(doc *asciidoc.Document, el asciidoc.Element, parent asciidoc.ParentElement, index int) (should parse.SearchShould) {
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

				var includedDoc *asciidoc.Document
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
					docGroup.SetSectionName(el, title.String())
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
*/
