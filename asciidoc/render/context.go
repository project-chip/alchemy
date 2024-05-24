package render

import (
	"context"
	"strings"

	"github.com/hasty/alchemy/asciidoc"
	"github.com/hasty/alchemy/internal/parse"
)

type Context struct {
	context.Context

	Doc InputDocument

	out      strings.Builder
	lastRune rune
}

type InputDocument interface {
	parse.HasElements

	Footnotes() []*asciidoc.Footnote
}

func NewContext(parent context.Context, doc InputDocument) *Context {
	return &Context{
		Context: parent,
		Doc:     doc,
	}
}

func (o *Context) WriteString(s string) {
	rs := []rune(s)
	if len(rs) > 0 {
		o.lastRune = rs[len(rs)-1]
		o.out.WriteString(s)
	}
}

func (o *Context) WriteRune(r rune) {
	o.out.WriteRune(r)
	o.lastRune = r
}

func (o *Context) EnsureNewLine() {
	if o.lastRune == '\n' {
		return
	}
	o.WriteRune('\n')
}

func (o *Context) String() string {
	return o.out.String()
}
