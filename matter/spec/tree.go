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

func buildTree(docs []*Doc) {

	tree := make(map[*Doc][]*asciidoc.FileInclude)
	docPaths := make(map[string]*Doc)

	for _, doc := range docs {

		path := doc.Path
		docPaths[path.Absolute] = doc

		parse.Search(doc.Elements(), func(link *asciidoc.FileInclude) parse.SearchShould {
			tree[doc] = append(tree[doc], link)
			return parse.SearchShouldContinue
		})
	}

	for doc, children := range tree {
		for _, link := range children {
			var p strings.Builder
			buildDataTypeString(doc, link.Set, &p)
			linkPath := filepath.Join(doc.Path.Dir(), p.String())
			slog.Debug("Link path", log.Path("from", doc.Path), slog.String("to", p.String()), slog.String("linkPath", linkPath))
			if cd, ok := docPaths[linkPath]; ok {
				cd.addParent(doc)
				doc.addChild(cd)
			} else {
				slog.Debug("unknown child path", log.Element("parent", doc.Path, link), "child", linkPath)
			}
		}
	}
}

func dumpTree(r *Doc, indent int) {
	fmt.Print(strings.Repeat("\t", indent))
	fmt.Printf("%s (%s)\n", r.Path.Absolute, r.Path.Relative)
	for _, c := range r.children {
		dumpTree(c, indent+1)
	}
}
