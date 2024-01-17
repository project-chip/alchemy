package files

import (
	"context"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/hasty/alchemy/ascii"
	"github.com/hasty/alchemy/matter"
)

// Hacky workaround: there's a nasty bit in one of the Door Lock tables where the name of the command has an asterisk reference in it
// We just cut that out for now
var doorLockPattern = regexp.MustCompile(`\n+\s*(Schedule ID|User ID)&#8224;\s+`)

func LoadSpec(cxt context.Context, specRoot string, filesOptions Options, asciiSettings []configuration.Setting) (spec *matter.Spec, docs []*ascii.Doc, err error) {
	var lock sync.Mutex
	asciiSettings = append(ascii.GithubSettings(), asciiSettings...)

	var specPaths []string
	specPaths, err = getSpecPaths(specRoot)
	if err != nil {
		return
	}

	err = Process(cxt, specPaths, func(cxt context.Context, path string, index, total int) error {

		var file []byte
		file, err = os.ReadFile(path)
		if err != nil {
			return err
		}
		contents := string(file)
		if filepath.Base(path) == "DoorLock.adoc" {
			contents = doorLockPattern.ReplaceAllString(contents, "")
		}

		doc, err := ascii.Read(contents, path, asciiSettings...)
		if err != nil {
			return err
		}

		lock.Lock()
		docs = append(docs, doc)
		lock.Unlock()
		if filesOptions.Serial {
			slog.InfoContext(cxt, "Parsed spec adoc", "file", path)
		}
		return nil
	}, filesOptions)

	if err != nil {
		return
	}

	slog.InfoContext(cxt, "Building spec...")
	spec, err = ascii.BuildSpec(docs)

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

func SplitSpec(docs []*ascii.Doc) (map[matter.DocType][]*ascii.Doc, error) {
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
