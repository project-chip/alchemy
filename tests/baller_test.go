package tests

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/disco"
)


func TestXrefStyleOnlyInRootWithFile(t *testing.T) {
	b := disco.NewBaller(nil, disco.DiscoOptions{XrefStyleOnlyInRoot: true})

	// Test Case 1: Root doc without xrefstyle (should add it)
	{
		path := asciidoc.Path{Relative: "../disco/testdata/root_without_xref.adoc"}
		in, err := os.ReadFile("../disco/testdata/root_without_xref.adoc")
		if err != nil {
			t.Fatalf("error reading file: %v", err)
		}

		doc, err := parse.Reader(path, bytes.NewReader(in))
		if err != nil {
			t.Fatalf("error parsing file: %v", err)
		}

		err = b.TestHelperDiscoBall(doc, true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Verify added after Copyright Notice
		found := false
		for i, el := range doc.Elements {
			if s, ok := el.(*asciidoc.Section); ok && len(s.Title) > 0 {
				if str, ok := s.Title[0].(*asciidoc.String); ok && strings.Contains(str.Value, "Copyright Notice") {
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
		path := asciidoc.Path{Relative: "../disco/testdata/non_root_with_xref.adoc"}
		in, err := os.ReadFile("../disco/testdata/non_root_with_xref.adoc")
		if err != nil {
			t.Fatalf("error reading file: %v", err)
		}

		doc, err := parse.Reader(path, bytes.NewReader(in))
		if err != nil {
			t.Fatalf("error parsing file: %v", err)
		}

		err = b.TestHelperDiscoBall(doc, false)
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
		path := asciidoc.Path{Relative: "../disco/testdata/root_with_xref.adoc"}
		in, err := os.ReadFile("../disco/testdata/root_with_xref.adoc")
		if err != nil {
			t.Fatalf("error reading file: %v", err)
		}

		doc, err := parse.Reader(path, bytes.NewReader(in))
		if err != nil {
			t.Fatalf("error parsing file: %v", err)
		}

		// Count existing xrefstyle
		countBefore := 0
		for _, el := range doc.Elements {
			if ae, ok := el.(*asciidoc.AttributeEntry); ok && ae.Name == "xrefstyle" {
				countBefore++
			}
		}

		err = b.TestHelperDiscoBall(doc, true)
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
