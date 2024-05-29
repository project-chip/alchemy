package parse

import "github.com/hasty/alchemy/asciidoc"

func copyPosition[T asciidoc.HasPosition](source any, element T) T {
	switch source := source.(type) {
	case asciidoc.HasPosition:
		element.SetPath(source.Path())
		element.SetPosition(source.Position())
	case asciidoc.Set:
		if len(source) > 0 {
			return copyPosition(source[0], element)
		}
	}
	return element
}
