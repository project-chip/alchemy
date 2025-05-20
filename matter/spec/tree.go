package spec

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/log"
	"github.com/project-chip/alchemy/internal/parse"
)

func buildTree(specRoot string, docs []*Doc) error {

	tree := make(map[*Doc][]*asciidoc.FileInclude)
	docPaths := make(map[string]*Doc)

	for _, doc := range docs {

		path := doc.Path
		docPaths[path.Absolute] = doc

		parse.Traverse(doc, doc.Elements(), func(link *asciidoc.FileInclude, parent parse.HasElements, index int) parse.SearchShould {
			tree[doc] = append(tree[doc], link)
			return parse.SearchShouldContinue
		})
	}

	for doc, children := range tree {
		for _, link := range children {
			var p strings.Builder
			buildDataTypeString(doc, link.Set, &p)
			linkFullPath := filepath.Join(doc.Path.Dir(), p.String())
			linkPath, err := asciidoc.NewPath(linkFullPath, specRoot)
			if err != nil {
				return err
			}
			slog.Debug("Link path", log.Path("from", doc.Path), slog.String("to", p.String()), log.Path("linkPath", linkPath))
			if cd, ok := docPaths[linkPath.Absolute]; ok {
				cd.addParent(doc)
				doc.addChild(cd)
			} else {
				if strings.HasPrefix(linkPath.Relative, "src/") {
					slog.Warn("unknown child path", log.Element("parent", doc.Path, link), "child", linkPath.Relative)
				}
			}
		}
	}
	return nil
}

func dumpTree(r *Doc, indent int) {
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Printf("%s (%s)\n", r.Path.Absolute, r.Path.Relative)
	for _, c := range r.children {
		dumpTree(c, indent+1)
	}
}
