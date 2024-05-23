package parse

import (
	"unicode"
	"unicode/utf8"
)

func (c *current) previousRune() (r rune, ok bool) {
	offset := c.pos.offset
	if offset < 1 {
		return
	}
	var size int
	r, size = utf8.DecodeLastRune(c.parser.data[:offset-1])
	if r == utf8.RuneError {
		return
	}
	if size == 0 {
		return
	}
	ok = true
	return
}

func (c *current) previousRuneIsWhitespace() bool {
	r, ok := c.previousRune()
	if !ok {
		return false
	}
	return unicode.IsSpace(r)
}

func (c *current) currentColumn() int {
	return c.pos.col
}

func (c *current) currentPosition() (line, col, offset int) {
	line = c.parser.offset.line + c.pos.line
	col = c.pos.col
	if c.pos.line <= 1 {
		col += c.parser.offset.col
	}
	offset = c.parser.offset.offset + c.pos.offset
	return
}
