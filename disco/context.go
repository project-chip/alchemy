package disco

import "context"

type Context struct {
	context.Context

	potentialDataTypes map[string][]*potentialDataType
}

func NewContext(parent context.Context) *Context {
	return &Context{
		Context:            parent,
		potentialDataTypes: make(map[string][]*potentialDataType),
	}
}
