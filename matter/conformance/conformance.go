package conformance

type Conformance interface {
	ASCIIDocString() string
	Description() string

	Type() Type
	Eval(context Context) (ConformanceState, error)
	Equal(oc Conformance) bool
	Clone() Conformance
}

type HasConformance interface {
	GetConformance() Set
}
