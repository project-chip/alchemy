package parse

import "github.com/project-chip/alchemy/asciidoc"

func copyPosition[T asciidoc.HasPosition](source any, element T) T {
	switch source := source.(type) {
	case asciidoc.HasPosition:
		element.SetPath(source.Path())
		element.SetPosition(source.Position())
	case asciidoc.Elements:
		if len(source) > 0 {
			return copyPosition(source[0], element)
		}
	}
	return element
}
