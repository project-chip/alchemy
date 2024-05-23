package parse

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/hasty/alchemy/asciidoc"
)

func File(path string) (*asciidoc.Document, error) {
	fmt.Printf("path: %s\n", path)
	//v, err := os.ReadFile(path)
	//	fmt.Printf("file: %s\n", string(v))
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error reading: %v\n", err)
		return nil, err
	}
	return Reader(path, file)
}

func Reader(path string, reader io.Reader) (*asciidoc.Document, error) {
	start := time.Now()
	vals, err := ParseReader(path, reader)
	if err != nil {
		fmt.Printf("error parsing: %v\n", err)
		return nil, err
	}
	elapsed := time.Since(start)

	switch vals := vals.(type) {
	case asciidoc.Set:
		//		fmt.Printf("coalescing asciidoc...\n")
		var d *asciidoc.Document
		var els asciidoc.Set
		err = reparseTables(vals)
		if err != nil {
			return nil, err
		}
		els, err = coalesce(vals)
		if err != nil {
			return nil, err
		}
		d, err = buildDoc(els)
		if err != nil {
			fmt.Printf("error building doc: %v\n", err)
			return nil, err
		}
		if debugParser {
			fmt.Printf("\n\n\n\n\n\n")
			dump(0, d.Elements()...)
			fmt.Printf("elapsed: %s\n", elapsed.String())
		}
		return d, nil
	default:
		return nil, fmt.Errorf("unexpected type in File: %T", vals)
	}
}
