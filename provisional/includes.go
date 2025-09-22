package provisional

import (
	"log/slog"
	"path/filepath"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func patchIncludes(s *spec.Specification, violations map[string][]Violation) (alteredDocs []*asciidoc.Document, err error) {
	docs := make(map[*asciidoc.Document][]Violation)
	for _, vs := range violations {
		for _, v := range vs {
			var includeViolation bool
			switch v.Entity.(type) {
			case *matter.Cluster:
				includeViolation = v.Type.Has(ViolationTypeNotIfDefd)
			}
			if !includeViolation {
				continue
			}
			if doc, ok := s.DocRefs[v.Entity]; ok {
				docs[doc] = append(docs[doc], v)
			} else {
				slog.Error("entity with missing doc", matter.LogEntity("entity", v.Entity))
			}
		}
	}
	for d := range docs {
		library, ok := s.LibraryForDocument(d)
		if !ok {
			continue
		}
		parents := library.Parents(d)
		for _, parent := range parents {
			sections := make(map[*asciidoc.Section]struct{})
			includes := make(map[*asciidoc.FileInclude]struct{})
			for _, el := range parent.Children() {
				switch el := el.(type) {
				case *asciidoc.Section:
					for _, titleElement := range el.Title {
						switch titleElement := titleElement.(type) {
						case *asciidoc.LinkMacro:
							if includeMatchesPath(s, d, titleElement, &titleElement.URL.Path, titleElement.URL.Path) {
								sections[el] = struct{}{}
							}
						}
					}
				case *asciidoc.FileInclude:
					if includeMatchesPath(s, d, el, el, el.Elements) {
						includes[el] = struct{}{}
					}
				}
			}
			if len(sections) == 0 && len(includes) == 0 {
				continue
			}
			children := make([]asciidoc.Element, 0, len(parent.Children()))
			for _, el := range parent.Children() {
				switch el := el.(type) {
				case *asciidoc.Section:
					if _, ok := sections[el]; ok {
						children = append(children, asciidoc.NewIfDef(inProgressAttributes, asciidoc.ConditionalUnionAny))
						children = append(children, el)
						children = append(children, asciidoc.NewEndIf(nil, asciidoc.ConditionalUnionAny))
					} else {
						children = append(children, el)

					}
				case *asciidoc.FileInclude:
					if _, ok := includes[el]; ok {
						children = append(children, asciidoc.NewIfDef(inProgressAttributes, asciidoc.ConditionalUnionAny))
						children = append(children, el)
						children = append(children, asciidoc.NewEndIf(nil, asciidoc.ConditionalUnionAny))
					} else {
						children = append(children, el)
					}
				default:
					children = append(children, el)
				}
			}
			parent.SetChildren(children)
			alteredDocs = append(alteredDocs, parent)
		}
	}
	return
}

func includeMatchesPath(s *spec.Specification, doc *asciidoc.Document, element log.Source, pathParent asciidoc.Parent, pathElements asciidoc.Elements) bool {
	path, err := asciidoc.RawReader.StringValue(pathParent, pathElements)
	if err != nil {
		slog.Warn("Error rendering file include path", slog.Any("error", err), log.Path("source", element))
		return false
	}
	path = filepath.Join(doc.Path.Dir(), path)
	specPath, err := spec.NewSpecPath(path, s.Root)
	if err != nil {
		slog.Warn("Error joining file include path", slog.Any("error", err), log.Path("source", element))
		return false
	}
	if specPath.Relative != doc.Path.Relative {
		return false
	}
	return true
}
