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

func File(path string, opts ...Option) (*asciidoc.Document, error) {
	file, err := os.Open(path)
	if err != nil {
		slog.Error("error reading file for parse", slog.String("path", path), slog.Any("error", err))
		return nil, err
	}
	return Reader(path, file, opts...)
}

func Reader(path string, reader io.Reader, opts ...Option) (*asciidoc.Document, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return Bytes(path, b, opts...)
}

func Bytes(path string, b []byte, opts ...Option) (*asciidoc.Document, error) {
	start := time.Now()
	vals, err := Parse(path, b, opts...)
	if err != nil {
		slog.Error("error parsing file", slog.String("path", path), slog.Any("error", err))
		return nil, err
	}
	elapsed := time.Since(start)

	switch vals := vals.(type) {
	case asciidoc.Elements:
		//		fmt.Printf("coalescing asciidoc...\n")
		d, err := elementsToDoc(vals)
		if err != nil {
			return nil, err
		}
		if debugParser {
			fmt.Printf("\n\n\n\n\n\n")
			dump(0, d.Children()...)
			fmt.Printf("elapsed: %s\n", elapsed.String())
		}
		return d, nil
	default:
		return nil, fmt.Errorf("unexpected type in File: %T", vals)
	}
}

func elementsToDoc(vals asciidoc.Elements) (d *asciidoc.Document, err error) {
	var els asciidoc.Elements
	err = reparseTables(vals)
	if err != nil {
		return nil, err
	}
	els, err = coalesce(vals)
	if err != nil {
		return nil, err
	}
	d = buildDoc(els)

	return d, nil
}

func String(path string, s string) (*asciidoc.Document, error) {
	return Bytes(path, []byte(s))
}
