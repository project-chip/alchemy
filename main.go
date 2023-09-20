package main

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
	"github.com/hasty/matterfmt/render"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetLevel(logrus.ErrorLevel)

	files, err := filepath.Glob(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Rendering %d files from %s...\n", len(files), os.Args[1])

	for _, file := range files {
		fmt.Printf("Rendering %s...\n", file)
		doc, err := readFile(file)

		cxt := output.NewContext(context.Background(), doc)

		if 1 == 0 {
			parse.Dump(cxt)
			return
		}

		if err != nil {
			panic(err)
		}

		if doc == nil {
			continue
		}

		out := postProcessFile(render.Render(cxt, doc))
		//fmt.Printf("Result:\n%s\n", out)

		os.WriteFile(file, []byte(out), os.ModeAppend)
	}

}

func readFile(path string) (*ascii.Doc, error) {
	config := configuration.NewConfiguration(
		configuration.WithFilename(path),
		configuration.WithAttribute("second-ballot", false),
		//configuration.WithAttributes(attrs),
		//configuration.WithCSS(css),
		//configuration.WithBackEnd(backend),
		//configuration.WithHeaderFooter(!noHeaderFooter)

	)

	file, err := os.ReadFile(config.Filename)
	if err != nil {
		panic(err)
	}

	d, err := parse.ParseDocument(strings.NewReader(preprocessFile(string(file))), config)

	if err != nil {
		panic(fmt.Errorf("failed parse: %w", err))
	}
	doc := ascii.NewDoc(d)

	return doc, nil
}

func preprocessFile(file string) string {
	//file = strings.ReplaceAll(file, `\|`, "{ESCAPEDPIPE}")
	return file
}

func postProcessFile(file string) string {
	return file
}
