package spec

import (
	"log/slog"
	"os"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
)

func readIncludeFile(docGroup *Library, path asciidoc.Path, parent asciidoc.Parent, forceRead bool) (doc *asciidoc.Document, err error) {

	table, isTable := parent.(*asciidoc.Table)

	if docGroup != nil && !isTable && !forceRead {

		doc, ok := docGroup.index[path.Relative]
		if ok {
			return doc, nil
		}

	}

	var contents *os.File
	contents, err = os.Open(path.Absolute)
	if err != nil {
		return
	}
	defer contents.Close()

	var options []parse.Option

	if isTable {
		options = append(options, parse.Entrypoint("IncludedTableElements"))
		options = append(options, parse.GlobalStore("table", table))
	}

	doc, err = parse.Reader(path, contents, options...)
	if err != nil {
		slog.Error("error parsing file", slog.String("path", path.Absolute), slog.Any("error", err))
		return
	}
	if isTable {
		parse.ReparseTable(table, doc.Elements)
	}
	return
}

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
