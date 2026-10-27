package spec

func Patch(patcher Patcher) BuilderOption {
	return func(b *Builder) {
		b.patchers = append(b.patchers, patcher)
	}
}

type Patcher interface {
	BeforeParse(spec *Specification) error
	AfterParse(spec *Specification) error
	BeforeDataModelResolution(spec *Specification) error
	AfterDataModelResolution(spec *Specification) error
}
