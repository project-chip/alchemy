package zap

import (
	"context"
	"io/fs"
	"log/slog"
	"path/filepath"
	"strings"
	"sync"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/cmd/files"
	"github.com/hasty/alchemy/matter"
)

func loadSpec(cxt context.Context, specRoot string, filesOptions files.Options, asciiSettings []configuration.Setting) (docs []*ascii.Doc, err error) {
	var lock sync.Mutex
	asciiSettings = append(ascii.GithubSettings(), asciiSettings...)

	var specPaths []string
	specPaths, err = getSpecPaths(specRoot)
	if err != nil {
		return
	}

	err = files.Process(cxt, specPaths, func(cxt context.Context, file string, index, total int) error {

		doc, err := ascii.Open(file, asciiSettings...)
		if err != nil {
			return err
		}
		lock.Lock()
		docs = append(docs, doc)
		lock.Unlock()
		if filesOptions.Serial {
			slog.InfoContext(cxt, "Parsed spec adoc", "file", file)
		}
		return nil
	}, filesOptions)
	return
}

func getSpecPaths(specRoot string) (paths []string, err error) {
	srcRoot := filepath.Join(specRoot, "/src/")
	err = filepath.WalkDir(srcRoot, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".adoc" && !strings.HasSuffix(path, "-draft.adoc") {
			paths = append(paths, path)
		}
		return nil
	})
	return
}

func splitSpec(docs []*ascii.Doc) (map[matter.DocType][]*ascii.Doc, error) {
	byType := make(map[matter.DocType][]*ascii.Doc)
	for _, d := range docs {
		docType, err := d.DocType()
		if err != nil {
			return nil, err
		}
		byType[docType] = append(byType[docType], d)
	}
	return byType, nil
}
