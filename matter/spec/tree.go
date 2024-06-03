package spec

import (
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

		top := parse.FindFirst[*Section](doc.Elements())
		if top == nil {
			continue
		}

		parse.Search[*asciidoc.Section](top.Base.Elements(), func(t *asciidoc.Section) parse.SearchShould {
			link := parse.FindFirst[*asciidoc.Link](t.Title)
			if link != nil {
				var p strings.Builder
				doc.buildDataTypeString(link.URL.Path, &p)
				linkPath := filepath.Join(filepath.Dir(path), p.String())
				tree[doc] = append(tree[doc], linkPath)

			}
			return parse.SearchShouldContinue
		})
	}

	for doc, children := range tree {
		for _, child := range children {
			if cd, ok := docPaths[child]; ok {
				cd.addParent(doc)
			} else {
				slog.Debug("unknown child path", "parent", doc.Path, "child", child)
			}
		}
	}
}
