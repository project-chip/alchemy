package conformance

type Context struct {
	Values            map[string]any
	VisitedReferences map[string]struct{}
}
