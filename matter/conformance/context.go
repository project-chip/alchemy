package conformance

type Context interface {
	Value(identifier string) (any, bool)
	MarkVisit(identifier string) bool
	UnmarkVisit(identifier string)
}

type BasicContext struct {
	Values            map[string]any
	VisitedReferences map[string]struct{}
}

func (bc *BasicContext) Value(identifier string) (any, bool) {
	val, ok := bc.Values[identifier]
	return val, ok
}

func (bc *BasicContext) MarkVisit(identifier string) bool {
	if bc.VisitedReferences == nil {
		bc.VisitedReferences = make(map[string]struct{})
	} else if _, ok := bc.VisitedReferences[identifier]; ok {
		// We've already visited this reference in this evaluation chain, so it's recursive
		return false
	}
	bc.VisitedReferences[identifier] = struct{}{}
	return true
}

func (bc *BasicContext) UnmarkVisit(identifer string) {
	delete(bc.VisitedReferences, identifer)
}
