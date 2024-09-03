package asciidoc

type child struct {
	parent Element
}

func (c child) Parent() Element {
	return c.parent
}

func (c *child) SetParent(e Element) {
	c.parent = e
}
