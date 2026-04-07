package tests

import (
	"strings"
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/disco"
)

func TestXrefStyleOnlyInRoot(t *testing.T) {
	// Create a Baller with the option enabled
	b := disco.NewBaller(nil, disco.DiscoOptions{XrefStyleOnlyInRoot: true})

	// Test Case 1: Root doc without xrefstyle, with Copyright Notice
	{
		doc := &asciidoc.Document{
			Elements: asciidoc.Elements{
				&asciidoc.Section{
					Title: asciidoc.Elements{asciidoc.NewString("Copyright Notice")},
					Level: 1,
				},
				&asciidoc.Section{
					Title: asciidoc.Elements{asciidoc.NewString("First Section")},
					Level: 1,
				},
			},
		}

		err := b.TestHelperDiscoBall(doc, true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Verify added after Copyright Notice
		found := false
		for i, el := range doc.Elements {
			if s, ok := el.(*asciidoc.Section); ok && strings.Contains(s.Title[0].(*asciidoc.String).Value, "Copyright Notice") {
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
		if !found {
			t.Error("expected xrefstyle to be added after Copyright Notice section")
		}
	}

	// Test Case 2: Root doc without xrefstyle, without Copyright Notice
	{
		doc := &asciidoc.Document{
			Elements: asciidoc.Elements{
				&asciidoc.Section{
					Title: asciidoc.Elements{asciidoc.NewString("First Section")},
					Level: 1,
				},
			},
		}

		err := b.TestHelperDiscoBall(doc, true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Verify added before first section
		if len(doc.Elements) < 3 {
			t.Fatalf("expected at least 3 elements, got %d", len(doc.Elements))
		}
		if ae, ok := doc.Elements[0].(*asciidoc.AttributeEntry); !ok || ae.Name != "xrefstyle" {
			t.Error("expected xrefstyle to be added at the beginning")
		}
	}

	// Test Case 3: Root doc with existing xrefstyle somewhere
	{
		ae := asciidoc.NewAttributeEntry("xrefstyle")
		ae.Elements = asciidoc.Elements{asciidoc.NewString("full")}

		doc := &asciidoc.Document{
			Elements: asciidoc.Elements{
				&asciidoc.Section{
					Title: asciidoc.Elements{asciidoc.NewString("First Section")},
					Level: 1,
				},
				ae,
			},
		}

		err := b.TestHelperDiscoBall(doc, true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Verify not changed
		if s, ok := ae.Elements[0].(*asciidoc.String); !ok || s.Value != "full" {
			t.Errorf("expected xrefstyle to remain full, got %v", ae.Elements[0])
		}
	}

	// Test Case 4: Non-root doc with xrefstyle anywhere (should be removed)
	{
		ae1 := asciidoc.NewAttributeEntry("xrefstyle")
		ae2 := asciidoc.NewAttributeEntry("xrefstyle")

		doc := &asciidoc.Document{
			Elements: asciidoc.Elements{
				ae1,
				&asciidoc.Section{
					Title: asciidoc.Elements{asciidoc.NewString("First Section")},
					Level: 1,
				},
				ae2,
			},
		}

		err := b.TestHelperDiscoBall(doc, false)
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
			t.Error("expected all xrefstyle to be removed from non-root document")
		}
	}
}
