package python

type GeneratorOption func(g *PythonTestRenderer)

type GeneratorOptions struct {
	TemplateRoot string `aliases:"templateRoot" help:"the root of your local template files; if not specified, Alchemy will use an internal copy" group:"Test Script Options:"`
	Overwrite    bool   `default:"false" help:"overwrite existing test scripts"  group:"Test Script Options:"`
}

func (g GeneratorOptions) ToOptions() []GeneratorOption {
	return []GeneratorOption{
		TemplateRoot(g.TemplateRoot),
		Overwrite(g.Overwrite),
	}
}

func TemplateRoot(templateRoot string) func(*PythonTestRenderer) {
	return func(g *PythonTestRenderer) {
		g.templateRoot = templateRoot
	}
}

func Overwrite(overwrite bool) func(*PythonTestRenderer) {
	return func(g *PythonTestRenderer) {
		g.overwrite = overwrite
	}
}
