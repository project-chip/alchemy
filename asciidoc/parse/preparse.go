package parse

import (
	"github.com/project-chip/alchemy/asciidoc"
)

type PreParseContext interface {
	IsSet(name string) bool
	Get(name string) any
	Set(name string, value any)
	Unset(name string)
	GetCounterState(name string, initialValue string) (*asciidoc.CounterState, error)
	ResolvePath(root string, path string) (asciidoc.Path, error)
	ShouldIncludeFile(path asciidoc.Path) bool
}
