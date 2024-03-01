package ascii

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func readFile(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	contents := string(file)

	if filepath.Base(path) == "DoorLock.adoc" {
		var doorLockPattern = regexp.MustCompile(`\n+\s*[^&\n]+&#8224;\s+`)
		contents = doorLockPattern.ReplaceAllString(contents, " ")
	}
	return contents, nil
}

func ReadFile(path string, settings ...configuration.Setting) (*Doc, error) {

	contents, err := readFile(path)
	if err != nil {
		return nil, err
	}
	return Read(contents, path)
}

func Read(contents string, path string) (doc *Doc, err error) {

	config := configuration.NewConfiguration(configuration.WithFilename(path))
	config.IgnoreIncludes = true

	var d *types.Document

	d, err = ParseDocument(strings.NewReader(contents), config, parser.MaxExpressions(2000000))

	if err != nil {
		return nil, fmt.Errorf("failed parse: %w", err)
	}

	doc, err = NewDoc(d)
	if err != nil {
		return nil, err
	}
	doc.Path = path

	PatchUnrecognizedReferences(doc)

	return doc, nil
}
