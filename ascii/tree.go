package ascii

import (
	"log/slog"
	"path/filepath"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/internal/parse"
)

func buildTree(docs []*Doc) {

	tree := make(map[*Doc][]string)
	docPaths := make(map[string]*Doc)

	for _, doc := range docs {

		path := doc.Path
		docPaths[path] = doc

		top := parse.FindFirst[*Section](doc.Elements)
		if top == nil {
			continue
		}

		parse.Search[*types.Section](top.Base.Elements, func(t *types.Section) bool {
			link := parse.FindFirst[*types.InlineLink](t.Title)
			if link != nil {
				linkPath, ok := link.Location.Path.(string)
				if ok {
					linkPath = filepath.Join(filepath.Dir(path), linkPath)
					tree[doc] = append(tree[doc], linkPath)
				}
			}
			return false
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
