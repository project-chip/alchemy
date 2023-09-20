package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Printf("hello!!\n")

	logrus.SetLevel(logrus.ErrorLevel)

	files, err := filepath.Glob(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Rendering %d files from %s...\n", len(files), os.Args[1])

	for _, file := range files {
		fmt.Printf("Rendering %s...\n", file)
		doc, err := readFile(file)

		if err != nil {
			panic(err)
		}

		if doc == nil {
			continue
		}

		out := postProcessFile(doc.render())
		//fmt.Printf("Result:\n%s\n", out)

		os.WriteFile(file, []byte(out), os.ModeAppend)
	}

}

func readFile(path string) (*doc, error) {
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

	d, err := parseDocument(strings.NewReader(preprocessFile(string(file))), config)
	if err != nil {
		panic(fmt.Errorf("failed parse: %w", err))
	}
	doc := &doc{base: d, root: &section{}}
	if 1 == 0 {
		dump(doc, d)
		return nil, nil
	}
	for _, e := range d.BodyElements() {
		switch el := e.(type) {
		case *types.Section:
			doc.addSection(doc.root, el)
		default:
			doc.root.elements = append(doc.root.elements, e)
		}
	}
	return doc, nil
}

func preprocessFile(file string) string {
	//file = strings.ReplaceAll(file, `\|`, "{ESCAPEDPIPE}")
	return file
}

func postProcessFile(file string) string {
	return file
}
