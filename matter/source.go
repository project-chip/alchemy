package matter

type Source interface {
	Origin() (path string, line int)
}
