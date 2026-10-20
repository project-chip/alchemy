package overlay

import (
	"testing"

	"github.com/project-chip/alchemy/asciidoc"
)

func TestSpecialCharacterRemovalOccurrence(t *testing.T) {
	elements := asciidoc.Elements{
		asciidoc.NewSpecialCharacter("&"),
		asciidoc.NewSpecialCharacter("&"),
		asciidoc.NewSpecialCharacter("<"),
	}
	reader := &Reader{removedCharacters: map[elementOccurrence]bool{
		{parent: nil, child: &elements[0]}: true,
	}}
	for _, test := range []struct {
		name     string
		elements asciidoc.Elements
		want     int
	}{
		{"whole", elements, 2},
		{"tail", elements[1:], 2},
		{"removed-only", elements[:1], 0},
		{"copied-occurrences", append(asciidoc.Elements{}, elements...), 3},
	} {
		t.Run(test.name, func(t *testing.T) {
			if got := reader.Iterate(nil, test.elements).Count(); got != test.want {
				t.Fatalf("yielded %d elements, want %d", got, test.want)
			}
		})
	}
}
