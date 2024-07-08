package render

import (
	"github.com/project-chip/alchemy/asciidoc"
	"github.com/project-chip/alchemy/internal/parse"
)

type InputDocument interface {
	parse.HasElements

	Footnotes() []*asciidoc.Footnote
}
