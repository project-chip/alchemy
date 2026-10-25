package spec

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
	"github.com/project-chip/alchemy/config"
	"github.com/project-chip/alchemy/matter"
)

func TestBitmapIncludeOperatorPreservation(t *testing.T) {
	for _, operator := range []struct {
		name  string
		value string
	}{
		{"and", "&"},
		{"or", "|"},
		{"less", "<"},
		{"greater", ">"},
		{"logical-and", "&&"},
		{"less-equal", "<="},
		{"greater-equal", ">="},
		{"left-shift", "<<"},
		{"right-shift", ">>"},
	} {
		t.Run(operator.name, func(t *testing.T) {
			for _, test := range []struct {
				name       string
				includes   string
				attributes []asciidoc.AttributeName
			}{
				{"standalone", "include::visible.adoc[]\n", nil},
				{"disabled-after", "include::visible.adoc[]\ninclude::sibling.adoc[]\n", nil},
				{"disabled-before", "include::sibling.adoc[]\ninclude::visible.adoc[]\n", nil},
				{"enabled", "include::visible.adoc[]\ninclude::sibling.adoc[]\n", []asciidoc.AttributeName{"synthetic-feature"}},
			} {
				t.Run(test.name, func(t *testing.T) {
					root := t.TempDir()
					expression := "MASK_A " + operator.value + " MASK_B"
					visible := strings.ReplaceAll("Combine `"+expression+"`.", "|", "\\|")
					want := strings.ReplaceAll("Combine "+expression+".", "|", "\\|")
					for name, contents := range map[string]string{
						"main.adoc":    test.includes,
						"visible.adoc": "[cols=\"1,1,4a\"]\n|===\n|Bit |Name |Summary\n|3 |SyntheticFlag |" + visible + "\n|===\n",
						"sibling.adoc": "ifdef::synthetic-all,synthetic-feature[]\nHidden `MASK_C " + operator.value + " MASK_D`.\nendif::[]\n",
					} {
						if err := os.WriteFile(filepath.Join(root, name), []byte(contents), 0600); err != nil {
							t.Fatal(err)
						}
					}
					document, err := readFile(filepath.Join(root, "main.adoc"), root)
					if err != nil {
						t.Fatal(err)
					}
					library := NewLibrary(document, config.Library{}, nil, NewDocCache(root))
					if err := library.Preparse(root, test.attributes); err != nil {
						t.Fatal(err)
					}
					found := 0
					parse.Traverse(document, library, document, library.Children(document), func(doc *asciidoc.Document, table *asciidoc.Table, _ asciidoc.ParentElement, _ int) parse.SearchShould {
						info, err := ReadTable(doc, library, table)
						if err != nil {
							t.Fatal(err)
						}
						if len(info.Rows) != 2 {
							t.Fatalf("rows = %d, want 2", len(info.Rows))
						}
						for _, reader := range []struct {
							name  string
							value asciidoc.Reader
						}{
							{"raw", asciidoc.RawReader},
							{"preprocessed", library},
						} {
							value, err := info.ReadValue(reader.value, info.Rows[1], matter.TableColumnSummary)
							if err != nil {
								t.Fatal(err)
							}
							if value != want {
								t.Errorf("%s summary = %q, want %q", reader.name, value, want)
							}
						}
						found++
						return parse.SearchShouldContinue
					})
					if found != 1 {
						t.Fatalf("visible table count = %d, want 1", found)
					}
					hidden := 0
					parse.Traverse(document, library, document, library.Children(document), func(_ *asciidoc.Document, element *asciidoc.Monospace, _ asciidoc.ParentElement, _ int) parse.SearchShould {
						var text strings.Builder
						for child := range library.Iterate(element, library.Children(element)) {
							switch child := child.(type) {
							case *asciidoc.String:
								text.WriteString(child.Value)
							case *asciidoc.SpecialCharacter:
								text.WriteString(child.Character)
							default:
								t.Fatalf("unexpected synthetic expression child: %T", child)
							}
						}
						value := text.String()
						if strings.Contains(value, "MASK_C") {
							hidden++
							if value != "MASK_C "+operator.value+" MASK_D" {
								t.Errorf("enabled sibling expression = %q", value)
							}
						}
						return parse.SearchShouldContinue
					})
					wantHidden := 0
					if test.name == "enabled" {
						wantHidden = 1
					}
					if hidden != wantHidden {
						t.Errorf("sibling expression count = %d, want %d", hidden, wantHidden)
					}
				})
			}
		})
	}
}
