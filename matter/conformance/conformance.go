package conformance

type Conformance interface {
	AsciiDocString() string
	Description() string

	Type() Type
	Eval(context Context) (State, error)
	Equal(oc Conformance) bool
	Clone() Conformance
}

type HasConformance interface {
	GetConformance() Set
}
