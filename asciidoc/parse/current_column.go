//go:build !windows

package parse

func (c *current) currentColumn() int {
	return c.pos.col
}
