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

func (o *unwrappedTarget) WriteString(s string) (n int, err error) {
	rs := []rune(s)
	if len(rs) > 0 {
		o.lastRune = rs[len(rs)-1]
		n, err = o.out.WriteString(s)
	}
	return
}

func (o *unwrappedTarget) WriteRune(r rune) (n int, err error) {
	o.lastRune = r
	n, err = o.out.WriteRune(r)
	return
}

func (o *unwrappedTarget) EnsureNewLine() {
	if o.lastRune == '\n' || o.out.Len() == 0 {
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
