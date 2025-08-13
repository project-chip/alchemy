package spec

import (
	"log/slog"
	"os"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
)

func readIncludeFile(docGroup *DocGroup, path asciidoc.Path, parent asciidoc.Parent, forceRead bool) (doc *Doc, err error) {

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

	var ad *asciidoc.Document
	ad, err = parse.Reader(path.Relative, contents, options...)
	if err != nil {
		slog.Error("error parsing file", slog.String("path", path.Absolute), slog.Any("error", err))
		return
	}
	if isTable {
		parse.ReparseTable(table, ad.Elements)
	}
	doc, err = newDoc(ad, path)
	if err == nil {
		doc.group = docGroup
	}
	return
}

func includePath(context parse.PreParseContext, doc *Doc, include *asciidoc.FileInclude) (path asciidoc.Path, err error) {

	var rawPath string
	rawPath, err = renderPreParsedDoc(include.Children())
	if err != nil {
		return
	}
	path, err = context.ResolvePath(doc.Path.Dir(), rawPath)
	if err != nil {
		return
	}
	return
}
