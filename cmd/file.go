package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/output"
	"github.com/hasty/matterfmt/parse"
)

func getFilePaths(filepath string) ([]string, error) {
	paths, err := doublestar.FilepathGlob(filepath)
	if err != nil {
		return nil, err
	}
	filtered := make([]string, 0, len(paths))
	for _, p := range paths {
		/*if strings.HasSuffix(p, "secure_channel/Discovery.adoc") {
			continue
		}*/
		filtered = append(filtered, p)
	}
	return filtered, nil
}

func getOutputContext(cxt context.Context, path string) (*output.Context, error) {
	config := configuration.NewConfiguration(
		configuration.WithFilename(path),
		configuration.WithAttribute("second-ballot", false),
	)

	file, err := os.ReadFile(config.Filename)
	if err != nil {
		return nil, err
	}

	d, err := parse.ParseDocument(strings.NewReader(string(file)), config, parser.MaxExpressions(2000000))

	if err != nil {
		return nil, fmt.Errorf("failed parse: %w", err)
	}
	doc, err := ascii.NewDoc(d)
	if err != nil {
		return nil, err
	}
	doc.Path = path

	return output.NewContext(cxt, doc), nil
}
