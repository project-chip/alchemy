package render

import (
	"context"
	"strings"
)

type unwrappedTarget struct {
	context.Context

	out      strings.Builder
	lastRune rune
}

func NewUnwrappedTarget(parent context.Context) Target {
	return &unwrappedTarget{
		Context: parent,
	}
}

func (o *unwrappedTarget) WriteString(s string) {
	rs := []rune(s)
	if len(rs) > 0 {
		o.lastRune = rs[len(rs)-1]
		o.out.WriteString(s)
	}
}

func (o *unwrappedTarget) WriteRune(r rune) {
	o.out.WriteRune(r)
	o.lastRune = r
}

func (o *unwrappedTarget) EnsureNewLine() {
	if o.lastRune == '\n' {
		return
	}
	o.WriteRune('\n')
}

func (o *unwrappedTarget) String() string {
	return o.out.String()
}

func (o *unwrappedTarget) FlushWrap() {

}

func (o *unwrappedTarget) EnableWrap() {

}

func (o *unwrappedTarget) DisableWrap() {

}

func (o *unwrappedTarget) StartBlock() {

}

func (o *unwrappedTarget) EndBlock() {

}

func (o *unwrappedTarget) Subtarget() Target {
	return NewUnwrappedTarget(o.Context)
}
