package ascii

import (
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/hasty/alchemy/parse"
)

func BuildTree(docs []*Doc) {

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
					if strings.HasSuffix(linkPath, "-draft.adoc") {
						return false
					}
					linkPath = filepath.Join(filepath.Dir(path), linkPath)
					linkPath = strings.ReplaceAll(linkPath, "energy-management.adoc", "energy_management.adoc")
					slog.Info("linked file\n", "from", path, "to", linkPath)
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
				slog.Warn("unknown child path", "parent", doc.Path, "child", child)
			}
		}
	}
}
