package spec

import (
	"log/slog"
	"os"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/internal/pipeline"
)

type DocCache struct {
	root    string
	cache   pipeline.Map[string, *asciidoc.Document]
	utility pipeline.Map[string, *asciidoc.Document]
}

func NewDocCache(root string) *DocCache {
	return &DocCache{
		root:    root,
		cache:   pipeline.NewConcurrentMap[string, *asciidoc.Document](),
		utility: pipeline.NewConcurrentMap[string, *asciidoc.Document](),
	}
}

func cacheFromPipeline(root string, docs []*pipeline.Data[*asciidoc.Document]) *DocCache {
	dc := NewDocCache(root)
	for _, d := range docs {
		dc.cache.Store(d.Content.Path.Relative, d.Content)
	}
	return dc
}

func (dc DocCache) Add(doc *asciidoc.Document) {
	dc.cache.Store(doc.Path.Relative, doc)
}

func (dc DocCache) include(path asciidoc.Path, parent asciidoc.Parent) (doc *asciidoc.Document, err error) {

	table, isTable := parent.(*asciidoc.Table)

	var options []parse.Option

	if isTable {
		existing, ok := dc.utility.Load(path.Relative)
		if ok {
			doc = existing.Clone().(*asciidoc.Document)
			return
		}
		options = append(options, parse.Entrypoint("IncludedTableElements"))
		options = append(options, parse.GlobalStore("table", table))
	} else {
		existing, ok := dc.cache.LoadAndDelete(path.Relative)
		if ok {
			//	slog.Info("returning existing file", "path", path.Relative)
			doc = existing
			return
		}
	}

	//slog.Info("including non-existent file", "path", path.Relative)
	var contents *os.File
	contents, err = os.Open(path.Absolute)
	if err != nil {
		return
	}
	defer contents.Close()

	doc, err = parse.RawReader(path, contents, options...)
	if err != nil {
		slog.Error("error parsing file", slog.String("path", path.Absolute), slog.Any("error", err))
		return
	}
	if isTable {
		parse.ReparseTable(table, doc.Elements)
	}
	return
}
