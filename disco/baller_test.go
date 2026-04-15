package disco

import (
	"bytes"
	"context"
	"os"
	"reflect"
	"strings"
	"testing"
	"unsafe"

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

func TestNormalizeAnchorCrossFile(t *testing.T) {
	pathA := asciidoc.Path{Relative: "testdata/anchor_file.adoc"}
	inA, err := os.ReadFile("testdata/anchor_file.adoc")
	if err != nil {
		t.Fatalf("error reading file A: %v", err)
	}
	docA, err := parse.Reader(pathA, bytes.NewReader(inA))
	if err != nil {
		t.Fatalf("error parsing file A: %v", err)
	}

	pathB := asciidoc.Path{Relative: "testdata/reference_file.adoc"}
	inB, err := os.ReadFile("testdata/reference_file.adoc")
	if err != nil {
		t.Fatalf("error reading file B: %v", err)
	}
	docB, err := parse.Reader(pathB, bytes.NewReader(inB))
	if err != nil {
		t.Fatalf("error parsing file B: %v", err)
	}

	s := &spec.Specification{}
	lib := spec.NewLibrary(docA, config.Library{}, nil, nil)
	lib.Reader = asciidoc.RawReader
	
	// Use reflect to set s.libraryIndex
	v := reflect.ValueOf(s).Elem()
	f := v.FieldByName("libraryIndex")
	
	ptr := unsafe.Pointer(f.UnsafeAddr())
	mPtr := (*map[*asciidoc.Document]*spec.Library)(ptr)
	
	*mPtr = make(map[*asciidoc.Document]*spec.Library)
	(*mPtr)[docA] = lib
	(*mPtr)[docB] = lib

	an := AnchorNormalizer{
		spec:    s,
		options: DiscoOptions{NormalizeAnchors: true},
	}
	
	_, err = lib.Anchors(asciidoc.RawReader)
	if err != nil {
		t.Fatalf("error indexing anchors: %v", err)
	}
	
	// Run it on docB
	an.rewriteCrossReferences(docB)
	
	// Verify that docB retains its label!
	var foundXref *asciidoc.CrossReference
	parse.Search(docB, asciidoc.RawReader, nil, docB.Children(), func(doc *asciidoc.Document, el asciidoc.Element, parent asciidoc.ParentElement, index int) parse.SearchShould {
		if xref, ok := el.(*asciidoc.CrossReference); ok {
			foundXref = xref
			return parse.SearchShouldStop
		}
		return parse.SearchShouldContinue
	})
	
	if foundXref == nil {
		t.Fatal("expected to find cross reference in docB")
	}
	
	if len(foundXref.Elements) == 0 {
		t.Error("expected reference to retain its label, but it was removed")
	}
	
	label := labelText(foundXref.Elements)
	expectedLabel := "A Non-Normalized Replacement Text Different Than Section Name"
	if label != expectedLabel {
		t.Errorf("expected label %q, got %q", expectedLabel, label)
	}
}
