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

var bannedPaths map[string]string = map[string]string{
	"namespaces/Namespace-Common-Position.adoc":         "parser does not support nested tables",
	"service_device_management/PowerSourceCluster.adoc": "parser gets stuck parsing",
	"secure_channel/Discovery.adoc":                     "parser gets stuck parsing",
}

func getFilePaths(filepath string) ([]string, error) {
	paths, err := doublestar.FilepathGlob(filepath)
	if err != nil {
		return nil, err
	}
	filtered := make([]string, 0, len(paths))
	for _, p := range paths {
		var banned bool
		for ban, reason := range bannedPaths {
			if strings.HasSuffix(p, ban) {
				fmt.Printf("Skipping excluded file %s; %s...\n", p, reason)
				banned = true
			}
		}
		if banned {
			continue
		}
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
