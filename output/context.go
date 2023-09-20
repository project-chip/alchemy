package output

import (
	"context"
	"strings"

	"github.com/hasty/matterfmt/ascii"
)

type Context struct {
	context.Context

	Doc *ascii.Doc

	out      strings.Builder
	lastRune rune

	OrderedListDepth int
}

func NewContext(parent context.Context, doc *ascii.Doc) *Context {
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

func (o *Context) WriteNewline() {
	if o.lastRune == '\n' {
		return
	}
	o.WriteRune('\n')
}

func (o *Context) String() string {
	return o.out.String()
}
