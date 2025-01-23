package spec

type BuilderOption func(tg *Builder)

func IgnoreHierarchy(ignore bool) BuilderOption {
	return func(b *Builder) {
		b.ignoreHierarchy = ignore
	}
}
