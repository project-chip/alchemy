// Package parse provides readers for parsing Asciidoc
package parse

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/project-chip/alchemy/asciidoc"
)

func File(path string) (*asciidoc.Document, error) {
	file, err := os.Open(path)
	if err != nil {
		slog.Error("error reading file for parse", slog.String("path", path), slog.Any("error", err))
		return nil, err
	}
	return Reader(path, file)
}

func Reader(path string, reader io.Reader) (*asciidoc.Document, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return Bytes(path, b)
}

func Bytes(path string, b []byte) (*asciidoc.Document, error) {
	start := time.Now()
	vals, err := Parse(path, b)
	if err != nil {
		slog.Error("error parsing file", slog.String("path", path), slog.Any("error", err))
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
		d = buildDoc(els)
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

func String(path string, s string) (*asciidoc.Document, error) {
	return Bytes(path, []byte(s))
}
