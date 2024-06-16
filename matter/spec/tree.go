package spec

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/hasty/alchemy/asciidoc"
	"github.com/hasty/alchemy/internal/parse"
)

func buildTree(docs []*Doc) {

	tree := make(map[*Doc][]string)
	docPaths := make(map[string]*Doc)

	for _, doc := range docs {

		path := doc.Path
		docPaths[path] = doc

		parse.Search(doc.Elements(), func(link *asciidoc.FileInclude) parse.SearchShould {
			var p strings.Builder
			doc.buildDataTypeString(link.Set, &p)
			linkPath := filepath.Join(filepath.Dir(path), p.String())
			tree[doc] = append(tree[doc], linkPath)
			return parse.SearchShouldContinue
		})
	}

	for doc, children := range tree {
		for _, child := range children {
			if cd, ok := docPaths[child]; ok {
				cd.addParent(doc)
				doc.addChild(cd)
			} else {
				slog.Warn("unknown child path", "parent", doc.Path, "child", child)
			}
		}
	}

}

func dumpTree(r *Doc, indent int) {
	fmt.Printf(strings.Repeat("\t", indent))
	fmt.Printf("%s\n", r.Path)
	for _, c := range r.children {
		dumpTree(c, indent+1)
	}
}
