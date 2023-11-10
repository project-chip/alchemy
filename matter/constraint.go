package matter

type Constraint interface {
	AsciiDocString() string
	Equal(o Constraint) bool
	MinMax(c *ConstraintContext) (min ConstraintExtreme, max ConstraintExtreme)
}

type ConstraintLimit interface {
	AsciiDocString() string
	Equal(o ConstraintLimit) bool
	MinMax(c *ConstraintContext) (min ConstraintExtreme, max ConstraintExtreme)
}

type ConstraintContext struct {
	Fields            FieldSet
	VisitedReferences map[string]struct{}
}

type ConstraintExtremeType uint8

const (
	ConstraintExtremeTypeUndefined ConstraintExtremeType = iota
	ConstraintExtremeTypeInt64
	ConstraintExtremeTypeUInt64
)

type ConstraintExtreme struct {
	Type   ConstraintExtremeType
	Int64  int64
	UInt64 uint64
}

func (ce *ConstraintExtreme) Defined() bool {
	return ce.Type != ConstraintExtremeTypeUndefined
}
