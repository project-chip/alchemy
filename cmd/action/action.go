package action

type Action struct {
	Disco Disco `cmd:"" help:"GitHub action for Matter spec documents"`
	ZAP   ZAP   `cmd:"" help:"GitHub action for Matter SDK ZAP XML"`
}
