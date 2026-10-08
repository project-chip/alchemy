package asciidoc

type child struct {
	parent ParentElement
}

func (c child) Parent() ParentElement {
	return c.parent
}

func (c *child) SetParent(e ParentElement) {
	c.parent = e
}
