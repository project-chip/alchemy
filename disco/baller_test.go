package disco

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/config"
	"github.com/project-chip/alchemy/matter"
	"github.com/project-chip/alchemy/matter/spec"
)

func TestXrefStyleOnlyInRootWithFile(t *testing.T) {
	b := NewBaller(nil, DiscoOptions{XrefStyleOnlyInRoot: true})

	// Test Case 1: Root doc without xrefstyle (should add it)
	{
		path := asciidoc.Path{Relative: "testdata/root_without_xref.adoc"}
		in, err := os.ReadFile("testdata/root_without_xref.adoc")
		if err != nil {
			t.Fatalf("error reading file: %v", err)
		}

		doc, err := parse.Reader(path, bytes.NewReader(in))
		if err != nil {
			t.Fatalf("error parsing file: %v", err)
		}

		lib := spec.NewLibrary(doc, config.Library{}, nil, nil)
		lib.Reader = asciidoc.RawReader
		dc := &discoContext{
			Context:            context.Background(),
			doc:                doc,
			library:            lib,
			potentialDataTypes: make(map[string][]*DataTypeEntry),
		}

		top := parse.FindFirst[*asciidoc.Section](doc, asciidoc.RawReader, doc)
		err = b.discoBallTopLevelSection(dc, top, matter.DocTypeCluster)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Verify added after Copyright Notice
		found := false
		for i, el := range doc.Elements {
			if s, ok := el.(*asciidoc.Section); ok {
				name := lib.SectionName(s)
				if strings.Contains(strings.ToLower(name), "copyright notice") {
					// Check next element
					if i+1 < len(doc.Elements) {
						if _, ok := doc.Elements[i+1].(*asciidoc.NewLine); ok {
							if i+2 < len(doc.Elements) {
								if ae, ok := doc.Elements[i+2].(*asciidoc.AttributeEntry); ok && ae.Name == "xrefstyle" {
									found = true
									break
								}
							}
						}
					}
				}
			}
		}
		if !found {
			t.Error("expected xrefstyle to be added after Copyright Notice section loaded from file")
		}
	}

	// Test Case 2: Non-Root doc with xrefstyle (should remove it)
	{
		path := asciidoc.Path{Relative: "testdata/non_root_with_xref.adoc"}
		in, err := os.ReadFile("testdata/non_root_with_xref.adoc")
		if err != nil {
			t.Fatalf("error reading file: %v", err)
		}

		doc, err := parse.Reader(path, bytes.NewReader(in))
		if err != nil {
			t.Fatalf("error parsing file: %v", err)
		}

		lib := spec.NewLibrary(nil, config.Library{}, nil, nil) // Root is nil, so doc is not root
		lib.Reader = asciidoc.RawReader
		dc := &discoContext{
			Context:            context.Background(),
			doc:                doc,
			library:            lib,
			potentialDataTypes: make(map[string][]*DataTypeEntry),
		}

		top := parse.FindFirst[*asciidoc.Section](doc, asciidoc.RawReader, doc)
		err = b.discoBallTopLevelSection(dc, top, matter.DocTypeCluster)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Verify no xrefstyle anywhere
		found := false
		for _, el := range doc.Elements {
			if ae, ok := el.(*asciidoc.AttributeEntry); ok && ae.Name == "xrefstyle" {
				found = true
				break
			}
		}
		if found {
			t.Error("expected all xrefstyle to be removed from non-root document loaded from file")
		}
	}

	// Test Case 3: Root doc with xrefstyle (should do nothing)
	{
		path := asciidoc.Path{Relative: "testdata/root_with_xref.adoc"}
		in, err := os.ReadFile("testdata/root_with_xref.adoc")
		if err != nil {
			t.Fatalf("error reading file: %v", err)
		}

		doc, err := parse.Reader(path, bytes.NewReader(in))
		if err != nil {
			t.Fatalf("error parsing file: %v", err)
		}

		lib := spec.NewLibrary(doc, config.Library{}, nil, nil)
		lib.Reader = asciidoc.RawReader
		dc := &discoContext{
			Context:            context.Background(),
			doc:                doc,
			library:            lib,
			potentialDataTypes: make(map[string][]*DataTypeEntry),
		}

		// Count existing xrefstyle
		countBefore := 0
		for _, el := range doc.Elements {
			if ae, ok := el.(*asciidoc.AttributeEntry); ok && ae.Name == "xrefstyle" {
				countBefore++
			}
		}

		top := parse.FindFirst[*asciidoc.Section](doc, asciidoc.RawReader, doc)
		err = b.discoBallTopLevelSection(dc, top, matter.DocTypeCluster)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Count after
		countAfter := 0
		for _, el := range doc.Elements {
			if ae, ok := el.(*asciidoc.AttributeEntry); ok && ae.Name == "xrefstyle" {
				countAfter++
			}
		}

		if countAfter != countBefore {
			t.Errorf("expected number of xrefstyle to remain %d, got %d", countBefore, countAfter)
		}
	}
}

func TestAddDiscoballed(t *testing.T) {
	b := NewBaller(nil, DiscoOptions{AddDiscoballed: true})

	// Case 1: Without marker
	{
		doc, err := parse.Reader(asciidoc.Path{Relative: "test.adoc"}, strings.NewReader("= Title\n\nContent\n"))
		if err != nil {
			t.Fatalf("error parsing: %v", err)
		}
		lib := spec.NewLibrary(doc, config.Library{}, nil, nil)
		lib.Reader = asciidoc.RawReader
		dc := &discoContext{
			Context:            context.Background(),
			doc:                doc,
			library:            lib,
			potentialDataTypes: make(map[string][]*DataTypeEntry),
		}
		top := parse.FindFirst[*asciidoc.Section](doc, asciidoc.RawReader, doc)
		err = b.discoBallTopLevelSection(dc, top, matter.DocTypeCluster)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(doc.Elements) < 1 {
			t.Fatal("doc has no elements")
		}
		ae, ok := doc.Elements[0].(*asciidoc.AttributeEntry)
		if !ok || ae.Name != "alchemy-discoballed" {
			t.Errorf("expected marker at top, got %T", doc.Elements[0])
		}
	}

	// Case 2: Marker not at top
	{
		doc, err := parse.Reader(asciidoc.Path{Relative: "test.adoc"}, strings.NewReader("= Title\n\n:alchemy-discoballed:\nContent\n"))
		if err != nil {
			t.Fatalf("error parsing: %v", err)
		}
		lib := spec.NewLibrary(doc, config.Library{}, nil, nil)
		lib.Reader = asciidoc.RawReader
		dc := &discoContext{
			Context:            context.Background(),
			doc:                doc,
			library:            lib,
			potentialDataTypes: make(map[string][]*DataTypeEntry),
		}
		top := parse.FindFirst[*asciidoc.Section](doc, asciidoc.RawReader, doc)
		err = b.discoBallTopLevelSection(dc, top, matter.DocTypeCluster)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(doc.Elements) < 1 {
			t.Fatal("doc has no elements")
		}
		ae, ok := doc.Elements[0].(*asciidoc.AttributeEntry)
		if !ok || ae.Name != "alchemy-discoballed" {
			t.Errorf("expected marker at top, got %T", doc.Elements[0])
		}

		count := 0
		for _, el := range doc.Elements {
			if ae, ok := el.(*asciidoc.AttributeEntry); ok && ae.Name == "alchemy-discoballed" {
				count++
			}
		}
		if count != 1 {
			t.Errorf("expected 1 marker, got %d", count)
		}
	}

	// Case 3: Marker already at top
	{
		doc, err := parse.Reader(asciidoc.Path{Relative: "test.adoc"}, strings.NewReader(":alchemy-discoballed:\n= Title\n\nContent\n"))
		if err != nil {
			t.Fatalf("error parsing: %v", err)
		}
		lib := spec.NewLibrary(doc, config.Library{}, nil, nil)
		lib.Reader = asciidoc.RawReader
		dc := &discoContext{
			Context:            context.Background(),
			doc:                doc,
			library:            lib,
			potentialDataTypes: make(map[string][]*DataTypeEntry),
		}
		top := parse.FindFirst[*asciidoc.Section](doc, asciidoc.RawReader, doc)
		err = b.discoBallTopLevelSection(dc, top, matter.DocTypeCluster)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(doc.Elements) < 1 {
			t.Fatal("doc has no elements")
		}
		ae, ok := doc.Elements[0].(*asciidoc.AttributeEntry)
		if !ok || ae.Name != "alchemy-discoballed" {
			t.Errorf("expected marker at top, got %T", doc.Elements[0])
		}

		count := 0
		for _, el := range doc.Elements {
			if ae, ok := el.(*asciidoc.AttributeEntry); ok && ae.Name == "alchemy-discoballed" {
				count++
			}
		}
		if count != 1 {
			t.Errorf("expected 1 marker, got %d", count)
		}
	}
}
