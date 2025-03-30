package spec

import "github.com/spf13/pflag"

type BuilderOption func(tg *Builder)

func IgnoreHierarchy(ignore bool) BuilderOption {
	return func(b *Builder) {
		b.ignoreHierarchy = ignore
	}
}

type ParserOption func(p *Parser)

func ParserFlags(flags *pflag.FlagSet) {
	flags.String("specRoot", "connectedhomeip-spec", "the src root of your clone of CHIP-Specifications/connectedhomeip-spec")
	flags.Bool("inline", false, "use inline parser")
}

func ParserOptions(flags *pflag.FlagSet) (options []ParserOption) {
	specRoot, _ := flags.GetString("specRoot")
	inline, _ := flags.GetBool("inline")

	options = append(options, SpecRoot(specRoot))
	options = append(options, UseInlineParser(inline))
	return
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
