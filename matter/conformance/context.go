package conformance

type ValueStore interface {
	Reference(id string) HasConformance
}

type Context struct {
	Values            map[string]any
	Store             ValueStore
	VisitedReferences map[string]struct{}
}
