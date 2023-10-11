package disco

import "context"

type discoContext struct {
	context.Context

	potentialDataTypes map[string][]*DataTypeEntry
}

func newContext(parent context.Context) *discoContext {
	return &discoContext{
		Context:            parent,
		potentialDataTypes: make(map[string][]*DataTypeEntry),
	}
}
