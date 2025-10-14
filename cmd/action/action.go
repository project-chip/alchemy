package action

type Action struct {
	Comment    Comment    `cmd:""  help:"GitHub action for Matter spec documents"`
	Disco      Disco      `cmd:"" default:"" help:"GitHub action for Matter spec documents"`
	ZAP        ZAP        `cmd:"" help:"GitHub action for Matter SDK ZAP XML"`
	MergeGuard MergeGuard `cmd:"" help:"GitHub action to prevent Provisionality and Parse errors to be merged."`
}
