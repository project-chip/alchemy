package overlay

import (
	"fmt"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal"
	"github.com/project-chip/alchemy/internal/text"
)

type Variables interface {
	IsSet(name string) bool
	Get(name string) any
	Set(name string, value any)
	Unset(name string)
}

type OverlayContext interface {
	AddDocument(doc *asciidoc.Document)
	ResolvePath(root string, path string) (asciidoc.Path, error)
	ShouldIncludeFile(path asciidoc.Path) bool
	IncludeFile(path asciidoc.Path, parent asciidoc.Parent) (doc *asciidoc.Document, err error)

	MakeSectionName(reader asciidoc.Reader, section *asciidoc.Section, variables Variables) (string, error)
	SectionName(section *asciidoc.Section) string
	SetSectionName(section *asciidoc.Section, name string)
	SectionLevel(section *asciidoc.Section) int
	SetSectionLevel(section *asciidoc.Section, level int)
	SetParent(parent *asciidoc.Document, child *asciidoc.Document)
}

func Build(cxt OverlayContext, rootDoc *asciidoc.Document, rootPath string, attributes []asciidoc.AttributeName) (reader *Reader, err error) {
	pps := &overlayState{
		attributes: make(map[string]any),
		counters:   map[string]*asciidoc.CounterState{},
		overlays:   map[asciidoc.Element]*elementOverlay{},
	}

	for _, a := range attributes {
		pps.attributes[string(a)] = nil
	}

	ppfs := &overlayFileState{
		state:    pps,
		rootDoc:  rootDoc,
		docPath:  rootDoc.Path,
		rootPath: rootPath,
	}
	err = preparseFile(cxt, ppfs, rootDoc, rootDoc, rootDoc.Children())
	if err != nil {
		return
	}
	reader = &Reader{overlays: pps.overlays}
	buildDocTree(cxt, ppfs, rootDoc)
	return
}

func preparseFile(cxt OverlayContext, pps *overlayFileState, d *asciidoc.Document, parent asciidoc.ParentElement, els asciidoc.Elements) (err error) {
	var suppressStack internal.Stack[*conditionalBlock]
	var suppress bool
	var lastTableCell *asciidoc.TableCell
	var addToCell bool
	cxt.AddDocument(d)

	parse.Traverse(d, asciidoc.RawReader, parent, els, func(doc *asciidoc.Document, el asciidoc.Element, parent asciidoc.ParentElement, index int) (should parse.SearchShould) {
		var remove, replace bool
		var replaceElements asciidoc.Elements
		switch el := el.(type) {
		case *asciidoc.Section:
			if suppress {
				remove = true
			} else {
				var sectionName string
				sectionName, err = cxt.MakeSectionName(&Reader{overlays: pps.state.overlays}, el, pps)
				if err != nil {
					should = parse.SearchShouldStop
					return
				}
				cxt.SetSectionName(el, sectionName)
				sectionLevel := el.Level + pps.sectionLevelOffset
				cxt.SetSectionLevel(el, sectionLevel)
			}

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
				err = &OverlayError{
					error:  fmt.Errorf("unexpected endif"),
					Source: el,
				}
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
					err = preparseFile(cxt, pps, doc, el, el.Elements)
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
				replace = true
				replaceElements = el.Elements
				err = preparseFile(cxt, pps, doc, el, el.Elements)
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

			var rawPath string
			rawPath, err = renderSimpleElements(el.Children())
			if err != nil {
				should = parse.SearchShouldStop
				return
			}

			var path asciidoc.Path

			path, err = cxt.ResolvePath(doc.Path.Dir(), rawPath)
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
					}
				}
			}

			if !pps.ShouldIncludeFile(path) {
				remove = true
				break
			}

			var includedDoc *asciidoc.Document
			includedDoc, err = cxt.IncludeFile(path, parent)
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

			cxt.SetParent(doc, includedDoc)

			ippfs := &overlayFileState{
				state:              pps.state,
				docPath:            path,
				rootPath:           pps.rootPath,
				sectionLevelOffset: sectionLevelOffset,
			}
			err = preparseFile(cxt, ippfs, includedDoc, el, includedDoc.Children())
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
		}
		return parse.SearchShouldContinue
	})
	return
}

func buildDocTree(cxt OverlayContext, pps *overlayFileState, d *asciidoc.Document) {
	iterator := &Reader{overlays: pps.state.overlays}
	for el := range parse.Skim[asciidoc.Element](iterator, d, iterator.Children(d)) {
		switch el := el.(type) {
		case *asciidoc.Section:
			sectionLevel := cxt.SectionLevel(el)
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
						if sectionLevel > cxt.SectionLevel(parentSection) {
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
						pps.state.overlays.appendChild(pps.rootDoc, el)
						pps.state.lastSection = nil
						//pps.state.overlays.setParent(el, library.Root)
					}
				}
			} else {
				//slog.Info("adding root section as child of doc", "name", title.String(), log.Path("source", el))
				pps.state.overlays.appendChild(pps.rootDoc, el)
				//pps.state.overlays.setParent(el, library.Root)
			}
			pps.state.lastSection = el
			pps.state.lastSectionLevel = sectionLevel
		default:
			if pps.state.lastSection != nil {
				pps.state.overlays.appendChild(pps.state.lastSection, el)
			} else {
				pps.state.overlays.appendChild(pps.rootDoc, el)
			}
		}
	}
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
