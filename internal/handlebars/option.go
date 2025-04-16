package handlebars

type TemplateOptions struct {
	TemplateRoot string `default:"" help:"the root of your local template files; if not specified, Alchemy will use an internal copy"`
}
