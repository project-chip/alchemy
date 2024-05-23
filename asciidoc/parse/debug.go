package parse

import (
	"fmt"
	"os"
)

var debugParser = false
var debugParserStack = false

func debug(format string, a ...any) (n int, err error) {
	//return
	return fmt.Fprintf(os.Stderr, format, a...)
}

func debugPosition(c *current, format string, a ...any) (n int, err error) {
	//return
	fmt.Fprintf(os.Stderr, "[%d, %d-%d; %d]", c.pos.line, c.pos.col, c.pos.col+len(string(c.text))-1, c.pos.offset)
	fmt.Fprintf(os.Stderr, format, a...)
	if debugParserStack {
		for _, r := range c.parser.rstack {
			fmt.Fprintf(os.Stderr, "\t%s\n", r.name)
		}
	}
	return 0, nil
}
