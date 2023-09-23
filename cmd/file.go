package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/hasty/matterfmt/ascii"
	"github.com/hasty/matterfmt/output"
	"github.com/hasty/matterfmt/parse"
	"github.com/urfave/cli/v2"
)

func getFilePaths(cCtx *cli.Context) ([]string, error) {
	return filepath.Glob(cCtx.Args().First())
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

	d, err := parse.ParseDocument(strings.NewReader(string(file)), config)

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
