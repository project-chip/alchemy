package parse

import "github.com/project-chip/alchemy/asciidoc"

type delimitedBlockState map[asciidoc.DelimitedBlockType][]int

func (c *current) pushDelimitedLevel(bt asciidoc.DelimitedBlockType, level int) {
	//fmt.Fprintf(os.Stderr, "pushing %d on %d (currently %v)\n", level, bt, c.delimitedBlocks[bt])
	c.delimitedBlockState[bt] = append(c.delimitedBlockState[bt], level)
}

func (c *current) peekDelimitedLevel(bt asciidoc.DelimitedBlockType) int {
	var level int
	if levels, ok := c.delimitedBlockState[bt]; ok && len(levels) > 0 {
		level = levels[len(levels)-1]
		//fmt.Fprintf(os.Stderr, "peeked %d on %d (currently %v)\n", level, bt, c.delimitedBlocks[bt])
	}
	return level
}

func (c *current) popDelimitedLevel(bt asciidoc.DelimitedBlockType) {
	if levels, ok := c.delimitedBlockState[bt]; ok && len(levels) > 0 {
		//fmt.Fprintf(os.Stderr, "popped  %d (currently %v)\n", bt, c.delimitedBlocks[bt])
		c.delimitedBlockState[bt] = levels[:len(levels)-1]
	}
}

func delimitedLength(a any) (out int) {
	switch a := a.(type) {
	case []byte:
		out += len(a)
	case string:
		out += len(a)
	case []any:
		for _, a := range a {
			out += delimitedLength(a)
		}
	}
	return
}
