//go:build windows

package parse

func (c *current) currentColumn() int {
	if len(c.text) == 2 && c.text[0] == '\r' && c.text[1] == '\n' {
		// Special case on Windows; NewLine matches \r\n, but the parser only sees \n as a col reset, so we fake it
		return 0
	}
	return c.pos.col
}
