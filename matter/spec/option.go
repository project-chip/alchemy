package spec

type BuilderOption func(tg *Builder)

func IgnoreHierarchy(ignore bool) BuilderOption {
	return func(b *Builder) {
		b.ignoreHierarchy = ignore
	}
}

type BuilderOptions struct {
	IgnoreHierarchy bool `default:"false" help:"ignore hierarchy" group:"Spec:"`
}

type ParserOption func(p *Parser)

type ParserOptions struct {
	SpecRoot string `default:"connectedhomeip-spec" aliases:"specRoot" help:"the src root of your clone of CHIP-Specifications/connectedhomeip-spec"  group:"Spec:"`
	Inline   bool   `default:"false" help:"use inline parser"  group:"Spec:"`
}

func (po ParserOptions) ToOptions() []ParserOption {
	return []ParserOption{
		SpecRoot(po.SpecRoot),
		UseInlineParser(po.Inline),
	}
}

func UseInlineParser(useInline bool) ParserOption {
	return func(p *Parser) {
		p.inline = useInline
	}
}

func SpecRoot(specRoot string) ParserOption {
	return func(p *Parser) {
		p.Root = specRoot
	}
}
