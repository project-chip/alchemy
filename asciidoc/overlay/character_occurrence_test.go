package overlay

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/asciidoc/parse"
)

func TestSpecialCharacterRemovalOccurrence(t *testing.T) {
	elements := asciidoc.Elements{
		asciidoc.NewSpecialCharacter("&"),
		asciidoc.NewSpecialCharacter("&"),
		asciidoc.NewSpecialCharacter("<"),
	}
	reader := &Reader{removedCharacters: map[*asciidoc.Element]bool{
		&elements[0]: true,
	}}
	for _, test := range []struct {
		name     string
		elements asciidoc.Elements
		want     int
	}{
		{"whole", elements, 2},
		{"tail", elements[1:], 2},
		{"removed-only", elements[:1], 0},
		// Copies have new occurrence identities; overlay readers must iterate the original backing slices.
		{"copied-occurrences", append(asciidoc.Elements{}, elements...), 3},
	} {
		t.Run(test.name, func(t *testing.T) {
			if got := reader.Iterate(nil, test.elements).Count(); got != test.want {
				t.Fatalf("yielded %d elements, want %d", got, test.want)
			}
		})
	}
}

func TestSpecialCharacterRemovalReplacement(t *testing.T) {
	elements := asciidoc.Elements{asciidoc.NewSpecialCharacter("&")}
	parent := asciidoc.NewSection(nil, 1)
	replaced := asciidoc.NewSection(nil, 1)
	reader := &Reader{
		overlays: elementOverlays{
			replaced: {action: overlayActionReplace, replace: elements},
		},
		removedCharacters: map[*asciidoc.Element]bool{
			&elements[0]: true,
		},
	}
	if got := reader.Iterate(parent, elements).Count(); got != 0 {
		t.Fatalf("direct iteration yielded %d removed characters", got)
	}
	if got := reader.Iterate(nil, asciidoc.Elements{replaced}).Count(); got != 0 {
		t.Fatalf("replacement yielded %d removed characters", got)
	}
}

func TestSpecialCharacterTraversedOccurrences(t *testing.T) {
	section := asciidoc.NewSection(asciidoc.Elements{
		asciidoc.NewSpecialCharacter("&"),
		asciidoc.NewSpecialCharacter("<"),
	}, 1)
	section.Elements = asciidoc.Elements{asciidoc.NewSpecialCharacter("&")}
	document := &asciidoc.Document{Elements: asciidoc.Elements{section}}
	reader := &occurrenceReader{Reader: asciidoc.RawReader}
	removed := map[*asciidoc.Element]bool{}
	parse.Traverse(document, reader, document, document.Children(), func(doc *asciidoc.Document, element asciidoc.Element, parent asciidoc.ParentElement, index int) parse.SearchShould {
		if _, character := element.(asciidoc.SpecialCharacter); character {
			removed[reader.child] = true
		}
		return parse.SearchShouldContinue
	})
	if len(removed) != 3 || !removed[&section.Title[0]] || !removed[&section.Title[1]] || !removed[&section.Elements[0]] {
		t.Fatal("traversal did not record the actual title and body slots")
	}
	filtered := &Reader{removedCharacters: removed}
	if got := filtered.Iterate(section, section.Title).Count(); got != 0 {
		t.Fatalf("title yielded %d removed characters", got)
	}
	if got := filtered.Iterate(section, section.Elements).Count(); got != 0 {
		t.Fatalf("body yielded %d removed characters", got)
	}
}
