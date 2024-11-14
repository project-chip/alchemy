package pics

import "strings"

type Expression interface {
	String() string
	PythonString() string
	PythonBuilder(aliases map[string]string, builder *strings.Builder)
}
