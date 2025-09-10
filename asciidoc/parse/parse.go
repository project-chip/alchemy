// Package parse provides readers for parsing Asciidoc
package parse

import (
	"fmt"
	"io"
	"log/slog"
	"time"

	"github.com/project-chip/alchemy/asciidoc"
)

func Reader(path asciidoc.Path, reader io.Reader, opts ...Option) (*asciidoc.Document, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return Bytes(path, b, opts...)
}

func Bytes(path asciidoc.Path, b []byte, opts ...Option) (*asciidoc.Document, error) {

	d := &asciidoc.Document{Path: path}

	elements, err := parseBytes(d, b, opts...)
	if err != nil {
		return nil, err
	}
	buildDoc(d, elements)

	return d, nil
}

func parseBytes(document *asciidoc.Document, b []byte, opts ...Option) (elements asciidoc.Elements, err error) {

	opts = append(opts, setDocument(document))

	start := time.Now()
	var vals any
	vals, err = parse(document.Path.Absolute, b, opts...)
	if err != nil {
		slog.Error("error parsing file", slog.String("path", document.Path.Absolute), slog.Any("error", err))
		return
	}
	elapsed := time.Since(start)

	switch vals := vals.(type) {
	case asciidoc.Elements:
		//		fmt.Printf("coalescing asciidoc...\n")
		err = reparseTables(vals)
		if err != nil {
			return
		}
		elements, err = coalesce(vals)
		if err != nil {
			return
		}
		if debugParser {
			fmt.Printf("\n\n\n\n\n\n")
			dump(0, elements.Children()...)
			fmt.Printf("elapsed: %s\n", elapsed.String())
		}
		return
	default:
		return nil, fmt.Errorf("unexpected type in File: %T", vals)
	}
}

func Raw(path asciidoc.Path, b []byte, opts ...Option) (*asciidoc.Document, error) {

	d := &asciidoc.Document{Path: path}

	elements, err := parseBytes(d, b, opts...)
	if err != nil {
		return nil, err
	}
	d.Elements = elements

	return d, nil
}

func RawReader(path asciidoc.Path, reader io.Reader, opts ...Option) (*asciidoc.Document, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return Raw(path, b, opts...)
}

func String(path asciidoc.Path, s string) (*asciidoc.Document, error) {
	return Bytes(path, []byte(s))
}
