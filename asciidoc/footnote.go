package asciidoc

type Footnote struct {
	position
	raw

	ID    string
	Value any
}

func NewFootnote(id string, value any) Footnote {
	return Footnote{ID: id, Value: value}
}
