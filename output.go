package main

import "strings"

type output struct {
	out strings.Builder

	lastRune rune
}

func (o *output) WriteString(s string) {
	rs := []rune(s)
	if len(rs) > 0 {
		o.lastRune = rs[len(rs)-1]
		o.out.WriteString(s)
	}
}

func (o *output) WriteRune(r rune) {
	o.out.WriteRune(r)
	o.lastRune = r
}

func (o *output) WriteNewline() {
	if o.lastRune == '\n' {
		return
	}
	o.WriteRune('\n')
}

func (o *output) String() string {
	return o.out.String()
}
