package asciidoc

type position struct {
	path   string
	line   int
	column int
	offset int
}

func (p position) Path() string {
	return p.path
}

func (p *position) SetPath(path string) {
	p.path = path
}

func (p position) Origin() (path string, line int) {
	return p.path, p.line
}

func (p position) Position() (line int, column int, offset int) {
	line = p.line
	column = p.column
	offset = p.offset
	return
}

func (p *position) SetPosition(line int, column int, offset int) {
	p.line = line
	p.column = column
	p.offset = offset
}

type HasPosition interface {
	Path() string
	SetPath(path string)
	Position() (line int, column int, offset int)
	SetPosition(line int, column int, offset int)
}
