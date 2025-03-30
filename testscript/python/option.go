package python

import (
	"github.com/spf13/pflag"
)

type GeneratorOption func(g *PythonTestRenderer)

func Flags(flags *pflag.FlagSet) {
	flags.String("templateRoot", "", "the root of your local template files; if not specified, Alchemy will use an internal copy")
	flags.Bool("overwrite", true, "overwrite existing test scripts")
}

func GeneratorOptions(flags *pflag.FlagSet) (options []GeneratorOption) {
	overwrite, _ := flags.GetBool("overwrite")
	templateRoot, _ := flags.GetString("templateRoot")
	return []GeneratorOption{
		Overwrite(overwrite),
		TemplateRoot(templateRoot),
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
