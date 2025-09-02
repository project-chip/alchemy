package spec

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
)

func includePath(context parse.PreParseContext, doc *asciidoc.Document, include *asciidoc.FileInclude) (path asciidoc.Path, err error) {

	var rawPath string
	rawPath, err = renderPreParsedDoc(include.Children())
	if err != nil {
		return
	}
	path, err = context.ResolvePath(doc.Path.Dir(), rawPath)
	if err != nil {
		return
	}
	//slog.Info("resolving path", "dir", doc.Path.Dir(), "path", rawPath, "resolved", path)
	return
}
